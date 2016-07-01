package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.005.001.02 Document"`
	Message *RequestForAccountManagementStatusReportV02 `xml:"ReqForAcctMgmtStsRptV02"`
}

func (d *Document00500102) AddMessage() *RequestForAccountManagementStatusReportV02 {
	d.Message = new(RequestForAccountManagementStatusReportV02)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent, sends the RequestForAccountManagementStatusReport message to the account servicer, for example, a registrar, transfer agent or custodian bank to request the status of an AccountOpeningInstruction or an AccountModificationInstruction.
// Usage
// The RequestForAccountManagementStatusReport message is used to request the processing status of a previously sent AccountOpeningInstruction message or AccountModificationInstruction message for which a AccountDetailsConfirmation message has not yet been received.
type RequestForAccountManagementStatusReportV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Information to identify the account management instruction for which the status is requested.
	RequestDetails *iso20022.AccountManagementMessageReference1 `xml:"ReqDtls"`

}


func (r *RequestForAccountManagementStatusReportV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForAccountManagementStatusReportV02) AddRequestDetails() *iso20022.AccountManagementMessageReference1 {
	r.RequestDetails = new(iso20022.AccountManagementMessageReference1)
	return r.RequestDetails
}

