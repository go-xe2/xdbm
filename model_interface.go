package xdbm


type IModel interface {
	// 表名
	TableName() string
	// 表名别名
	AliasName() string
	// 表命名
	TableAlias(alias ...string) string
	// 字段列表
	Fields() map[string]interface{}
	// 带字段别名
	Field(fieldName string) string
}
