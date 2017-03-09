package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600103 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.006.001.03 Document"`
	Message *AmendmentAcceptanceNotificationV03 `xml:"AmdmntAccptncNtfctn"`
}

func (d *Document00600103) AddMessage() *AmendmentAcceptanceNotificationV03 {
	d.Message = new(AmendmentAcceptanceNotificationV03)
	return d.Message
}

// Scope
// The AmendmentAcceptanceNotification message is sent by the matching application to the requester of an amendment.
// This message is used to notify the acceptance of an amendment request.
// Usage
// The AmendmentAcceptanceNotification message can be sent by the matching application to pass on information about the acceptance of an amendment request that it has obtained through the receipt of an AmendmentAcceptance message.
// In order to pass on information about the rejection of an amendment request the matching application sends an AmendmentRejectionNotification message.
type AmendmentAcceptanceNotificationV03 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reference to the identification of the delta report that contained the amendment.
	DeltaReportReference *iso20022.MessageIdentification1 `xml:"DltaRptRef"`

	// Sequence number of the accepted baseline amendment.
	AcceptedAmendmentNumber *iso20022.Count1 `xml:"AccptdAmdmntNb"`

	// Party that has accepted the amendment.
	Initiator *iso20022.BICIdentification1 `xml:"Initr"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (a *AmendmentAcceptanceNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	a.NotificationIdentification = new(iso20022.MessageIdentification1)
	return a.NotificationIdentification
}

func (a *AmendmentAcceptanceNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	a.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AmendmentAcceptanceNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	a.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return a.EstablishedBaselineIdentification
}

func (a *AmendmentAcceptanceNotificationV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	a.TransactionStatus = new(iso20022.TransactionStatus4)
	return a.TransactionStatus
}

func (a *AmendmentAcceptanceNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	a.UserTransactionReference = append(a.UserTransactionReference, newValue)
	return newValue
}

func (a *AmendmentAcceptanceNotificationV03) AddDeltaReportReference() *iso20022.MessageIdentification1 {
	a.DeltaReportReference = new(iso20022.MessageIdentification1)
	return a.DeltaReportReference
}

func (a *AmendmentAcceptanceNotificationV03) AddAcceptedAmendmentNumber() *iso20022.Count1 {
	a.AcceptedAmendmentNumber = new(iso20022.Count1)
	return a.AcceptedAmendmentNumber
}

func (a *AmendmentAcceptanceNotificationV03) AddInitiator() *iso20022.BICIdentification1 {
	a.Initiator = new(iso20022.BICIdentification1)
	return a.Initiator
}

func (a *AmendmentAcceptanceNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	a.RequestForAction = new(iso20022.PendingActivity2)
	return a.RequestForAction
}
