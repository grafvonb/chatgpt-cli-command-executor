package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AppConfig struct {
	CfgFile string `mapstructure:"CCCE_CONFIG_FILE"`
	ApiKey  string `mapstructure:"CCCE_CHATGPT_API_KEY"`
}

var config AppConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccce",
	Short: "ChatGPT CLI Command Executor (CCCE)",
	Long:  `ChatGPT CLI Command Executor (CCCE)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags, global for the application
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is $HOME/.ccce.yaml)")
	rootCmd.PersistentFlags().StringVar(&config.ApiKey, "api-key", "", "your ChatGPT API key")

	// Local flags, run only when this action is called directly
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Bind flags to Viper
	_ = viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Bind ENVS to Viper
	_ = viper.BindEnv("api-key", "CCCE_CHATGPT_API_KEY")
	viper.AutomaticEnv() // read in environment variables that match

	if config.CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config.CfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ccce" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ccce")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			// Config file was found but another error occurred
			log.Fatalf("Fatal error reading config file: %s \n", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Fatal error unmarshalling config: %s \n", err)
	}

	fmt.Printf("%+v\n", viper.AllSettings())
	return
}
