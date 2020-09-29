package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//分类
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"` //主键且自增
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//相关方法
//查询分类是否存在
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name=?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED //2001
	}
	return errmsg.SUCCESS //200
}

//新增标签分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int, int) {
	var category []Category
	var total int //总量
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&category).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.INVALID_PARAMS, 0
	}
	return category, errmsg.SUCCESS, total
}

//编辑分类
func EditCategory(id int, data *Category) int {
	var category Category
	err := db.Model(&category).Where("id=?", id).Update(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//删除分类标签
func DeleteCategory(id int) int {
	var category Category

	err := db.Where("id=?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
