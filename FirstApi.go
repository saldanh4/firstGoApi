package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Função que usa o Gin Context para a requisição HTTP da API
// Optei por definir uma variável para pegar dia e hora e outra para uma formatação mais amigável
func getTime(c *gin.Context) {
	h := time.Now()
	fmt.Println(h)
	fomattedDate := h.Format("02/Jan/2006 15:04:05")
	c.IndentedJSON(http.StatusOK, fomattedDate)
}

// Na função main defini o router do gin
// Após isso a definição do GET sem especificar um caminho de rota
// e a porta 8080 para o localhost
func main() {
	router := gin.Default()
	router.GET("/", getTime)

	router.Run("localhost:8080")
}
