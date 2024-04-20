/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// SubtractionCmd represents the Subtraction command
var SubtractionCmd = &cobra.Command{
    Use:   "Subtraction",
    Short: "Performs the subtraction operation for two numbers",
    Long:  "Performs the subtraction operation for two numbers.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Subtraction called")

		var a,b int 
		var err error

		if len(args) != 2 {
			fmt.Println("Error : Please provide two numbers ")
		}

		a, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error : Invalid input for the first number. ")
			return 
		}

		b, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error : Invalid input for the second number. ")
		}

		result := a/b 
		fmt.Printf("The result is %d\n", result)
	},
}

func init() {
	rootCmd.AddCommand(SubtractionCmd)
    

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SubtractionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SubtractionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
