package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02500103 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.025.001.03 Document"`
	Message *StatusChangeNotificationV03 `xml:"StsChngNtfctn"`
}

func (d *Document02500103) AddMessage() *StatusChangeNotificationV03 {
	d.Message = new(StatusChangeNotificationV03)
	return d.Message
}

// Scope
// The StatusChangeNotification message is sent by the matching application to the parties involved in the change of the status of a transaction.
// This message is used to inform about the change of a status.
// Usage
// The StatusChangeNotification message can be sent by the matching application
// - to the submitter of the request to change the status of a transaction to pass on the information about the acceptance of the request that it has obtained through the receipt of an StatusChangeRequestAcceptance message.
// - to the accepter of a request to change the status of a transaction inform about the actual status of the transaction in response to a StatusChangeRequestAcceptance message.
// - to either party involved in the establishment of a transaction to inform about the change of the status of the transaction to the status close. This can be done when the limit of possible attempts to establish the transaction has been reached or when unilaterally requested by one of the parties involved in the not yet established transaction.
// - to either party involved in a transaction to inform about the change of the status of the transaction as announced in a TimeOutNotification message sent by the matching application prior to the StatusChangeNotification message.
type StatusChangeNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (s *StatusChangeNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	s.NotificationIdentification = new(iso20022.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusChangeNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusChangeNotificationV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	s.TransactionStatus = new(iso20022.TransactionStatus4)
	return s.TransactionStatus
}

func (s *StatusChangeNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusChangeNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	s.RequestForAction = new(iso20022.PendingActivity2)
	return s.RequestForAction
}
