package handler

import (
	"database/sql"
	"jarvishttp/config"
	"log"
)

type HTTPHandler struct {
	DB     *sql.DB
	Logger *log.Logger
	Config *config.AppConfig
}
