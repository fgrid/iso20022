package secl

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:secl.007.001.03 Document"`
	Message *BuyInNotificationV03 `xml:"BuyInNtfctn"`
}

func (d *Document00700103) AddMessage() *BuyInNotificationV03 {
	d.Message = new(BuyInNotificationV03)
	return d.Message
}

// Scope
// The BuyInNotification message is sent by the central counterparty (CCP) to a clearing member to notify the start of the buy in process.
// 
// The message definition is intended for use with the ISO 20022 Business Application Header.
// 
// Usage
// The buy in process is a process by which the CCP buys in stocks to cover failed transactions; the clearing member is notified that this process has started. Depending on each CCP internal rules, this message can also be sent, as a warning, by the central counterparty to the clearing member some days before the buy in process starts.
type BuyInNotificationV03 struct {

	// Unambiguous identification of the transaction as known by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId,omitempty"`

	// Provides the identification of the clearing member (individual clearing member or general clearing member).
	ClearingMember *iso20022.PartyIdentification35Choice `xml:"ClrMmb"`

	// Indicates if the message is a notification or a warning and gives the option to specify the buy in date.
	NotificationDetails *iso20022.BuyIn4 `xml:"NtfctnDtls,omitempty"`

	// Provides details about the original settlement obligation that did not settle and for which the buy in process will be launched.
	OriginalSettlementObligation *iso20022.SettlementObligation7 `xml:"OrgnlSttlmOblgtn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (b *BuyInNotificationV03) SetTransactionIdentification(value string) {
	b.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (b *BuyInNotificationV03) AddClearingMember() *iso20022.PartyIdentification35Choice {
	b.ClearingMember = new(iso20022.PartyIdentification35Choice)
	return b.ClearingMember
}

func (b *BuyInNotificationV03) AddNotificationDetails() *iso20022.BuyIn4 {
	b.NotificationDetails = new(iso20022.BuyIn4)
	return b.NotificationDetails
}

func (b *BuyInNotificationV03) AddOriginalSettlementObligation() *iso20022.SettlementObligation7 {
	b.OriginalSettlementObligation = new(iso20022.SettlementObligation7)
	return b.OriginalSettlementObligation
}

func (b *BuyInNotificationV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}

