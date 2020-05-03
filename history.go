package bmlgo

// HistoryResponse ...
type HistoryResponse struct {
	Success bool            `json:"success"`
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Payload *HistoryPayload `json:"payload"`
}

// HistoryPayload ...
type HistoryPayload struct {
	CurrentPage int            `json:"currentPage"`
	TotalPages  int            `json:"totalPages"`
	History     []*HistoryItem `json:"history"`
}

// HistoryItem ...
type HistoryItem struct {
	ID          string  `json:"id"`
	Reference   string  `json:"reference"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Minus       bool    `json:"minus"`
	BookingDate string  `json:"bookingDate"`
	ValueDate   string  `json:"valueDate"`
	Narrative1  string  `json:"narrative1"`
	Narrative2  string  `json:"narrative2"`
	Narrative3  string  `json:"narrative3"`
	Narrative4  string  `json:"narrative4"`
	Description string  `json:"description"`
}
