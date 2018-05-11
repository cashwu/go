//https://www.kancloud.cn/digest/batu-go/153529
package main

import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"

func main() {

	resp, err := http.Get("https://www.google.com.tw")
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))

}

func httpPost() {
	resp, err := http.PostForm("", url.Values{"key": {"Value"}})

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))

}
