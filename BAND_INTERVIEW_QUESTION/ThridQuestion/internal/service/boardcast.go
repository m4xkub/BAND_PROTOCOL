package service

import (
	"ThridQuestion/internal/config"
	"ThridQuestion/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionStatus struct {
	ID     string
	Status string
}

var TransactionMemory []TransactionStatus

func BoardCast(c *gin.Context) {
	var boardCastRequest model.BoardCastRequest

	if err := c.ShouldBindJSON(&boardCastRequest); err != nil {
		fmt.Println("unable to bind data")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Binding failed... Pls check your request",
		})
		return
	}

	boardcastPayload := map[string]interface{}{
		"symbol":    boardCastRequest.Symbol,
		"price":     boardCastRequest.Price,
		"timestamp": time.Now().Unix(),
	}
	jsonData, err := json.Marshal(boardcastPayload)
	if err != nil {
		fmt.Println("unable to marshal data")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server fail to marshal data, pls try again...",
		})
		return
	}

	res, err := http.Post(config.EndPointUrl+"/broadcast", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("unable to send request to endpint")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request, pls try again..."})
		return
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var boardCastResponse model.BoardCastResponse

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("unable to read response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server is currently not available..."})
			return
		}
		bodyString := string(bodyBytes)

		err = json.Unmarshal([]byte(bodyString), &boardCastResponse)
		if err != nil {
			fmt.Println("unable to unmarshal data")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server is currently not available..."})
			return
		}
		TransactionMemory = append(TransactionMemory, TransactionStatus{ID: boardCastResponse.Tx_hash, Status: config.STATUSPENDING})

		c.JSON(http.StatusOK, boardCastResponse)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server fail to boardcast message...",
		})
	}
}
