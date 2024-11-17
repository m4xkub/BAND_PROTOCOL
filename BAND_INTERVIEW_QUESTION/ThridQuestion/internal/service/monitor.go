package service

import (
	"ThridQuestion/internal/config"
	"ThridQuestion/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Monitor(transactionStatus *TransactionStatus, isChange *bool) {
	var monitorResponse model.MonitorResponse

	// check current state of transaction
	// if status != pending, it means that this transaction is in finalization state.
	// so dont send request to server to decrease server's traffic

	previousStatus := transactionStatus.Status
	if previousStatus != config.STATUSPENDING {
		return
	}

	res, err := http.Get(config.EndPointUrl + "/check/" + transactionStatus.ID)
	if err != nil {
		fmt.Printf("Error fetching status for tx %s: %v\n", transactionStatus.ID, err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Server is currently not available\n")
		fmt.Printf("unexpected response status: %d", res.StatusCode)
		return
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response for tx %s: %v\n", transactionStatus.ID, err)
		return
	}
	bodyString := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyString), &monitorResponse)

	if monitorResponse.Tx_status == config.STATUSPENDING {
		*isChange = false
	}

	if err != nil {
		fmt.Printf("Error unmarshal response for tx %s: %v\n", transactionStatus.ID, err)
		return
	}

	if monitorResponse.Tx_status != previousStatus || monitorResponse.Tx_status == config.STATUSPENDING {
		fmt.Printf("Status of transaction id : %s is %s\n", transactionStatus.ID, monitorResponse.Tx_status)
		transactionStatus.Status = monitorResponse.Tx_status
	}

}
