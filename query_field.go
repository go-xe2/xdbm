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
	// 获取关系表达式
	GetJoin() []string
}


type queryFieldImp struct {
	// 规则，MD,ML,UD,UL,GD,GL,S
	rules 		[]string
	// 字段查询表达式
	express 	string
	// 如果是关联表字段，则有该值
	joinTable	string
	// 关系表达式
	join []string
}

// 创建查询字段
func NewQueryField(rules string, express string, joinTable string, join []string) QueryField {
	arr := strings.Split(strings.ToUpper(rules), ",")
	return &queryFieldImp{
		rules: 			arr,
		express: 		express,
		joinTable: 		joinTable,
		join: 			join,
	}
}

func (f *queryFieldImp) GetRules() []string {
	return f.rules
}

func (f *queryFieldImp) HasRule(rule string) bool {
	s := strings.Join(f.rules, "")
	if s == "" || s == "S" {
		return true
	}
	for _, s := range f.rules {
		if strings.ToUpper(s) == strings.ToUpper(rule) {
			return true
		}
	}
	return false
}

func (f *queryFieldImp) GetExpress() string {
	return f.express
}

func (f *queryFieldImp) GetJoinTable() string {
	return f.joinTable
}

func (f *queryFieldImp) GetJoin() []string {
	return f.join
}