package Services

import (
	"MyCart/cmd/Models"
	"MyCart/cmd/Utils"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pquerna/ffjson/ffjson"
)

func ViewCategories() error {
	//Get Catlog Details from file
	catlog, err := Utils.GetCatlog()
	if err != nil {
		fmt.Println("Error while reading catlog1")
		return err
	}
	fmt.Println("We have the following Categories:")
	for i := 0; i < len(catlog); i++ {
		fmt.Println(catlog[i].CategoryName)
	}
	return nil
}

func ViewSubCategories(userValue string) error {
	isPresent := false
	//Get Catlog Details from file
	catlog, err := Utils.GetCatlog()
	if err != nil {
		fmt.Println("Error while reading catlog")
		return err
	}
	fmt.Println("We have the following Sub-Categories for ", userValue, " category")
	//iterate to find sub category
	for i := 0; i < len(catlog); i++ {
		if catlog[i].CategoryName == userValue {
			isPresent = true
			for _, subCategory := range catlog[i].SubCategory {
				fmt.Println(subCategory.SubCategoryName)
			}
			break
		}

	}
	if !isPresent {
		fmt.Println("Sorry!! Entered Category is not present")
	}
	return nil
}
func ViewProducts(userValue string) error {
	isPresent := false
	//Get Catlog Details from file
	catlog, err := Utils.GetCatlog()
	if err != nil {
		fmt.Println("Error while reading catlog")
		return err
	}
	fmt.Println("We have the following Prodcuts for ", userValue, " category")
	//iterate to find products under sub-category
	for i := 0; i < len(catlog); i++ {
		if catlog[i].CategoryName == userValue {
			isPresent = true
			for _, subCategory := range catlog[i].SubCategory {
				for _, product := range subCategory.Products {
					fmt.Println("Product Name:", product.ProductName)
				}
			}
			break
		}

	}
	if !isPresent {
		fmt.Println("Sorry!! Entered Category is not present")
	}
	return nil
}
func ViewProductDetails(userValue string) error {
	// get details of a particular product
	product, err := Utils.GetProductDetails(userValue)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Name:", product.ProductName)
	fmt.Println("Price:", product.Price, "Rs")
	if product.Specification != "" {
		fmt.Println("Specification:", product.Specification)
	}
	return nil
}

func AddProduct(userValue, userName string) error {
	//check user input is empty
	if userValue == "" || userName == "" {
		return errors.New("Data Not Found")
	}
	filePath := "FileDB/Users/" + userName + "/Cart.json"
	//get the product obj based on product name
	product, err := Utils.GetProduct(userValue)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//check if cart already present
	_, osErr := os.Stat(filePath)
	if osErr == nil {
		// Already Existing User-cart
		cart, cErr := Utils.GetUserCart(filePath)
		if cErr != nil {
			fmt.Println(cErr)
			return cErr
		}
		//append product to cart
		cart = append(cart, product)
		//write in file
		marshalledCartObj, mErr := ffjson.Marshal(cart)
		if mErr != nil {
			fmt.Println("Marshal Error : ", mErr)
			return mErr
		}
		wErr := ioutil.WriteFile(filePath, marshalledCartObj, 0644)
		if wErr != nil {
			fmt.Println("Write Error : ", wErr)
			return wErr
		}
		fmt.Println(userValue, " Added To Cart")
	}
	if os.IsNotExist(osErr) {
		//user-cart not present
		_, osErr := os.Stat("FileDB/Users/" + userName)
		if os.IsNotExist(osErr) {
			//if new user
			derr := os.Mkdir("FileDB/Users/"+userName, 0755)
			if derr != nil {
				fmt.Println(derr)
				return derr
			}
		}
		//New User...Prepare Object
		cart := []Models.Cart{}
		cart = append(cart, product)
		//write in file
		marshalledCartObj, mErr := ffjson.Marshal(cart)
		if mErr != nil {
			fmt.Println("Marshal Error : ", mErr)
			return mErr
		}
		wErr := ioutil.WriteFile(filePath, marshalledCartObj, 0644)
		if wErr != nil {
			fmt.Println("Write Error : ", wErr)
			return wErr
		}
		fmt.Println(userValue, " Added To Cart")
	}
	return nil
}
func RemoveProduct(userValue, userName string) error {
	if userValue == "" || userName == "" {
		return errors.New("Data Not Found")
	}
	filePath := "FileDB/Users/" + userName + "/Cart.json"
	isPresent := false
	//add obj in cart
	_, osErr := os.Stat(filePath)
	if osErr == nil {
		// Already Existing User
		cart, cErr := Utils.GetUserCart(filePath)
		if cErr != nil {
			fmt.Println(cErr)
			return cErr
		}
		for index, cartObj := range cart {
			if cartObj.Product.ProductName == userValue {
				isPresent = true
				cart = append(cart[:index], cart[index+1:]...)
			}
		}
		if !isPresent {
			fmt.Println("This Product was not present in cart")
			return nil
		}
		//write in file
		marshalledCartObj, mErr := ffjson.Marshal(cart)
		if mErr != nil {
			fmt.Println("Marshal Error : ", mErr)
			return mErr
		}
		wErr := ioutil.WriteFile(filePath, marshalledCartObj, 0644)
		if wErr != nil {
			fmt.Println("Write Error : ", wErr)
			return wErr
		}
		fmt.Println(userValue, " Removed from Cart")
	}
	if os.IsNotExist(osErr) {
		fmt.Println("No Items In Cart")
	}
	return nil
}

func GenarateBill(userName string) error {
	cartFilePath := "FileDB/Users/" + userName + "/Cart.json"
	billFilePath := "FileDB/Users/" + userName + "/Bill.json"

	billObj := Models.Bill{}

	_, osErr := os.Stat(cartFilePath)
	if osErr == nil {
		// Already Existing User
		cart, cErr := Utils.GetUserCart(cartFilePath)
		if cErr != nil {
			fmt.Println(cErr)
			return cErr
		}
		for _, cartItem := range cart {
			billObj.Product = append(billObj.Product, cartItem.Product)
			billObj.TotalAmount += cartItem.Product.Price
		}
		if billObj.TotalAmount > 10000 {
			billObj.IsDiscountApplied = true
			billObj.AmountWithDiscount = billObj.TotalAmount - 500
		}
		//write in file
		marshalledBillObj, mErr := ffjson.Marshal(billObj)
		if mErr != nil {
			fmt.Println("Marshal Error : ", mErr)
			return mErr
		}
		wErr := ioutil.WriteFile(billFilePath, marshalledBillObj, 0644)
		if wErr != nil {
			fmt.Println("Write Error : ", wErr)
			return wErr
		}
		fmt.Println("Bill Generated!!")
		os.Remove(cartFilePath)

	}
	if os.IsNotExist(osErr) {
		fmt.Println("No Items in the Cart")
	}
	return nil
}
