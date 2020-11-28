package bmlgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Session ...
type Session struct {
	client *http.Client
}

// Authenticate ...
func (s *Session) Authenticate(username, password string) error {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	res, err := s.client.PostForm("https://www.bankofmaldives.com.mv/internetbanking/api/login", form)
	if err != nil {
		return err
	}

	// parse the response
	if res.StatusCode != 200 {
		return errors.New("non 200 status code received")
	}

	authResponse := &Response{}
	json.NewDecoder(res.Body).Decode(authResponse)

	if authResponse.Success == false {
		return errors.New("authentication failed")
	}

	if authResponse.Code != 0 {
		return fmt.Errorf("unexpected code: %d", authResponse.Code)
	}

	// the following is some weird gymnastics to get the server to allow subsequent requests to the history endpoint.
	_, err = s.client.Get("https://www.bankofmaldives.com.mv/internetbanking/api/profile")
	if err != nil {
		return err
	}

	return nil
}

// GetTodayHistory ...
func (s *Session) GetTodayHistory(accountID string) (*HistoryPayload, error) {
	res, err := s.client.Get(fmt.Sprintf("https://www.bankofmaldives.com.mv/internetbanking/api/account/%s/history/today", accountID))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("non 200 status code")
	}

	historyResponse := &Response{Payload: &HistoryPayload{}}
	err = json.NewDecoder(res.Body).Decode(historyResponse)
	if err != nil {
		return nil, err
	}

	if historyResponse.Success == false {
		return nil, errors.New("get history unsuccessful at the application level")
	}

	if historyResponse.Code != 0 {
		return nil, fmt.Errorf("unexpected code: %d", historyResponse.Code)
	}

	return historyResponse.Payload.(*HistoryPayload), nil
}

// GetStatement ...
func (s *Session) GetStatement(accountID string, from, to time.Time, page int) (*HistoryPayload, error) {
	fromDate := from.Format("20060102")
	toDate := to.Format("20060102")

	res, err := s.client.Get(fmt.Sprintf("https://www.bankofmaldives.com.mv/internetbanking/api/account/%s/history/%s/%s/%d", accountID, fromDate, toDate, page))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("non 200 status code")
	}

	historyResponse := &Response{Payload: &HistoryPayload{}}
	err = json.NewDecoder(res.Body).Decode(historyResponse)
	if err != nil {
		return nil, err
	}

	if historyResponse.Success == false {
		return nil, errors.New("get history unsuccessful at the application level")
	}

	if historyResponse.Code != 0 {
		return nil, fmt.Errorf("unexpected code: %d", historyResponse.Code)
	}

	return historyResponse.Payload.(*HistoryPayload), nil
}

// GetActivity ...
func (s *Session) GetActivity(from, to time.Time, page int) (*ActivityPayload, error) {
	fromDate := from.Format("01/02/2006")
	toDate := to.Format("01/02/2006")

	res, err := s.client.Get(fmt.Sprintf("https://www.bankofmaldives.com.mv/internetbanking/api/activities?min_date=%s&max_date=%s&page=%d", fromDate, toDate, page))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("non 200 status code")
	}

	activityResponse := &Response{Payload: &ActivityPayload{}}
	err = json.NewDecoder(res.Body).Decode(activityResponse)
	if err != nil {
		return nil, err
	}

	if activityResponse.Success == false {
		return nil, errors.New("get activity unsuccessful at the application level")
	}

	if activityResponse.Code != 0 {
		return nil, fmt.Errorf("unexpected code: %d", activityResponse.Code)
	}

	return activityResponse.Payload.(*ActivityPayload), nil
}

// PostTransferRequest ...
func (s *Session) PostTransferRequest(requestForm url.Values) error {
	res, err := s.client.PostForm("https://www.bankofmaldives.com.mv/internetbanking/api/transfer", requestForm)
	if err != nil {
		return err
	}

	// parse the response
	if res.StatusCode != 200 {
		return errors.New("non 200 status code received")
	}

	response := &Response{}
	json.NewDecoder(res.Body).Decode(response)

	if response.Success == false {
		return errors.New("transfer request failed")
	}

	if response.Code != 22 {
		return fmt.Errorf("unexpected code: %d", response.Code)
	}

	return nil
}

// PostTransferCompletion ...
func (s *Session) PostTransferCompletion(completionForm url.Values) (*TransferCompletionPayload, error) {
	res, err := s.client.PostForm("https://www.bankofmaldives.com.mv/internetbanking/api/transfer", completionForm)
	if err != nil {
		return nil, err
	}

	// parse the response
	if res.StatusCode != 200 {
		return nil, errors.New("non 200 status code received")
	}

	response := &Response{Payload: &TransferCompletionPayload{}}
	json.NewDecoder(res.Body).Decode(response)

	if response.Success == false {
		return nil, errors.New("transfer request failed")
	}

	if response.Code != 0 {
		return nil, fmt.Errorf("unexpected code: %d", response.Code)
	}

	return response.Payload.(*TransferCompletionPayload), nil
}

// NewSession ...
func NewSession() *Session {
	return &Session{
		client: NewClient(),
	}
}
