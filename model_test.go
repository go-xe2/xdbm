package xdbm

import "testing"

type userModel struct {
	Model
}

func newUserModel() IModel {
	return new(userModel)
}

func (m *userModel) TableName() string {
	return "users"
}

func (m *userModel) AliasName() string {
	return "u"
}

var userModelFields = map[string]interface{}{
	"user_id": 1,
	"sex": "",
	"product_id": "-",
	"name": "name",
	"is_delete": "-",
	"status": []interface{}{"MD,ML,S", `case $1.status when 0 then '正常' when 1 then '删除' end`, "status_name"},
	"up_date": []interface{}{"ML,UL,GL", `from_unixtime($1.up_date,'%Y-%m-%d %H:%i:%s')`, "up_date"},
	"product_name":  NewModelRelation("MD", newProductModel(), "$1.product_id=$2.product_id","left join", "$2.name + $1.name"),
}

func (m *userModel) Fields() map[string]interface{} {
	return userModelFields
}

func (m *userModel) AppendRules() []string {
	return nil
}

// 更新规则
func (m *userModel) UpdateRules() []string {
	return nil
}

// 允许排序字段
func (m *userModel) AllowSortFields() map[string]bool {
	return nil
}


type productModel struct {
	Model
}

func newProductModel() IModel {
	return new(productModel)
}

func (m *productModel) TableName() string {
	return "products"
}

func (m *productModel) AliasName() string {
	return "p"
}

var productModelFields = map[string]interface{}{
	"product_id": "",
	"name": "",
	"price": "sale_price",
	"cr_date": "form_unixtime($1.cr_date)",
}

func (m *productModel) Fields() map[string]interface{} {
	return productModelFields
}

func (m *productModel) AppendRules() []string {
	return nil
}

// 更新规则
func (m *productModel) UpdateRules() []string {
	return nil
}

// 允许排序字段
func (m *productModel) AllowSortFields() map[string]bool {
	return nil
}

func TestModel_Select(t *testing.T) {
	u := newUserModel()

	fields, joins := GetQueryFields(u, "MD", )
	t.Log("fields:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "", "product_name")
	t.Log("fields:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "", "user_id,cr_date,product_name")
	t.Log("fields:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "GD")
	t.Log("fields:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "MD")
	t.Log("fields:", fields, ",joins:", joins)

	// Output:

}

