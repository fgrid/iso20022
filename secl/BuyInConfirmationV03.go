package secl

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:secl.009.001.03 Document"`
	Message *BuyInConfirmationV03 `xml:"BuyInConf"`
}

func (d *Document00900103) AddMessage() *BuyInConfirmationV03 {
	d.Message = new(BuyInConfirmationV03)
	return d.Message
}

// Scope
// The Buy In Confirmation message is sent by the central counterparty (CCP) to the clearing member to confirm the details of the transaction resulting from the buy in.
// 
// The message definition is intended for use with the ISO 20022 Business Application Header.
// 
// Usage
// The Buy In Confirmation message is sent by the central counterparty (CCP) to confirm the details of the buy in transaction.
type BuyInConfirmationV03 struct {

	// Unambiguous identification of the transaction as known by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId,omitempty"`

	// Provides the identification of the clearing member (individual clearing member or general clearing member).
	ClearingMember *iso20022.PartyIdentification35Choice `xml:"ClrMmb"`

	// Provides the buy-in details.
	BuyInDetails *iso20022.BuyIn2 `xml:"BuyInDtls"`

	// Provides details about the original settlement obligation that did not settle and for which the buy in process will be launched.
	OriginalSettlementObligation *iso20022.SettlementObligation7 `xml:"OrgnlSttlmOblgtn,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (b *BuyInConfirmationV03) SetTransactionIdentification(value string) {
	b.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (b *BuyInConfirmationV03) AddClearingMember() *iso20022.PartyIdentification35Choice {
	b.ClearingMember = new(iso20022.PartyIdentification35Choice)
	return b.ClearingMember
}

func (b *BuyInConfirmationV03) AddBuyInDetails() *iso20022.BuyIn2 {
	b.BuyInDetails = new(iso20022.BuyIn2)
	return b.BuyInDetails
}

func (b *BuyInConfirmationV03) AddOriginalSettlementObligation() *iso20022.SettlementObligation7 {
	b.OriginalSettlementObligation = new(iso20022.SettlementObligation7)
	return b.OriginalSettlementObligation
}

func (b *BuyInConfirmationV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}

