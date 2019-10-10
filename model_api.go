package xdbm

import (
	"fmt"
	"github.com/go-xe2/xorm"
	"strings"
)

var (
	modelFieldCaches = make(map[string]map[string]QueryField)
)

// 获取模型查询字段定义
func GetModelFields(m IModel) map[string]QueryField {
	modelName := m.TableName()
	modelAlias := m.AliasName()
	if q, ok := modelFieldCaches[modelName]; ok {
		return q
	}

	fields := m.Fields()
	var results = make(map[string]QueryField)
	for key, cfg := range fields {
		switch val := cfg.(type) {
		case IModelRelation:
			// 关系型字段
			model := val.GetModel()
			joinType 	:= val.GetJoin()
			joinExpr 	:= val.GetExpress()
			rule 		:= val.GetRule()
			joinOn 		:= val.GetRelation()
			joinTable	:= model.TableName()
			joinAlias	:= model.AliasName()

			var sqlExpr = ""
			if joinExpr == key || joinExpr == "" {
				sqlExpr = fmt.Sprintf("%s.%s", joinAlias, key)
			} else {
				joinExpr = strings.Replace(joinExpr, "$1", modelAlias, -1)
				joinExpr = strings.Replace(joinExpr, "$2", joinAlias, -1)
				sqlExpr = fmt.Sprintf("%s as %s", joinExpr, key)
			}
			joinOn = strings.Replace(joinOn, "$1", modelAlias, -1)
			joinOn = strings.Replace(joinOn, "$2", joinAlias, -1)
			join := []string{
				fmt.Sprintf("%s %s", joinTable, joinAlias),
				joinOn,
				joinType,
			}
			item := NewQueryField(rule, sqlExpr, joinTable, join)
			results[key] = item
			break
		case string:
			if val == "-" {
				// 过滤
				continue
			}
			fdName := fmt.Sprintf("%s.%s", modelAlias, key)
			if val != "" && val != key {
				fdName = fmt.Sprintf("%s.%s as %s", modelAlias, key, val)
			}
			item := NewQueryField("", fdName, "", nil)
			results[key] = item
			break
		case []interface{}:
			aLen := len(val)
			switch aLen {
			case 3:
				rule, ok1 := val[0].(string)
				expr, ok2 := val[1].(string)
				alias, ok3 := val[2].(string)
				if !ok3 {
					alias = ""
				}
				if ok1 && ok2 {
					sqlExpr := ""
					if sqlExpr == key {
						sqlExpr = fmt.Sprintf("%s.%s", modelAlias, key)
					} else {
						expr = strings.Replace(expr, "$1", modelAlias, -1)
						sqlExpr = fmt.Sprintf("%s as %s", expr, key)
						if alias != "" {
							sqlExpr = fmt.Sprintf("%s as %s", expr, alias)
						}
					}
					item := NewQueryField(rule, sqlExpr, "", nil)
					results[key] = item
				}
				break
			case 2:
				rule, ok1 := val[0].(string)
				expr, ok2 := val[1].(string)
				if ok1 && ok2 {
					expr = strings.Replace(expr, "$1", modelAlias, -1)
					sqlExpr := fmt.Sprintf("%s.%s", modelAlias, key)
					if expr != key {
						sqlExpr = fmt.Sprintf("%s as %s", expr, key)
					}
					item := NewQueryField(rule, sqlExpr, "", nil)
					results[key] = item
				}
				break
			case 1:
				var item QueryField
				if s, ok := val[0].(string); ok {
					item = NewQueryField(s, fmt.Sprintf("%s.%s",modelAlias, key), "", nil)
				} else {
					item = NewQueryField("", fmt.Sprintf("%s.%s", modelAlias, key), "", nil)
				}
				results[key] = item
				break
			}
			break
		}
	}
	modelFieldCaches[modelName] = results
	return results
}


// 选择查询模型字段列表
func GetQueryFields(m IModel, rule string, selectFields ...interface{}) (string, [][]string) {
	var selects = make(map[string]interface{})

	if len(selectFields) > 0 {
		switch val := selectFields[0].(type) {
		case string:
			arr := strings.Split(val, ",")
			for _, k := range arr {
				k = strings.Trim(k, " ")
				selects[k] = true
			}
			break
		case []string:
			for _, s := range val {
				arr := strings.Split(s, ",")
				for _, k := range arr {
					selects[k] = true
				}
			}
		}
	}
	hasSelect := len(selects) > 0
	var joinTables = make(map[string][]string)
	var results []string
	queryFields := GetModelFields(m)

	for key, q := range queryFields {
		if hasSelect {
			if _, ok := selects[key]; ok {
				joinTable := q.GetJoinTable()
				if joinTable != "" {
					joinTables[joinTable] = q.GetJoin()
				}
				results = append(results, q.GetExpress())
			}
			continue
		}
		if rule == "" || q.HasRule(rule) {
			joinTable := q.GetJoinTable()
			if joinTable != "" {
				joinTables[joinTable] = q.GetJoin()
			}
			results = append(results, q.GetExpress())
		}
		if q.HasRule("S") && q.GetExpress() != key {
			results = append(results, fmt.Sprintf("%s.%s", m.AliasName(), key))
		}
	}
	var joins [][]string
	for _, item := range joinTables {
		joins = append(joins, item)
	}
	if len(results) > 0 {
		return strings.Join(results, ","), joins
	} else {
		return "*", joins
	}
}

// 查询模型字段
func Select(db xorm.IOrm, m IModel, rule string, selects ...interface{}) xorm.IOrm  {
	query := db.Table(m.TableAlias())
	return Query(query, m, rule, selects...)
}

// 生成查询字段及join配置
func Query(query xorm.IOrm, m IModel, rule string, selects ...interface{}) xorm.IOrm  {
	fields, joins := GetQueryFields(m, rule, selects...)
	query = query.Fields(fields)
	if len(joins) > 0 {
		for _, join := range joins {
			query = query.Join(join[2], join[0], join[1])
		}
	}
	return query
}
