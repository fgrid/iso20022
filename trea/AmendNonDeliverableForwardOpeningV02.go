package trea

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:trea.002.001.02 Document"`
	Message *AmendNonDeliverableForwardOpeningV02 `xml:"AmdNDFOpngV02"`
}

func (d *Document00200102) AddMessage() *AmendNonDeliverableForwardOpeningV02 {
	d.Message = new(AmendNonDeliverableForwardOpeningV02)
	return d.Message
}

// Scope
// The AmendNonDeliverableForwardOpening message is sent by a participant to a central system or to a counterparty to notify the amendment of the opening of a non deliverable trade previously confirmed by the sender.
// Usage
// The message is sent from a participant to a central settlement system to advise of the update of a previously sent notification and it contains a "Related Reference" to link it to the previous notification.
type AmendNonDeliverableForwardOpeningV02 struct {

	// Provides references and date of the non deliverable trade which is amended.
	TradeInformation *iso20022.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is amended.
	TradingSideIdentification *iso20022.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the non deliverable trade which is amended.
	CounterpartySideIdentification *iso20022.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the non deliverable trade which is amended.
	TradeAmounts *iso20022.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the rate of the non deliverable trade which is amended.
	AgreedRate *iso20022.AgreedRate1 `xml:"AgrdRate"`

	// Specifies the valuation conditions of the non deliverable trade which is amended.
	ValuationConditions *iso20022.NonDeliverableForwardValuationConditions2 `xml:"ValtnConds"`

}


func (a *AmendNonDeliverableForwardOpeningV02) AddTradeInformation() *iso20022.TradeAgreement2 {
	a.TradeInformation = new(iso20022.TradeAgreement2)
	return a.TradeInformation
}

func (a *AmendNonDeliverableForwardOpeningV02) AddTradingSideIdentification() *iso20022.TradePartyIdentification3 {
	a.TradingSideIdentification = new(iso20022.TradePartyIdentification3)
	return a.TradingSideIdentification
}

func (a *AmendNonDeliverableForwardOpeningV02) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification3 {
	a.CounterpartySideIdentification = new(iso20022.TradePartyIdentification3)
	return a.CounterpartySideIdentification
}

func (a *AmendNonDeliverableForwardOpeningV02) AddTradeAmounts() *iso20022.AmountsAndValueDate1 {
	a.TradeAmounts = new(iso20022.AmountsAndValueDate1)
	return a.TradeAmounts
}

func (a *AmendNonDeliverableForwardOpeningV02) AddAgreedRate() *iso20022.AgreedRate1 {
	a.AgreedRate = new(iso20022.AgreedRate1)
	return a.AgreedRate
}

func (a *AmendNonDeliverableForwardOpeningV02) AddValuationConditions() *iso20022.NonDeliverableForwardValuationConditions2 {
	a.ValuationConditions = new(iso20022.NonDeliverableForwardValuationConditions2)
	return a.ValuationConditions
}

