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
ResponseData ...Json structure
*/
type ResponseData struct {
	Numbers []int `json:"numbers"`
}

/*
AppContext ...
*/
type AppContext struct {
}

/*
Query ...This function invokes a goroutine for each request url to fetch data from different URL.
Slow server response is ignored i.e. if an URL takes too long to respond, it is ignored.
If an URL returns an error, its response is ignored. It also merges the responses fetched from
different URLs and maintains uniqueness. Finally, returns array of sorted integer.
*/
func (appCtx *AppContext) Query(urls []string) []int {
	// Create a channel to store numbers fetched from different URLs
	c := make(chan Result)
	// Define a set data structure to keep unique key stored
	set := make(map[int]bool)

	for _, requesturl := range urls {
		// Invokes a go routine for each request url
		go func() { c <- appCtx.Get(requesturl) }()
		// Response is ignored after timeout (500ms)
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
Get ...This function makes a REST call to the given URL, and returns 'Result'
type Result struct {
	Numbers []int
	Error   error
}
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

	var response ResponseData
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return Result{Numbers: nil, Error: err}
	}

	return Result{Numbers: response.Numbers, Error: nil}
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

	// For any other response than OK, return an empty array as response body and error
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("Unexpected server response status code: %s", resp.Status)
	}

	return body, nil
}

/*
SortKeys ...This function sorts all keys of the map in ascending order
*/
func SortKeys(data map[int]bool) []int {
	var keys []int
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
