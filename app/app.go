package app

import (
	"net/http"

	"github.com/20326/vega/app/config"
	"github.com/20326/vega/app/handler"
	"github.com/20326/vega/app/middleware"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/graceful"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
//	//Permissions      model.PermissionService
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
//	//perms model.PermissionService,
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
//		//Permissions:      perms,
//		//Roles:      roles,
//		//Users:      users,
//	}
//}
var log = logrus.New()

func StartHttpServer(configPath string, pidFile string) {
	var err error

	log.Info("===> Vega starting ... <===")
	cfg, err := config.LoadConfig(configPath, log)
	log.SetFormatter(&logrus.JSONFormatter{})
	if nil != err {
		log.WithError(err).Fatal("Load config has some errors!")
	}

	r := gin.Default()

	// init service
	srv := service.NewService(cfg, log)
	r.Use(middleware.ServiceMiddleware(srv))

	// init session
	sessionStore := middleware.NewSessionsStore(cfg)
	r.Use(sessions.Sessions("session", sessionStore))

	r.Use(middleware.LoggerWithRequestID(log))
	r.Use(middleware.Cors())
	r.Use(middleware.AdmitMiddleware())

	log.Info("init service")
	// use middleware

	//init handler
	handler.NewHandlers(r)

	//app, err := InitializeApplication(AppConfig)
	if nil != err {
		log.WithFields(logrus.Fields{
			"service": srv,
		}).Info("init service")
		log.WithError(err).Fatal("Init application has some errors!")
	}

	log.Info("Init application ok")

	server := &http.Server{
		// Set timeouts, etc.
		Addr:    cfg.Addr,
		Handler: r,
	}
	graceful.StartGracefulServer(server, pidFile)
}
