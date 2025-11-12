// Package main .
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/crispgm/kickertool-entries/pkg/kickertool"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	api := kickertool.New(os.Getenv("KICKERTOOL_ACCESS_TOKEN"))

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.tmpl", nil)
	})

	r.GET("/api/tournamentEntries", func(c *gin.Context) {
		tournamentID := c.Query("tournamentID")
		if !strings.HasPrefix(tournamentID, "tio:") {
			c.JSON(http.StatusOK, gin.H{
				"error": "invalid tournament ID",
			})
			return
		}
		tournament, err := api.GetTournament(tournamentID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}

		allEntries := make(map[string][]kickertool.Entry)
		for _, disc := range tournament.Disciplines {
			entries, err := api.GetDisciplineEntries(tournamentID, disc.ID)
			if err != nil {
				fmt.Println("loading", disc.Name, "failed")
			}
			allEntries[disc.Name] = entries
			time.Sleep(20 * time.Millisecond)
		}
		c.JSON(http.StatusOK, gin.H{
			"disciplines": allEntries,
		})
	})

	r.Run(":8080") // TODO: dotenv
}
