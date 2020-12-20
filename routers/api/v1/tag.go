package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"learngo/models"
	"learngo/pkg/e"
	"learngo/pkg/setting"
	"learngo/pkg/util"
	"learngo/requests/validation"
)

// GetTags 获取多个标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTag 添加标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	_tag := models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}

	errors := validation.TagForm(_tag)

	code := e.INVALID_PARAMS
	if len(errors) == 0 {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": errors,
	})
}

// EditTag 修改标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	state := com.StrTo(c.Query("state")).MustInt()

	_tag := models.Tag{
		Name: name,
		ModifiedBy: modifiedBy,
		State: state,
	}

	code := e.INVALID_PARAMS
	errors := validation.EditTagForm(_tag)

	if len(errors) == 0 {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
	
			models.EditTag(id, data)
		}else{
			code = e.ERROR_NOT_EXIST_TAG
		}
	}


	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": errors,
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var errors map[string]string
	if id < 0 {
		errors["id"] = "ID必须大于0"
	}
	code := e.INVALID_PARAMS

	if len(errors) == 0{
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		}else{
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": errors,
	})
}
