package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03600103 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.036.001.03 Document"`
	Message *StatusExtensionRequestNotificationV03 `xml:"StsXtnsnReqNtfctn"`
}

func (d *Document03600103) AddMessage() *StatusExtensionRequestNotificationV03 {
	d.Message = new(StatusExtensionRequestNotificationV03)
	return d.Message
}

// Scope
// The StatusExtensionRequestNotification message is sent by the matching application to the party requested to accept or reject a request to extend the status of a transaction.
// This message is used to notify the request to extend the status of a transaction.
// Usage
// The StatusExtensionRequestNotification message can be sent by the matching application to pass on information about the request to extend the status of a transaction that it has obtained through the receipt of a StatusExtensionRequest message.
type StatusExtensionRequestNotificationV03 struct {

	// Identifies the notification.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Identifies the status for which an extension of the validity period has been requested.
	StatusToBeExtended *iso20022.TransactionStatus5 `xml:"StsToBeXtnded"`

	// Party that has requested the status extension.
	Initiator *iso20022.BICIdentification1 `xml:"Initr"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *StatusExtensionRequestNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	s.NotificationIdentification = new(iso20022.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusExtensionRequestNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionRequestNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusExtensionRequestNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusExtensionRequestNotificationV03) AddStatusToBeExtended() *iso20022.TransactionStatus5 {
	s.StatusToBeExtended = new(iso20022.TransactionStatus5)
	return s.StatusToBeExtended
}

func (s *StatusExtensionRequestNotificationV03) AddInitiator() *iso20022.BICIdentification1 {
	s.Initiator = new(iso20022.BICIdentification1)
	return s.Initiator
}

func (s *StatusExtensionRequestNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	s.RequestForAction = new(iso20022.PendingActivity2)
	return s.RequestForAction
}
