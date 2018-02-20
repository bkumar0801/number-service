package appcontext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"taudience.com/number-service/constant"
)

/*
Data ...
*/
type Data struct {
	Numbers []int `json:"numbers"`
}

/*
AppContext ...
*/
type AppContext struct {
}

/*
Query ...
*/
func (appCtx *AppContext) Query(urls []string) []int {
	c := make(chan Result)
	set := make(map[int]bool)

	for _, requesturl := range urls {
		go func() { c <- appCtx.Get(requesturl) }()
		timeout := time.After(constant.Timeout * time.Millisecond)
		select {
		case result := <-c:
			for _, v := range result.Numbers {
				set[v] = true
			}
		case <-timeout:
			return nil
		}
	}
	return SortKeys(set)
}

/*
Get ...
*/
func (appCtx *AppContext) Get(requesturl string) Result {
	request, err := http.NewRequest("GET", requesturl, nil)
	if err != nil {
		return Result{Numbers: nil, Error: err}
	}

	request.Header.Add("Content-Type", "application/json")

	bytes, err := do(request)
	if err != nil {
		return Result{Numbers: nil, Error: err}
	}

	var data Data
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return Result{Numbers: nil, Error: err}
	}

	return Result{Numbers: data.Numbers, Error: nil}
}

func do(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("Unexpected server response status code: %s", resp.Status)
	}

	return body, nil
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
