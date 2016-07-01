package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.04 Document"`
	Message *RequestForAccountManagementStatusReportV04 `xml:"ReqForAcctMgmtStsRpt"`
}

func (d *Document00500104) AddMessage() *RequestForAccountManagementStatusReportV04 {
	d.Message = new(RequestForAccountManagementStatusReportV04)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent, sends the RequestForAccountManagementStatusReport message to the account servicer, for example, a registrar, transfer agent, custodian bank or securities depository  to request the status of an AccountOpeningInstruction,  GetAccountDetails or an AccountModificationInstruction.
// Usage
// The RequestForAccountManagementStatusReport message is used to request the processing status of a previously sent AccountOpeningInstruction message or AccountModificationInstruction message for which a AccountDetailsConfirmation message has not yet been received.
type RequestForAccountManagementStatusReportV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Identifies the account for which the status of the account management instruction is requested.
	RequestDetails *iso20022.AccountManagementMessageReference3 `xml:"ReqDtls"`

}


func (r *RequestForAccountManagementStatusReportV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForAccountManagementStatusReportV04) AddRequestDetails() *iso20022.AccountManagementMessageReference3 {
	r.RequestDetails = new(iso20022.AccountManagementMessageReference3)
	return r.RequestDetails
}

