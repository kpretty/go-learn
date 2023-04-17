package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v] from [%v]", r.Method, r.RemoteAddr)
	_, _ = fmt.Fprintln(w, "hello go")
}

func login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("[%v] from [%v] and body [%v]", r.Method, r.RemoteAddr, string(body))
	_, _ = fmt.Fprintln(w, "this is login website")
}

// Server 服务端
func Server() {
	// 注册处理器
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}

func Client() {
	resp, err := http.Get("http://127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	post, err := http.Post("http://127.0.0.1:9000/login", "", strings.NewReader(`{"name":"张三","age":12}`))
	if err != nil {
		panic(err)
	}
	defer func() { _ = post.Body.Close() }()
	body, err = ioutil.ReadAll(post.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
