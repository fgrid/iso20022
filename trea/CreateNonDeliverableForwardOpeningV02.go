package trea

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:trea.001.001.02 Document"`
	Message *CreateNonDeliverableForwardOpeningV02 `xml:"CretNDFOpngV02"`
}

func (d *Document00100102) AddMessage() *CreateNonDeliverableForwardOpeningV02 {
	d.Message = new(CreateNonDeliverableForwardOpeningV02)
	return d.Message
}

// Scope
// The CreateNonDeliverableForwardOpening message is sent by a participant to a central system or to a counterparty to notify the opening of a non deliverable trade.
// Usage
// The trading parties will send similar messages to the central settlement system and the central settlement system will send notifications to both parties.
type CreateNonDeliverableForwardOpeningV02 struct {

	// Provides identification and date of the non deliverable trade which is created.
	TradeInformation *iso20022.TradeAgreement1 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is created.
	TradingSideIdentification *iso20022.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the non deliverable trade which is created.
	CounterpartySideIdentification *iso20022.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the non deliverable trade which is created.
	TradeAmounts *iso20022.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the rate agreed at the opening of a non deliverable trade.
	AgreedRate *iso20022.AgreedRate1 `xml:"AgrdRate"`

	// Specifies the valuation conditions of the non deliverable trade which is created.
	ValuationConditions *iso20022.NonDeliverableForwardValuationConditions2 `xml:"ValtnConds"`
}

func (c *CreateNonDeliverableForwardOpeningV02) AddTradeInformation() *iso20022.TradeAgreement1 {
	c.TradeInformation = new(iso20022.TradeAgreement1)
	return c.TradeInformation
}

func (c *CreateNonDeliverableForwardOpeningV02) AddTradingSideIdentification() *iso20022.TradePartyIdentification3 {
	c.TradingSideIdentification = new(iso20022.TradePartyIdentification3)
	return c.TradingSideIdentification
}

func (c *CreateNonDeliverableForwardOpeningV02) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification3 {
	c.CounterpartySideIdentification = new(iso20022.TradePartyIdentification3)
	return c.CounterpartySideIdentification
}

func (c *CreateNonDeliverableForwardOpeningV02) AddTradeAmounts() *iso20022.AmountsAndValueDate1 {
	c.TradeAmounts = new(iso20022.AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *CreateNonDeliverableForwardOpeningV02) AddAgreedRate() *iso20022.AgreedRate1 {
	c.AgreedRate = new(iso20022.AgreedRate1)
	return c.AgreedRate
}

func (c *CreateNonDeliverableForwardOpeningV02) AddValuationConditions() *iso20022.NonDeliverableForwardValuationConditions2 {
	c.ValuationConditions = new(iso20022.NonDeliverableForwardValuationConditions2)
	return c.ValuationConditions
}
