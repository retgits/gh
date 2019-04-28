// Package cmd defines and implements command-line commands and flags
// used by gh. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "Git Helper",
	Long: `
A collection of git helper commands to make my life a little easier`,
}

// The constants
const (
	// The name of the config file
	ConfigName = ".ghconfig"
)

var (
	// Version number of gh
	Version = "N/A"
	// BuildTime of gh
	BuildTime = "N/A"
	cfgFile   string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Specify the configuration file and init method
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ghconfig.yml)")

	// Set the version and template function to render the version text
	rootCmd.Version = fmt.Sprintf("%s built on %s", Version, BuildTime)
	rootCmd.SetVersionTemplate("\nYou're running gh version {{ .Version }}\n\n")
}

func initConfig() {
	// Read the configuration file
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Printf("fatal locating $HOME directory: %s\n", err)
			os.Exit(1)
		}

		// Search .ghconfig.yml in home directory
		viper.AddConfigPath(home)
		viper.SetConfigName(ConfigName)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file: %s\nrelying on flags for configuration\n\n", err)
	}
}
