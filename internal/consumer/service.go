package consumer

import (
	"fmt"
	"sync"
)

type Item struct {
	By, Title, Url               string
	ItemType                     string `json:"type"`
	Descendants, Id, Score, Time int
	Kids                         []int
}

func Consume(httpService DataService) {
	ids := getTopStories(httpService)

	idChan := make(chan int)
	itemChan := make(chan Item)

	go populateIdChan(idChan, ids)

	go fanOutIds(idChan, itemChan, httpService)

	for i := range itemChan {
		fmt.Println(i)
	}
}

func fanOutIds(idChan chan int, itemChan chan Item, httpService DataService) {
	var wg sync.WaitGroup
	const goRoutines = 10
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			for id := range idChan {
				func(id2 int) {
					item := getItem(httpService, id2)
					itemChan <- item
				}(id)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(itemChan)
}

func populateIdChan(c chan int, ids []int) {
	for _, id := range ids {
		c <- id
	}
	close(c)
}
