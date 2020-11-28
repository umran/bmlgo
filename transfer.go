package bmlgo

import "net/url"

type transfer struct {
	request url.Values
}

func (t *transfer) getRequestForm() url.Values {
	return t.request
}

func (t *transfer) generateCompletionForm(otp string) url.Values {
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

func newTransfer(debitAmount, debitAccount, creditAccount string) *transfer {
	request := url.Values{}
	request.Add("transfertype", "IAT")
	request.Add("channel", "email")
	request.Add("debitAmount", debitAmount)
	request.Add("currency", "MVR")
	request.Add("creditAccount", creditAccount)
	request.Add("debitAccount", debitAccount)

	return &transfer{
		request: request,
	}
}
