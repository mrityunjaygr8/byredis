package cobra

import (
	"fmt"

	"github.com/mrityunjaygr8/byredis/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use:   "del [KEY]",
	Short: "Delete key from the redis server",
	Run: func(cmd *cobra.Command, args []string) {
		// log := utils.GetLogger()
		encoded := utils.EncodeCommand(append([]string{utils.DEL}, args...))
		resp := utils.SendCommand(encoded)

		fmt.Println(resp)
	},
}
