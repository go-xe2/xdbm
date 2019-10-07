package xdbm

import (
	"encoding/json"
	"testing"
)

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
	"user_id": 		"user_id",
	"login_name": 	"login_name",
	"nick_name": 	[]interface{}{"", "from_base64($1.nick_name)", "nick_name"},
	"pwd": 			"-",
	"enc": 			"-",
	"sex": 			[]interface{}{ "S",`case $1.sex when 0 then "未知" when 1 then "男" when 2 then "女" end`, "sex_name" },
	"mobile": 		"mobile",
	"qq": 			"qq",
	"province": 	"province",
	"city": 		"city",
	"county": 		"county",
	"town": 		"town",
	"province_id": 	"-",
	"city_id": 		"-",
	"county_id": 	"-",
	"town_id": 		"-",
	"address": 		"address",
	"head": 		"head",
	"cr_date": 		[]interface{}{ "ML,UL,GL,MD", "from_unixtime($1.cr_date)" },
	"up_date": 		[]interface{}{ "ML,UL,GL,MD", "from_unixtime($1.up_date)" },
	"last_login": 	[]interface{}{ "ML,MD", "from_unixtime($1.last_login)" },
	"last_ip": 		[]interface{}{ "ML,MD" },
	"status": 		[]interface{}{ "ML,MD", "case $1.status when 0 then '未审核' when 1 then '审核失败' when 2 then '审核通过' when 3 then '锁定' end", "status_name" },
	"fav_count": 	"fav_count",
	"visit_count": 	"visit_count",
	"audit_id":		[]interface{}{ "ML,MD" },
	"audit_summery":[]interface{}{ "ML,MD" },
	"audit_date":	[]interface{}{ "ML,MD", "case when $1.audit_date > 0 then from_unixtime($1.audit_date) else '' end", "audit_date" },
	"is_expert": 	"is_expert",
	"product_name":  NewModelRelation("MD", newProductModel(), "$1.product_id=$2.product_id","left", "$2.name + $1.name"),
}

func (m *userModel) Fields() map[string]interface{} {
	return userModelFields
}

func (m *userModel) Field(fieldName string) string {
	return m.AliasName() + "." + fieldName
}

func (m *userModel) AppendRules() []string {
	return []string{
		"^login_name?loginName!required#登录名不能为空",
		"^nick_name?nickName!required#昵称不能为空",
		"^mobile!required|phone#手机号不能为空|不是有效的手机号",
		"^pwd?pwd!password2#请输入6-18数字和大小写字母组合的密码",
		"^sex~0",
		"^head~",
		"^province&string!required#省份不能为空",
		"^city&string!required#城市不能为空",
		"^county&string!required#区县不能为空",
	}
}

// 更新规则
func (m *userModel) UpdateRules() []string {
	return []string{
		"^login_name?loginName!required#登录名不能为空",
		"^nick_name?nickName!required#昵称不能为空",
		"^mobile!required|phone#手机号不能为空|不是有效的手机号",
		"^pwd?pwd!password2#请输入6-18数字和大小写字母组合的密码",
		"^sex~0",
		"^head~",
		"^province&string!required#省份不能为空",
		"^city&string!required#城市不能为空",
		"^county&string!required#区县不能为空",
	}
}

// 允许排序字段
func (m *userModel) AllowSortFields() map[string]bool {
	return map[string]bool {
		"cr_date": true,
		"up_date": true,
		"fav_count": true,
		"visit_count": true,
		"discuss_count": true,
		"province": true,
		"city": true,
		"county": true,
	}
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

func (m *productModel) Field(fieldName string) string {
	return m.AliasName() + "." + fieldName
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

	modelFields := GetModelFields(u)
	bytes, err := json.Marshal(modelFields)
	t.Log("model fields:", string(bytes), ", error:", err, "\n")

	fields, joins := GetQueryFields(u, "", )
	t.Log("fields1:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "", "product_name")
	t.Log("fields2:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "", "user_id,cr_date,product_name")
	t.Log("fields3:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "GD")
	t.Log("fields4:", fields, ",joins:", joins)

	fields, joins = GetQueryFields(u, "MD")
	t.Log("fields5:", fields, ",joins:", joins)

	if i,ok := u.(IModelMutation); ok {
		rule1 := i.AppendRules()
		t.Log("append rules:", rule1)

		rule2 := i.UpdateRules()
		t.Log("update rules:", rule2)
	}
	if i, ok := u.(IModelQuery); ok {
		fields := i.AllowSortFields()
		t.Log("allow sorts:", fields)
	}
	// Output:

}

