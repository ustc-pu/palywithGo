package main

import (
	"fmt"
	"pcheng/playwithGo/retriver/mock"
	"pcheng/playwithGo/retriver/real"
	"time"
)

type Retriever interface{
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.ucsd.edu")
}

type fakeType struct {

}

func (f fakeType) Get(url string) string {
	return "fake implementation"
}

func main() {
	var r Retriever
	r = mock.Retrieve{"it's a fake ucsd.edu"}
	//fmt.Println(download(r))
	inspect(r)

	r = &real.Retrieve{"real retriever", time.Millisecond}
	inspect(r)
	//fmt.Println(download(r))

	if mockRetrieve, ok := r.(mock.Retrieve); ok {
		fmt.Println(mockRetrieve.Contents)
	} else {
		fmt.Println("not a mock.Retrieve")
	}
}

func inspect(r Retriever) {
	fmt.Printf("%T, %v\ntype switch:\n", r, r)
	switch v := r.(type) {
	case mock.Retrieve:
		fmt.Println(v.Contents)
	case *real.Retrieve:
		fmt.Println("user agent: ", v.UserAgent)

	}
}