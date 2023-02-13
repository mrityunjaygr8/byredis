package cobra

import (
	"fmt"

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
		encoded := utils.EncodeCommand(append([]string{utils.GET}, args...))
		resp := utils.SendCommand(encoded)

		fmt.Println(resp)
	},
}
