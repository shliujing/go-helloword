package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
)

func main() {

	postParam := url.Values{
		"mobile":      {"xxxxxx"},
		"isRemberPwd": {"1"},
	}

	resp, err := http.PostForm("http://www.maimaiche.com/loginRegister/login.do", postParam)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}