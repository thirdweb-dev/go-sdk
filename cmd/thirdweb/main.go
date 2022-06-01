package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	privateKey  string
	chainRpcUrl string

	rootCmd = &cobra.Command{
		Use:   "thirdweb",
		Short: "A CLI for the thirdweb go SDK",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&privateKey, "privateKey", "k", "", "private key used to sign transactions")
	rootCmd.PersistentFlags().StringVarP(&chainRpcUrl, "chainRpcUrl", "u", "mumbai", "chain url where all rpc requests will be sent")
	_ = viper.BindPFlag("privateKey", rootCmd.PersistentFlags().Lookup("privateKey"))
	_ = viper.BindPFlag("chainRpcUrl", rootCmd.PersistentFlags().Lookup("chainRpcUrl"))
	viper.SetDefault("chainRpcUrl", "mumbai")

	rootCmd.AddCommand(nftCmd)
	rootCmd.AddCommand(editionCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.AddCommand(nftDropCmd)
	rootCmd.AddCommand(editionDropCmd)
	rootCmd.AddCommand(multiwrapCmd)
	rootCmd.AddCommand(storageCmd)
	rootCmd.AddCommand(customCmd)
}

func initConfig() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	if err := Execute(); err != nil {
		panic(err)
	}
}
