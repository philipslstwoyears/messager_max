package apiServer

import (
	"awesomeProject2/configs"
	"awesomeProject2/repository"
	"awesomeProject2/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ApiServer struct {
	config  *configs.Config
	logger  *logrus.Logger
	storage *repository.Storage
	user    *UserHandler
}

func New(config *configs.Config, storage *repository.Storage) (*ApiServer, error) {
	userService := service.NewUserService(storage)
	userHandler := NewUserHandler(userService)
	apiServer := &ApiServer{
		config:  config,
		logger:  logrus.New(),
		storage: storage,
		user:    userHandler,
	}
	err := apiServer.ConfigLogger()
	if err != nil {
		return nil, err
	}
	return apiServer, nil
}

func (s *ApiServer) Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/AddMessage", LogMidellware(RecoverMiddleware(AuthMiddleware(s.AddMessageHandler)))).Methods("POST")
	r.HandleFunc("/DeleteMessage", LogMidellware(RecoverMiddleware(AuthMiddleware(s.DeleteMessageHandler)))).Methods("DELETE")
	r.HandleFunc("/GetMessage", LogMidellware(RecoverMiddleware(s.GetMessageHandler))).Methods("GET")
	r.HandleFunc("/GetAllRecieaved", LogMidellware(RecoverMiddleware(AuthMiddleware(s.GetAllHandlerRecieaved)))).Methods("GET")
	r.HandleFunc("/GetAllSend", LogMidellware(RecoverMiddleware(AuthMiddleware(s.GetAllHandlerSend)))).Methods("GET")
	r.HandleFunc("/register", LogMidellware(RecoverMiddleware(s.user.RegisterUser))).Methods("POST")
	r.HandleFunc("/login", LogMidellware(RecoverMiddleware(s.user.LoginUser))).Methods("POST")
	s.logger.Info("Starting API Server")

	return http.ListenAndServe(s.config.BindAddr, r)
}

func (s *ApiServer) ConfigLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
