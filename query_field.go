package xdbm

import (
	"strings"
)

type QueryField interface {
	// 获取规则
	GetRules() []string
	// 检查是否存在规则
	HasRule(rule string) bool
	// 获取字段查询表达式
	GetExpress() string
	// 获取关系表
	GetJoinTable() string
	GetTableAlias() string
	// 获取关系表达式
	GetJoin() []string
}


type queryFieldImp struct {
	// 规则，MD,ML,UD,UL,GD,GL,S
	Rules 		[]string
	// 字段查询表达式
	Express 	string
	// 如果是关联表字段，则有该值
	JoinTable	string
	TableAlias string
	// 关系表达式
	Join []string
}

// 创建查询字段
func NewQueryField(rules string, express string, joinTable string, alias string, join []string) QueryField {
	arr := strings.Split(strings.ToUpper(rules), ",")
	return &queryFieldImp{
		Rules: 			arr,
		Express: 		express,
		JoinTable: 		joinTable,
		Join: 			join,
		TableAlias: alias,
	}
}

func (f *queryFieldImp) GetRules() []string {
	return f.Rules
}

func (f *queryFieldImp) HasRule(rule string) bool {
	s := strings.Join(f.Rules, "")
	if rule == "S" {
		return strings.Index(s, rule) >= 0
	}
	return s == "" || s == "S" || strings.Index(s, rule) >= 0
}

func (f *queryFieldImp) GetExpress() string {
	return f.Express
}

func (f *queryFieldImp) GetJoinTable() string {
	return f.JoinTable
}

func (f *queryFieldImp) GetTableAlias() string {
	return f.TableAlias
}

func (f *queryFieldImp) GetJoin() []string {
	return f.Join
}