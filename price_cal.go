package CourierService

import (
	"fmt"
	"log"
	"strconv"
)

// conevrt string to float
func convertToFloat(in string) float64 {
	float_value, err := strconv.ParseFloat(in, 64)
	if err != nil {
		log.Fatalln("Error while converting to float64: ", err)
	}
	return float_value
}

// discount_Calculate will calculate the amount to be detectud after discount
func discount_Calculate(total_cost_before_discount float64, offer_code string, pkg_weight float64, pkg_distance float64) float64 {
	var discount float64
	// check for the offer_code and apply the formula
	if offer_code == "OFR001" {
		if 70 <= pkg_weight && pkg_weight <= 200 && pkg_distance < 200 {
			discount = (total_cost_before_discount * 10) / 100
		}
	} else if offer_code == "OFR002" {
		if 100 <= pkg_weight && pkg_weight <= 250 && pkg_distance <= 150 && pkg_distance >= 50 {
			discount = (total_cost_before_discount * 7) / 100
		}
	} else if offer_code == "OFR003" {
		if 10 <= pkg_weight && pkg_weight <= 100 && pkg_distance <= 250 && pkg_distance >= 50 {
			discount = (total_cost_before_discount * 5) / 100
		}
	} else {
		discount = 0
	}
	return discount
}

// calulate the total cost of a pkg
func Calculate_cost(order_list map[int][]string, base_delivery_cost float64) map[int][]string {
	orders_cost := make(map[int][]string)
	// iterate through each order and calculate the cost
	for key, pkg_details := range order_list {
		// convert weight and distance to float
		pkg_weight := convertToFloat(pkg_details[1])
		pkg_distance := convertToFloat(pkg_details[2])
		// calculate the total distance before discount using the formula
		total_cost_before_discount := base_delivery_cost + (pkg_weight * float64(10)) + (pkg_distance * float64(5))
		// calcuate the discount
		discount := discount_Calculate(total_cost_before_discount, pkg_details[3], pkg_weight, pkg_distance)
		total_cost_after_discount := total_cost_before_discount - discount
		// store the calculated data in a map and return the map
		orders_cost[key] = append(orders_cost[key], pkg_details[0])
		orders_cost[key] = append(orders_cost[key], fmt.Sprintf("%f", discount))
		orders_cost[key] = append(orders_cost[key], fmt.Sprintf("%f", total_cost_after_discount))
	}
	return orders_cost
}

// display_Order_Cost will just display the total cost after discount
func Display_Order_Cost(orders_cost map[int][]string, number_of_packages int) {
	fmt.Println("\n ORDERS COST:- ")
	for i := 0; i < number_of_packages; i++ {
		fmt.Println(orders_cost[i][0], " ", orders_cost[i][1], " ", orders_cost[i][2])
	}
}
