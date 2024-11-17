package model

type MonitorResponse struct {
	Tx_status string `json:"tx_status" binding:"required"`
}
