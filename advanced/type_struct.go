package advanced

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// type 关键字
// 1. 自定义类型
// type 新类型名 类型名

type MyInt int

// 自定义类型是全新的类型，只不过底层存储使用原先类型，但新类型与原类型没有关系
func customType() {
	var myInt MyInt = 1
	fmt.Printf("%v %T\n", myInt, myInt)
}

// 2. 类型别名
// type aliasType = type
// 与自定义类型区别在于，类型别名与原类型是同一个类型，例如 rune 和 byte

// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
// type byte = uint8

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
// type rune = int32

// any is an alias for interface{} and is equivalent to interface{} in all ways.
// type any = interface{}

type newInt = int

// 注：类型别名在编译完成后是不存在的，例如上面的 byte、rune、any 编译完成后会被替换成对应的类型，自定义类型则不会，是真实存在的
func typeNature() {
	var a MyInt
	var b newInt
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
}

// 3. 结构体
// 对多个基本数据类型的封装，类比其他语言的对象
type myFirstStruct struct {
	a int
	b float64
	c string
	d []int
	e map[string]int
}

// 结构体如果多个字段是同个类型，可以写在一行有助于编译器内存对齐
type person struct {
	name, city string
	age, high  int
}

// 结构体的实例化
func initStruct() {
	// 1.因为结构体是值类型，会给其所有字段都赋零值
	var s1 myFirstStruct
	fmt.Printf("%v %T\n", s1, s1)
	// 2.匿名结构体，该类型的结构体只会被用一次
	var user struct {
		Name string
		Age  int
	}
	user.Name = "张三"
	user.Age = 18
	fmt.Printf("%v %T\n", user, user)
	// 3.创建并赋值，每行用 , 最后一个字段也需要
	p := person{
		name: "张三",
		city: "杭州",
		age:  18,
		high: 180,
	}
	fmt.Printf("%v %T\n", p, p)
	// 4.使用 new
	p2 := new(person)
	// p2 是指针类型
	p2.name = "李四"
	fmt.Printf("%v %T\n", p2, p2)
	// 5.new 的等价用法
	p3 := &person{}
	p3.name = "王五"
	fmt.Printf("%v %T\n", p3, p3)
	// 6.将3和5结合起来
	p4 := &person{
		name: "赵六",
		city: "杭州",
		age:  18,
		high: 180,
	}
	fmt.Printf("%v %T\n", p4, p4)
}

// 方法和接受者
// 类比其他语言的对象和方法，接受者本质上是this或者self
// 可以是指针类型也可以时值类型，区别在于是否会在接受者对应的变量调用时复制变量

// 为 person 提供set方法
// 因为接受者是值类型，因此当该方法被调用时会复制一份数据给 p，因此原变量的age并没有被改变
func (p person) setAge(age int) {
	p.age = age
}

// go 中不推荐同一个结构体既有非指针类型的接受者有存在指针类型的结构体

type Person person

func (p *Person) SetAge(age int) {
	p.age = age
}

func method() {
	p1 := person{}
	p2 := new(Person)
	fmt.Println(p1.age)
	fmt.Println(p2.age)
	// 修改年龄
	p1.setAge(18)
	p2.SetAge(18)
	fmt.Println(p1.age)
	fmt.Println(p2.age)
}

// 什么时候应该使用指针类型接收者
//	1. 需要修改接收者中的值
//	2. 接收者是拷贝代价比较大的大对象
//	3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// 结构体匿名字段
// 结构体允许成员只有类型没有字段名
type anonymous struct {
	string
	int
}

func anonymousDemo() {
	a1 := anonymous{
		"张三",
		18,
	}
	fmt.Println(a1)
	// 匿名字段本质上会有一个与类型名一致的字段名，因此匿名字段类型不能重复
	fmt.Println(a1.int, a1.string)
}

// 嵌套结构体中使用匿名字段

type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名字段
}

func anonymousDemo1() {
	user := User{
		"张三",
		"男",
		Address{
			"浙江",
			"杭州",
		},
	}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Gender)
	// 匿名字段可以省略
	fmt.Println(user.Province)
	fmt.Println(user.City)
}

// 实现学生信息管理系统

// Student 学生信息结构体
type Student struct {
	id    int
	name  string
	age   int
	score float64
}

// StudentSystem 存储学生信息
type StudentSystem struct {
	students map[int]*Student
}

// 添加一个学生
func (ssy *StudentSystem) addStudent(stu *Student) {
	// 判断id是否存在
	_, ok := ssy.students[stu.id]
	if ok {
		log.Printf("当前学生id[%d]已存在，请勿重复添加", stu.id)
		return
	}
	ssy.students[stu.id] = stu
}

// 删除一个学生
func (ssy *StudentSystem) removeStudent(id int) {
	// 判断id是否存在
	_, ok := ssy.students[id]
	if !ok {
		log.Printf("当前学生id[%d]不存在无法删除", id)
	} else {
		delete(ssy.students, id)
		log.Println("成功删除id ", id)
	}
}

// 修改学生信息
func (ssy *StudentSystem) modifyStudent(stu *Student) {
	// 判断id是否存在
	_, ok := ssy.students[stu.id]
	if !ok {
		log.Printf("当前学生id[%d]不存在无法修改", stu.id)
		return
	}
	ssy.students[stu.id] = stu
}

// 查看学生信息
func (ssy *StudentSystem) showStudent(ids ...int) {
	if ids == nil || len(ids) == 0 {
		// 查询所有
		for _, student := range ssy.students {
			fmt.Println(student)
		}
	} else {
		for _, id := range ids {
			fmt.Println(ssy.students[id])
		}
	}
}

// Bootstrap 启动
func Bootstrap() {
	// 初始化系统
	ssy := &StudentSystem{
		students: make(map[int]*Student),
	}
	// 接收控制台输入
	scanner := bufio.NewScanner(os.Stdin)
	// 将日志输出到标准错误流中
	log.SetOutput(os.Stderr)

	fmt.Println("1. 添加")
	fmt.Println("2. 删除")
	fmt.Println("3. 修改")
	fmt.Println("4. 查询")
	fmt.Println("5. 退出")
Loop:
	for {
		fmt.Print("输入选项[1-4]> ")
		scanner.Scan()
		op := scanner.Text()
		switch op {
		case "1", "3":
			fmt.Println("依次输入学号，姓名，年龄，成绩(用空格隔开)")
			fmt.Print("> ")
			scanner.Scan()
			line := scanner.Text()
			studentInfo := strings.Split(line, " ")
			id, _ := strconv.ParseInt(studentInfo[0], 10, 64)
			age, _ := strconv.ParseInt(studentInfo[2], 10, 64)
			score, _ := strconv.ParseFloat(studentInfo[3], 64)
			student := &Student{
				id:    int(id),
				name:  studentInfo[1],
				age:   int(age),
				score: score,
			}
			if op == "1" {
				ssy.addStudent(student)
			} else {
				ssy.modifyStudent(student)
			}
		case "2":
			fmt.Print("输入学生学号> ")
			scanner.Scan()
			id, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			ssy.removeStudent(int(id))
		case "4":
			// 查询所有
			ssy.showStudent()
		case "5":
			break Loop
		default:
			log.Println("无效输入")
		}
	}
}
