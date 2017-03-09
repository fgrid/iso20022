package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04500102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.045.001.02 Document"`
	Message *ForwardIntentToPayNotificationV02 `xml:"FwdInttToPayNtfctn"`
}

func (d *Document04500102) AddMessage() *ForwardIntentToPayNotificationV02 {
	d.Message = new(ForwardIntentToPayNotificationV02)
	return d.Message
}

// Scope
// The ForwardIntentToPayNotification message is forwarded by the matching application from one primary bank to the other primary bank in order to provide details about a future payment.
// This message contains details about an intention to pay a certain amount, on a certain date, in relation to one or several transactions known to the matching application.
// Usage
// The ForwardIntentToPayNotification message is a copy of the IntentToPayNotification message received by the matching application and forwarded to the other primary bank for information. No response is expected.
type ForwardIntentToPayNotificationV02 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for the financial institutions involved in this transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *iso20022.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *iso20022.BICIdentification1 `xml:"SellrBk"`

	// Provides the details of the intention to pay.
	IntentToPay *iso20022.IntentToPay2 `xml:"InttToPay"`

	// Next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *ForwardIntentToPayNotificationV02) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	f.NotificationIdentification = new(iso20022.MessageIdentification1)
	return f.NotificationIdentification
}

func (f *ForwardIntentToPayNotificationV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	f.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return f.TransactionIdentification
}

func (f *ForwardIntentToPayNotificationV02) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	f.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return f.EstablishedBaselineIdentification
}

func (f *ForwardIntentToPayNotificationV02) AddTransactionStatus() *iso20022.TransactionStatus4 {
	f.TransactionStatus = new(iso20022.TransactionStatus4)
	return f.TransactionStatus
}

func (f *ForwardIntentToPayNotificationV02) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	f.UserTransactionReference = append(f.UserTransactionReference, newValue)
	return newValue
}

func (f *ForwardIntentToPayNotificationV02) AddBuyerBank() *iso20022.BICIdentification1 {
	f.BuyerBank = new(iso20022.BICIdentification1)
	return f.BuyerBank
}

func (f *ForwardIntentToPayNotificationV02) AddSellerBank() *iso20022.BICIdentification1 {
	f.SellerBank = new(iso20022.BICIdentification1)
	return f.SellerBank
}

func (f *ForwardIntentToPayNotificationV02) AddIntentToPay() *iso20022.IntentToPay2 {
	f.IntentToPay = new(iso20022.IntentToPay2)
	return f.IntentToPay
}

func (f *ForwardIntentToPayNotificationV02) AddRequestForAction() *iso20022.PendingActivity2 {
	f.RequestForAction = new(iso20022.PendingActivity2)
	return f.RequestForAction
}
