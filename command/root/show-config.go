package root

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xueqianLu/twitter-bee/config"
)

var showConfigCmd = &cobra.Command{
	Use:   "show-config",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configJSON, err := json.MarshalIndent(config.Global, "", "  ")
		if err != nil {
			log.WithField("error", err).Error("Error marshaling config")
			return
		}
		fmt.Println(string(configJSON))
	},
}

func init() {
	rootCmd.AddCommand(showConfigCmd)
}
