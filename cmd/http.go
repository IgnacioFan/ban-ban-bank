/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-bank-express/internal/app/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var port int

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:           "http",
	Short:         "Start Http Server",
	SilenceUsage:  true,
	SilenceErrors: true,
	Run:           runHttpCmd,
}

func init() {
	rootCmd.AddCommand(httpCmd)

	httpCmd.Flags().IntVarP(&port, "port", "", 3005, "set port number")
}

func runHttpCmd(cmd *cobra.Command, args []string) {
	server, _ := http.Initialize()
	if err := server.Start(port); err != nil {
		logrus.Panicf("server.Run failed, err: %v", err)
	}
}
