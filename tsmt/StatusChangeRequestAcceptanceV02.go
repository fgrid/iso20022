package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02700102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.027.001.02 Document"`
	Message *StatusChangeRequestAcceptanceV02 `xml:"StsChngReqAccptnc"`
}

func (d *Document02700102) AddMessage() *StatusChangeRequestAcceptanceV02 {
	d.Message = new(StatusChangeRequestAcceptanceV02)
	return d.Message
}

// Scope
// The StatusChangeRequestAcceptance message is sent by the party requested to accept or reject the request of a change in the status of a transaction to the matching application.
// This message is used to inform about the acceptance of a request to change the status of a transaction.
// Usage
// The StatusChangeRequestAcceptance message can be sent by the party requested to accept or reject a request to change the status of a transaction to inform that it accepts the request.
// The message can be sent in response to a StatusChangeRequestNotification message.
// The rejection of a request to change the status of a transaction can be achieved by sending a StatusChangeRequestRejection message.
type StatusChangeRequestAcceptanceV02 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *iso20022.MessageIdentification1 `xml:"AccptncId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the status accepted.
	AcceptedStatus *iso20022.TransactionStatus3 `xml:"AccptdSts"`

}


func (s *StatusChangeRequestAcceptanceV02) AddAcceptanceIdentification() *iso20022.MessageIdentification1 {
	s.AcceptanceIdentification = new(iso20022.MessageIdentification1)
	return s.AcceptanceIdentification
}

func (s *StatusChangeRequestAcceptanceV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestAcceptanceV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusChangeRequestAcceptanceV02) AddAcceptedStatus() *iso20022.TransactionStatus3 {
	s.AcceptedStatus = new(iso20022.TransactionStatus3)
	return s.AcceptedStatus
}

