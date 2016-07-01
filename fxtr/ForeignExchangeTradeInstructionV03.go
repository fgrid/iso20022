package fxtr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.014.001.03 Document"`
	Message *ForeignExchangeTradeInstructionV03 `xml:"FXTradInstr"`
}

func (d *Document01400103) AddMessage() *ForeignExchangeTradeInstructionV03 {
	d.Message = new(ForeignExchangeTradeInstructionV03)
	return d.Message
}

// Scope
// The ForeignExchangeTradeInstruction message is sent by a participant to a central settlement system to notify the creation of the foreign exchange trade agreed by both trading parties.
// Usage
// The ForeignExchangeTradeInstruction message is sent from a participant to a central settlement system to advise of the creation of a foreign exchange trade.
type ForeignExchangeTradeInstructionV03 struct {

	// General information related to the trade.
	TradeInformation *iso20022.TradeAgreement10 `xml:"TradInf"`

	// Party(ies) on the trading side of the trade.
	TradingSideIdentification *iso20022.TradePartyIdentification6 `xml:"TradgSdId"`

	// Party(ies) on the counterparty side of the trade.
	CounterpartySideIdentification *iso20022.TradePartyIdentification6 `xml:"CtrPtySdId"`

	// Amounts of the trade.
	TradeAmounts *iso20022.AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *iso20022.AgreedRate1 `xml:"AgrdRate"`

	// Settlement instructions for the amounts received by the trading side.
	TradingSideSettlementInstructions *iso20022.SettlementParties29 `xml:"TradgSdSttlmInstrs,omitempty"`

	// Settlement instructions for the amounts received by the counterparty.
	CounterpartySideSettlementInstructions *iso20022.SettlementParties29 `xml:"CtrPtySdSttlmInstrs,omitempty"`

	// Specifies whether the trade is a block or an individual trade. It also contains supplementary information such as free format information, broker's identification, dealing branches and references.
	OptionalGeneralInformation *iso20022.GeneralInformation4 `xml:"OptnlGnlInf,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *iso20022.RegulatoryReporting4 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (f *ForeignExchangeTradeInstructionV03) AddTradeInformation() *iso20022.TradeAgreement10 {
	f.TradeInformation = new(iso20022.TradeAgreement10)
	return f.TradeInformation
}

func (f *ForeignExchangeTradeInstructionV03) AddTradingSideIdentification() *iso20022.TradePartyIdentification6 {
	f.TradingSideIdentification = new(iso20022.TradePartyIdentification6)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeTradeInstructionV03) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification6 {
	f.CounterpartySideIdentification = new(iso20022.TradePartyIdentification6)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeTradeInstructionV03) AddTradeAmounts() *iso20022.AmountsAndValueDate1 {
	f.TradeAmounts = new(iso20022.AmountsAndValueDate1)
	return f.TradeAmounts
}

func (f *ForeignExchangeTradeInstructionV03) AddAgreedRate() *iso20022.AgreedRate1 {
	f.AgreedRate = new(iso20022.AgreedRate1)
	return f.AgreedRate
}

func (f *ForeignExchangeTradeInstructionV03) AddTradingSideSettlementInstructions() *iso20022.SettlementParties29 {
	f.TradingSideSettlementInstructions = new(iso20022.SettlementParties29)
	return f.TradingSideSettlementInstructions
}

func (f *ForeignExchangeTradeInstructionV03) AddCounterpartySideSettlementInstructions() *iso20022.SettlementParties29 {
	f.CounterpartySideSettlementInstructions = new(iso20022.SettlementParties29)
	return f.CounterpartySideSettlementInstructions
}

func (f *ForeignExchangeTradeInstructionV03) AddOptionalGeneralInformation() *iso20022.GeneralInformation4 {
	f.OptionalGeneralInformation = new(iso20022.GeneralInformation4)
	return f.OptionalGeneralInformation
}

func (f *ForeignExchangeTradeInstructionV03) AddRegulatoryReporting() *iso20022.RegulatoryReporting4 {
	f.RegulatoryReporting = new(iso20022.RegulatoryReporting4)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeInstructionV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}

