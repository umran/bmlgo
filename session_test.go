package bmlgo

import "testing"

func TestAuth(t *testing.T) {
	s := NewSession()
	if err := s.Authenticate("USERNAME", "PASSWORD"); err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}

func TestGetTodayHistory(t *testing.T) {
	s := NewSession()
	err := s.Authenticate("USERNAME", "PASSWORD")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}

	_, err = s.GetTodayHistory("ACCOUNT_ID")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}

func TestPostTransferRequest(t *testing.T) {
	s := NewSession()
	err := s.Authenticate("USERNAME", "PASSWORD")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}

	transfer := NewTransfer("1", "FROM_ACCOUNT_NUMBER", "TO_ACCOUNT_NUMBER")
	err = s.PostTransferRequest(transfer.GetRequestForm())
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}

func TestPostTransferCompletion(t *testing.T) {
	s := NewSession()
	err := s.Authenticate("USERNAME", "PASSWORD")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}

	transfer := NewTransfer("1", "FROM_ACCOUNT_NUMBER", "TO_ACCOUNT_NUMBER")
	_, err = s.PostTransferCompletion(transfer.GenerateCompletionForm("OTP"))
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}
