package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}
type AutoGenerated struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []interface{} `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{"en2zh", word}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"jack"}`)
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	data := bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var autoGenerated AutoGenerated
	err = json.Unmarshal(bodyText, &autoGenerated)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v\n", autoGenerated)
	fmt.Println(autoGenerated.Dictionary.Prons.En, autoGenerated.Dictionary.Prons.EnUs)
	for _, item := range autoGenerated.Dictionary.Explanations {
		fmt.Println(item)
	}
}
func main() {
	var s string
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			log.Fatal(err)
		}
		query(s)
	}
}