package template

var TABLE_RESOLVER_GO = `
package resolver
import (
	"{{.Conf.ProjectName}}/internal/db"
	"{{.Conf.ProjectName}}/{{.Conf.ModuleDir}}/graph/generated"
	"{{.Conf.ProjectName}}/{{.Conf.ModuleDir}}/graph/model"
	model1 "{{.Conf.ProjectName}}/{{.Conf.ModuleDir}}/model"
	"{{.Conf.ProjectName}}/pkg/graphql/util"
	util2 "{{.Conf.ProjectName}}/pkg/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) Delete{{.Params.ModelName}}(ctx context.Context, where model.{{.Params.ModelName}}BoolExp) (*model.{{.Params.ModelName}}MutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.{{.Params.ModelName}}{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.{{.Params.ModelName}}
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.{{.Params.ModelName}}MutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) Delete{{.Params.ModelName}}ByPk(ctx context.Context, {{.Params.PrimaryKeyName}} {{.Params.PrimaryKeyType}}) (*model1.{{.Params.ModelName}}, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.{{.Params.ModelName}}
	tx := db.DB.Model(&model1.{{.Params.ModelName}}{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("{{.Params.PrimaryKeyColumnName}} = ?", {{.Params.PrimaryKeyName}}).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) Insert{{.Params.ModelName}}(ctx context.Context, objects []*model.{{.Params.ModelName}}InsertInput) (*model.{{.Params.ModelName}}MutationResponse, error) {
	rs := make([]*model1.{{.Params.ModelName}}, 0)
	for _, object := range objects{
		v := &model1.{{.Params.ModelName}}{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.{{.Params.ModelName}}{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.{{.Params.ModelName}}MutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) Insert{{.Params.ModelName}}One(ctx context.Context, object model.{{.Params.ModelName}}InsertInput) (*model1.{{.Params.ModelName}}, error) {
	rs := &model1.{{.Params.ModelName}}{

	}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.{{.Params.ModelName}}{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) Update{{.Params.ModelName}}(ctx context.Context, inc *model.{{.Params.ModelName}}IncInput, set *model.{{.Params.ModelName}}SetInput, where model.{{.Params.ModelName}}BoolExp) (*model.{{.Params.ModelName}}MutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.{{.Params.ModelName}}{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.{{.Params.ModelName}}MutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.{{.Params.ModelName}}MutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) Update{{.Params.ModelName}}ByPk(ctx context.Context, inc *model.{{.Params.ModelName}}IncInput, set *model.{{.Params.ModelName}}SetInput, {{.Params.PrimaryKeyName}} {{.Params.PrimaryKeyType}}) (*model1.{{.Params.ModelName}}, error) {
	tx := db.DB.Where("{{.Params.PrimaryKeyColumnName}} = ?", {{.Params.PrimaryKeyName}})
	qt := util.NewQueryTranslator(tx, &model1.{{.Params.ModelName}}{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.{{.Params.ModelName}}
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) {{.Params.ModelName}}(ctx context.Context, distinctOn []model.{{.Params.ModelName}}SelectColumn, limit *int, offset *int, orderBy []*model.{{.Params.ModelName}}OrderBy, where *model.{{.Params.ModelName}}BoolExp) ([]*model1.{{.Params.ModelName}}, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.{{.Params.ModelName}}{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.{{.Params.ModelName}}
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) {{.Params.ModelName}}Aggregate(ctx context.Context, distinctOn []model.{{.Params.ModelName}}SelectColumn, limit *int, offset *int, orderBy []*model.{{.Params.ModelName}}OrderBy, where *model.{{.Params.ModelName}}BoolExp) (*model.{{.Params.ModelName}}Aggregate, error) {
	var rs model.{{.Params.ModelName}}Aggregate

	qt := util.NewQueryTranslator(db.DB, &model1.{{.Params.ModelName}}{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Error
	return &rs, err
}

func (r *queryResolver) {{.Params.ModelName}}ByPk(ctx context.Context, {{.Params.PrimaryKeyName}} {{.Params.PrimaryKeyType}}) (*model1.{{.Params.ModelName}}, error) {
	var rs model1.{{.Params.ModelName}}
	tx := db.DB.Model(&model1.{{.Params.ModelName}}{}).First(&rs, {{.Params.PrimaryKeyName}})
	err := tx.Error
	return &rs, err
}
{{if .Params.DeclareResolver}}
// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
{{end}}
`
