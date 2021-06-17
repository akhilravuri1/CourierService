# courierservice

Calculates order amount after discount and time to deliver the order.

## Prerequisite

Golang should be installed.

## How to install

```
go get github.com/akhilravuri1/courierservice

```

## How to use this package

### example.go:-

```go
package main

import (
	"github.com/akhilravuri1/courierservice"
	"fmt"
)

func main() {
    //pkg_name : weight, distance, offer_code
    ordersList := map[string][]string{
		"PKG1": {"50", "50", "OFR001"},
		"PKG2": {"149", "125", "OFR002"},
		"PKG3": {"100", "140", "OFR003"},
		"PKG4": {"100", "125", "OFR004"},
	}
    baseDeliveryCost := float64(10)
    // you can pass only int's
    numberOfVehicles := 2
    // you can pass float values
    maxSpeed := float64(70)
    maxWeight :=  float64(200)
    courierservice.Calculate(ordersList, base_delivery_cost, numberOfVehicles, maxSpeed, maxWeight)
    orderWithDiscount := courierservice.CalulateCost(ordersList, baseDeliveryCost)
    orderWithDeliveryTime := courierservice.CalculateTime(ordersList,numberOfVehicles,maxSpeed,maxWeight)
    DisplayOrderCostWithTime(orderWithDiscount,orderWithDeliveryTime)

}
```
To run the sample:- go run example.go

### output:-
```
ORDERS COST WITH TIME:-
PKG1   0.000000   760.000000   0.714286hr
PKG2   148.750000   1976.250000   1.785714hr
PKG3   85.500000   1624.500000   2.000000hr
PKG4   0.000000   1635.000000   1.785714hr
```
