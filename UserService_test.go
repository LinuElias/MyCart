package main

import (
	"MyCart/cmd/Services"
	"errors"
	"fmt"
	"testing"
)

func TestAddProduct(t *testing.T) {
	var tests = []struct {
		userValue string
		userName  string
		expected  error
	}{

		{"Shirt", "Linu", nil}, //Already Prsent User
		{"qwert", "Linu", errors.New("Sorry!! Entered Product is not present")}, //Product wrong(not in catlog)
		{"", "Tina", errors.New("Data Not Found")},                              //Product empty
		{"LG", "", errors.New("Data Not Found")},                                //user empty
		{"Sony", "Rahul", nil},                                                  //New User
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.userValue)
		t.Run(testname, func(t *testing.T) {
			ans := Services.AddProduct(test.userValue, test.userName)
			if ans != test.expected {
				t.Fail()
				t.Error("Test Failed: recieved:", ans)
			} else {
				t.Log("Test Pass for ", test)
			}
		})
	}
}

func TestRemoveProduct(t *testing.T) {
	var tests = []struct {
		userValue string
		userName  string
		expected  error
	}{

		{"Shirt", "Linu", nil}, //Already Prsent User
		{"qwert", "Linu", errors.New("This Product was not present in cart")}, //Product wrong(not in catlog)
		{"", "Tina", errors.New("Data Not Found")},                            //Product empty
		{"LG", "", errors.New("Data Not Found")},                              //user empty
		{"Sony", "Rahul", nil},                                                //New User
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.userValue)
		t.Run(testname, func(t *testing.T) {
			ans := Services.AddProduct(test.userValue, test.userName)
			if ans != test.expected {
				t.Fail()
				t.Error("Test Failed: recieved:", ans)
			} else {
				t.Log("Test Pass for ", test)
			}
		})
	}
}
