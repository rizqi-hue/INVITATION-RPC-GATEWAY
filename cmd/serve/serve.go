/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package serve

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/INVITATION-RPC-GATEWAY/api"
	def "github.com/INVITATION-RPC-GATEWAY/cmd/config"
	"github.com/INVITATION-RPC-GATEWAY/pkg/config"
	"github.com/INVITATION-RPC-GATEWAY/pkg/logger"
)

// serveCmd represents the serve command
var (
	Config string

	ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "running gateway",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
)

func startServer() {
	if err := config.Load(def.DefaultConfig, Config); err != nil {
		log.Fatal(err)
	}

	logger.Configure()

	api.NewFiberServer()
}
