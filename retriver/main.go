package main

import (
	"fmt"
	"retriver/mock"
	real2 "retriver/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func main() {
	url := "https://www.baidu.com"
	var r Retriever = mock.Retriever{Contents: "retrieve the content from url " + url}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r, url))

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	//fmt.Println(download(r, url))

	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.Timeout)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}
