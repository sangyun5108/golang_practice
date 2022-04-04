package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

type jobData struct {
	title string
	location string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python" 

func main(){

	c1:=make(chan []jobData)
	jobs:=[]jobData{}

	totalPages:=getPages()
	
	for i:=0; i<totalPages; i++ {
		go getPage(i,c1)
		job := <- c1
		jobs = append(jobs, job...)
	}

	writeJobs(jobs)
}

func writeJobs(jobs []jobData){
	csv,err := ccsv.NewCsvWriter("jobs.csv")
	checkErr(err)
	done:=make(chan bool)
	

	// 데이터가 파일에 입력된다.
	defer csv.Close()


	for i:=0; i<len(jobs); i++ {
		go func(i int){
			csv.Write([]string{jobs[i].title,jobs[i].location})
			done <- true
		}(i)
	}

	for i:=0; i<len(jobs); i++ {
		<- done
	}
}

func getPage(page int, c1 chan <- []jobData)  {

	jobs:=[]jobData{}
	c:=make(chan jobData)

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

	cards := doc.Find(".job_seen_beacon")

	cards.Each(func(i int, s *goquery.Selection){
		go getData(s,c)
	})

	for i:=0; i<cards.Length(); i++ {
		job:= <-c
		fmt.Println(job)
		jobs = append(jobs,job)
	}

	c1 <- jobs
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