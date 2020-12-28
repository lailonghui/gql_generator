package template

import (
	"gqlgen-generator/config"
	"gqlgen-generator/gen/table"
)

// 模块上下文
type TemplateContext struct {
	// 配置信息
	Conf *config.Conf
	// 所有表信息
	AllTable []table.TableInfo
	// 当前表信息
	CurrentTable table.TableInfo
	// 其他信息
	Params map[string]interface{}
}
