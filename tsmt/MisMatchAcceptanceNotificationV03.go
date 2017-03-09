package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.021.001.03 Document"`
	Message *MisMatchAcceptanceNotificationV03 `xml:"MisMtchAccptncNtfctn"`
}

func (d *Document02100103) AddMessage() *MisMatchAcceptanceNotificationV03 {
	d.Message = new(MisMatchAcceptanceNotificationV03)
	return d.Message
}

// Scope
// The MisMatchAcceptanceNotification message is sent by the matching application to the parties involved in the transaction.
// This message is used to notify the acceptance of mis-matched data sets.
// Usage
// The MisMatchAcceptanceNotification message can be sent by the matching application to pass on information about the acceptance of mis-matched data sets that it has obtained through the receipt of an MisMatchAcceptance message.
// In order to pass on information about the rejection of mis-matched data sets the matching application sends a MisMatchRejectionNotification message.
type MisMatchAcceptanceNotificationV03 struct {

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

	// Reference to the identification of the report that contained the difference.
	DataSetMatchReportReference *iso20022.MessageIdentification1 `xml:"DataSetMtchRptRef"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (m *MisMatchAcceptanceNotificationV03) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	m.NotificationIdentification = new(iso20022.MessageIdentification1)
	return m.NotificationIdentification
}

func (m *MisMatchAcceptanceNotificationV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	m.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return m.TransactionIdentification
}

func (m *MisMatchAcceptanceNotificationV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	m.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return m.EstablishedBaselineIdentification
}

func (m *MisMatchAcceptanceNotificationV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	m.TransactionStatus = new(iso20022.TransactionStatus4)
	return m.TransactionStatus
}

func (m *MisMatchAcceptanceNotificationV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	m.UserTransactionReference = append(m.UserTransactionReference, newValue)
	return newValue
}

func (m *MisMatchAcceptanceNotificationV03) AddDataSetMatchReportReference() *iso20022.MessageIdentification1 {
	m.DataSetMatchReportReference = new(iso20022.MessageIdentification1)
	return m.DataSetMatchReportReference
}

func (m *MisMatchAcceptanceNotificationV03) AddRequestForAction() *iso20022.PendingActivity2 {
	m.RequestForAction = new(iso20022.PendingActivity2)
	return m.RequestForAction
}
