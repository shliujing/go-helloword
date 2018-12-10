package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
)

func main() {

	postParam := "{\"email\": \"wangzeyi@qiniu.com\", \"password\": \"nji90okm\" }"

	req, err := http.NewRequest("POST","https://qiniu.udesk.cn/open_api_v1/log_in", bytes.NewBuffer([]byte(postParam)))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

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
