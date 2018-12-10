package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type RetToken struct {
	Code  int    `json:"code"`
	Token string `json:"open_api_auth_token"`
}

func main() {

	postParam := "{\"email\": \"wangzeyi@qiniu.com\", \"password\": \"nji90okm\" }"

	req, err := http.NewRequest("POST", "https://qiniu.udesk.cn/open_api_v1/log_in", bytes.NewBuffer([]byte(postParam)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	ret := RetToken{}
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
