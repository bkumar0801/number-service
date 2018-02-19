package handlers

import (
	"encoding/json"
	"net/http"

	"taudience.com/number-service/filterurl"
	"taudience.com/number-service/web"
)

/*
Payload ...
*/
type Payload struct {
	Numbers []int `json:"numbers"`
}

/*
HandleNumbers ...
*/
func HandleNumbers(w http.ResponseWriter, r *http.Request) {
	urls := filterurl.Filter(r.URL.Query()["u"])
	numbers := web.Query(urls)
	jsonData, err := MarshalJSON(numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

/*
MarshalJSON ...
*/
func MarshalJSON(keys []int) ([]byte, error) {
	jsonData, err := json.Marshal(CreatePayload(keys))
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*
CreatePayload ...
*/
func CreatePayload(data []int) Payload {
	return Payload{Numbers: data}
}
