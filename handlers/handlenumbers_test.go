package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	h "taudience.com/number-service/handlers"
)

func TestHandleNumber(t *testing.T) {
	rootRequest, err := http.NewRequest("GET", "/numbers?u=http://example.com/primes", nil)
	if err != nil {
		t.Errorf("Root request error: %s", err)
	}

	response := httptest.NewRecorder()
	response.WriteHeader(http.StatusInternalServerError)

	cases := []struct {
		w                    *httptest.ResponseRecorder
		r                    *http.Request
		expectedResponseCode int
		expectedResponseBody []byte
	}{
		{
			w:                    httptest.NewRecorder(),
			r:                    rootRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(`{"numbers":null}`),
		},
		{
			w:                    response,
			r:                    rootRequest,
			expectedResponseCode: http.StatusInternalServerError,
			expectedResponseBody: []byte(`{"numbers":null}`),
		},
	}

	for _, c := range cases {

		h.HandleNumbers(c.w, c.r)

		if c.expectedResponseCode != c.w.Code {
			t.Errorf("\nStatus Code didn't match:\n\t%q\n\t%q", c.expectedResponseCode, c.w.Code)
		}

		if !bytes.Equal(c.expectedResponseBody, c.w.Body.Bytes()) {
			t.Errorf("\nBody didn't match:\n\t%q\n\t%q", string(c.expectedResponseBody), c.w.Body.String())
		}
	}
}
