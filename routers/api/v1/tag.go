package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

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
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

// AddTag 添加标签
func AddTag(c *gin.Context) {
	/*name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation*/
}

// EditTag 修改标签
func EditTag(c *gin.Context) {

}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {

}