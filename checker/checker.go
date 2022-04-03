package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main(){

	results := map[string]string{}



	urls:=[]string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.instargram.com",
		"https://www.facebook.com",
	}


	for _,value:=range urls {

		ok:="OK"
		err:= hitURL(value)

		if err!=nil {
			fmt.Println(err)
		}else {
			results[value] = ok
		}
	}

	for url,result:=range results {
		fmt.Println(url,result)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {

	fmt.Println("Checking:",url)

	resp,err:=http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
	}

	fmt.Println("완벽합니다")
	
	return nil
}
