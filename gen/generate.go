package gen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"gqlgen-generator/config"
	"gqlgen-generator/gen/graphql"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/types"
	"gqlgen-generator/gen/table"
	"gqlgen-generator/gen/template"
	"gqlgen-generator/gen/tools"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ColumnType uint

func Generate() {
	exist, err := PathExists(filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir))
	if err != nil {
		log.Fatal(err)
	}
	if exist {
		log.Fatal("文件夹不为空，不能执行generate")
	}
	// 解析表信息
	tableInfos, err := table.ResolverTableInfos(config.CONF_INSTANCE.GenerateConf.Tables)
	if err != nil {
		log.Fatal(err)
	}
	err = ensureAllFile()
	if err != nil {
		log.Fatal(err)
	}

	// 生成用于生成model的gqlgen.yml
	err = generateGqlgenYmlFile()
	if err != nil {
		log.Fatal(err)
	}
	// 生成只包含model的graphql文件
	for _, tableInfo := range tableInfos {
		err = generateModelFile(tableInfo)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = ensureAllFile()
	if err != nil {
		log.Fatal(err)
	}
	// 生成graphql文件
	for _, tableInfo := range tableInfos {
		err = generateGraphqlFile(tableInfo)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = generateOtherGraphqlFile()
	if err != nil {
		log.Fatal(err)
	}
	// 执行gqlgen generate
	_, err = tools.RunInDir("go run github.com/99designs/gqlgen generate.", filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir))
	if err != nil {
		log.Fatal(err)
	}
	// 覆盖resolver函数
	err = regenerateResolverFiles(tableInfos)
	if err != nil {
		log.Fatal(err)
	}
	// 生成endpoint.go文件
	err = generateEndpointFile()
	if err != nil {
		log.Fatal(err)
	}
}

func generateGqlgenYmlFile() error {
	templateContext := template.TemplateContext{
		Conf: config.CONF_INSTANCE,
	}
	gqlgenYmlPath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "gqlgen.yml")
	return tools.WriteTemplate(template.GQLGEN_YML, gqlgenYmlPath, templateContext)
}

//生成所有文件夹
func ensureAllFile() error {
	// 生成工程目录
	err := ensureFile(config.CONF_INSTANCE.ProjectDir)
	if err != nil {
		return err
	}
	// 生成模块目录
	err = ensureFile(filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir))
	if err != nil {
		return err
	}
	// 生成graphql schema文件目录
	err = ensureFile(filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph/graphqls"))
	if err != nil {
		return err
	}
	// 生成model目录
	err = ensureFile(filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "model"))
	if err != nil {
		return err
	}
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 生成文件夹
func ensureFile(filePath string) error {
	return os.MkdirAll(filePath, 0777)
}

// 生成其他公用的graphql文件
func generateOtherGraphqlFile() error {
	graphqlPath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph/graphqls", "common.graphqls")
	return ioutil.WriteFile(graphqlPath, []byte(template.COMMON_GRAPHQLS), 0777)
}

// 生成graphql文件
func generateGraphqlFile(info table.TableInfo) error {
	graphqlText := graphql.GenerateGraphqlText(&info)
	graphqlPath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph/graphqls", info.TableName+".graphqls")
	return ioutil.WriteFile(graphqlPath, []byte(graphqlText), 0777)
}

// 生成用于生成model的graphql文件
func generateModelGraphqlFile(info table.TableInfo) error {
	graphqlText := graphql.GenerateModelGraphqlText(&info)
	graphqlPath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph/graphqls", info.TableName+"_model.graphqls")
	return ioutil.WriteFile(graphqlPath, []byte(graphqlText), 0777)
}

// 重新生成resolver文件
func regenerateResolverFiles(infos []table.TableInfo) error {
	for i, info := range infos {
		// 是否声明resolver struct
		declareResolver := false
		if i == 0 {
			declareResolver = true
		}
		context := template.TemplateContext{
			Conf:         config.CONF_INSTANCE,
			CurrentTable: info,
			Params: map[string]interface{}{
				// model的golang结构名称
				"ModelName": types.TableModelName(info.TableName),
				// 主键的golang类型
				"PrimaryKeyType": calculateGolangTypeFromGqlType(scalars.GetPrimaryKeyType(info)),
				// 主键的数据库名称
				"PrimaryKeyColumnName": info.PrimaryKeyColumn().Columnname,
				// 主键的golang对象名称
				"PrimaryKeyName": strcase.ToCamel(info.PrimaryKeyColumn().Columnname),
				// 是否声明resolver struct
				"DeclareResolver": declareResolver,
			},
		}
		resolverFilePath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph/resolver", info.TableName+".resolvers.go")
		err := tools.WriteTemplate(template.TABLE_RESOLVER_GO, resolverFilePath, context)
		if err != nil {
			return err
		}
	}
	return nil
}

// 生成endpoint文件
func generateEndpointFile() error {
	context := template.TemplateContext{
		Conf: config.CONF_INSTANCE,
	}
	resolverFilePath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "endpoint.go")
	err := tools.WriteTemplate(template.ENDPOINT_GO, resolverFilePath, context)
	return err
}

// 生成model文件
func generateModelFile(tableInfo table.TableInfo) error {
	err := ensureAllFile()
	if err != nil {
		return err
	}
	err = generateModelGraphqlFile(tableInfo)
	if err != nil {
		return err
	}
	err = generateOtherGraphqlFile()
	if err != nil {
		return err
	}
	// 执行gqlgen generate 生成model
	_, err = tools.RunInDir("go run github.com/99designs/gqlgen generate.", filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir))
	if err != nil {
		return err
	}
	// 移动model文件
	sourceModelFilePath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph/model/models_gen.go")
	targetModelFilePath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "model/"+tableInfo.TableName+".go")
	err = os.Rename(sourceModelFilePath, targetModelFilePath)
	if err != nil {
		return err
	}
	// 生成dataloader文件
	tableModuleName := types.TableModelName(tableInfo.TableName)
	generateDataLoaderFileCmd := fmt.Sprintf("go run github.com/vektah/dataloaden %sLoader string *%s/%s/model.%s", tableModuleName, config.CONF_INSTANCE.ProjectName, config.CONF_INSTANCE.ModuleDir, tableModuleName)
	fmt.Println(generateDataLoaderFileCmd)
	out, err := tools.RunInDir(generateDataLoaderFileCmd, filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "model"))
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	// 生成model ext文件
	modelExtFilePath := filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "model/"+tableInfo.TableName+"_ext.go")
	context := template.TemplateContext{
		Conf: config.CONF_INSTANCE,
		Params: map[string]interface{}{
			"DataloaderCmd": generateDataLoaderFileCmd,
			"ModelName":     tableModuleName,
		},
	}
	err = tools.WriteTemplate(template.MODEL_EXT_GO, modelExtFilePath, context)
	if err != nil {
		return err
	}
	// 删除生成的文件
	err = os.RemoveAll(filepath.Join(config.CONF_INSTANCE.ProjectDir, config.CONF_INSTANCE.ModuleDir, "graph"))
	if err != nil {
		return err
	}
	return nil
}

// 翻译gqlType到golangType
func calculateGolangTypeFromGqlType(gqlType string) string {
	switch gqlType {
	case scalars.SCALAR_INT, scalars.SCALAR_BIGINT:
		return "int64"
	}
	return "string"
}

// 追加内容到文件
func appendFile(filePath, content string) error {
	fd, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	buf := []byte(content)
	_, err := fd.Write(buf)
	if err != nil {
		return err
	}
	err = fd.Close()
	return err
}
