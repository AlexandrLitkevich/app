package apiserver

import (
	"io"
	"net/http"

	"github.com/AlexandrLitkevich/app/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer API Server ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	//  добавляем роутер
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	// Initial APIServer
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	//  Если при запуске ошибка то выходим
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//Конфигурираем логгер

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

//Конфигурираем router

func (s *APIServer) configureRouter() {
	// TODO Как сделать несколько роутов?
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	// Open database
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// тут можно определить какие либо переменные
	return func(w http.ResponseWriter, r *http.Request) {
		// Что за аргументы  w
		io.WriteString(w, "Hello Alex")
	}
}
