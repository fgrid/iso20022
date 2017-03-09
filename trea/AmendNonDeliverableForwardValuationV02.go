package trea

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:trea.005.001.02 Document"`
	Message *AmendNonDeliverableForwardValuationV02 `xml:"AmdNDFValtnV02"`
}

func (d *Document00500102) AddMessage() *AmendNonDeliverableForwardValuationV02 {
	d.Message = new(AmendNonDeliverableForwardValuationV02)
	return d.Message
}

// Scope
// The AmendNonDeliverableForwardValuation message is sent by a participant to a central system or to a counterparty to notify the amendment of the valuation of a non deliverable trade previously confirmed by the sender.
// Usage
// The message is sent from a participant to a central settlement system to advise of the update of a previously sent notification and it contains a "Related Reference" to link it to the previous notification.
type AmendNonDeliverableForwardValuationV02 struct {

	// Provides references and date of the valuation of the non deliverable trade which is amended.
	TradeInformation *iso20022.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the valuation of the non deliverable trade which is amended.
	TradingSideIdentification *iso20022.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the valuation of the non deliverable trade which is amended.
	CounterpartySideIdentification *iso20022.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the valuation of the non deliverable trade which is amended.
	TradeAmounts *iso20022.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the valuation rate of the valuation of the non deliverable trade which is amended.
	ValuationRate *iso20022.AgreedRate1 `xml:"ValtnRate"`

	// Specifies the valuation information of the valuation of the non deliverable trade which is amended.
	ValuationInformation *iso20022.ValuationData2 `xml:"ValtnInf"`
}

func (a *AmendNonDeliverableForwardValuationV02) AddTradeInformation() *iso20022.TradeAgreement2 {
	a.TradeInformation = new(iso20022.TradeAgreement2)
	return a.TradeInformation
}

func (a *AmendNonDeliverableForwardValuationV02) AddTradingSideIdentification() *iso20022.TradePartyIdentification3 {
	a.TradingSideIdentification = new(iso20022.TradePartyIdentification3)
	return a.TradingSideIdentification
}

func (a *AmendNonDeliverableForwardValuationV02) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification3 {
	a.CounterpartySideIdentification = new(iso20022.TradePartyIdentification3)
	return a.CounterpartySideIdentification
}

func (a *AmendNonDeliverableForwardValuationV02) AddTradeAmounts() *iso20022.AmountsAndValueDate1 {
	a.TradeAmounts = new(iso20022.AmountsAndValueDate1)
	return a.TradeAmounts
}

func (a *AmendNonDeliverableForwardValuationV02) AddValuationRate() *iso20022.AgreedRate1 {
	a.ValuationRate = new(iso20022.AgreedRate1)
	return a.ValuationRate
}

func (a *AmendNonDeliverableForwardValuationV02) AddValuationInformation() *iso20022.ValuationData2 {
	a.ValuationInformation = new(iso20022.ValuationData2)
	return a.ValuationInformation
}
