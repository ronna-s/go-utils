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


func ToInterfaceArray(v interface{}) (res []interface{}) {
        switch reflect.TypeOf(v).Kind() {
        case reflect.Slice:
                s := reflect.ValueOf(v)
                res = make([]interface{}, s.Len())
                for i := 0; i < s.Len(); i++ {
                        res[i] = s.Index(i).Interface()
                }
        default:
                panic(fmt.Sprintf("unexpected type %T", v))
        }
        return res
}

func Included(arr interface{}, i interface{}) bool {
        for _, item := range ToInterfaceArray(arr){
                if item == i {
                        return true
                }
        }
	return false
}
