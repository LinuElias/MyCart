package cmd

import (
	"MyCart/cmd/Services"

	"github.com/spf13/cobra"
)

// viewCategoriesCmd represents the viewCategories command
var viewCategoriesCmd = &cobra.Command{
	Use:   "viewCategories",
	Short: "You can view all the categories for shopping",
	Long:  `You can view all the categories for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("viewCategories called")
		return Services.ViewCategories()

	},
}

// viewSubCategoriesCmd represents the viewSubCategories command
var viewSubCategoriesCmd = &cobra.Command{
	Use:   "viewSubCategories",
	Short: "You can view all the sub-categories under a category for shopping",
	Long:  `You can view all the sub-categories under a category for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Services.ViewSubCategories(userValue)

	},
}

// viewProductsCmd represents the viewProducts command
var viewProductsCmd = &cobra.Command{
	Use:   "viewProducts",
	Short: "You can view all the Products under a sub-category for shopping",
	Long:  `You can view all the Products under a sub-category for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Services.ViewProducts(userValue)

	},
}

// viewProductDetailsCmd represents the viewProductDetails command
var viewProductDetailsCmd = &cobra.Command{
	Use:   "viewProductDetails",
	Short: "You can view all the Product Details of a particular product",
	Long:  `You can view all the Product Details of a particular product`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Services.ViewProductDetails(userValue)

	},
}

// addProductCmd represents the addProduct command
var addProductCmd = &cobra.Command{
	Use:   "addProduct",
	Short: "You can add Products to cart for shopping",
	Long:  `You can add Products to cart for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Services.AddProduct(userValue, userName)

	},
}

// removeProductCmd represents the removeProduct command
var removeProductCmd = &cobra.Command{
	Use:   "removeProduct",
	Short: "You can remove a Product from cart",
	Long:  `You can remove a Product from cart`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Services.RemoveProduct(userValue, userName)

	},
}

// generateBillCmd represents the generateBill command
var generateBillCmd = &cobra.Command{
	Use:   "generateBill",
	Short: "This Command Generats a bill ",
	Long:  `This Command Generats a bill`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Services.GenarateBill(userName)

	},
}

func init() {
	rootCmd.AddCommand(viewCategoriesCmd)
	rootCmd.AddCommand(viewSubCategoriesCmd)
	rootCmd.AddCommand(viewProductsCmd)
	rootCmd.AddCommand(viewProductDetailsCmd)
	rootCmd.AddCommand(addProductCmd)
	rootCmd.AddCommand(removeProductCmd)
	rootCmd.AddCommand(generateBillCmd)

}
