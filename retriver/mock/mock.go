package mock

type Retrieve struct {
	Contents string
}

func (r Retrieve) Get(url string) string {
	return r.Contents
}

