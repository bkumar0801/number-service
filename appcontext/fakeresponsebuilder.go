package appcontext

import (
	"errors"
	"sort"
	"strings"
	"time"
)

/*
FakeAppContext ...
*/
type FakeAppContext struct {
}

/*
Query ...
*/
func (app *FakeAppContext) Query(urls []string) []int {
	c := make(chan Result)
	set := make(map[int]bool)
	for _, requesturl := range urls {
		go func() { c <- app.Get(requesturl) }()
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
	return sortKeys(set)
}

/*
Get ...
*/
func (app *FakeAppContext) Get(requesturl string) Result {
	if strings.Contains(requesturl, "primes") {
		return Result{Numbers: []int{2, 3, 5, 7, 11, 13}, Error: nil}
	} else if strings.Contains(requesturl, "odd") {
		return Result{Numbers: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, Error: nil}
	} else if strings.Contains(requesturl, "fibo") {
		return Result{Numbers: []int{1, 1, 2, 3, 5, 8, 13, 21}, Error: nil}
	} else if strings.Contains(requesturl, "rand") {
		return Result{Numbers: []int{50, 77, 93, 30}, Error: nil}
	}
	return Result{Numbers: nil, Error: errors.New("mock error")}
}

func sortKeys(data map[int]bool) []int {
	var keys []int
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}