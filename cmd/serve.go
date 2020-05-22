package cmd

import (
	"fmt"
	"os"

	"github.com/20326/vega/app"
	"github.com/spf13/cobra"
)

var (
	configPathFlag string
	pidFileFlag    string
)

var rootCmd = &cobra.Command{
	Use:   "vega",
	Short: "Vega is a very flexible and friendly scaffold",
	Long: `A Flexible web application scaffold built with gin by spf13 and friends in Go.
Complete documentation is available at http://vega.run`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var runCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Long:  `Run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		app.StartHttpServer(configPathFlag, pidFileFlag)
	},
}

var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Init database",
	Long:  `Init admin,roles,settings,permissions server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// run flags: config, config-url, pidfile
	runCmd.PersistentFlags().StringVar(&configPathFlag, "config", "", "config file path")
	runCmd.PersistentFlags().StringVar(&pidFileFlag, "pidfile", "", "PID file")
	rootCmd.AddCommand(runCmd)

	// read config and init
	initdbCmd.PersistentFlags().StringVar(&configPathFlag, "config", "", "config file path")
	rootCmd.AddCommand(initdbCmd)

	cobra.OnInitialize()
}
