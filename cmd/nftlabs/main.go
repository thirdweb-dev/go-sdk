package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	privateKey string
	chainRpcUrl string

	rootCmd = &cobra.Command{
		Use:   "nftlabs",
		Short: "A CLI for the nftlabs-protocols go SDK",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&privateKey, "privateKey", "k", "", "private key used to sign transactions")
	rootCmd.PersistentFlags().StringVarP(&chainRpcUrl, "chainRpcUrl", "u", "https://rpc-mumbai.maticvigil.com", "chain url where all rpc requests will be sent")
	_ = viper.BindPFlag("privateKey", rootCmd.PersistentFlags().Lookup("privateKey"))
	_ = viper.BindPFlag("chainRpcUrl", rootCmd.PersistentFlags().Lookup("chainRpcUrl"))
	viper.SetDefault("chainRpcUrl", "https://rpc-mumbai.maticvigil.com")

	rootCmd.AddCommand(nftCmd)
}

func initConfig() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
Execute()
}