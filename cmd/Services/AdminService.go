package Services

import (
	"MyCart/cmd/Models"
	"MyCart/cmd/Utils"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pquerna/ffjson/ffjson"
)

func AddProductsToCatlog(category, subcategory, productname, specification string, price float64) error {
	isCategoryPresent, isSubCategoryPresent := false, false
	product := Models.Products{}
	product.ProductName = productname
	product.Price = price
	product.Specification = specification
	catlog, err := Utils.GetCatlog()
	if err != nil {
		fmt.Println("Error while reading catlog")
		return err
	}
	for i := 0; i < len(catlog); i++ {
		if catlog[i].CategoryName == category {
			isCategoryPresent = true
			for _, subCategry := range catlog[i].SubCategory {
				if subCategry.SubCategoryName == subcategory {
					isSubCategoryPresent = true
					subCategry.Products = append(subCategry.Products, product)
				}
			}
		}
	}
	if isCategoryPresent && !isSubCategoryPresent {
		//if category presnt but not sub category
		for i := 0; i < len(catlog); i++ {
			if catlog[i].CategoryName == category {
				subCategry := Models.SubCategory{}
				subCategry.SubCategoryName = subcategory
				subCategry.Products = append(subCategry.Products, product)
				catlog[i].SubCategory = append(catlog[i].SubCategory, subCategry)
			}
		}
	}
	if !isCategoryPresent {
		//if category is not presnt
		catlogObj := Models.Catlog{}
		catlogObj.CategoryName = category

		subCategry := Models.SubCategory{}
		subCategry.SubCategoryName = subcategory
		subCategry.Products = append(subCategry.Products, product)
		catlogObj.SubCategory = append(catlogObj.SubCategory, subCategry)
		catlog = append(catlog, catlogObj)
	}
	//write in file
	marshalledCatlogObj, mErr := ffjson.Marshal(catlog)
	if mErr != nil {
		fmt.Println("Marshal Error : ", mErr)
		return mErr
	}
	wErr := ioutil.WriteFile("FileDB/catlog.json", marshalledCatlogObj, 0644)
	if wErr != nil {
		fmt.Println("Write Error : ", wErr)
		return wErr
	}
	fmt.Println("Product Added to Catlog")
	return nil
}

func ViewCartDetails(username string) error {
	filePath := "FileDB/Users/" + username + "/Cart.json"
	_, osErr := os.Stat(filePath)
	if osErr == nil {
		// Already Existing User
		cart, cErr := Utils.GetUserCart(filePath)
		if cErr != nil {
			fmt.Println(cErr)
			return cErr
		}
		if len(cart) == 0 {
			fmt.Println("Cart is Empty!")
			return nil
		}
		for index, cartObj := range cart {
			fmt.Println("Product", index+1, ":")
			fmt.Println("Product Name:", cartObj.Product.ProductName)
			fmt.Println("Price:", cartObj.Product.Price)
			fmt.Println("Specification:", cartObj.Product.Specification)
			fmt.Println("################################")
		}

	}
	if os.IsNotExist(osErr) {
		fmt.Println("No Items In Cart")
	}
	return nil
}
func ViewBillDetails(username string) error {
	filePath := "FileDB/Users/" + username + "/Bill.json"
	_, osErr := os.Stat(filePath)
	if osErr == nil {
		// Already Existing User
		billObj, cErr := Utils.GetUserBill(filePath)
		if cErr != nil {
			fmt.Println(cErr)
			return cErr
		}
		for index, product := range billObj.Product {
			fmt.Println("Product", index+1, ":")
			fmt.Println("Product Name:", product.ProductName)
			fmt.Println("Price:", product.Price)
			fmt.Println("Specification:", product.Specification)
			fmt.Println("################################")
		}
		fmt.Println("Total Amount:", billObj.TotalAmount)
		fmt.Println("Is Discount Applied:", billObj.IsDiscountApplied)
		fmt.Println("Discounted Amount :", billObj.AmountWithDiscount)

	}
	if os.IsNotExist(osErr) {
		fmt.Println("Bill not Generated!")
	}
	return nil
}
