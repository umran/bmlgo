package bmlgo

import "net/url"

// TransferRequestResponse ...
type TransferRequestResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// TransferCompletionResponse ...
type TransferCompletionResponse struct {
	Success bool                       `json:"success"`
	Code    int                        `json:"code"`
	Message string                     `json:"message"`
	Payload *TransferCompletionPayload `json:"payload"`
}

// TransferCompletionPayload ...
type TransferCompletionPayload struct {
	Reference string `json:"reference"`
	Timestamp string `json:"timestamp"`
}

// Transfer ...
type Transfer struct {
	request url.Values
}

// GetRequestForm ...
func (t *Transfer) GetRequestForm() url.Values {
	return t.request
}

// GenerateCompletionForm ...
func (t *Transfer) GenerateCompletionForm(otp string) url.Values {
	form := url.Values{}

	form["currency"] = t.request["currency"]
	form["transfertype"] = t.request["transfertype"]
	form["debitAmount"] = t.request["debitAmount"]
	form.Add("narrative2", "")
	form.Add("m_ref", "")
	form.Add("additionalInstruction", "")
	form.Add("remarks", "")
	form["channel"] = t.request["channel"]
	form["debitAccount"] = t.request["debitAccount"]
	form["creditAccount"] = t.request["creditAccount"]
	form.Add("otp", otp)

	return form
}

// NewTransfer ...
func NewTransfer(debitAmount, debitAccount, creditAccount string) *Transfer {
	request := url.Values{}
	request.Add("transfertype", "IAT")
	request.Add("channel", "email")
	request.Add("debitAmount", debitAmount)
	request.Add("currency", "MVR")
	request.Add("creditAccount", creditAccount)
	request.Add("debitAccount", debitAccount)

	return &Transfer{
		request: request,
	}
}
