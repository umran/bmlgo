package bmlgo

import (
	"net/url"
	"time"

	"github.com/umran/decimal"
)

// Client ...
type Client struct {
	username string
	password string
	session  *session
}

// NewClient ...
func NewClient(username, password string) (*Client, error) {
	session := newSession()
	if err := session.authenticate(username, password); err != nil {
		return nil, err
	}

	return &Client{
		username,
		password,
		session,
	}, nil
}

// GetNewStatementItems ...
func (c *Client) GetNewStatementItems(accountID, cursor string) ([]*HistoryItem, string, error) {
	to := time.Now()
	from := to.AddDate(-1, 0, 0)
	data := make([]*HistoryItem, 0)
	currentPage := 1

search:
	for {
		statement, err := c.session.getStatement(accountID, from, to, currentPage)
		if err == ErrorNotAuthenticated {
			if err := c.reauthenticate(); err != nil {
				return nil, "", err
			}
			time.Sleep(time.Second)
			continue
		} else if err != nil {
			return nil, "", err
		}

		if len(statement.History) == 0 {
			if cursor == "" {
				break search
			}
			return nil, "", ErrorCursorUnreachable
		}

		for _, item := range statement.History {
			if item.ID == cursor {
				break search
			}
			data = append(data, item)
		}

		currentPage++
		time.Sleep(time.Second)
	}

	nextCursor := cursor
	if len(data) > 0 {
		nextCursor = data[0].ID
	}

	return data, nextCursor, nil
}

// InitiateTransfer ...
func (c *Client) InitiateTransfer(amount int, debitAccount, creditAccount string) (url.Values, error) {
	t := newTransfer(amountAsRufiyaaString(amount), debitAccount, creditAccount)
	for {
		err := c.session.postTransferRequest(t.getRequestForm())
		if err == ErrorNotAuthenticated {
			if err := c.reauthenticate(); err != nil {
				return nil, err
			}
			time.Sleep(time.Second)
			continue
		} else if err != nil {
			return nil, err
		}
		return t.getRequestForm(), nil
	}
}

// CompleteTransfer ...
func (c *Client) CompleteTransfer(request url.Values, otp string) (*TransferCompletionPayload, error) {
	t := &transfer{request}
	for {
		completionPayload, err := c.session.postTransferCompletion(t.generateCompletionForm(otp))
		if err == ErrorNotAuthenticated {
			if err := c.reauthenticate(); err != nil {
				return nil, err
			}
			time.Sleep(time.Second)
			continue
		} else if err != nil {
			return nil, err
		}
		return completionPayload, nil
	}
}

// helper method to reauthenticate session
func (c *Client) reauthenticate() error {
	return c.session.authenticate(c.username, c.password)
}

func amountAsRufiyaaString(amount int) string {
	return decimal.New(int64(amount), int32(-2)).
		StringFixedBank(2)
}
