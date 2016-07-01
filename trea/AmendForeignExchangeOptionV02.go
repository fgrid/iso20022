package trea

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:trea.010.001.02 Document"`
	Message *AmendForeignExchangeOptionV02 `xml:"AmdFXOptnV02"`
}

func (d *Document01000102) AddMessage() *AmendForeignExchangeOptionV02 {
	d.Message = new(AmendForeignExchangeOptionV02)
	return d.Message
}

// Scope
// The AmendForeignExchangeOption message is sent by a participant to a central system or to a counterparty to notify the amendment of a foreign currency option contract.
// Usage
// The message contains a Related Reference to link it to the previously sent notification and may contain an reason for amendment.
// This message is only suitable for Simple (i.e. not Barrier) Vanilla (i.e. not Binary, Digital, Notouch) Foreign Exchange Options.
type AmendForeignExchangeOptionV02 struct {

	// Provides reference and date of the foreign exchange option trade which is amended.
	TradeInformation *iso20022.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the currency option trade which is amended.
	TradingSideIdentification *iso20022.TradePartyIdentification4 `xml:"TradgSdId"`

	// Specifies the counterparty of the currency option trade which is amended.
	CounterpartySideIdentification *iso20022.TradePartyIdentification4 `xml:"CtrPtySdId"`

	// Specifies the parameters of the currency option which is bought by the trading side.
	Option *iso20022.Option3 `xml:"Optn"`

}


func (a *AmendForeignExchangeOptionV02) AddTradeInformation() *iso20022.TradeAgreement2 {
	a.TradeInformation = new(iso20022.TradeAgreement2)
	return a.TradeInformation
}

func (a *AmendForeignExchangeOptionV02) AddTradingSideIdentification() *iso20022.TradePartyIdentification4 {
	a.TradingSideIdentification = new(iso20022.TradePartyIdentification4)
	return a.TradingSideIdentification
}

func (a *AmendForeignExchangeOptionV02) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification4 {
	a.CounterpartySideIdentification = new(iso20022.TradePartyIdentification4)
	return a.CounterpartySideIdentification
}

func (a *AmendForeignExchangeOptionV02) AddOption() *iso20022.Option3 {
	a.Option = new(iso20022.Option3)
	return a.Option
}

