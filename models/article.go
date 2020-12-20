package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 文章模型
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy	string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `josn:"state"`
}

// ExistArticleByID 文章是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

// GetArticleTotal 获取文章总数
func GetArticleTotal(maps interface{}) (count int)  {
	db.Model(&Article{}).Where(maps).Count(&count)

	return count
}

// GetArticles 分页获取文章
func GetArticles(pageNum int, pageSize int, maps interface{}) (article []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)

	return
}

// GetArticle 获取指定文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).Find(&article)

	return
}
                                                 
// BeforeCreate 创建前置钩子
func (a *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil                      
}

// BeforeUpdate 更新前置钩子
func (a *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}