package main

import (
	"MyCart/cmd/Services"
	"errors"
	"fmt"
	"testing"
)

func TestAddProductsToCatlog(t *testing.T) {
	var tests = []struct {
		category      string
		subCategory   string
		productName   string
		price         float64
		specification string
		expected      error
	}{

		{"HomeApplications", "KitchenVessels", "Pan", 100, "", nil},    //New Category
		{"HomeApplications", "Furniture", "Bed", 7000, "Large", nil},   //New Sub-Category
		{"HomeApplications", "KitchenVessels", "Mixer", 4000, "", nil}, //New Product
	}

	for _, test := range tests {
		//testname := fmt.Sprintf("%s", test.userValue)
		t.Run("TestAddProductsToCatlog", func(t *testing.T) {
			ans := Services.AddProductsToCatlog(test.category, test.subCategory, test.productName, test.specification, test.price)
			if ans != test.expected {
				t.Fail()
				t.Error("Test Failed: recieved:", ans)
			} else {
				t.Log("Test Pass for ", test)
			}
		})
	}
}

func TestViewCartDetails(t *testing.T) {
	var tests = []struct {
		userName string
		expected error
	}{

		{"Linu", nil}, //Already Prsent User
		{"qwert", errors.New("No Items In Cartt")}, //No Cart
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.userName)
		t.Run(testname, func(t *testing.T) {
			ans := Services.ViewCartDetails(test.userName)
			if ans != test.expected {
				t.Fail()
				t.Error("Test Failed: recieved:", ans)
			} else {
				t.Log("Test Pass for ", test)
			}
		})
	}
}
