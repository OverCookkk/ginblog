package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category                       //声明一个模型，告诉数据库，要操作的是这个模型
	db.Where("name = ?", name).First(&cate) //First：取查询出的第一个值
	if cate.Name == "" {
		return errmsg.SUCCSE
	} else {
		return errmsg.ERROR_CATENAME_USER
	}
}

// 新增分类
func CreateCategory(cate *Category) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询分类列表
func GetCategorys(pageSize int, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类
func EditCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	var cate []Category
	err := db.Model(cate).Where("id = ?", id).Updates(maps).Error //Updates + map同事修改多个字段
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var cate []Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
