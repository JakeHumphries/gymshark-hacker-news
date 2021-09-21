package consumer

import "net/http"

type DataService interface {
	Get(url string) (resp *http.Response, err error)
}

type HttpService struct {
}

func (hs HttpService) Get(url string) (resp *http.Response, err error) {
	r, e := http.Get(url)
	return r, e
}
