package template

var MODEL_EXT_GO = `
package model
import (
	"{{.Conf.ProjectName}}/internal/db"
)

// go:generate {{.Params.DataloaderCmd}}

func (t {{.Params.ModelName}}) TableName() string {
	return "{{.Params.ModelName}}"
}

func (t *{{.Params.ModelName}}) NewLoader() *{{.Params.ModelName}}Loader {
	return &{{.Params.ModelName}}Loader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*{{.Params.ModelName}}, []error) {
			var rs []*{{.Params.ModelName}}
			// TODO 按实际需要实现
			db.DB.Model(&{{.Params.ModelName}}{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

`
