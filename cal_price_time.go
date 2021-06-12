package CourierService

import "fmt"

// Calculate will calculate and diaplay the discount amount, price to be paid after discount and time to deliver
func Calculate(orders_list map[string][]string, base_delivery_cost float64, number_of_vehicals int, max_speed float64, max_weight float64) {
	for id, pkg_details := range orders_list {
		if convertToFloat(pkg_details[1]) > max_weight {
			fmt.Println("Removing", pkg_details[0], "from the orders because it's weight is graterthan the vehical maximum carring weight")
			delete(orders_list, id)
		}
	}
	orders_cost := calculate_cost(orders_list, base_delivery_cost)
	orders_time := calculate_time(orders_list, number_of_vehicals, max_speed, max_weight)
	display_order_cost_with_time(orders_cost, orders_time)
}
