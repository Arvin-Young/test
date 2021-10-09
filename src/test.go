package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Sever() {
	http.HandleFunc("/", HelloHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	Sever()
	//region, _ := getRegion("cr.registry.sd-1.res.sgmc.sgcc.com.cn")
	//fmt.Println(region)
}

func getRegion(endpoint string) (string, error) {
	result := strings.Split(endpoint, ".")
	if len(result) < 3 {
		return "", nil
	}
	return result[2], nil
}

func RequestTest() {
	time.Sleep(2 * time.Second)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://localhost:8081/")
	if err != nil {
		panic(err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
}