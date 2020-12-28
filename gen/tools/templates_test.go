package tools

import (
	"fmt"
	"gqlgen-generator/config"
	"gqlgen-generator/gen/table"
	"gqlgen-generator/gen/template"
	"os"
	"testing"
)

func TestWriteTemplate(t *testing.T) {
	content := template.GQLGEN_YML
	context := template.TemplateContext{
		Conf: &config.Conf{
			GenerateConf: config.GenerateConf{
				ProjectDir:  "test001",
				ModuleDir:   "module",
				ModelDir:    "",
				ProjectName: "test001",
			},
		},
		AllTable:     nil,
		CurrentTable: table.TableInfo{},
		Params:       map[string]interface{}{"ModelFileName": "model.enterprise"},
	}
	err := os.MkdirAll(`D:/generator/template-test001`, 0777)
	if err != nil {
		fmt.Println(err)
	}
	filePath := `D:/generator/template-test001/gqlgen.yml`
	err = WriteTemplate(content, filePath, context)
	if err != nil {
		fmt.Println(err)
	}
}
