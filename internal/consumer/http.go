package consumer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type DataService interface {
	Get(url string) (resp *http.Response, err error)
}

type HttpService struct {
}

func (hs HttpService) Get(url string) (resp *http.Response, err error) {
	r, e := http.Get(url)
	return r, e
}

const hnUrl string = "https://hacker-news.firebaseio.com/"

func getTopStories(httpService DataService) ([]int, error) {
	url := fmt.Sprintf("%sv0/topstories.json?print=pretty", hnUrl)
	resp, err := httpService.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Get top stories: ")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Get top stories: ")
	}

	var ids = []int{}
	json.Unmarshal(responseData, &ids)

	return ids, nil
}

func getItem(httpService DataService, id int) (*Item, error) {
	url := fmt.Sprintf("%sv0/item/%d.json?print=pretty", hnUrl, id)
	resp, err := httpService.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Get item: ")
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Get item: ")
	}

	var item Item
	json.Unmarshal(responseData, &item)

	return &item, nil
}
