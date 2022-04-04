package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type jobData struct {
	title string
	location string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python" 

func main(){

	c2:=make(chan []jobData)
	jobs:=[]jobData{}

	totalPages:=getPages()
	
	for i:=0; i<totalPages; i++ {
		go getPage(i,c2)
		job:= <- c2
		fmt.Println(job)
	}

	writeJobs(jobs)
}

func writeJobs(jobs []jobData){
	file,err := os.Create("jobs.csv")

	checkErr(err)

	// 파일 생성
	w:=csv.NewWriter(file)

	// 데이터가 파일에 입력된다.
	defer w.Flush()

	headers := []string{"Title","Location"}

	wrError := w.Write(headers)

	checkErr(wrError)

	for _,job := range jobs {
		jobSlice := []string{job.title, job.location}
		wrError = w.Write(jobSlice)
		checkErr(wrError)
	}
}

func getPage(page int,c2 chan []jobData) {

	c:=make(chan jobData)
	jobs:=[]jobData{}

	// strconv.Itoa : int를 str로 형변환
	pageUrl := baseURL + "&start=" + strconv.Itoa(page*10)

	fmt.Println("Requesting:",pageUrl)

	res,err := http.Get(pageUrl)

	if err != nil {
		fmt.Println("Error")
	}

	// 메모리 누수 방지
	defer res.Body.Close()

	doc,err := goquery.NewDocumentFromReader(res.Body)

	if err!=nil {
		fmt.Println("Error")
	}

	doc.Find(".job_seen_beacon").Each(func(i int, s *goquery.Selection){
		go getData(s,c)
	})

	for i:=0; i<page; i++ {
		job:= <-c
		fmt.Println(job)
		jobs = append(jobs,job)
		fmt.Println(jobs)
		c2 <- jobs // page별 목록
	}
}

func getData(s *goquery.Selection, c chan jobData) {
	title:= s.Find(".jobTitle>span").Text()
	location:=s.Find(".companyLocation").Text()

	c <- jobData{title: title,location: location}
}

func getPages() int {

	pages:=0

	res,err := http.Get(baseURL)

	// check Error
	checkErr(err)
	// check StatusCode
	checkStatus(res)

	// leak memory
	defer res.Body.Close()

	doc,err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)

	// class 명을 이용해서 찾기
	doc.Find(".pagination").Each(func(i int,s *goquery.Selection){
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err!=nil {
		log.Fatal(err)
	}
}

func checkStatus(res *http.Response){
	if res.StatusCode != 200 {
		fmt.Println("Status Code:",res.StatusCode)
	}
}