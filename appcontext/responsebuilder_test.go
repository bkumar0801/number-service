package appcontext_test

import (
	"reflect"
	"testing"

	appCtx "taudience.com/number-service/appcontext"
)

func TestQueryShouldReturnEmptyList(t *testing.T) {
	appContext := &appCtx.AppContext{}
	urls := []string{"/numbers?u=x.y/primes", "/numbers?u=x.y/odd", "/numbers?u=x.y/rand"}
	actual := appContext.Query(urls)
	expected := []int{}
	if !reflect.DeepEqual(len(actual), len(expected)) {
		t.Errorf("Not equal = actual %v expected %v", actual, expected)
	}
}

func TestSortKeysShouldSortKeys(t *testing.T) {
	m := make(map[int]bool)
	m[10] = true
	m[3] = true
	m[70] = true
	actual := appCtx.SortKeys(m)
	expected := []int{3, 10, 70}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Not equal = actual %v expected %v", actual, expected)
	}

}
