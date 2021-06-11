package CourierService

import (
	"fmt"
	"strconv"
)

func getKeysOfMap(myMap map[string]float64) []string {
	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	return keys
}

// all_combinations returns all combinations for a given string array with weight less than given weight
func getMaxSubset(pkg_weight map[string]float64, max_weight float64) (subsets []string) {
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
		if subset_sum <= max_weight {
			if subset_sum > max_subset_sum {
				max_subset_sum = subset_sum
				subsets = subset
			}
		}
	}
	return subsets
}

func create_vehicals(number_of_vehicals int) (vehicals_list map[string]float64) {
	vehicals_list = make(map[string]float64)
	for i := 0; i < number_of_vehicals; i++ {
		vehical_name := "vehical" + strconv.Itoa(i)
		vehicals_list[vehical_name] = 0
	}
	return vehicals_list
}

func create_map_with_pkgname_distance(order_list map[int][]string) (order_with_distance map[string]float64) {
	order_with_distance = make(map[string]float64)
	for i := 0; i < len(order_list); i++ {
		order_with_distance[order_list[i][0]] = convertToFloat(order_list[i][2])
	}
	return order_with_distance
}

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

func calculate_duration(combinations [][]string, order_list map[int][]string, number_of_vehicals int, max_speed float64, max_weight float64) map[string]float64 {
	vehicals_list := create_vehicals(number_of_vehicals)
	order_with_distance := create_map_with_pkgname_distance(order_list)
	result_map_with_time := make(map[string]float64)
	for _, combination := range combinations {
		vehical, time_availability := check_fastest_available_vehical(vehicals_list)
		present_order_delivery_time := float64(0)
		for _, pkg_name := range combination {
			temp := order_with_distance[pkg_name] / max_speed
			result_map_with_time[pkg_name] = temp + time_availability
			if present_order_delivery_time < temp {
				present_order_delivery_time = temp
			}
		}
		// for _, pkg_name := range combination {
		// 	result_map_with_time[pkg_name] = time_availability + (present_order_delivery_time * 2)
		// }
		vehicals_list[vehical] = time_availability + (present_order_delivery_time * 2)
	}
	return result_map_with_time
}

// calculate time
func Calculate_time(order_list map[int][]string, number_of_vehicals int, max_speed float64, max_weight float64) map[int][]string {
	temp_orders := make(map[string]float64)
	var combinations [][]string
	for _, value := range order_list {
		temp_orders[value[0]] = convertToFloat(value[1])
	}
	for len(temp_orders) != 0 {
		combination := getMaxSubset(temp_orders, max_weight)
		combinations = append(combinations, combination)
		for _, pkg_name := range combination {
			delete(temp_orders, pkg_name)
		}
	}
	pkg_with_duration := calculate_duration(combinations, order_list, number_of_vehicals, max_speed, max_weight)
	for key, value := range order_list {
		for key1, value1 := range pkg_with_duration {
			if value[0] == key1 {
				order_list[key] = append(order_list[key], fmt.Sprintf("%f", value1))
			}
		}
	}
	return order_list
}

// display_Order_Cost_with_time will just display the total cost after discount
func Display_Order_Cost_with_time(orders_cost, orders_time map[int][]string) {
	fmt.Println("\nORDERS COST WITH TIME:- ")
	for i := 0; i < len(orders_time); i++ {
		orders_cost[i] = append(orders_cost[i], orders_time[i][4])
	}
	for i := 0; i < len(orders_cost); i++ {
		fmt.Println(orders_cost[i][0], " ", orders_cost[i][1], " ", orders_cost[i][2], " ", orders_cost[i][3])
	}
}
