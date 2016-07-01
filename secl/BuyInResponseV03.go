package secl

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:secl.008.001.03 Document"`
	Message *BuyInResponseV03 `xml:"BuyInRspn"`
}

func (d *Document00800103) AddMessage() *BuyInResponseV03 {
	d.Message = new(BuyInResponseV03)
	return d.Message
}

// Scope
// The BuyInResponse message is sent by the clearing member to the central counterparty as a response to the previous buy-in notification message.
// 
// The message definition is intended for use with the ISO 20022 Business Application Header.
// 
// Usage
// The BuyInResponse may be sent in response to the BuyInNotification message. However, the use of this message in the buy in process is optional and depends on the rules set by each central counterparty.
type BuyInResponseV03 struct {

	// Unambiguous identification of the transaction as known by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId,omitempty"`

	// Provides response details such as a request for delay and the number of days associated to that request.
	BuyInResponseDetails *iso20022.BuyIn3 `xml:"BuyInRspnDtls"`

	// Provides details about the original settlement obligation that did not settle and for which the buy in process will be launched.
	OriginalSettlementObligationDetails *iso20022.SettlementObligation7 `xml:"OrgnlSttlmOblgtnDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (b *BuyInResponseV03) SetTransactionIdentification(value string) {
	b.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (b *BuyInResponseV03) AddBuyInResponseDetails() *iso20022.BuyIn3 {
	b.BuyInResponseDetails = new(iso20022.BuyIn3)
	return b.BuyInResponseDetails
}

func (b *BuyInResponseV03) AddOriginalSettlementObligationDetails() *iso20022.SettlementObligation7 {
	b.OriginalSettlementObligationDetails = new(iso20022.SettlementObligation7)
	return b.OriginalSettlementObligationDetails
}

func (b *BuyInResponseV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}

