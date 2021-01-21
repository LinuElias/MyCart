# MyCart

MyCart is a CLI (Command Line Interface) based E-commerce app written in Go Programming Language.

## Requirements
* Go
* Cobra Library

## Installation

Clone the Project

```bash
git clone url
```


## Usage

```bash
#User Commands

#To view the categories in the application
myCart viewCategories

#To view the sub-categories in the application
#category_name is a category name
myCart viewSubCategories -v category_name

#To view the Products in the application
myCart viewProducts -v category_name

#To view the Product Details of a particular Product
#product_name is a category name
myCart viewProductDetails -v product_name

#To Add Product to Cart
#product_name is a product name
#user_name is a user name
myCart addProduct -v product_name -n user_name

#To Remove Product from Cart
#user_name is a user name
myCart removeProduct -v product_name -n user_name

#To Generate Bill 
#user_name is a user name
myCart genearateBill -n user_name

#Admin Commands
#To Add Products To Catlog
myCart addProductsToCatlog -c category_name -s subCategory_name -p productName -d specification -r price

#To View Cart of a User
myCart viewCartDetails -n user_name

# To View Bill of a User
myCart viewBillDetails -n user_name

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
