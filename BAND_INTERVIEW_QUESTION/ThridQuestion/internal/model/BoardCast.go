package model

type BoardCastRequest struct {
	Symbol string `json:"symbol" binding:"required"`
	Price  uint64 `json:"price" binding:"required"`
}

type BoardCastResponse struct {
	Tx_hash string `json:"tx_hash"`
}
