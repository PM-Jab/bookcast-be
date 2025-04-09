package servers

import (
	"bookcast/configs"
	"log"

	"bookcast/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	App *gin.Engine
	Cfg *configs.Configs
	Db  *pgx.Conn
}

func NewServer(cfg *configs.Configs, db *pgx.Conn) *Server {
	return &Server{
		App: gin.Default(),
		Cfg: cfg,
		Db:  db,
	}
}

func (s *Server) Start() {
	s.MapHandlers()

	ginConnURL, err := utils.ConnectionURLBuilder("gin", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Run(ginConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
