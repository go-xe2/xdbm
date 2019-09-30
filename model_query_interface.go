package xdbm

type IModelQuery interface {
	// 允许排序字段
	AllowSortFields() map[string]bool
}
