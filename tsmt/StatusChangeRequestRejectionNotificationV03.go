package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03000103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.030.001.03 Document"`
	Message *StatusChangeRequestRejectionNotificationV03 `xml:"StsChngReqRjctnNtfctn"`
}

func (d *Document03000103) AddMessage() *StatusChangeRequestRejectionNotificationV03 {
	d.Message = new(StatusChangeRequestRejectionNotificationV03)
	return d.Message
}

// Scope
// The StatusChangeRequestRejectionNotification message is sent by the matching application to the submitter of a request to change the status of a transaction.
// This message is used to inform about the rejection of a request to change the status of a transaction.
// Usage
// The StatusChangeRequestRejectionNotification message can be sent by the matching application to pass on information about the rejection of a request to change the status of a transaction that it has obtained through the receipt of a StatusChangeRequestRejection message.
// In order to pass on information about the acceptance of a request to change the status of a transaction the matching application sends a StatusChangeNotification message.
type StatusChangeRequestRejectionNotificationV03 struct {

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

	// Specifies the status that was rejected by the other party.
	RejectedStatusChange *iso20022.TransactionStatus3 `xml:"RjctdStsChng"`

	// Reason why the user cannot accept the request.
	RejectionReason *iso20022.Reason2 `xml:"RjctnRsn"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`

}


func (s *StatusChangeRequestRejectionNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	s.NotificationIdentification = new(iso20022.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusChangeRequestRejectionNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestRejectionNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusChangeRequestRejectionNotificationV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	s.TransactionStatus = new(iso20022.TransactionStatus4)
	return s.TransactionStatus
}

func (s *StatusChangeRequestRejectionNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new (iso20022.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusChangeRequestRejectionNotificationV03) AddRejectedStatusChange() *iso20022.TransactionStatus3 {
	s.RejectedStatusChange = new(iso20022.TransactionStatus3)
	return s.RejectedStatusChange
}

func (s *StatusChangeRequestRejectionNotificationV03) AddRejectionReason() *iso20022.Reason2 {
	s.RejectionReason = new(iso20022.Reason2)
	return s.RejectionReason
}

func (s *StatusChangeRequestRejectionNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	s.RequestForAction = new(iso20022.PendingActivity2)
	return s.RequestForAction
}

