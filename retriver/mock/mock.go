package mock

import "fmt"

type Retrieve struct {
	Contents string
}

/**
we do need to use the key word implements(we do this all the time in
other languages). Define a struct and implement the function Get for
Retriever and Post for Poster. Note that the Retrieve value type is the receiver.
Then the compiler knows that Retrieve has implemented Retriever and Poster.
**/

func (r *Retrieve) Get(url string) string {
	return r.Contents
}

func (r *Retrieve) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok."
}

//implements Stringer interface
func (r Retrieve) String() string {
	return fmt.Sprintf("Retriever: {contents=%s}", r.Contents)
}

