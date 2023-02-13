package main

import (
	cobra "github.com/mrityunjaygr8/byredis/cmd/client/cobra"

	"github.com/mrityunjaygr8/byredis/utils"
)

func main() {
	utils.SetupLogger(utils.SERVER_HOST, utils.SERVER_PORT)

	cobra.Execute()
}
