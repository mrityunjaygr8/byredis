package cobra

import (
	"fmt"

	"github.com/mrityunjaygr8/byredis/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [KEY] [DATA]",
	Short: "Set data for key on the redis server",
	Run: func(cmd *cobra.Command, args []string) {
		// log := utils.GetLogger()
		encoded := utils.EncodeCommand(append([]string{utils.SET}, args...))

		resp := utils.SendCommand(encoded)

		fmt.Println(resp)
	},
}
