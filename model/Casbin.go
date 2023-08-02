package model

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityName string       `json:"authority_name"`
	Group         string       `json:"group"`
	CasbinInfos   []CasbinInfo `json:"casbinInfos"`
}

func UpdateCasbin(AuthorityName string, casbinInfos []CasbinInfo) error {
	rules := [][]string{}
	for _, casbinInfo := range casbinInfos {
		rules = append(rules, []string{AuthorityName, casbinInfo.Path, casbinInfo.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicy(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func UpdateGroupCasbin(AuthorityName string, Group string) error {
	e := Casbin()
	success, _ := e.AddGroupingPolicy(AuthorityName, Group)
	if !success {
		return errors.New("存在相同group,添加失败,请联系管理员")
	}
	return nil
}

func DeleteRoleForUser(AuthorityName string, Group string) error {
	e := Casbin()
	success, _ := e.DeleteRoleForUser(AuthorityName, Group)
	if !success {
		return errors.New("不存在该role,添加失败,请联系管理员")
	}
	return nil
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		// 创建gorm适配器
		db.Table("ginblog_casbin_rule").AutoMigrate(&gormadapter.CasbinRule{})
		a, err := gormadapter.NewAdapterByDBUseTableName(db, "", "ginblog_casbin_rule")
		if err != nil {
			return
		}
		// model配置
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
		// 创建model
		m, err := model.NewModelFromString(text)
		if err != nil {
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60) // 设置缓存中的数据过期时间，过期了再从数据库中拿
		syncedCachedEnforcer.LoadPolicy()           // 加载数据库的策略到变量中
		return
	})
	return syncedCachedEnforcer
}
