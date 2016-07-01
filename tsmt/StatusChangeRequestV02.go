package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02600102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.026.001.02 Document"`
	Message *StatusChangeRequestV02 `xml:"StsChngReq"`
}

func (d *Document02600102) AddMessage() *StatusChangeRequestV02 {
	d.Message = new(StatusChangeRequestV02)
	return d.Message
}

// Scope
// The StatusChangeRequest message is sent by a party involved in a transaction to the matching application.
// This message is used to request a change in the status of a transaction.
// Usage
// The StatusChangeRequest message can be sent by either party involved in a transaction to the matching application to request a change in the status of a transaction.
// The matching application will pass on the request by sending a StatusChangeRequestNotification message to the counterparty which can accept or reject the request by sending a SatausChangeRequestAcceptance or StatusChangeRequestRejection message.
type StatusChangeRequestV02 struct {

	// Identifies the request message.
	RequestIdentification *iso20022.MessageIdentification1 `xml:"ReqId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the baseline status requested by the submitter by means of a code. 
	RequestedStatus *iso20022.TransactionStatus3 `xml:"ReqdSts"`

	// Specifies the reason for the request to change status.
	RequestReason *iso20022.Reason2 `xml:"ReqRsn,omitempty"`

}


func (s *StatusChangeRequestV02) AddRequestIdentification() *iso20022.MessageIdentification1 {
	s.RequestIdentification = new(iso20022.MessageIdentification1)
	return s.RequestIdentification
}

func (s *StatusChangeRequestV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusChangeRequestV02) AddRequestedStatus() *iso20022.TransactionStatus3 {
	s.RequestedStatus = new(iso20022.TransactionStatus3)
	return s.RequestedStatus
}

func (s *StatusChangeRequestV02) AddRequestReason() *iso20022.Reason2 {
	s.RequestReason = new(iso20022.Reason2)
	return s.RequestReason
}

