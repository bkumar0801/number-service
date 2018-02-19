package filterurl_test

import (
	"reflect"
	"testing"

	f "taudience.com/number-service/filterurl"
)

func TestIsValidURLToReturnTrueForValidURL(t *testing.T) {
	var tests = []struct {
		url      string
		expected bool
	}{
		{"http://travelaudience.com", true},
		{"https://travelaudience.com", true},
		{"http://www.travelaudience.com", true},
		{"http://foobar.com/fibo HTTP/1.0", true},
		//etc
	}

	for _, test := range tests {
		actual := f.IsValidURL(test.url)
		if actual != test.expected {
			t.Errorf("Not equal = actual %v expected %v", actual, test.expected)
		}
	}
}

func TestIsValidURLToReturnFalseForInvalidURL(t *testing.T) {
	var tests = []struct {
		url      string
		expected bool
	}{
		{"travelaudience.com", false},
		{"http//travelaudience.com", false},
		{"travelaudience", false},
		//etc
	}

	for _, test := range tests {
		actual := f.IsValidURL(test.url)
		if actual != test.expected {
			t.Errorf("Not equal = actual %v expected %v", actual, test.expected)
		}
	}
}

func TestFilterShouldFilterValidURLs(t *testing.T) {
	urls := []string{"http://travelaudience.com", "travelaudience.com",
		"travelaudience", "https://travelaudience.com", "http://foobar.com/fibo HTTP/1.0"}
	expected := []string{"http://travelaudience.com", "https://travelaudience.com", "http://foobar.com/fibo"}
	actual := f.Filter(urls)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Not equal = actual %v expected %v", actual, expected)
	}
}
