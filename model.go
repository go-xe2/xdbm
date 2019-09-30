package xdbm

import (
	"fmt"
)

type Model struct {
	IModel
	IModelMutation
	IModelQuery
}

func (m *Model) TableAlias(alias ...string) string {
	if len(alias) > 0 {
		return fmt.Sprintf("%s %s", m.TableName(), alias[0])
	}
	return fmt.Sprintf("%s %s", m.TableName(), m.AliasName())
}
