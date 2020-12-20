package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Tag 标签模型
type Tag struct{
	Model

	Name string `json:"name" valid:"name"`
	CreatedBy string `json:"created_by" valid:"created_by"`
	ModifiedBy string `json:"modified_by" valid:"modified_by"`
	State int `json:"state" valid:"state"`
}

// GetTags 获取标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return tags
}

// GetTagTotal 获取总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName 判断标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

// ExistTagByID 通过ID查找
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

// DeleteTag 删除
func DeleteTag(id int) bool {
    db.Where("id = ?", id).Delete(&Tag{})

    return true
}

// EditTag 编辑
func EditTag(id int, data interface {}) bool {
    db.Model(&Tag{}).Where("id = ?", id).Updates(data)

    return true
}

// AddTag 添加标签
func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag {
		Name: name,
		State: state,
		CreatedBy: createBy,
	})

	return true
}

// BeforeCreate 创建前回调
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate 更新回调
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}