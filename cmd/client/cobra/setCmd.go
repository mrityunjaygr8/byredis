package cobra

import (
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
		log := utils.GetLogger()
		log.Println("get")
		log.Println(args)
		encoeded := utils.EncodeCommand(append([]string{"get"}, args...))
		log.Println(encoeded)
		log.Println(string(encoeded))

		words := utils.DecodeCommand(encoeded)
		log.Println(words)
	},
}
