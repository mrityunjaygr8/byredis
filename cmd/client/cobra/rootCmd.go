package cobra

import (
	"os"

	"github.com/mrityunjaygr8/byredis/utils"
	"github.com/spf13/cobra"
)

var log = utils.GetLogger()
var rootCmd = &cobra.Command{
	Use:   "byredis",
	Short: "A short redis client written in Golang",
	Long:  "A short redis client written in Golang",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("yo")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorln(err)
		os.Exit(1)
	}
}
