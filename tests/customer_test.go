package tests

import (
	"data_manager/data"
	"reflect"
	"testing"
)

func Test_GetCustomers(t *testing.T) {
	var c []data.Customer
	customers := data.GetCustomers()

	if reflect.TypeOf(customers) != reflect.TypeOf(c) && len(customers) <= 0 {
		t.Error("Incorrect result: expected []Customers, got", reflect.TypeOf(customers))
	}
}
