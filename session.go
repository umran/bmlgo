package bmlgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type session struct {
	client *http.Client
}

func (s *session) authenticate(username, password string) error {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	res, err := s.client.PostForm("https://www.bankofmaldives.com.mv/internetbanking/api/login", form)
	if err != nil {
		return err
	}

	// parse the response
	if res.StatusCode != 200 {
		return interpretStatusCode(res.StatusCode)
	}

	response := &Response{}
	json.NewDecoder(res.Body).Decode(response)

	if response.Success == false {
		return interpretCode(response.Code)
	}

	// the following is some weird gymnastics to get the server to allow subsequent requests to the history endpoint.
	_, err = s.client.Get("https://www.bankofmaldives.com.mv/internetbanking/api/profile")
	if err != nil {
		return err
	}

	return nil
}

func (s *session) getTodaysHistory(accountID string) (*HistoryPayload, error) {
	res, err := s.client.Get(fmt.Sprintf("https://www.bankofmaldives.com.mv/internetbanking/api/account/%s/history/today", accountID))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, interpretStatusCode(res.StatusCode)
	}

	response := &Response{Payload: &HistoryPayload{}}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Success == false {
		return nil, interpretCode(response.Code)
	}

	return response.Payload.(*HistoryPayload), nil
}

func (s *session) getStatement(accountID string, from, to time.Time, page int) (*HistoryPayload, error) {
	fromDate := from.Format("20060102")
	toDate := to.Format("20060102")

	res, err := s.client.Get(fmt.Sprintf("https://www.bankofmaldives.com.mv/internetbanking/api/account/%s/history/%s/%s/%d", accountID, fromDate, toDate, page))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, interpretStatusCode(res.StatusCode)
	}

	response := &Response{Payload: &HistoryPayload{}}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Success == false {
		return nil, interpretCode(response.Code)
	}

	return response.Payload.(*HistoryPayload), nil
}

func (s *session) getActivity(from, to time.Time, page int) (*ActivityPayload, error) {
	fromDate := from.Format("01/02/2006")
	toDate := to.Format("01/02/2006")

	res, err := s.client.Get(fmt.Sprintf("https://www.bankofmaldives.com.mv/internetbanking/api/activities?min_date=%s&max_date=%s&page=%d", fromDate, toDate, page))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, interpretStatusCode(res.StatusCode)
	}

	response := &Response{Payload: &ActivityPayload{}}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Success == false {
		return nil, interpretCode(response.Code)
	}

	return response.Payload.(*ActivityPayload), nil
}

func (s *session) postTransferRequest(requestForm url.Values) error {
	res, err := s.client.PostForm("https://www.bankofmaldives.com.mv/internetbanking/api/transfer", requestForm)
	if err != nil {
		return err
	}

	// parse the response
	if res.StatusCode != 200 {
		return interpretStatusCode(res.StatusCode)
	}

	response := &Response{}
	json.NewDecoder(res.Body).Decode(response)

	if response.Success == false {
		return interpretCode(response.Code)
	}

	return nil
}

func (s *session) postTransferCompletion(completionForm url.Values) (*TransferCompletionPayload, error) {
	res, err := s.client.PostForm("https://www.bankofmaldives.com.mv/internetbanking/api/transfer", completionForm)
	if err != nil {
		return nil, err
	}

	// parse the response
	if res.StatusCode != 200 {
		return nil, interpretStatusCode(res.StatusCode)
	}

	response := &Response{Payload: &TransferCompletionPayload{}}
	json.NewDecoder(res.Body).Decode(response)

	if response.Success == false {
		return nil, interpretCode(response.Code)
	}

	return response.Payload.(*TransferCompletionPayload), nil
}

func newSession() *session {
	jar, _ := cookiejar.New(nil)
	return &session{
		client: &http.Client{
			Jar: jar,
		},
	}
}
