package app

import (
	"net/http"

	"github.com/20326/vega/app/config"
	// "github.com/20326/vega/app/model"
	"github.com/20326/vega/app/handler"
	"github.com/20326/vega/app/middleware"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/graceful"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

//var (
//	AppConfig *config.Config
//)
//
//// application is the main struct for the Drone server.
//type Application struct {
//	Engine     *gin.Engine
//	Config     *config.Config
//	Server     *http.Server
//	//Actions    model.ActionService
//	//Admissions model.AdmissionService
//	//Resources  model.ResourceService
//	//Perms      model.PermService
//	//Roles      model.RoleService
//	//Users      model.UserService
//}
//
//// newApplication creates a new application struct.
//func newApplication(
//	engine *gin.Engine,
//	config *config.Config,
//	server *http.Server,
//	//actions model.ActionService,
//	//admissions model.AdmissionService,
//	//resources model.ResourceService,
//	//perms model.PermService,
//	//roles model.RoleService,
//	//users model.UserService,
//) Application {
//	return Application{
//		// Engine:     engine,
//		Config:     config,
//		Server:     server,
//		//Actions:    actions,
//		//Admissions: admissions,
//		//Resources:  resources,
//		//Perms:      perms,
//		//Roles:      roles,
//		//Users:      users,
//	}
//}

func StartHttpServer(configPath string, pidFile string) {
	var err error

	cfg, err := config.LoadConfig(configPath)
	if nil != err {
		log.Fatal().Err(err).Msg("Load config has some errors!")
	}

	r := gin.Default()

	// init service
	srv := service.NewService(cfg)
	r.Use(middleware.ServiceMiddleware(srv))

	// init session
	sessionStore := middleware.NewSessionsStore(cfg)
	r.Use(sessions.Sessions("session", sessionStore))

	log.Info().Msgf("Init ctx %+v", srv)
	// use midleware

	//init handler
	handler.NewHandlers(r)

	//app, err := InitializeApplication(AppConfig)
	if nil != err {
		log.Fatal().Err(err).Msg("Init application has some errors!")
	}

	log.Info().Msg("Init application ok")

	server := &http.Server{
		// Set timeouts, etc.
		Addr:    cfg.Addr,
		Handler: r,
	}
	graceful.StartGracefulServer(server, pidFile)
}
