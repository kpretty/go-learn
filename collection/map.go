package collection

import (
	"fmt"
	"strings"
)

// 定义map，map[key的类型]value的类型
var m map[string]int

// 初始化 map 的方式
func initMap() {
	// 区分切片，切片也有类似的初始化方式
	m = map[string]int{"k1": 1, "k2": 2}
	// map 为散列，没有 cap 容量的概念，但是有长度
	fmt.Printf("%v ptr %p type %T len %d\n", m, m, m, len(m))
	// 通过 make 进行初始化，有长度的概念，但是可以不指定
	m = make(map[string]int)
	fmt.Printf("%v ptr %p type %T len %d\n", m, m, m, len(m))
	m["张三"] = 20
	m["李四"] = 18
	fmt.Printf("%v ptr %p type %T len %d\n", m, m, m, len(m))
}

// 判断某个 key 存不存在
func isExist() {
	m1 := make(map[string]int)
	m1["张三"] = 20
	m1["李四"] = 18
	result := m1["张三"]
	fmt.Println(result)
	// 上述方式存在问题，若 key 不存在返回 value 类型的默认值，存在歧义
	// 正确方式
	v, ok := m1["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key不存在")
	}
	v1, ok := m1["张三1"]
	if ok {
		fmt.Println(v1)
	} else {
		fmt.Println("key不存在")
	}
}

// 遍历map
func iteratorMap() {
	m1 := make(map[string]int)
	m1["k1"] = 1
	m1["k2"] = 2
	m1["k3"] = 3
	m1["k4"] = 4
	// for range
	for k, v := range m1 {
		fmt.Printf("key %v value %v\n", k, v)
	}
}

// 删除map的key
func deleteMapKey() {
	m1 := make(map[string]int)
	m1["k1"] = 1
	m1["k2"] = 2
	m1["k3"] = 3
	m1["k4"] = 4
	delete(m1, "k2")
	// 试图删除不存在的 key
	delete(m1, "kk")
	// for range
	for k, v := range m1 {
		fmt.Printf("key %v value %v\n", k, v)
	}
}

func wordCount(line string) {
	words := strings.Split(line, " ")
	// 初始化一个 map
	result := make(map[string]int)
	for _, word := range words {
		// 如果 key 不存在返回默认值 0
		result[word] = result[word] + 1
	}
	for k, v := range result {
		fmt.Printf("%v %v\n", k, v)
	}
}
