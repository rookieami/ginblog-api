package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//文章表
type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //分类
	gorm.Model
	Title string `gorm:"type:"varchar(100);not null" json:"title"` //标题

	Cid     int    `gorm:"type:int ;not null" json:"cid"`
	Desc    string `gorm:"type:"varchar(200)" json:"desc"` //简述
	Content string `gorm:"type:"longtext" json:"content"`  //文章内容
	Img     string `gorm:"type:"varchar(100)" json:"img"`  //文章内图片
}

//新增文章
func CreateAtr(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询文章列表
func GetAtr(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int //总量
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

//查询分类下所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtList []Article
	var total int //总量
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

//查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS

}

//编辑文章
func EditArt(id int, data *Article) int {
	var article []Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&article).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//删除文章
func DeleteArt(id int) int {
	var article Article

	err := db.Where("id=?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
