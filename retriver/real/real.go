package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retrieve struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Retrieve) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	res, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(res)
}
