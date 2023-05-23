package main

import (
	pgdatabase "github.com/Rajendro1/AccuKnox/pgDatabase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	pgdatabase.Connect(&gin.Context{})
	// router.HandleRequest()
}
