package helloworld

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type HelloEndpoint struct {
}

func InitHelloEndpoint() (*HelloEndpoint, error) {
	return &HelloEndpoint{}, nil
}

func (he *HelloEndpoint) RegisterEndpoint(router *gin.Engine) {

	turingRegex := regexp.MustCompile(`(?i)^(turing|alan[ _-]?turing)$`)
	router.GET("/:path", func(c *gin.Context) {
		path := c.Param("path")

		// Nettoyage du chemin (supprime les éventuels slashs et espaces superflus)
		path = strings.Trim(path, "/")

		if turingRegex.MatchString(path) {
			he.HelloWorld(c)
			return
		}

		// Si ce n'est pas une route valide, Gin retournera une 404
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.GET("/reponse", he.NotResponse)
}

func (he *HelloEndpoint) HelloWorld(c *gin.Context) {
	c.JSON(200, "Félicitations, vous avez terminé l'initiation ! J'espère que le jeu de piste vous a plu. Le code pour la suite est l'année de naissance de Alan Turing")
}

func (he *HelloEndpoint) NotResponse(c *gin.Context) {
	c.JSON(200, "Il faut changer le mot reponse par le bon nom")
}
