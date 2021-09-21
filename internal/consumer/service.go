package consumer

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Consume(httpService DataService) {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	resp, err := httpService.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }
    fmt.Println(string(responseData))
}
