package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/node"
)

// define user command to get user profile.
var (
	userCmd = &cobra.Command{
		Use:   "user [username]",
		Short: "Query user profile.",
		Run:   userRun,
	}
)

func init() {
	rootCmd.AddCommand(userCmd)
}

func userRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}
	config.InitConfig(cfgFile)
	n, err := node.NewNode(config.Global)
	if err != nil {
		log.WithError(err).Error("NewNode failed")
		return
	}
	balancer := n.Balancer()
	userGetter := balancer.GetRandomUserGetter()
	for _, getter := range userGetter {
		info, err := getter.GetUserInfo(args[0])
		if err != nil {
			log.WithError(err).Error("GetUserInfo failed")
		} else {
			log.WithField("info", info).Info("GetUserInfo success")
			break
		}
	}
}
