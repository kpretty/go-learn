package project

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// 使用 sql 和 http 实现学生管理系统

// student 学生对象
type student struct {
	Id    int    `json:"id"`    // 学号
	Age   int8   `json:"age"`   // 年龄
	Name  string `json:"name"`  // 姓名
	Class string `json:"class"` // 班级
}

// marshal 结构体 -> json str
func marshal(stu *student) string {
	str, err := json.Marshal(stu)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func marshalList(stu []*student) string {
	str, err := json.Marshal(stu)
	if err != nil {
		panic(err)
	}
	return string(str)
}

// unmarshal json str -> 结构体
func unmarshal(str string) *student {
	stu := &student{}
	err := json.Unmarshal([]byte(str), stu)
	if err != nil {
		panic(err)
	}
	return stu
}

// 封装对数据库的增删改查
var serverDbDDL = "create database if not exists demo"
var studentDDL = "create table if not exists demo.student (" +
	"id int," +
	"age int," +
	"name varchar(100)," +
	"class varchar(100)" +
	")"
var db *sql.DB

// dbInit 初始化表
func dbInit() {
	_db, err := sql.Open("mysql", "root:980729@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	db = _db
	exec, err := db.Exec(serverDbDDL)
	if err != nil {
		panic(err)
	}
	log.Println("初始化DB：", exec)
	exec, err = db.Exec(studentDDL)
	if err != nil {
		panic(err)
	}
	log.Println("初始化TB：", exec)
}

// getStudentList 查询学生列表
func getStudentList() []*student {
	stmt, _ := db.Prepare("select * from demo.student where 1=1")
	defer func() { _ = stmt.Close() }()
	query, _ := stmt.Query()
	defer func() { _ = query.Close() }()
	students := make([]*student, 0)
	var counts int
	for query.Next() {
		var stu student
		err := query.Scan(
			&stu.Id,
			&stu.Age,
			&stu.Name,
			&stu.Class,
		)
		if err != nil {
			panic(err)
		}
		students = append(students, &stu)
		counts++
	}
	return students[:counts]
}

// addStudent 添加学生信息
func addStudent(stu *student) {
	stmt, _ := db.Prepare("insert into demo.student values (?,?,?,?)")
	defer func() { _ = stmt.Close() }()
	result, _ := stmt.Exec(stu.Id, stu.Age, stu.Name, stu.Class)
	log.Println("添加学生信息完成：", result)
}

// deleteStudent 删除学生信息
func deleteStudent(id int) {
	stmt, _ := db.Prepare("delete from demo.student where id = ?")
	defer func() { _ = stmt.Close() }()
	result, _ := stmt.Exec(id)
	log.Println("删除学生信息完成：", result)
}

// modifyStudent 修改学生信息
func modifyStudent(stu *student) {
	stmt, _ := db.Prepare("update demo.student set age = ?, name = ?, class = ? where id = ?")
	defer func() { _ = stmt.Close() }()
	result, _ := stmt.Exec(stu.Age, stu.Name, stu.Class, stu.Id)
	log.Println("修改学生信息完成：", result)
}

// getOneStudent 查询单个
func getOneStudent(id int) *student {
	stmt, _ := db.Prepare("select * from demo.student where id = ?")
	defer func() { _ = stmt.Close() }()
	var stu student
	_ = stmt.QueryRow(id).Scan(&stu.Id, &stu.Age, &stu.Name, &stu.Class)
	return &stu
}

// closeDB 关闭数据库连接
func closeDB() {
	_ = db.Close()
}

// 使用 http 封装对数据库的操作

func addStudentHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	body, _ := ioutil.ReadAll(r.Body)
	stu := unmarshal(string(body))
	addStudent(stu)
	_, _ = fmt.Fprintln(w, "ok")
}

func getStudentListHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	_, _ = fmt.Fprintln(w, marshalList(getStudentList()))
}

func deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	// 获取 get 请求参数
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	deleteStudent(int(id))
	_, _ = fmt.Fprintln(w, "ok")
}

func modifyStudentHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	// 获取 POST 参数
	body, _ := ioutil.ReadAll(r.Body)
	// body -> struct
	stu := unmarshal(string(body))
	modifyStudent(stu)
	_, _ = fmt.Fprintln(w, "ok")
}

func getOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	// 获取 get 请求参数
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	oneStudent := getOneStudent(int(id))
	_, _ = fmt.Fprintln(w, marshal(oneStudent))
}

func Server() {
	// 初始化服务
	dbInit()
	http.HandleFunc("/getAll", getStudentListHandler)
	http.HandleFunc("/getOne", getOneStudentHandler)
	http.HandleFunc("/add", addStudentHandler)
	http.HandleFunc("/delete", deleteStudentHandler)
	http.HandleFunc("/modify", modifyStudentHandler)

	srv := &http.Server{
		Addr: ":9000",
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// 注册鉴定
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 监听，无数据阻塞
	<-quit
	// 关闭服务器
	_ = srv.Shutdown(context.Background())
	closeDB()

}
