/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"go-bank-express/internal/broker"
	"go-bank-express/pkg/mq/reader"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var createUserCmd = &cobra.Command{
	Use:           "user",
	Short:         "Start User Consummer",
	SilenceUsage:  true,
	SilenceErrors: true,
	Run:           runCreateUserCmd,
}

func init() {
	rootCmd.AddCommand(createUserCmd)
}

func runCreateUserCmd(cmd *cobra.Command, args []string) {
	createUserReader := reader.NewReader("create-user-group", "create-user")
	defer createUserReader.Close()
	broker.NewUserConsumer(createUserReader).CreateUser(context.Background())
}
