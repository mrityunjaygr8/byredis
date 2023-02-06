package cobra

import (
	"github.com/mrityunjaygr8/byredis/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [KEY]",
	Short: "Get data of key from the redis server",
	Run: func(cmd *cobra.Command, args []string) {
		log := utils.GetLogger()
		log.Println("get")
	},
}
