package courierservice

import (
	"fmt"
	"strconv"
)

// getKeysOfMap returns keys of map.
func getKeysOfMap(myMap map[string]float64) []string {
	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	return keys
}

// getMaxSubset returns a max subset for a given string array with total weight less than or equal to passed weight
func getMaxSubset(pkgWeight map[string]float64, maxWeight float64) (maxSubset []string) {
	maxSubsetSum := float64(0)
	set := getKeysOfMap(pkgWeight)
	length := uint(len(set))
	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsetSum := float64(0)

		for _, pkgName := range subset {
			subsetSum += pkgWeight[pkgName]
		}
		if subsetSum <= maxWeight && subsetSum > maxSubsetSum {
			maxSubsetSum = subsetSum
			maxSubset = subset
		}
	}
	return maxSubset
}

// createVehicles returns a list of vehicles with time
func createVehicles(numberOfVehicles int) (vehiclesList map[string]float64) {
	vehiclesList = make(map[string]float64)
	for i := 0; i < numberOfVehicles; i++ {
		vehicleName := "vehicle" + strconv.Itoa(i)
		vehiclesList[vehicleName] = 0
	}
	return vehiclesList
}

// createMapWithPkgnameDistance returns map with pkgname and distance
func createMapWithPkgnameDistance(orderList map[string][]string) (orderWithDistance map[string]float64) {
	orderWithDistance = make(map[string]float64)
	for key := range orderList {
		orderWithDistance[key] = convertToFloat(orderList[key][1])
	}
	return orderWithDistance
}

// checkFastestAvailablevehicle will return the earliest available vehicle
func checkFastestAvailablevehicle(vehiclesList map[string]float64) (vehicle string, timeAvailability float64) {
	i := 0
	for key, value := range vehiclesList {
		if i == 0 {
			timeAvailability = value
			vehicle = key
			i++
		} else {
			if timeAvailability > value {
				timeAvailability = value
				vehicle = key
			}
		}
	}
	return vehicle, timeAvailability
}

// calculateDuration will return a map with pkg_name as key and time as value
func calculateDuration(combinations [][]string, orderList map[string][]string, numberOfVehicles int, maxSpeed float64, maxWeight float64) map[string]float64 {
	vehiclesList := createVehicles(numberOfVehicles)
	orderWithDistance := createMapWithPkgnameDistance(orderList)
	resultMapWithTime := make(map[string]float64)
	for _, combination := range combinations {
		vehicle, timeAvailability := checkFastestAvailablevehicle(vehiclesList)
		presentOrderDeliveryTime := float64(0)
		for _, pkgName := range combination {
			temp := orderWithDistance[pkgName] / maxSpeed
			resultMapWithTime[pkgName] = temp + timeAvailability
			if presentOrderDeliveryTime < temp {
				presentOrderDeliveryTime = temp
			}
		}
		vehiclesList[vehicle] = timeAvailability + (presentOrderDeliveryTime * 2)
	}
	return resultMapWithTime
}

// calculateTime will calculate the time to be taken to deliver the order
func calculateTime(orderList map[string][]string, numberOfVehicles int, maxSpeed float64, maxWeight float64) map[string][]string {
	tempOrders := make(map[string]float64)
	var pkgDelSeq [][]string
	for key, value := range orderList {
		tempOrders[key] = convertToFloat(value[0])
	}
	for len(tempOrders) != 0 {
		combination := getMaxSubset(tempOrders, maxWeight)
		pkgDelSeq = append(pkgDelSeq, combination)
		for _, pkgName := range combination {
			delete(tempOrders, pkgName)
		}
	}
	pkgWithDuration := calculateDuration(pkgDelSeq, orderList, numberOfVehicles, maxSpeed, maxWeight)
	for key := range orderList {
		for key1, value1 := range pkgWithDuration {
			if key == key1 {
				orderList[key] = append(orderList[key], fmt.Sprintf("%f", value1))
			}
		}
	}
	return orderList
}
