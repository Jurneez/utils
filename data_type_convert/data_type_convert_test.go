package data_type_convert

import "testing"

func Test_StringToInt(t *testing.T) {
	s := "123"
	t.Log(StringToInt(s))
}

func Test_PrintSpeicalString(t *testing.T) {
	cs1 := "abc\n123"
	t.Log(PrintSpeicalString(cs1))
}
