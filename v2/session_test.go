package bmlgo

import "testing"

func TestAuth(t *testing.T) {
	s := newSession()
	if err := s.authenticate("USERNAME", "PASSWORD"); err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}

func TestGetTodaysHistory(t *testing.T) {
	s := newSession()
	err := s.authenticate("USERNAME", "PASSWORD")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}

	_, err = s.getTodaysHistory("ACCOUNT_ID")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}

func TestPostTransferRequest(t *testing.T) {
	s := newSession()
	err := s.authenticate("USERNAME", "PASSWORD")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}

	transfer := newTransfer("1", "FROM_ACCOUNT_NUMBER", "TO_ACCOUNT_NUMBER")
	err = s.postTransferRequest(transfer.getRequestForm())
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}

func TestPostTransferCompletion(t *testing.T) {
	s := newSession()
	err := s.authenticate("USERNAME", "PASSWORD")
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}

	transfer := newTransfer("1", "FROM_ACCOUNT_NUMBER", "TO_ACCOUNT_NUMBER")
	_, err = s.postTransferCompletion(transfer.generateCompletionForm("OTP"))
	if err != nil {
		t.Error("not so lucky")
		t.Log(err)
	}
}
