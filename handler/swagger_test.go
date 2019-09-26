package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_HttpTestSwagger(t *testing.T) {
	token := ""
	//form表单数据
	data := `{}`
	req, err := http.NewRequest("POST", "http://localhost:1323/api/activity/earning/GetUserBalance", strings.NewReader(data))
	if err != nil {
		t.Error("http.NewRequest err:", err)
		return
	}
	//header设置
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", token)
	req.Header.Set("X-SITE-ID", "127")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("http.DefaultClient.Do err:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
