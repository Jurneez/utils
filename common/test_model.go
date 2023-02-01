package common

type People struct {
	Name string
	Age  int
}

func NewPeople(name string, age int) *People {
	return &People{
		Name: name,
		Age:  age,
	}
}

var P1 = NewPeople("Jone", 13)
var P2 = NewPeople("jury", 15)

type Student struct {
	P     []*People
	Grade string
	Tags  []string
}

func NewStudent(ps []*People, grade string, tags []string) *Student {
	return &Student{
		P:     ps,
		Grade: grade,
		Tags:  tags,
	}
}

var S1 = NewStudent([]*People{P1, P2}, "高一", []string{"先进班集体", "光荣少年班"})
