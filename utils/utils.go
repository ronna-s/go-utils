package utils

import (
	"reflect"
)

func MapFunc(name string, arr ...interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	for idx, item := range arr {
		v := reflect.ValueOf(interface{}(item))
		if field_by_name := v.FieldByName(name); field_by_name.IsValid() {
			res[idx] = field_by_name.Interface()
		} else { //method
			m := v.MethodByName(name).Call([]reflect.Value{})
			inter := make ([]interface{}, len(m))
			for i, j:= range m{
								    inter[i] = j.Interface()
			}
			res[idx] = inter
		}
	}
	return res
}

func ToInterface(v interface{}) []interface{} {
	// var news []*news_items.NewsItem
	var res []interface{}
	switch t := v.(type) {
	default:
		panic(fmt.Sprintf("unexpected type %T", t))

	}
	return res
}

