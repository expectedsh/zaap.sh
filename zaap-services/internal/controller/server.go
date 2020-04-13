package controller

import (
	"context"
	"database/sql"
	"github.com/streadway/amqp"
	"net/http"

	_ "github.com/lib/pq"
)

type Config struct {
	PostgresURL string `envconfig:"POSTGRES_URL" default:"postgres://zaap:zaap@localhost/zaap"`
	RabbitURL   string `envconfig:"RABBITMQ_URL" default:"amqp://localhost"`
	Addr        string `envconfig:"ADDR" default:"localhost:9090"`
}

type Server struct {
	config             *Config
	amqpConnection     *amqp.Connection
	postgresConnection *sql.DB
	httpServer         *http.Server
}

func New(config Config) *Server {
	return &Server{
		config: &config,
	}
}

func (s *Server) Start() error {
	amqpConnection, err := amqp.Dial(s.config.RabbitURL)
	if err != nil {
		return err
	}
	defer amqpConnection.Close()
	s.amqpConnection = amqpConnection

	postgresConnection, err := sql.Open("postgres", s.config.PostgresURL)
	if err != nil {
		return err
	}
	defer postgresConnection.Close()
	s.postgresConnection = postgresConnection

	s.httpServer = &http.Server{
		Addr:    s.config.Addr,
		Handler: s.HttpHandler(),
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
