package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04400102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.044.001.02 Document"`
	Message *IntentToPayNotificationV02 `xml:"InttToPayNtfctn"`
}

func (d *Document04400102) AddMessage() *IntentToPayNotificationV02 {
	d.Message = new(IntentToPayNotificationV02)
	return d.Message
}

// Scope
// The IntentToPayNotification message is sent by a party to the matching application in order to provide details about a future payment.
// This message contains details about an intention to pay a certain amount, on a certain date, in relation to one or several transactions known to the matching application.
// Usage
// The IntentToPayNotification message can be sent by a party to the transaction at any time as long as the transaction is established and not yet closed.
// The message is unsolicited, that is, it is not sent in response to another message.
type IntentToPayNotificationV02 struct {

	// Identifies the notification message.
	NotificationIdentification *iso20022.MessageIdentification1 `xml:"NtfctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	//
	BuyerBank *iso20022.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	//
	SellerBank *iso20022.BICIdentification1 `xml:"SellrBk"`

	// Provides the details of the intention to pay.
	IntentToPay *iso20022.IntentToPay2 `xml:"InttToPay"`
}

func (i *IntentToPayNotificationV02) AddNotificationIdentification() *iso20022.MessageIdentification1 {
	i.NotificationIdentification = new(iso20022.MessageIdentification1)
	return i.NotificationIdentification
}

func (i *IntentToPayNotificationV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	i.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return i.TransactionIdentification
}

func (i *IntentToPayNotificationV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	i.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return i.SubmitterTransactionReference
}

func (i *IntentToPayNotificationV02) AddBuyerBank() *iso20022.BICIdentification1 {
	i.BuyerBank = new(iso20022.BICIdentification1)
	return i.BuyerBank
}

func (i *IntentToPayNotificationV02) AddSellerBank() *iso20022.BICIdentification1 {
	i.SellerBank = new(iso20022.BICIdentification1)
	return i.SellerBank
}

func (i *IntentToPayNotificationV02) AddIntentToPay() *iso20022.IntentToPay2 {
	i.IntentToPay = new(iso20022.IntentToPay2)
	return i.IntentToPay
}
