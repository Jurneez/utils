package reflect_struct

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

// func traverse(v interface{}) {
// 	p := reflect.ValueOf(v)
// 	t := reflect.TypeOf(v)
// 	if t.Kind() == reflect.Ptr {
// 		p = p.Elem()
// 		t = t.Elem()
// 	}
// 	// 获取结构体字段数
// 	num := p.NumField()
// 	for i := 0; i < num; i++ {
// 		if p.Field(i).Kind() == reflect.Struct || p.Field(i).Kind() == reflect.Ptr || p.Field(i).Elem().Kind() == reflect.Struct {
// 			traverse(p.Field(i).Interface())
// 		} else {
// 			f := t.Field(i)
// 			val := p.Field(i).Interface()
// 			fmt.Printf("%s  %v=%v \n", f.Name, f.Type, val)
// 		}

// 	}
// }
