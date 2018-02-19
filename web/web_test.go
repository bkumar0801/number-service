package web_test

import (
	"reflect"
	"testing"

	"taudience.com/number-service/web"
)

func TestSortKeysShouldSortKeys(t *testing.T) {
	m := make(map[int]bool)
	m[10] = true
	m[3] = true
	m[70] = true
	actual := web.SortKeys(m)
	expected := []int{3, 10, 70}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Not equal = actual %v expected %v", actual, expected)
	}

}
