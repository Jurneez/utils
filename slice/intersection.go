package slice

import "reflect"

// 交集
func Intersection(s1, s2 interface{}) interface{} {
	inArr1 := reflect.ValueOf(s1)
	if inArr1.Kind() != reflect.Slice && inArr1.Kind() != reflect.Array {
		return s1
	}

	s1Map := make(map[interface{}]bool)
	for i := 0; i < inArr1.Len(); i++ {
		iVal := inArr1.Index(i)
		s1Map[iVal.Interface()] = true
	}

	inArr2 := reflect.ValueOf(s2)
	if inArr2.Kind() != reflect.Slice && inArr2.Kind() != reflect.Array {
		return s1
	}

	if inArr1.Kind() != inArr2.Kind() {
		return s1
	}

	outArr := reflect.MakeSlice(inArr1.Type(), 0, inArr2.Len())
	for i := 0; i < inArr2.Len(); i++ {
		iVal := inArr2.Index(i)
		if _, ok := s1Map[iVal.Interface()]; !ok {
			outArr = reflect.Append(outArr, inArr2.Index(i))
		}
	}

	return outArr.Interface()
}
