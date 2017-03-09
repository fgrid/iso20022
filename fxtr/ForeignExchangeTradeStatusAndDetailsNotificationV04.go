package fxtr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700104 struct {
	XMLName xml.Name                                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.017.001.04 Document"`
	Message *ForeignExchangeTradeStatusAndDetailsNotificationV04 `xml:"FXTradStsAndDtlsNtfctn"`
}

func (d *Document01700104) AddMessage() *ForeignExchangeTradeStatusAndDetailsNotificationV04 {
	d.Message = new(ForeignExchangeTradeStatusAndDetailsNotificationV04)
	return d.Message
}

// Scope
// The ForeignExchangeTradeStatusAndDetails message is sent by a central system to the participant to provide notification of the status and details of a foreign exchange trade.
// Usage
// The notification is sent by a central settlement system to the two trading parties after it has received foreign exchange trade instructions from both.
type ForeignExchangeTradeStatusAndDetailsNotificationV04 struct {

	// Provides information on the status of a foreign exchange trade in the central system.
	StatusDetails *iso20022.TradeData14 `xml:"StsDtls"`

	// General information related to the foreign exchange trade.
	TradeInformation *iso20022.TradeAgreement12 `xml:"TradInf"`

	// Party(ies) on the trading side of the foreign exchange trade.
	TradingSideIdentification *iso20022.TradePartyIdentification6 `xml:"TradgSdId"`

	// Party(ies) on the counterparty side of the foreign exchange trade.
	CounterpartySideIdentification *iso20022.TradePartyIdentification6 `xml:"CtrPtySdId"`

	// Amounts of the foreign exchange trade.
	TradeAmounts *iso20022.AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *iso20022.AgreedRate3 `xml:"AgrdRate"`

	// Provides the opening and fixing information for an NDF trade.
	NonDeliverableForwardConditions *iso20022.NonDeliverableForwardConditions2 `xml:"NDFConds,omitempty"`

	// Settlement instructions for the amounts received by the trading side.
	TradingSideSettlementInstructions *iso20022.SettlementParties29 `xml:"TradgSdSttlmInstrs,omitempty"`

	// Settlement instructions for the amounts received by the counterparty.
	CounterpartySideSettlementInstructions *iso20022.SettlementParties29 `xml:"CtrPtySdSttlmInstrs,omitempty"`

	// Additional Information about the foreign exchange trade.
	GeneralInformation *iso20022.GeneralInformation5 `xml:"GnlInf,omitempty"`

	// Details of the split trade.
	SplitTradeInformation []*iso20022.SplitTradeDetails3 `xml:"SpltTradInf,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *iso20022.RegulatoryReporting6 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddStatusDetails() *iso20022.TradeData14 {
	f.StatusDetails = new(iso20022.TradeData14)
	return f.StatusDetails
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradeInformation() *iso20022.TradeAgreement12 {
	f.TradeInformation = new(iso20022.TradeAgreement12)
	return f.TradeInformation
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradingSideIdentification() *iso20022.TradePartyIdentification6 {
	f.TradingSideIdentification = new(iso20022.TradePartyIdentification6)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification6 {
	f.CounterpartySideIdentification = new(iso20022.TradePartyIdentification6)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradeAmounts() *iso20022.AmountsAndValueDate1 {
	f.TradeAmounts = new(iso20022.AmountsAndValueDate1)
	return f.TradeAmounts
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddAgreedRate() *iso20022.AgreedRate3 {
	f.AgreedRate = new(iso20022.AgreedRate3)
	return f.AgreedRate
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddNonDeliverableForwardConditions() *iso20022.NonDeliverableForwardConditions2 {
	f.NonDeliverableForwardConditions = new(iso20022.NonDeliverableForwardConditions2)
	return f.NonDeliverableForwardConditions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradingSideSettlementInstructions() *iso20022.SettlementParties29 {
	f.TradingSideSettlementInstructions = new(iso20022.SettlementParties29)
	return f.TradingSideSettlementInstructions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddCounterpartySideSettlementInstructions() *iso20022.SettlementParties29 {
	f.CounterpartySideSettlementInstructions = new(iso20022.SettlementParties29)
	return f.CounterpartySideSettlementInstructions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddGeneralInformation() *iso20022.GeneralInformation5 {
	f.GeneralInformation = new(iso20022.GeneralInformation5)
	return f.GeneralInformation
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddSplitTradeInformation() *iso20022.SplitTradeDetails3 {
	newValue := new(iso20022.SplitTradeDetails3)
	f.SplitTradeInformation = append(f.SplitTradeInformation, newValue)
	return newValue
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddRegulatoryReporting() *iso20022.RegulatoryReporting6 {
	f.RegulatoryReporting = new(iso20022.RegulatoryReporting6)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
