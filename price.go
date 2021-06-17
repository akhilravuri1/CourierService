package courierservice

import (
	"fmt"
	"log"
	"strconv"
)

// convertToFloat converts string to float
func convertToFloat(in string) float64 {
	floatValue, err := strconv.ParseFloat(in, 64)
	if err != nil {
		log.Fatalln("Error while converting to float64: ", err)
	}
	return floatValue
}

// discountCalculate will calculate the amount to be deducted after discount
func discountCalculate(totalCostBeforeDiscount float64, offerCode string, pkgWeight float64, pkgDistance float64) float64 {
	var discount float64
	if offerCode == "OFR001" {
		if 70 <= pkgWeight && pkgWeight <= 200 && pkgDistance < 200 {
			discount = (totalCostBeforeDiscount * 10) / 100
		}
	} else if offerCode == "OFR002" {
		if 100 <= pkgWeight && pkgWeight <= 250 && pkgDistance <= 150 && pkgDistance >= 50 {
			discount = (totalCostBeforeDiscount * 7) / 100
		}
	} else if offerCode == "OFR003" {
		if 10 <= pkgWeight && pkgWeight <= 100 && pkgDistance <= 250 && pkgDistance >= 50 {
			discount = (totalCostBeforeDiscount * 5) / 100
		}
	} else {
		// here we can add new offer codes
		discount = 0
	}
	return discount
}

// calculateCost will calculate the total cost of a pkg
func calculateCost(orderList map[string][]string, baseDeliveryCost float64) map[string][]string {
	ordersCost := make(map[string][]string)
	// iterate through each order and calculate the cost
	for key, pkgDetails := range orderList {
		// convert weight and distance to float
		pkgWeight := convertToFloat(pkgDetails[0])
		pkgDistance := convertToFloat(pkgDetails[1])
		// calculate the total distance before discount using the formula
		totalCostBeforeDiscount := baseDeliveryCost + (pkgWeight * float64(10)) + (pkgDistance * float64(5))
		// calcuate the discount
		discount := discountCalculate(totalCostBeforeDiscount, pkgDetails[2], pkgWeight, pkgDistance)
		totalCostAfterDiscount := totalCostBeforeDiscount - discount
		// store the calculated data in a map and return the map
		ordersCost[key] = append(ordersCost[key], fmt.Sprintf("%f", discount))
		ordersCost[key] = append(ordersCost[key], fmt.Sprintf("%f", totalCostAfterDiscount))
	}
	return ordersCost
}
