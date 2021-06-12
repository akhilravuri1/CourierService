# CourierService

Calulates order amount after discount and time to deliver the order.

## Prerequisite

Golang should be installed.

## How to install

```
go get github.com/akhilravuri1/CourierService

```

## How to use this package

### example1.go:-

```go
package main

import (
	"database/sql"
	"fmt"
	"github.com/akhilravuri1/CourierService"
)

func main() {
    //pkg_name, weight, distance, offer_code
    orders_list := map[string][]string{
		"PKG1": {"50", "50", "OFR001"},
		"PKG2": {"149", "125", "OFR002"},
		"PKG3": {"100", "140", "OFR003"},
		"PKG4": {"100", "125", "OFR004"},
	}
    base_delivery_cost := float64(10)
    // you can pass only int's
    number_of_vehicals := 2
    // you can pass float values
    max_speed := float64(70)
    max_weight :=  float64(200)
    CourierService.Calculate(orders_list, base_delivery_cost, number_of_vehicals, max_speed, max_weight)

}
```
To run the sample:- go run example1.go

