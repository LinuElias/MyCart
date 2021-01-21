package Utils

import (
	"MyCart/cmd/Models"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/pquerna/ffjson/ffjson"
)

func ReadFile(filePath string) ([]byte, error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error in Reading File")
		return nil, err
	}
	return data, nil
}
func GetCatlog() ([]Models.Catlog, error) {
	catlog := []Models.Catlog{}
	catlogBytes, err := ReadFile("FileDB/catlog.json")
	if err != nil {
		fmt.Println("Error while reading catlog1")
		return catlog, err
	}
	err = ffjson.Unmarshal(catlogBytes, &catlog)
	if err != nil {
		fmt.Println("Error while reading catlog2")
		return catlog, err
	}
	return catlog, nil
}

func GetProductDetails(userValue string) (Models.Products, error) {
	isPresent := false
	productObj := Models.Products{}
	catlog, err := GetCatlog()
	if err != nil {
		fmt.Println("Error while reading catlog")
		return productObj, err
	}
	for i := 0; i < len(catlog); i++ {
		for _, subCategory := range catlog[i].SubCategory {
			for _, product := range subCategory.Products {
				if product.ProductName == userValue {
					isPresent = true
					productObj = product

					break
				}
			}
		}
	}
	if !isPresent {
		fmt.Println("Sorry!! Entered Product is not present")
	}
	return productObj, nil
}
func GetProduct(userValue string) (Models.Cart, error) {
	isPresent := false
	cartObj := Models.Cart{}
	catlog, err := GetCatlog()
	if err != nil {
		fmt.Println("Error while reading catlog")
		return cartObj, err
	}
	for i := 0; i < len(catlog); i++ {
		for _, subCategory := range catlog[i].SubCategory {
			for _, product := range subCategory.Products {
				if product.ProductName == userValue {
					isPresent = true
					cartObj.Product = product
					cartObj.CategoryName = catlog[i].CategoryName
					cartObj.SubCategoryName = subCategory.SubCategoryName
					break
				}
			}
		}
	}
	if !isPresent {
		fmt.Println("Sorry!! Entered Product is not present")
		return cartObj, errors.New("Sorry!! Entered Product is not present")
	}
	return cartObj, nil
}

func GetUserCart(filePath string) ([]Models.Cart, error) {
	cart := []Models.Cart{}
	cartBytes, err := ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while reading cart")
		return cart, err
	}
	err = ffjson.Unmarshal(cartBytes, &cart)
	if err != nil {
		fmt.Println("Error while reading cart2")
		return cart, err
	}
	return cart, nil
}
func GetUserBill(filePath string) (Models.Bill, error) {
	bill := Models.Bill{}
	billBytes, err := ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while reading cart")
		return bill, err
	}
	err = ffjson.Unmarshal(billBytes, &bill)
	if err != nil {
		fmt.Println("Error while reading cart2")
		return bill, err
	}
	return bill, nil
}
