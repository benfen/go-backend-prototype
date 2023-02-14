package main

import (
	"log"

	"github.com/benfen/go-prototype-backend/api"
	"github.com/benfen/go-prototype-backend/db"
	"github.com/benfen/go-prototype-backend/web"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *db.DB
	engine *gin.Engine
}

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()

	api.CreateAPI(engine, db)
	web.AttachWeb(engine)

	engine.Run("localhost:4000")
}
