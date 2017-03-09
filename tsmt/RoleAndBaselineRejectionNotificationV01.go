package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05200101 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.052.001.01 Document"`
	Message *RoleAndBaselineRejectionNotificationV01 `xml:"RoleAndBaselnRjctnNtfctn"`
}

func (d *Document05200101) AddMessage() *RoleAndBaselineRejectionNotificationV01 {
	d.Message = new(RoleAndBaselineRejectionNotificationV01)
	return d.Message
}

// Scope
// The RoleAndBaselineRejectionNotification message is sent by the matching application to the primary banks to inform about role and baseline rejection by a secondary bank.
// Usage
// The RoleAndBaselineRejectionNotification message is used to inform that a secondary bank has rejected the role and baseline. No response is expected.
type RoleAndBaselineRejectionNotificationV01 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction that is not extended.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Party that has rejected.
	Initiator *iso20022.BICIdentification1 `xml:"Initr"`

	// Reason why the user cannot accept the request.
	RejectionReason *iso20022.Reason2 `xml:"RjctnRsn,omitempty"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (r *RoleAndBaselineRejectionNotificationV01) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	r.NotificationIdentification = new(iso20022.MessageIdentification1)
	return r.NotificationIdentification
}

func (r *RoleAndBaselineRejectionNotificationV01) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	r.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return r.TransactionIdentification
}

func (r *RoleAndBaselineRejectionNotificationV01) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	r.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return r.EstablishedBaselineIdentification
}

func (r *RoleAndBaselineRejectionNotificationV01) AddTransactionStatus() *iso20022.TransactionStatus4 {
	r.TransactionStatus = new(iso20022.TransactionStatus4)
	return r.TransactionStatus
}

func (r *RoleAndBaselineRejectionNotificationV01) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	r.UserTransactionReference = append(r.UserTransactionReference, newValue)
	return newValue
}

func (r *RoleAndBaselineRejectionNotificationV01) AddInitiator() *iso20022.BICIdentification1 {
	r.Initiator = new(iso20022.BICIdentification1)
	return r.Initiator
}

func (r *RoleAndBaselineRejectionNotificationV01) AddRejectionReason() *iso20022.Reason2 {
	r.RejectionReason = new(iso20022.Reason2)
	return r.RejectionReason
}

func (r *RoleAndBaselineRejectionNotificationV01) AddRequestForAction() *iso20022.PendingActivity2 {
	r.RequestForAction = new(iso20022.PendingActivity2)
	return r.RequestForAction
}
