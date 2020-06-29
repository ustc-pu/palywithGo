package main

import (
	"fmt"
	"pcheng/playwithGo/retriver/mock"
	"pcheng/playwithGo/retriver/real"
	"time"
)
const url = "http://www.ucsd.edu"

type Retriever interface{
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(p Poster)  {
	p.Post(url, map[string]string{
		"course" : "golang",
		"name": "Patrcik",
	})
}

//interface combination
type RetrieverPoster interface {
	Poster
	Retriever
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked ucsd.edu!",
	})
	return s.Get(url)
}


/*
this is how we declare an interface
and implement the interface
*/
type fakeType struct {

}

func (f fakeType) Get(url string) string {
	return "fake implementation"
}

func main() {
	var r Retriever
	retriever := &mock.Retrieve{"it's a fake ucsd.edu"}
	r = retriever
	fmt.Println(download(r))
	inspect(r)
	fmt.Println("try a session:")
	fmt.Println(session(retriever))

	r = &real.Retrieve{"real retriever", time.Millisecond}
	inspect(r)
	//fmt.Println(download(r))

	if mockRetrieve, ok := r.(*mock.Retrieve); ok {
		fmt.Println(mockRetrieve.Contents)
	} else {
		fmt.Printf("it is a %T, not a %T\n", r, mockRetrieve)
	}
}

func inspect(r Retriever) {
	fmt.Println("Inspecting:")
	fmt.Printf("> %T, %v\n", r, r)
	fmt.Printf("> type switch: ")
	switch v := r.(type) {
	case *mock.Retrieve:
		fmt.Println("Contents:", v.Contents)
	case *real.Retrieve:
		fmt.Println("user agent:", v.UserAgent)
	}
	fmt.Println()
}