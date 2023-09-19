package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// correctCmd represents the correct command
var correctCmd = &cobra.Command{
	Use:   "correct",
	Short: "Corrects a given text",
	Long:  `Corrects a given text`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("correct called")
		fmt.Println(viper.GetString("api-key"))
		fmt.Println(config.ApiKey)
		fmt.Printf("%+v\n", viper.AllSettings())
	},
}

func init() {
	rootCmd.AddCommand(correctCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// correctCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// correctCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
