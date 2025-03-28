package root

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/xueqianLu/twitter-bee/config"
	"github.com/xueqianLu/twitter-bee/node"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "tbee [command]",
		Short: "Tbee is a simulator to act as a twitter user.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			bindFlags(cmd)
			return config.InitConfig(cfgFile)
		},
		Run: runCommand,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().String("server", config.DefaultConfig.ServiceUrl, "service address")
	rootCmd.PersistentFlags().String("log-level", config.DefaultConfig.LogLevel, "log level, default is info")
	rootCmd.PersistentFlags().String("proxy", config.DefaultConfig.Proxy, "http proxy address")
}

func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, ".") {
			envVar := strings.ToUpper(strings.ReplaceAll(f.Name, ".", "_"))
			viper.BindEnv(f.Name, envVar)
		}

		viper.BindPFlag(f.Name, f)
	})
}

func runCommand(cmd *cobra.Command, _ []string) {
	cfg := config.Global
	setlog(cfg.LogLevel)
	log.WithField("config", cfg).Info("load config success")

	server, err := node.NewNode(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		fmt.Printf("Received signal: %v\n", sig)
	}
	//server.Stop()

	log.Info("bee exit, bye bye !!!")
	return
}

func setlog(level string) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}
