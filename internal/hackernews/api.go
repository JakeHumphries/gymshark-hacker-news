package hackernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
)

// Api to get data from the hacker news api
type Api struct{}

const hackerNewsUrl string = "https://hacker-news.firebaseio.com/"

// GetTopStories gets the top stories from the hacker news api
func (a Api) GetTopStories() ([]int, error) {
	url := fmt.Sprintf("%sv0/topstories.json?print=pretty", hackerNewsUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get top stories: ")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("err: bad status back from hacker news, status: %d", resp.StatusCode))
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

// GetItem gets a specific item from the hacker news api based on an item id
func (a Api) GetItem(id int) (*models.Item, error) {
	url := fmt.Sprintf("%sv0/item/%d.json?print=pretty", hackerNewsUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get item: ")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("err: bad status back from hacker news, status: %d", resp.StatusCode))
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "get item: ")
	}

	var item models.Item
	err = json.Unmarshal(responseData, &item)
	if err != nil {
		return nil, errors.Wrap(err, "get item: ")
	}

	return &item, nil
}
