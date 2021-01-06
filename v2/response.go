package bmlgo

// Response ...
type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// AuthPayload ...
type AuthPayload int

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

// ActivityPayload ...
type ActivityPayload struct {
	Content *ActivityContent `json:"content"`
}

// ActivityContent ...
type ActivityContent struct {
	Data         []*ActivityData `json:"data"`
	From         int             `json:"from"`
	To           int             `json:"to"`
	PerPage      int             `json:"per_page"`
	Total        int             `json:"total"`
	CurrentPage  int             `json:"current_page"`
	LastPage     int             `json:"last_page"`
	FirstPageURL string          `json:"first_page_url"`
	LastPageURL  string          `json:"last_page_url"`
	NextPageURL  string          `json:"next_page_url"`
	PrevPageURL  string          `json:"prev_page_url"`
	Path         string          `json:"path"`
}

// ActivityData ...
type ActivityData struct {
	Content         []*ActivityDataContent
	CreditName      string `json:"creditName"`
	DateTime        string `json:"dateTime"`
	FormattedAmount string `json:"formattedAmount"`
	Message         string `json:"message"`
	Status          string `json:"status"`
	Type            string `json:"type"`
}

// ActivityDataContent ...
type ActivityDataContent struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// TransferCompletionPayload ...
type TransferCompletionPayload struct {
	Reference string `json:"reference"`
	Timestamp string `json:"timestamp"`
}
