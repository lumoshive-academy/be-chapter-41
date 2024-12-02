package middleware

import (
	"golang-chapter-41/implem-redis/database"
	"golang-chapter-41/implem-redis/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Cacher database.Cacher
}

func NewMiddleware(cacher database.Cacher) Middleware {
	return Middleware{
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id := c.GetHeader("ID-KEY")
		val, err := m.Cacher.Get(id)
		if err != nil {
			helper.BadResponse(c, "server error", http.StatusInternalServerError)
			return
		}

		if val == "" || val != token {
			helper.BadResponse(c, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// before request
		c.Next()

	}
}
