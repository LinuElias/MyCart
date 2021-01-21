package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var userValue, userName, category, subcategory, productname, specification string
var price float64

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myCart",
	Short: "A brief description of your application",
	Long:  `Root`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//User
	rootCmd.PersistentFlags().StringVarP(&userValue, "Value", "v", "", "This is UserInput")
	rootCmd.PersistentFlags().StringVarP(&userName, "name", "n", "", "This is User Name")

	//Admin
	rootCmd.PersistentFlags().StringVarP(&category, "category", "c", "", "This is User Name")
	rootCmd.PersistentFlags().StringVarP(&subcategory, "subcategory", "s", "", "This is User Name")
	rootCmd.PersistentFlags().StringVarP(&productname, "productname", "p", "", "This is User Name")
	rootCmd.PersistentFlags().StringVarP(&specification, "specification", "d", "", "This is User Name")
	rootCmd.PersistentFlags().Float64VarP(&price, "price", "r", 0, "This is User Name")
}
