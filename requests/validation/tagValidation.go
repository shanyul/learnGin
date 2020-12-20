package validation

import (
	"learngo/models"
	"github.com/thedevsaddam/govalidator"
)


// TagForm 标签表单验证
func TagForm(data models.Tag) map[string][]string {
	rules := govalidator.MapData{
		"name":       []string{"required", "min:1", "max:40"},
		"state":      []string{"required"},
		"created_by": []string{"required", "max:40"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:标题为必填项",
			"min:标题长度需大于 1",
			"max:标题长度需小于 40",
		},
		"state": []string{
			"required:状态为必填项",
		},
		"created_by": []string{
			"required:创建人为必填项",
			"min:长度需小于 40",
		},
	}


	// 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // struct 标签标识符
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}

// EditTagForm 标签表单验证
func EditTagForm(data models.Tag) map[string][]string {
	rules := govalidator.MapData{
		"name":       []string{"required", "min:1", "max:40"},
		"state":      []string{"required"},
		"modified_by": []string{"required", "max:40"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:标题为必填项",
			"min:标题长度需大于 1",
			"max:标题长度需小于 40",
		},
		"state": []string{
			"required:状态为必填项",
		},
		"modified_by": []string{
			"required:修改人为必填项",
			"min:长度需小于 40",
		},
	}


	// 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // struct 标签标识符
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}