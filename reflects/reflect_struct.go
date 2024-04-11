package reflects

import (
	"fmt"
	"reflect"
)

// 遍历
func simpleTraverse(v interface{}) {
	p := reflect.ValueOf(v) // 获取v中的值
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr { // t.Kind() 可以获取到结构体的名称
		p = p.Elem() // 这里，如果t.Kind() 是指针类型的，那么p.Elem()函数会返回指针指向的值
		t = t.Elem()
	}

	// 获取结构体字段数
	num := p.NumField()
	for i := 0; i < num; i++ {
		f := t.Field(i)
		val := p.Field(i).Interface()
		fmt.Printf("%s  %v=%v \n", f.Name, f.Type, val)
	}
}

// 将数组型结构体 转为 数组型map
func StructToMapSlice(in interface{}, tagName string) (out []map[string]interface{}, err error) {
	if tagName == "" {
		tagName = "json"
	}
	out = make([]map[string]interface{}, 0)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice { // 非数组报错
		return out, fmt.Errorf("StructToMapSlice only accepts array or slice pointer; got %T", v)
	}
	length := v.Len()
	for k := 0; k < length; k++ {
		each := v.Index(k)

		if each.Kind() == reflect.Ptr {
			each = each.Elem()
		}

		if each.Kind() != reflect.Struct { // 非结构体返回错误提示
			return out, fmt.Errorf("v only accepts struct or struct pointer; got %T", v)
		}
		t := each.Type()
		num := each.NumField()
		outSingle := make(map[string]interface{})
		for i := 0; i < num; i++ {
			fi := t.Field(i)
			if tagValue := fi.Tag.Get(tagName); tagValue != "" {
				outSingle[tagValue] = each.Field(i).Interface()
			}
		}
		out = append(out, outSingle)
	}

	return out, nil
}
