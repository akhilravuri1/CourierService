package courierservice

import (
	"reflect"
	"testing"
)

func assert_equal(actual_result, returned_result map[string][]string) bool {
	return reflect.DeepEqual(actual_result, returned_result)
}

func TestDiscountWithValidData(t *testing.T) {
	//pkg_name : weight, distance, offer_code
	orderList := map[string][]string{
		"PKG1": {"50", "50", "OFR001"},
		"PKG2": {"149", "125", "OFR002"},
	}
	actualResult := map[string][]string{
		"PKG1": {"0.000000", "760.000000"},
		"PKG2": {"148.750000", "1976.250000"},
	}
	baseDeliveryCost := float64(10)
	returnedResult := CalculateCost(orderList, baseDeliveryCost)
	result := assert_equal(actualResult, returnedResult)

	if result != true {
		t.Errorf("DiscountWithValidData failed")
	} else {
		t.Logf("DiscountWithValidData success")
	}
}

func TestValidOfferCode(t *testing.T) {
	discountAmount := discountCalculate(100, "OFR001", 100, 170)
	if discountAmount == 10 {
		t.Logf("ValidOfferCode success")
	} else {
		t.Errorf("ValidOfferCode failed and the Discount Amount is %v", discountAmount)
	}
}

func TestInValidOfferCode(t *testing.T) {
	discountAmount := discountCalculate(100, "OFR004", 100, 170)
	if discountAmount == 0 {
		t.Logf("InValidOfferCode success")
	} else {
		t.Errorf("InValidOfferCode failed and the Discount Amount is %v", discountAmount)
	}
}

func TestInValidOfferCriteria(t *testing.T) {
	discountAmount := discountCalculate(100, "OFR001", 10, 210)
	if discountAmount == 0 {
		t.Logf("InValidOfferCriteria success")
	} else {
		t.Errorf("InValidOfferCriteria failed and the Discount Amount is %v", discountAmount)
	}
}
