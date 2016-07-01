package trea

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:trea.003.001.02 Document"`
	Message *CancelNonDeliverableForwardOpeningV02 `xml:"CclNDFOpngV02"`
}

func (d *Document00300102) AddMessage() *CancelNonDeliverableForwardOpeningV02 {
	d.Message = new(CancelNonDeliverableForwardOpeningV02)
	return d.Message
}

// Scope
// The CancelNonDeliverableForwardOpening message is sent by a participant to a central system or to a counterparty to notify the cancellation of the opening of a non deliverable trade previously confirmed by the sender.
// Usage
// The message will contain a Related Reference to link it to the previously sent notification. It may contain a reason for cancellation.
type CancelNonDeliverableForwardOpeningV02 struct {

	// Provides references and date of the non deliverable trade which is cancelled.
	TradeInformation *iso20022.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is cancelled.
	TradingSideIdentification *iso20022.TradePartyIdentification3 `xml:"TradgSdId,omitempty"`

	// Specifies the counterparty of the non deliverable trade which is cancelled.
	CounterpartySideIdentification *iso20022.TradePartyIdentification3 `xml:"CtrPtySdId,omitempty"`

	// Specifies the amounts of the non deliverable trade which is cancelled.
	TradeAmounts *iso20022.AmountsAndValueDate1 `xml:"TradAmts,omitempty"`

	// Specifies the rate of the non deliverable trade which is cancelled.
	AgreedRate *iso20022.AgreedRate1 `xml:"AgrdRate,omitempty"`

	// Specifies the valuation conditions of the non deliverable trade which is cancelled.
	ValuationConditions *iso20022.NonDeliverableForwardValuationConditions2 `xml:"ValtnConds,omitempty"`

}


func (c *CancelNonDeliverableForwardOpeningV02) AddTradeInformation() *iso20022.TradeAgreement2 {
	c.TradeInformation = new(iso20022.TradeAgreement2)
	return c.TradeInformation
}

func (c *CancelNonDeliverableForwardOpeningV02) AddTradingSideIdentification() *iso20022.TradePartyIdentification3 {
	c.TradingSideIdentification = new(iso20022.TradePartyIdentification3)
	return c.TradingSideIdentification
}

func (c *CancelNonDeliverableForwardOpeningV02) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification3 {
	c.CounterpartySideIdentification = new(iso20022.TradePartyIdentification3)
	return c.CounterpartySideIdentification
}

func (c *CancelNonDeliverableForwardOpeningV02) AddTradeAmounts() *iso20022.AmountsAndValueDate1 {
	c.TradeAmounts = new(iso20022.AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *CancelNonDeliverableForwardOpeningV02) AddAgreedRate() *iso20022.AgreedRate1 {
	c.AgreedRate = new(iso20022.AgreedRate1)
	return c.AgreedRate
}

func (c *CancelNonDeliverableForwardOpeningV02) AddValuationConditions() *iso20022.NonDeliverableForwardValuationConditions2 {
	c.ValuationConditions = new(iso20022.NonDeliverableForwardValuationConditions2)
	return c.ValuationConditions
}

