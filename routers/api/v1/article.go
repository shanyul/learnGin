package v1

import (
	"learngo/models"
	"learngo/pkg/e"
	"learngo/pkg/setting"
	"learngo/pkg/util"
	"learngo/requests/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	var errors map[string]interface{}
	id := com.StrTo(c.Param("id")).MustInt()
	if id < 1 {
		errors["id"] = "ID必须大于零"
	}

	code := e.INVALID_PARAMS
	var data interface{}
	if len(errors) == 0 {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"error": errors,
		"data":  data,
	})
}

//GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	if arg := c.Query("tag_id"); arg != "" {
		tagID := com.StrTo(arg).MustInt()
		maps["tag_id"] = tagID
	}

	data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	_article := models.Article{
		TagID:     tagID,
		Title:     title,
		Desc:      desc,
		Content:   content,
		CreatedBy: createdBy,
		State:     state,
	}

	errors := validation.ArticleForm(_article)

	// var errors map[string][]string

	code := e.INVALID_PARAMS
	if len(errors) == 0 {
		if models.ExistTagByID(tagID) {
			data := make(map[string]interface{})
			data["tag_id"] = tagID
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"error": errors,
		"data":  make(map[string]interface{}),
	})
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	article := models.Article{
		TagID:      tagID,
		Title:      title,
		Desc:       desc,
		Content:    content,
		ModifiedBy: modifiedBy,
	}

	errors := validation.EditArticleForm(article)
	code := e.INVALID_PARAMS
	if models.ExistArticleByID(id) {
		if models.ExistTagByID(tagID) {
			data := make(map[string]interface{})
			if tagID > 0 {
				data["tag_id"] = tagID
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}

			data["modified_by"] = modifiedBy

			models.EditArticle(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"error": errors,
		"data":  make(map[string]interface{}),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.INVALID_PARAMS
	if id > 0 {
		code = e.SUCCESS
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
