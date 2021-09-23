package consumer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// DataService - Interface for getting hackernews data
type DataService interface {
	getTopStories() ([]int, error)
	getItem(id int) (*Item, error)
}

// HttpService - Implementation of the Dataservice for http
type HttpService struct {}

const hnUrl string = "https://hacker-news.firebaseio.com/"

func (hs HttpService) getTopStories() ([]int, error) {
	url := fmt.Sprintf("%sv0/topstories.json?print=pretty", hnUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get top stories: ")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "get top stories: ")
	}

	var ids = []int{}
	err = json.Unmarshal(responseData, &ids)
	if err != nil {
		return nil, errors.Wrap(err, "get top stories: ")
	}

	return ids, nil
}

func (hs HttpService) getItem(id int) (*Item, error) {
	url := fmt.Sprintf("%sv0/item/%d.json?print=pretty", hnUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get item: ")
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "get item: ")
	}

	var item Item
	err = json.Unmarshal(responseData, &item)
	if err != nil {
		return nil, errors.Wrap(err, "get item: ")
	}


	return &item, nil
}
