package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500103 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.035.001.03 Document"`
	Message *StatusExtensionRequestV03 `xml:"StsXtnsnReq"`
}

func (d *Document03500103) AddMessage() *StatusExtensionRequestV03 {
	d.Message = new(StatusExtensionRequestV03)
	return d.Message
}

// Scope
// The StatusExtensionRequest message is sent by either party involved in a transaction to the matching application.
// This message is used to request the extension of the status of a transaction.
// Usage
// The StatusExtensionRequest message can be sent by either party involved in a transaction to the matching application to request the extension of the status of a transaction.
// The matching application will pass on the request by sending a StatusExtensionRequestNotification message to the counterparty which can accept or reject the status extension request by sending a StatusExtensionAcceptance or StatusExtensionRejection message.
type StatusExtensionRequestV03 struct {

	// Identifies the request message.
	RequestIdentification *iso20022.MessageIdentification1 `xml:"ReqId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Identifies the status of the transaction by means of a code.
	StatusToBeExtended *iso20022.TransactionStatus5 `xml:"StsToBeXtnded"`
}

func (s *StatusExtensionRequestV03) AddRequestIdentification() *iso20022.MessageIdentification1 {
	s.RequestIdentification = new(iso20022.MessageIdentification1)
	return s.RequestIdentification
}

func (s *StatusExtensionRequestV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionRequestV03) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *StatusExtensionRequestV03) AddStatusToBeExtended() *iso20022.TransactionStatus5 {
	s.StatusToBeExtended = new(iso20022.TransactionStatus5)
	return s.StatusToBeExtended
}
