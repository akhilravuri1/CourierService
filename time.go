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
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		subsetSum := float64(0)

		for _, pkgName := range subset {
			subsetSum += pkgWeight[pkgName]
		}
		// add subset to subsets
		// check whether subsetSum is less that maxWeight
		// grater than existing max subset weight
		if subsetSum <= maxWeight && subsetSum > maxSubsetSum {
			maxSubsetSum = subsetSum
			maxSubset = subset
		}
	}
	return maxSubset
}

// createVehicals returns a list of vehicals with time
func createVehicals(numberOfVehicals int) (vehicalsList map[string]float64) {
	vehicalsList = make(map[string]float64)
	for i := 0; i < numberOfVehicals; i++ {
		vehicalName := "vehical" + strconv.Itoa(i)
		vehicalsList[vehicalName] = 0
	}
	return vehicalsList
}

// createMapWithPkgnameDistance returns map with pkgname and distance
func createMapWithPkgnameDistance(orderList map[string][]string) (orderWithDistance map[string]float64) {
	orderWithDistance = make(map[string]float64)
	for key := range orderList {
		orderWithDistance[key] = convertToFloat(orderList[key][1])
	}
	return orderWithDistance
}

// checkFastestAvailableVehical will return the earliest available vehical
func checkFastestAvailableVehical(vehicalsList map[string]float64) (vehical string, timeAvailability float64) {
	i := 0
	for key, value := range vehicalsList {
		if i == 0 {
			timeAvailability = value
			vehical = key
			i++
		} else {
			if timeAvailability > value {
				timeAvailability = value
				vehical = key
			}
		}
	}
	return vehical, timeAvailability
}

// calculateDuration will return a map with pkg_name as key and time as value
func calculateDuration(combinations [][]string, orderList map[string][]string, numberOfVehicals int, maxSpeed float64, maxWeight float64) map[string]float64 {
	vehicalsList := createVehicals(numberOfVehicals)
	orderWithDistance := createMapWithPkgnameDistance(orderList)
	resultMapWithTime := make(map[string]float64)
	for _, combination := range combinations {
		vehical, timeAvailability := checkFastestAvailableVehical(vehicalsList)
		presentOrderDeliveryTime := float64(0)
		for _, pkgName := range combination {
			// temp contains the time of current package
			temp := orderWithDistance[pkgName] / maxSpeed
			resultMapWithTime[pkgName] = temp + timeAvailability
			// This condition to set the total time taken to deliver all the packages
			// in a single trips
			if presentOrderDeliveryTime < temp {
				presentOrderDeliveryTime = temp
			}
		}
		// total time*2 is beacuse the vehical has to return to the store point
		// to pick another package
		vehicalsList[vehical] = timeAvailability + (presentOrderDeliveryTime * 2)
	}
	return resultMapWithTime
}

// calculateTime will calculate the time to be taken to deliver the order
func calculateTime(orderList map[string][]string, numberOfVehicals int, maxSpeed float64, maxWeight float64) map[string][]string {
	tempOrders := make(map[string]float64)
	// pkgDelSeq is an array of sequence in which packages will be delivered
	var pkgDelSeq [][]string
	// temp map is created with pkg name as key and weight as value
	for key, value := range orderList {
		tempOrders[key] = convertToFloat(value[0])
	}
	for len(tempOrders) != 0 {
		combination := getMaxSubset(tempOrders, maxWeight)
		// store all the set of packages which need to be delivered
		pkgDelSeq = append(pkgDelSeq, combination)
		// removing the max set of packages to find the max set from other packages
		for _, pkgName := range combination {
			delete(tempOrders, pkgName)
		}
	}
	pkgWithDuration := calculateDuration(pkgDelSeq, orderList, numberOfVehicals, maxSpeed, maxWeight)
	// adding time_duration to the order details
	for key := range orderList {
		for key1, value1 := range pkgWithDuration {
			if key == key1 {
				orderList[key] = append(orderList[key], fmt.Sprintf("%f", value1))
			}
		}
	}
	return orderList
}
