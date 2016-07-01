package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02800103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.028.001.03 Document"`
	Message *StatusChangeRequestNotificationV03 `xml:"StsChngReqNtfctn"`
}

func (d *Document02800103) AddMessage() *StatusChangeRequestNotificationV03 {
	d.Message = new(StatusChangeRequestNotificationV03)
	return d.Message
}

// Scope
// The StatusChangeRequestNotification message is sent by the matching application to the party requested to accept or reject the request of a change in the status of a transaction.
// This message is used to notify the request of a change in the status of a transaction.
// Usage
// The StatusChangeRequestNotification message can be sent by the matching application to pass on the information about the request of a change in the status of a transaction that it has obtained through the receipt of a StatusChangeRequest message.
type StatusChangeRequestNotificationV03 struct {

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

	// Specifies the status that is proposed by the other party.
	ProposedStatusChange *iso20022.TransactionStatus3 `xml:"PropsdStsChng"`

	// Specifies the reason for the request to change status.
	RequestReason *iso20022.Reason2 `xml:"ReqRsn,omitempty"`

	// Party that has requested the status change.
	Initiator *iso20022.BICIdentification1 `xml:"Initr"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`

}


func (s *StatusChangeRequestNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	s.NotificationIdentification = new(iso20022.MessageIdentification1)
	return s.NotificationIdentification
}

func (s *StatusChangeRequestNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	s.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *StatusChangeRequestNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	s.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return s.EstablishedBaselineIdentification
}

func (s *StatusChangeRequestNotificationV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	s.TransactionStatus = new(iso20022.TransactionStatus4)
	return s.TransactionStatus
}

func (s *StatusChangeRequestNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new (iso20022.DocumentIdentification5)
	s.UserTransactionReference = append(s.UserTransactionReference, newValue)
	return newValue
}

func (s *StatusChangeRequestNotificationV03) AddProposedStatusChange() *iso20022.TransactionStatus3 {
	s.ProposedStatusChange = new(iso20022.TransactionStatus3)
	return s.ProposedStatusChange
}

func (s *StatusChangeRequestNotificationV03) AddRequestReason() *iso20022.Reason2 {
	s.RequestReason = new(iso20022.Reason2)
	return s.RequestReason
}

func (s *StatusChangeRequestNotificationV03) AddInitiator() *iso20022.BICIdentification1 {
	s.Initiator = new(iso20022.BICIdentification1)
	return s.Initiator
}

func (s *StatusChangeRequestNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	s.RequestForAction = new(iso20022.PendingActivity2)
	return s.RequestForAction
}

