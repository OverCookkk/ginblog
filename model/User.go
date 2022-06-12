package model

import (
	"encoding/base64"
	"fmt"
	"ginblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"` //validate配合validate包做用户验证
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var users User                               //声明一个模型，告诉数据库，要操作的是这个模型
	db.Where("username = ?", name).First(&users) //First：取查询出的第一个值
	if users.Username == "" {
		return errmsg.SUCCSE
	} else {
		return errmsg.ERROR_USERNAME_USER
	}
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	fmt.Println(data.Username)

	var users []User
	err := db.Model(users).Where("id = ?", id).Updates(maps).Error //Updates + map同事修改多个字段
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user []User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	// var code int
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USESR_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
