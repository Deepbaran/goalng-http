package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Client

func main() {
	// Plain old URL - http://localhost
	res, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

	// http://localhost/url?name=Deep
	res, err = http.PostForm("http://localhost/url",
		url.Values{"name": {"Deep"}})
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

	// http://localhost/body
	// with body in json: {"name": "Deep"}
	type deep struct {
		Name string
	}
	deep1 := deep{"Deep"}
	deepJson, _ := json.Marshal(deep1)
	res, err = http.Post(
		"http://localhost/body",
		"application/json",
		bytes.NewBuffer(deepJson),
	)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

	client := http.Client{}
	req, er := http.NewRequest(
		"GET",
		"http://localhost/body",
		bytes.NewBuffer(deepJson),
	)
	if er != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}
}
