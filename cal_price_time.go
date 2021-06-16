// Package courierservice provides support to calculate the discount ammount and time to deliver the order.
package courierservice

import "fmt"

// Calculate function will calculate and diaplay the discount amount, price to be paid after discount and time to deliver
func Calculate(orders_list map[string][]string, base_delivery_cost float64, number_of_vehicals int, max_speed float64, max_weight float64) {
	for id, pkg_details := range orders_list {
		if convertToFloat(pkg_details[0]) > max_weight {
			fmt.Println("Removing", id, "from the orders because it's weight is graterthan the vehical maximum carring weight")
			delete(orders_list, id)
		}
	}
	orders_cost := calculateCost(orders_list, base_delivery_cost)
	orders_time := calculateTime(orders_list, number_of_vehicals, max_speed, max_weight)
	DisplayOrderCostWithTime(orders_cost, orders_time)
}

// CalculateCost exposed API to return discount amount and cost after discount
func CalculateCost(orders_list map[string][]string, base_delivery_cost float64) map[string][]string {
	orders_cost := calculateCost(orders_list, base_delivery_cost)
	return orders_cost
}

// CalculateTime exposed API to return estimated time
func CalculateTime(orders_list map[string][]string, number_of_vehicals int, max_speed float64, max_weight float64) map[string][]string {
	for id, pkg_details := range orders_list {
		if convertToFloat(pkg_details[0]) > max_weight {
			fmt.Println("Removing", id, "from the orders because it's weight is greater than the vehical maximum carring weight")
			delete(orders_list, id)
		}
	}
	orders_time := calculateTime(orders_list, number_of_vehicals, max_speed, max_weight)
	return orders_time
}

// DisplayOrderCostWithTime will just display the total cost after discount
func DisplayOrderCostWithTime(orders_cost, orders_time map[string][]string) {
	fmt.Println("\nORDERS COST WITH TIME:- ")
	for key := range orders_time {
		orders_cost[key] = append(orders_cost[key], orders_time[key][3])
	}
	for pkg_name := range orders_cost {
		fmt.Println(pkg_name, " ", orders_cost[pkg_name][0], " ", orders_cost[pkg_name][1], " ", orders_cost[pkg_name][2]+"hr")
	}
}
