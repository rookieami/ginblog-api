package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

//配置数据库初始化参数

//数据库初始化连接
var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"` //id
	CreatedOn  int `json:"created_on"`            //创建时间
	ModifiedOn int `json:"modified_on"`           //修改
}

func Init() {
	var (
		err                                             error
		dbType, dbName, user, passwd, host, tablePrefix string
	)

	dbType = utils.Conf.Type
	dbName = utils.Conf.Name
	user = utils.Conf.User
	passwd = utils.Conf.Password
	host = utils.Conf.Host
	tablePrefix = utils.Conf.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		passwd,
		host,
		dbName))
	if err != nil {
		log.Printf("数据库连接失败,err:%v\n", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//禁用默认表的附属形式
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Article{}, &Category{}) //迁移模型
	//db.LogMode(true)//打印数据库操作
	//连接池最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	//最大连接数
	db.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	defer db.Close()
}
