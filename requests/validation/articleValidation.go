package validation

import (
	"learngo/models"

	"github.com/thedevsaddam/govalidator"
)

// ArticleForm 文章表单验证
func ArticleForm(data models.Article) map[string][]string {
	rules := govalidator.MapData{
		"tag_id":     []string{"required"},
		"title":      []string{"required"},
		"desc":       []string{"required"},
		"content":    []string{"required"},
		"created_by": []string{"required"},
		"state":      []string{"required"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"tag_id": []string{
			"required:必须选择所属标签",
		},
		"title": []string{
			"required:标题为必填项",
		},
		"desc": []string{
			"required:描述为必填项",
		},
		"content": []string{
			"required:内容为必填项",
		},
		"created_by": []string{
			"required:创建人为必填项",
		},
		"state": []string{
			"required:状态为必填项",
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

// EditArticleForm 标签表单验证
func EditArticleForm(data models.Article) map[string][]string {
	rules := govalidator.MapData{
		"tag_id":      []string{"required"},
		"title":       []string{"required", "max: 64"},
		"desc":        []string{"required"},
		"content":     []string{"required"},
		"modified_by": []string{"required", "max:40"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"tag_id": []string{
			"required:必须选择所属标签",
		},
		"modified_by": []string{
			"required:创建人为必填项",
			"min:长度需小于 40",
		},
		"title": []string{
			"required:标题为必填项",
			"min:长度需小于 40",
		},
		"desc": []string{
			"required:描述为必填项",
		},
		"content": []string{
			"required:内容为必填项",
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
