package CourierService

// Calculate will calculate and diaplay the discount amount, price to be paid after discount and time to deliver
func Calculate(orders_list map[int][]string, base_delivery_cost float64, number_of_vehicals int, max_speed float64, max_weight float64) {
	orders_cost := Calculate_cost(orders_list, base_delivery_cost)
	orders_time := Calculate_time(orders_list, number_of_vehicals, max_speed, max_weight)
	Display_Order_Cost_with_time(orders_cost, orders_time)
}
