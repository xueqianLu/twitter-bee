package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xueqianLu/twitter-bee/client"
)

var (
	beeAddr   string
	queryUser string
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query with client",
	Run: func(cmd *cobra.Command, args []string) {
		cli := client.NewBeeClient(beeAddr)
		res, err := cli.GetFollowerCount(queryUser)
		if err != nil {
			log.WithError(err).Error("get follower count failed")
		} else {
			log.WithFields(log.Fields{
				"user":  queryUser,
				"count": res.Count,
			}).Info("query follower count")
		}

		resList, err := cli.GetFollowerList(queryUser, "")
		if err != nil {
			log.WithError(err).Error("get follower list failed")
		} else {
			log.WithFields(log.Fields{
				"user": queryUser,
				"list": resList.List,
				"next": resList.Next,
			}).Info("query latest follower list")
		}
	},
}

func init() {
	queryCmd.PersistentFlags().StringVar(&beeAddr, "bee", "127.0.0.1:8088", "bee service url")
	queryCmd.PersistentFlags().StringVar(&queryUser, "user", "", "user name to query")
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
