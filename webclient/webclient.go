package webclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"taudience.com/number-service/ds"
)

/*
Data ...
*/
type Data struct {
	Numbers []int `json:"numbers"`
}

/*
RequestWeb ...
*/
func RequestWeb(requesturls string) ds.Result {
	request, err := http.NewRequest("GET", requesturls, nil)
	if err != nil {
		return ds.Result{Numbers: nil, Error: err}
	}

	request.Header.Add("Content-Type", "application/json")

	bytes, err := do(request)
	if err != nil {
		return ds.Result{Numbers: nil, Error: err}
	}

	var data Data
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return ds.Result{Numbers: nil, Error: err}
	}

	return ds.Result{Numbers: data.Numbers, Error: nil}
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
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
