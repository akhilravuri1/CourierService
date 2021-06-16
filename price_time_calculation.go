// Package courierservice provides support to calculate the discount ammount and time to deliver the order.
package courierservice

import "fmt"

// Calculate function will calculate and diaplay the discount amount, price to be paid after discount and time to deliver
func Calculate(ordersList map[string][]string, baseDeliveryCost float64, numberOfVehicals int, maxSpeed float64, maxWeight float64) {
	for id, pkgDetails := range ordersList {
		if convertToFloat(pkgDetails[0]) > maxWeight {
			fmt.Println("Removing", id, "from the orders because it's weight is graterthan the vehical maximum carring weight")
			delete(ordersList, id)
		}
	}
	ordersCost := calculateCost(ordersList, baseDeliveryCost)
	ordersTime := calculateTime(ordersList, numberOfVehicals, maxSpeed, maxWeight)
	DisplayOrderCostWithTime(ordersCost, ordersTime)
}

// CalculateCost exposed API to return discount amount and cost after discount
func CalculateCost(ordersList map[string][]string, baseDeliveryCost float64) map[string][]string {
	ordersCost := calculateCost(ordersList, baseDeliveryCost)
	return ordersCost
}

// CalculateTime exposed API to return estimated time
func CalculateTime(ordersList map[string][]string, numberOfVehicals int, maxSpeed float64, maxWeight float64) map[string][]string {
	for id, pkgDetails := range ordersList {
		if convertToFloat(pkgDetails[0]) > maxWeight {
			fmt.Println("Removing", id, "from the orders because it's weight is greater than the vehical maximum carring weight")
			delete(ordersList, id)
		}
	}
	ordersTime := calculateTime(ordersList, numberOfVehicals, maxSpeed, maxWeight)
	return ordersTime
}

// DisplayOrderCostWithTime will just display the total cost after discount
func DisplayOrderCostWithTime(ordersCost, ordersTime map[string][]string) {
	fmt.Println("\nORDERS COST WITH TIME:- ")
	for key := range ordersTime {
		ordersCost[key] = append(ordersCost[key], ordersTime[key][3])
	}
	for pkgName := range ordersCost {
		fmt.Println(pkgName, " ", ordersCost[pkgName][0], " ", ordersCost[pkgName][1], " ", ordersCost[pkgName][2]+"hr")
	}
}
