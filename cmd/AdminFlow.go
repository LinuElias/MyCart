package cmd

import (
	"MyCart/cmd/Services"

	"github.com/spf13/cobra"
)

var addProductsToCatlogCmd = &cobra.Command{
	Use:   "addProductsToCatlog",
	Short: "You can view all the categories for shopping",
	Long:  `You can view all the categories for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("viewCategories called")
		return Services.AddProductsToCatlog(category, subcategory, productname, specification, price)

	},
}
var viewCartDetailsCmd = &cobra.Command{
	Use:   "viewCartDetails",
	Short: "You can view all the categories for shopping",
	Long:  `You can view all the categories for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("viewCategories called")
		return Services.ViewCartDetails(userName)

	},
}
var viewBillDetailsCmd = &cobra.Command{
	Use:   "viewBillDetails",
	Short: "You can view all the categories for shopping",
	Long:  `You can view all the categories for shopping`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("viewCategories called")
		return Services.ViewBillDetails(userName)

	},
}

func init() {
	rootCmd.AddCommand(addProductsToCatlogCmd)
	rootCmd.AddCommand(viewCartDetailsCmd)
	rootCmd.AddCommand(viewBillDetailsCmd)
}
