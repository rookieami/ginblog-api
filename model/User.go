package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

//用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT :2" json:"role" validate:"required,gte=2" label:"角色码"` //角色 1管理员,2用户
}

//相关方法

//查询用户是否存在
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username=?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS //200
}

//新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int, int) {
	var users []User
	var total int //总量
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.INVALID_PARAMS, 0
	}
	return users, errmsg.SUCCESS, total
}

//编辑用户信息,只限定密码之外的可以修改
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//删除用户
func DeleteUser(id int) int {
	var user User
	//UPDATE `ginblog_user` SET `deleted_at`='2020-09-28 22:45:17'
	//WHERE `ginblog_user`.`deleted_at` IS NULL AND ((id=2))
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//gorm钩子函数,在保存前对密码进行加密
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

//密码加密
func ScryptPw(password string) string {
	const keyLen = 10
	//加盐
	salt := make([]byte, 8)
	salt = []byte{12, 31, 22, 4, 5, 54, 232, 43}
	//加密
	HansPs, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatalln(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HansPs)
	return fpw
}

//登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		//不存在
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG //密码错误
	}
	if user.Role == 1 {
		return errmsg.ERROR_USER_NO_RIGHT //没权限
	}
	return errmsg.SUCCESS
}
