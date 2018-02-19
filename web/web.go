package web

import (
	"fmt"
	"sort"
	"time"

	"taudience.com/number-service/ds"
	"taudience.com/number-service/webclient"
)

/*
Query ...
*/
func Query(urls []string) []int {
	c := make(chan ds.Result)
	set := make(map[int]bool)

	for _, requesturl := range urls {
		go func() { c <- webclient.RequestWeb(requesturl) }()
		timeout := time.After(500 * time.Millisecond)
		select {
		case result := <-c:
			for _, v := range result.Numbers {
				set[v] = true
			}
		case <-timeout:
			return nil
		}
	}
	
	fmt.Println("set :", SortKeys(set))
	return SortKeys(set)
}

/*
SortKeys ...
*/
func SortKeys(data map[int]bool) []int {
	var keys []int
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
