package internal

import (
	"VentureCookie1/internal/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TRANSFERUSER struct {
	USERID  string `json:"userid"`
	VISITED string `json:"visited"`
}

func PostUser(c *gin.Context) {
	var user TRANSFERUSER
	c.BindJSON(&user)
	id := primitive.NewObjectID()
	c.JSON(200, id.Hex())
	mongodb.Create(user.VISITED, id)
}

func UpdateUser(c *gin.Context) {
	var user TRANSFERUSER
	c.BindJSON(&user)
	c.Writer.WriteHeader(200)
	mongodb.AddVisited(user.USERID, user.VISITED)
}
