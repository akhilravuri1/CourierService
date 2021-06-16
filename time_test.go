package courierservice

import (
	"testing"
)

func TestTimeWithValidWeight(t *testing.T) {
	//pkg_name : weight, distance, offer_code
	orderList := map[string][]string{
		"PKG1": {"50", "50", "OFR001"},
		"PKG2": {"149", "125", "OFR002"},
	}
	actualResult := map[string][]string{
		"PKG1": {"50", "50", "OFR001", "0.714286"},
		"PKG2": {"149", "125", "OFR002", "1.785714"},
	}
	numberOfVehicals := 2
	maxSpeed := float64(70)
	maxWeight := float64(200)
	returnedResult := CalculateTime(orderList, numberOfVehicals, maxSpeed, maxWeight)
	result := assert_equal(actualResult, returnedResult)

	if result != true {
		t.Errorf("TimeWithValidWeight failed")
	} else {
		t.Logf("TimeWithValidWeight success")
	}
}

func TestTimeWithInValidWeight(t *testing.T) {
	//pkg_name : weight, distance, offer_code
	orderList := map[string][]string{
		"PKG1": {"50", "50", "OFR001"},
		"PKG2": {"149", "125", "OFR002"},
		"PKG3": {"201", "140", "OFR003"},
	}
	actualResult := map[string][]string{
		"PKG1": {"50", "50", "OFR001", "0.714286"},
		"PKG2": {"149", "125", "OFR002", "1.785714"},
	}
	numberOfVehicals := 2
	maxSpeed := float64(70)
	maxWeight := float64(200)
	returnedResult := CalculateTime(orderList, numberOfVehicals, maxSpeed, maxWeight)
	result := assert_equal(actualResult, returnedResult)

	if result != true {
		t.Errorf("TestTimeWithInValidWeight failed")
	} else {
		t.Logf("TestTimeWithInValidWeight success")
	}
}
