package consumer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type DataService interface {
	getTopStories() ([]int, error)
	getItem(id int) (*Item, error)
}

type HttpService struct {
}

const hnUrl string = "https://hacker-news.firebaseio.com/"

func (hs HttpService) getTopStories() ([]int, error) {
	url := fmt.Sprintf("%sv0/topstories.json?print=pretty", hnUrl)
	resp, err := http.Get(url)
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

func (hs HttpService) getItem(id int) (*Item, error) {
	url := fmt.Sprintf("%sv0/item/%d.json?print=pretty", hnUrl, id)
	resp, err := http.Get(url)
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
