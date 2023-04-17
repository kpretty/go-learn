package project

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	student := &student{
		Id:    1,
		Name:  "张三",
		Age:   20,
		Class: "三年二班",
	}
	marshal := marshal(student)
	fmt.Println(marshal)
}

func TestUnmarshal(t *testing.T) {
	str := `{"id":1,"age":20,"name":"张三","class":"三年二班"}`
	student := unmarshal(str)
	fmt.Println(student)
}

func TestMarshalList(t *testing.T) {
	stu := &student{
		Id:    1,
		Name:  "张三",
		Age:   20,
		Class: "三年二班",
	}
	list := marshalList([]*student{stu, stu})
	fmt.Println(list)
}

func TestServer(t *testing.T) {
	Server()
}
