package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04000103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.040.001.03 Document"`
	Message *TimeOutNotificationV03 `xml:"TmOutNtfctn"`
}

func (d *Document04000103) AddMessage() *TimeOutNotificationV03 {
	d.Message = new(TimeOutNotificationV03)
	return d.Message
}

// Scope
// The TimeOutNotification message is sent by the matching application to a party involved in a transaction.
// This message is used to inform that a transaction will be closed within a given span of time if no action is taken.
// Usage
// The TimeOutNotification message can be sent by the matching application to inform the parties that a transaction will be closed within a given span of time if no action is taken by either party because
// - no activity for the transaction has taken place within a specified span of time,or
// - a significant date is reached, for example latest shipment date.
type TimeOutNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	// 
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established. 
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Describes the time-out consequences.
	TimeOutDescription *iso20022.TimeOutResult2 `xml:"TmOutDesc"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`

}


func (t *TimeOutNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	t.NotificationIdentification = new(iso20022.MessageIdentification1)
	return t.NotificationIdentification
}

func (t *TimeOutNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	t.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return t.TransactionIdentification
}

func (t *TimeOutNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	t.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return t.EstablishedBaselineIdentification
}

func (t *TimeOutNotificationV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	t.TransactionStatus = new(iso20022.TransactionStatus4)
	return t.TransactionStatus
}

func (t *TimeOutNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new (iso20022.DocumentIdentification5)
	t.UserTransactionReference = append(t.UserTransactionReference, newValue)
	return newValue
}

func (t *TimeOutNotificationV03) AddTimeOutDescription() *iso20022.TimeOutResult2 {
	t.TimeOutDescription = new(iso20022.TimeOutResult2)
	return t.TimeOutDescription
}

func (t *TimeOutNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	t.RequestForAction = new(iso20022.PendingActivity2)
	return t.RequestForAction
}

