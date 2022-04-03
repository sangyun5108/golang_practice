package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

func main(){

	results:=map[string]string{}

	c:=make(chan requestResult)

	urls:=[]string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.instargram.com",
		"https://www.facebook.com",
	}


	for _,value:=range urls {
		go hitURL(value,c)
	}

	for i:=0 ; i<len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for key,value := range results {
		fmt.Println(key, value)
	}

}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string, c chan requestResult) {

	fmt.Println("Checking:",url)

	resp,err:=http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "OK"}
	}
	c <- requestResult{url: url, status:"Fail"}
}
