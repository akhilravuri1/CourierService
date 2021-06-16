package main

import (
	"fmt"
	"reflect"

	"github.com/akhilravuri1/CourierService"
)

func assert_equal(actual_result, returned_result map[string][]string) bool {
	return reflect.DeepEqual(actual_result, returned_result)
}

func DiscountWithValidData() bool {
	//pkg_name : weight, distance, offer_code
	orders_list := map[string][]string{
		"PKG1": {"50", "50", "OFR001"},
		"PKG2": {"149", "125", "OFR002"},
	}
	actual_result := map[string][]string{
		"PKG1": {"0.000000", "760.000000"},
		"PKG2": {"148.750000", "1976.250000"},
	}
	base_delivery_cost := float64(10)
	returned_result := CourierService.CalculateCost(orders_list, base_delivery_cost)
	//fmt.Println(returned_result)
	return assert_equal(actual_result, returned_result)
}

func main() {
	result := DiscountWithValidData()
	if result != true {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
