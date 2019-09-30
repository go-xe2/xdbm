package xdbm

import "strings"

// 实体模型关系接口
type IModelRelation interface {
	// 获取连接类型
	GetJoin() string
	// 获取关联模型
	GetModel() IModel
	// 获取关系
	GetRelation() string
	// 获取查询表达式
	GetExpress() string
	// rule
	GetRule() string
}

// 实体模型关联关系
type modelRelation struct {
	// 连接类型,join, left join, right join, inner join
	joinType string
	// 关联的模型
	model IModel
	// 实体模型关系
	relation string
	// 查询表达式，如case when $1.status > 0 then $2.name when $2.alias end
	// $1表示主表别名，$2表示关系表别名
	// 如果为空，表示关系表该键值的名称
	express string
	// MD,ML,UD,UL,GD,GL
	rule string
}

func NewModelRelation(rule string, model IModel, on string, join string, express ...string) IModelRelation {
	s := ""
	if len(express) > 0 {
		s = strings.Join(express, " ")
	}
	return &modelRelation{
		joinType: join,
		relation: on,
		model: model,
		express: s,
		rule: rule,
	}
}

// 获取连接类型
func (r *modelRelation) GetJoin() string {
	return r.joinType
}

// 获取关联模型
func (r *modelRelation) GetModel() IModel {
	return r.model
}

// 获取连接关系
func (r *modelRelation) GetRelation() string {
	return r.relation
}

func (r *modelRelation) GetExpress() string {
	return r.express
}

func (r *modelRelation) GetRule() string {
	return r.rule
}
