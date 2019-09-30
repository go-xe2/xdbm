package xdbm

// 数据变动规则
type IModelMutation interface {
	// 添加规则
	AppendRules() []string
	// 更新规则
	UpdateRules() []string
}