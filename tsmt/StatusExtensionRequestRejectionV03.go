package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03300103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.033.001.03 Document"`
	Message *StatusExtensionRequestRejectionV03 `xml:"StsXtnsnReqRjctn"`
}

func (d *Document03300103) AddMessage() *StatusExtensionRequestRejectionV03 {
	d.Message = new(StatusExtensionRequestRejectionV03)
	return d.Message
}

// Scope
// The StatusExtensionRequestRejection message is sent by the party requested to accept or reject a request to extend the status of a transaction to the matching application.
// This message is used to inform about the rejection of a request to extend the status of a transaction.
// Usage
// The StatusExtensionRequestRejection message can be sent by the party requested to accept or reject the request to extend the status of a transaction to inform that it rejects the request.
// The message can be sent in response to a StatusExtensionRequestNotification message.
// The acceptance of a request to extend the status of a transaction can be achieved by sending a StatusExtensionRequestAcceptance message.
type StatusExtensionRequestRejectionV03 struct {

	// Identifies the rejection message.
	RejectionIdentification *iso20022.MessageIdentification1 `xml:"RjctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	// 
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Identifies the status that the submitter does not want to be extended.
	StatusNotToBeExtended *iso20022.TransactionStatus4 `xml:"StsNotToBeXtnded"`

	// Reason why the user cannot accept the request.
	RejectionReason *iso20022.Reason2 `xml:"RjctnRsn"`

}


func (s *StatusExtensionRequestRejectionV03) AddRejectionIdentification() *iso20022.MessageIdentification1 {
	s.RejectionIdentification = new(iso20022.MessageIdentification1)
	return s.RejectionIdentification
}

func (s *StatusExtensionRequestRejectionV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionRequestRejectionV03) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusExtensionRequestRejectionV03) AddStatusNotToBeExtended() *iso20022.TransactionStatus4 {
	s.StatusNotToBeExtended = new(iso20022.TransactionStatus4)
	return s.StatusNotToBeExtended
}

func (s *StatusExtensionRequestRejectionV03) AddRejectionReason() *iso20022.Reason2 {
	s.RejectionReason = new(iso20022.Reason2)
	return s.RejectionReason
}

