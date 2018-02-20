package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	appCtx "taudience.com/number-service/appcontext"
	h "taudience.com/number-service/handler"
)

func TestHandleNumbersWithAppContext(t *testing.T) {
	rootRequest, err := http.NewRequest("GET", "/numbers?u=http://odd.com/odd&u=http://primes.com/primes", nil)
	if err != nil {
		t.Errorf("Root request error: %s", err)
	}

	cases := []struct {
		description          string
		w                    *httptest.ResponseRecorder
		r                    *http.Request
		expectedResponseCode int
		expectedResponseBody []byte
		appContext           appCtx.ResponseBuilder
	}{
		{
			description:          "Test status is ok",
			w:                    httptest.NewRecorder(),
			r:                    rootRequest,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(`{"numbers":null}`),
			appContext:           &appCtx.AppContext{},
		},
	}

	for _, c := range cases {

		h.HandleNumbers(c.appContext, c.w, c.r)
		actualResponseBody := bytes.Trim(c.w.Body.Bytes(), "\n")
		if c.expectedResponseCode != c.w.Code {
			t.Errorf("\nHandler returned wrong status code::\n\t%d\n\t%d", c.expectedResponseCode, c.w.Code)
		}

		if !bytes.Equal(c.expectedResponseBody, actualResponseBody) {
			t.Errorf("\nHandler returned unexpected body:\n\t%q\n\t%q", string(c.expectedResponseBody), string(actualResponseBody))
		}
	}
}

func TestHandleNumbersWithFakeAppContext(t *testing.T) {
	request, err := http.NewRequest("GET", "/numbers?u=http://example.com/x", nil)
	if err != nil {
		t.Errorf("Root request error: %s", err)
	}

	requestForPrimeAndFibNumber, err := http.NewRequest("GET", "/numbers?u=http://example.com/primes&u=http://foobar.com/fibo HTTP/1.0", nil)
	if err != nil {
		t.Errorf("Root request error: %s", err)
	}

	requestForOddAndRandNumber, err := http.NewRequest("GET", "/numbers?u=http://example.com/odd&u=http://foobar.com/rand", nil)
	if err != nil {
		t.Errorf("Root request error: %s", err)
	}

	cases := []struct {
		description          string
		w                    *httptest.ResponseRecorder
		r                    *http.Request
		expectedResponseCode int
		expectedResponseBody []byte
		appContext           appCtx.ResponseBuilder
	}{
		{
			description:          "Test InternalServerError status and response body for request to host url fails",
			w:                    httptest.NewRecorder(),
			r:                    request,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(`{"numbers":null}`),
			appContext:           &appCtx.FakeAppContext{},
		},
		{
			description:          "Test status OK and response body for prime number and fibonacci numbers request",
			w:                    httptest.NewRecorder(),
			r:                    requestForPrimeAndFibNumber,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(`{"numbers":[1,2,3,5,7,8,11,13,21]}`),
			appContext:           &appCtx.FakeAppContext{},
		},
		{
			description:          "Test status OK and response body for odd number and random numbers request",
			w:                    httptest.NewRecorder(),
			r:                    requestForOddAndRandNumber,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: []byte(`{"numbers":[1,3,5,7,9,11,13,15,17,19,21,23,30,50,77,93]}`),
			appContext:           &appCtx.FakeAppContext{},
		},
	}

	for _, c := range cases {

		h.HandleNumbers(c.appContext, c.w, c.r)
		actualResponseBody := bytes.Trim(c.w.Body.Bytes(), "\n")

		if c.expectedResponseCode != c.w.Code {
			t.Errorf("\nHandler returned wrong status code::\n\t%d\n\t%d", c.expectedResponseCode, c.w.Code)
		}

		if !bytes.Equal(c.expectedResponseBody, actualResponseBody) {
			t.Errorf("\nHandler returned unexpected body:\n\t%q\n\t%q", string(c.expectedResponseBody), string(actualResponseBody))
		}
	}
}

func TestMiddlewareIsCalled(t *testing.T) {
	request, err := http.NewRequest("GET", "/numbers?u=http://odd.com/odd&u=http://primes.com/primes", nil)
	if err != nil {
		t.Errorf("Root request error: %s", err)
	}
	response := httptest.NewRecorder()
	appContext := &appCtx.AppContext{}
	h.Middleware(appContext, h.HandleNumbers).ServeHTTP(response, request)
}
