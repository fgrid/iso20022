package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.034.001.03 Document"`
	Message *StatusExtensionRejectionNotificationV03 `xml:"StsXtnsnRjctnNtfctn"`
}

func (d *Document03400103) AddMessage() *StatusExtensionRejectionNotificationV03 {
	d.Message = new(StatusExtensionRejectionNotificationV03)
	return d.Message
}

// Scope
// The StatusExtensionRejectionNotification message is sent by the matching application to the submitter of a request to extend the status of a transaction.
// This message is used to inform about the rejection of a request to extend the status of a transaction.
// Usage
// The StatusExtensionRejectionNotification message can be sent by the matching application to pass on information about the rejection of a request to extend the status of a transaction that it has obtained through the receipt of a StatusExtensionRejection message.
// In order to pass on information about the acceptance of a request to extend the status of a transaction the matching application sends a StatusExtensionNotification message
type StatusExtensionRejectionNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	// 
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established. 
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction that is not extended.
	NonExtendedStatus *iso20022.TransactionStatus4 `xml:"NonXtndedSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reason why the user cannot accept the request.
	RejectionReason *iso20022.Reason2 `xml:"RjctnRsn"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`

}


func (s *StatusExtensionRejectionNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	s.NotificationIdentification = new(iso20022.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusExtensionRejectionNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusExtensionRejectionNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusExtensionRejectionNotificationV03) AddNonExtendedStatus() *iso20022.TransactionStatus4 {
	s.NonExtendedStatus = new(iso20022.TransactionStatus4)
	return s.NonExtendedStatus
}

func (s *StatusExtensionRejectionNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new (iso20022.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusExtensionRejectionNotificationV03) AddRejectionReason() *iso20022.Reason2 {
	s.RejectionReason = new(iso20022.Reason2)
	return s.RejectionReason
}

func (s *StatusExtensionRejectionNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	s.RequestForAction = new(iso20022.PendingActivity2)
	return s.RequestForAction
}

