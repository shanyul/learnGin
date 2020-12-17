package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"github.com/unknwon/com"

	"learngo/models"
	"learngo/pkg/e"
	"learngo/pkg/setting"
	"learngo/pkg/util"
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
		Name:  name,
		State:  state,
		CreatedBy: createdBy,
	}

	rules := govalidator.MapData{
		"name": []string{"required", "min:1", "max:40"},
		"created_by":  []string{"required", "min:40"},
		"state":  []string{"required"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:标题为必填项",
			"min:标题长度需大于 1",
			"max:标题长度需小于 40",
		},
		"created_by": []string{
			"required:创建人为必填项",
			"min:长度需大于 40",
		},
		"state": []string{
			"required:状态为必填项",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          &_tag,
		Rules:         rules,
		TagIdentifier: "valid", // struct 标签标识符
		Messages:      messages,
	}

	errors := govalidator.New(opts).ValidateStruct()

	if len(errors) != 0 {
		log.Fatal(errors)
	}

	code := e.INVALID_PARAMS
	if !models.ExistTagByName(name) {
		code = e.SUCCESS
		models.AddTag(name, state, createdBy)
	} else {
		code = e.ERROR_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag 修改标签
func EditTag(c *gin.Context) {

}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {

}
