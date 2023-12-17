package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 对象的属性值必须是大写的，否则无法被序列化（大小写用于访问控制）
// 以下` `json:"name"“指定的是对象实例化和反序列化的字段名称，如果没有这个标记，默认情况下取的就是字段的名称（首字母是大写的）
type BaseResult struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func index(w http.ResponseWriter, r *http.Request) {
	printLog(r)
	// output: http://localhost:9090/index?id=1
	var resultObj BaseResult
	resultObj.Code = 200
	resultObj.Success = true
	resultObj.Data = "{\"message\": \"hello go web\"}"
	resultObj.Message = "success"
	result, err := json.Marshal(resultObj)
	if err != nil {
		fmt.Print(err)
		result = []byte(`{"code": -1, "success": flase, "message": "json failed"}`)
	}

	writeResult(w, result)
}

func printLog(r *http.Request) {
	log.Println("--------------------")
	log.Println(r.Proto)
	// output:HTTP/1.1
	log.Println(r.TLS)
	// output: <nil>
	log.Println(r.Host)
	// output: localhost:9090
	log.Println(r.RequestURI)
	// output: /index?id=1

	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	log.Println(strings.Join([]string{scheme, r.Host, r.RequestURI}, ""))
}

func writeResult(writer http.ResponseWriter, reuslt []byte) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(reuslt)
}

func GetURL(r *http.Request) (Url string) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
}

func main() {
	http.HandleFunc("/index", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
