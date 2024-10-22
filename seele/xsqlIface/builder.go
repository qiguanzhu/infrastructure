/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/9/30 -- 15:15
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: dao.go
*/

package xsqlIface

import "context"

// BuilderProxy 构造 查询语句 & 查询参数
type BuilderProxy interface {
	BuildSelect(tableName string, where map[string]interface{}, selectedField []string) (query string, args []interface{}, err error)
	BuildSelectWithContext(ctx context.Context, tableName string, where map[string]interface{}, selectedField []string) (query string, args []interface{}, err error)
	BuildUpdate(tableName string, where map[string]interface{}, update map[string]interface{}) (string, []interface{}, error)
	BuildDelete(tableName string, where map[string]interface{}) (string, []interface{}, error)
	BuildInsert(tableName string, data []map[string]interface{}) (string, []interface{}, error)
	BuildInsertIgnore(tableName string, data []map[string]interface{}) (string, []interface{}, error)
	BuildReplaceIgnore(tableName string, data []map[string]interface{}) (string, []interface{}, error)
	BuildCustomize(tableName string, f func(string) (string, []interface{}, error)) (string, []interface{}, error)
	AggregateQuery(ctx context.Context, db XDB, tableName string, where map[string]interface{}, aggregate AggregateSymbolBuilder) (ResultResolver, error)
}

// Comparable requires type implements the Build method
type Comparable interface {
	Build() ([]string, []interface{})
}

// AggregateSymbolBuilder need to be implemented so that executor can
// get what should be put into `select Symbol() from xxx where yyy`
type AggregateSymbolBuilder interface {
	Symbol() string
}

// ResultResolver is a helper for retrieving data
// caller should know the type and call the responding method
type ResultResolver interface {
	Int64() int64
	Float64() float64
}