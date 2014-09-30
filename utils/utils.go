package utils

import (
	"reflect"
        "fmt"
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
func ToInterface(v interface{}) (res []interface{}) {
        switch reflect.TypeOf(v).Kind() {
        case reflect.Slice:
                s := reflect.ValueOf(v)
                res = make([]interface{}, s.Len())
                for i := 0; i < s.Len(); i++ {
                        res[i] = s.Index(i).Interface()
                }
        default:
                panic(fmt.Sprintf("unexpected type %T", reflect.TypeOf(v).Kind()))

        }
	return res
}
