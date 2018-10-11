package main

import (
	"fmt"
	"go-learning/retriver/mock"
	"go-learning/retriver/real"
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

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}

	fmt.Println(download(r, url))

	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)
}

func inspect(r Retriever) {

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}
