package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.maimaiche.com/loginRegister/login.do",
		strings.NewReader("mobile=xxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(resp.Header.Get("Content-Type")) //application/json;charset=UTF-8

	type Result struct {
		Msg    string
		Status string
		Obj    string
	}

	result := &Result{}
	json.Unmarshal(body, result) //解析json字符串

	if result.Status == "1" {
		fmt.Println(result.Msg)
	} else {
		fmt.Println("login error")
	}
	fmt.Println(result)
}