package cmd

import (
	"context"

	"github.com/20326/vega/app/config"
	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initdbCmd represents the init db command
var initdbCmd = &cobra.Command{
	Use:   "initdata",
	Short: "Init database",
	Long:  `Init admin,roles,settings,permissions server`,
	Run: func(cmd *cobra.Command, args []string) {
		startInitDB()
	},
}

var initDataPathFlag string

func init() {
	runCmd.PersistentFlags().StringVar(&initDataPathFlag, "initdata", "", "init data file path")
	rootCmd.AddCommand(initdbCmd)

}

func startInitDB() {
	var err error
	var log = logrus.New()

	log.Info("===> Vega init db ... <===")
	cfg, err := config.LoadConfig(configPathFlag, log)
	log.SetFormatter(&logrus.JSONFormatter{})
	if nil != err {
		log.WithError(err).Fatal("Load config has some errors!")
	}

	// init service
	srv := service.NewService(cfg, log)

	data, err := config.LoadInitData(initDataPathFlag, log)
	initData(srv, data)

}

func initData(srv *service.Service, data *config.InitData) {
	ctx := context.Background()

	permID := uint64(1)
	actionID := uint64(1)
	resID := uint64(1)

	for _, permItem := range data.Permissions {
		perm := &model.Permission{
			Name:     permItem.Name,
			Label:    permItem.Label,
			Describe: permItem.Describe,
			Icon:     permItem.Icon,
			Path:     permItem.Path,
			Status:   permItem.Status,
			Deleted:  permItem.Deleted,
		}
		perm.ID = permID

		for _, actionItem := range permItem.Actions {
			action := model.Action{
				Name:     actionItem.Name,
				Describe: actionItem.Describe,
			}
			action.ID = actionID
			for _, resItem := range actionItem.Resources {
				res := &model.Resource{
					ActionID: actionID,
					Method:   resItem.Method,
					Path:     resItem.Path,
				}
				res.ID = resID
				action.Resources = append(action.Resources, res)

				resID++
			}
			perm.Actions = append(perm.Actions, action)
			actionID++
		}
		_ = srv.Permissions.Create(ctx, perm)

		permID++
	}
}
