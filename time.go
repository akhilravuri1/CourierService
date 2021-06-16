package CourierService

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
func getMaxSubset(pkg_weight map[string]float64, max_weight float64) (max_subset []string) {
	max_subset_sum := float64(0)
	set := getKeysOfMap(pkg_weight)
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
		subset_sum := float64(0)

		for _, pkg_name := range subset {
			subset_sum += pkg_weight[pkg_name]
		}
		// add subset to subsets
		// check whether subset_sum is less that max_weight
		// grater than existing max subset weight
		if subset_sum <= max_weight && subset_sum > max_subset_sum {
			max_subset_sum = subset_sum
			max_subset = subset
		}
	}
	return max_subset
}

// create vehicals
func create_vehicals(number_of_vehicals int) (vehicals_list map[string]float64) {
	vehicals_list = make(map[string]float64)
	for i := 0; i < number_of_vehicals; i++ {
		vehical_name := "vehical" + strconv.Itoa(i)
		vehicals_list[vehical_name] = 0
	}
	return vehicals_list
}

// create_map_with_pkgname_distance returns map with pkgname and distance
func create_map_with_pkgname_distance(order_list map[string][]string) (order_with_distance map[string]float64) {
	order_with_distance = make(map[string]float64)
	for key := range order_list {
		order_with_distance[key] = convertToFloat(order_list[key][1])
	}
	return order_with_distance
}

// check_fastest_available_vehical will return the earliest available vehical
func check_fastest_available_vehical(vehicals_list map[string]float64) (vehical string, time_availability float64) {
	i := 0
	for key, value := range vehicals_list {
		if i == 0 {
			time_availability = value
			vehical = key
			i++
		} else {
			if time_availability > value {
				time_availability = value
				vehical = key
			}
		}
	}
	return vehical, time_availability
}

// calculate_duration will return a map with pkg_name as key and time as value
func calculate_duration(combinations [][]string, order_list map[string][]string, number_of_vehicals int, max_speed float64, max_weight float64) map[string]float64 {
	// create vehicals with initial value as zero
	vehicals_list := create_vehicals(number_of_vehicals)
	// returns a map with key of pkg_name and value of distance
	order_with_distance := create_map_with_pkgname_distance(order_list)
	result_map_with_time := make(map[string]float64)
	for _, combination := range combinations {
		// return the vehical name and time of availability
		vehical, time_availability := check_fastest_available_vehical(vehicals_list)
		present_order_delivery_time := float64(0)
		for _, pkg_name := range combination {
			// temp contains the time of current package
			temp := order_with_distance[pkg_name] / max_speed
			result_map_with_time[pkg_name] = temp + time_availability
			// This condition to set the total time taken to deliver all the packages
			// in a single trips
			if present_order_delivery_time < temp {
				present_order_delivery_time = temp
			}
		}
		// total time*2 is beacuse the vehical has to return to the store point
		// to pick another package
		vehicals_list[vehical] = time_availability + (present_order_delivery_time * 2)
	}
	return result_map_with_time
}

// calculate_time will calculate the time to be taken to deliver the order
func calculate_time(order_list map[string][]string, number_of_vehicals int, max_speed float64, max_weight float64) map[string][]string {
	temp_orders := make(map[string]float64)
	// pkg_del_seq is an array of sequence in which packages will be delivered
	var pkg_del_seq [][]string
	// temp map is created with pkg name as key and weight as value
	for key, value := range order_list {
		temp_orders[key] = convertToFloat(value[0])
	}
	for len(temp_orders) != 0 {
		// getMaxSubset will return the set of packages with max weight
		combination := getMaxSubset(temp_orders, max_weight)
		// store all the set of packages which need to be delivered
		pkg_del_seq = append(pkg_del_seq, combination)
		// removing the max set of packages to find the max set from other packages
		for _, pkg_name := range combination {
			delete(temp_orders, pkg_name)
		}
	}
	// calculate_duration will return the time taken to deliver each package
	pkg_with_duration := calculate_duration(pkg_del_seq, order_list, number_of_vehicals, max_speed, max_weight)
	// adding time_duration to the order details
	for key := range order_list {
		for key1, value1 := range pkg_with_duration {
			if key == key1 {
				order_list[key] = append(order_list[key], fmt.Sprintf("%f", value1))
			}
		}
	}
	return order_list
}
