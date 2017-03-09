package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02700101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.027.001.01 Document"`
	Message *SecuritiesTradeConfirmationV01 `xml:"SctiesTradConf"`
}

func (d *Document02700101) AddMessage() *SecuritiesTradeConfirmationV01 {
	d.Message = new(SecuritiesTradeConfirmationV01)
	return d.Message
}

// Scope
// Sent by an executing party to an instructing party directly or through Central Matching Utility (CMU) to provide trade confirmation on a per-account basis based on instructions provided by the instructing party in the SecuritiesAllocationInstruction message.
// It may also be used to provide trade confirmation on the trade level from an executing party or an instructing party to the custodian or an affirming party directly or through CMU.
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager or of other categories of investors.
// The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The custodian or the affirming party is typically the custodian, trustee, financial institution, intermediary system/vendor communicating on behalf of them, or their agent.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: In local matching, the initiator of this message is always the executing party. In central matching the initiator may be either the executing party or instructing party.
// Respondent: instructing party, a custodian or an affirming party responds with SecuritiesTradeConfirmationResponse (accept or reject) message.
type SecuritiesTradeConfirmationV01 struct {

	// Information that unambiguously identifies an SecuritiesTradeConfirmation message as known by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.TransactiontIdentification4 `xml:"Id"`

	// Count of the number of transactions linked.
	NumberCount *iso20022.NumberCount1Choice `xml:"NbCnt,omitempty"`

	// Reference to the transaction identifier issued by a business party and/or market infrastructure. It may also be used to reference a previous transaction, for example, a block/allocation instruction, or tie a set of messages together.
	References []*iso20022.Linkages15 `xml:"Refs,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.Order14 `xml:"TradDtls"`

	// Unique and unambiguous identifier of a financial instrument, assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes31 `xml:"FinInstrmAttrbts,omitempty"`

	// Underlying financial instrument to which an trade confirmation is related.
	UnderlyingFinancialInstrument []*iso20022.UnderlyingFinancialInstrument1 `xml:"UndrlygFinInstrm,omitempty"`

	// Additional restrictions on the financial instrument, related to the stipulation.
	Stipulations *iso20022.FinancialInstrumentStipulations2 `xml:"Stiptns,omitempty"`

	// Parties involved in the confirmation of the details of a trade.
	ConfirmationParties []*iso20022.ConfirmationParties2 `xml:"ConfPties"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails43 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstruction *iso20022.StandingSettlementInstruction9 `xml:"StgSttlmInstr,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties23 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties23 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the specific transaction.
	CashParties *iso20022.CashParties6 `xml:"CshPties,omitempty"`

	// Provides clearing member information.
	ClearingDetails *iso20022.Clearing3 `xml:"ClrDtls,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.  The amount includes the principal with any commissions and fees or accrued interest.
	SettlementAmount *iso20022.AmountAndDirection28 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts []*iso20022.OtherAmounts16 `xml:"OthrAmts,omitempty"`

	// Other prices than the deal price.
	OtherPrices []*iso20022.OtherPrices1 `xml:"OthrPrics,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties18 `xml:"OthrBizPties,omitempty"`

	// Identifies a transaction that the trading parties are agreeing to repurchase, sell back or return the same or similar securities at a later time.
	// The two leg transaction details defines the closing leg conditions of a two leg transaction. It is also used to define the anticipated closing leg conditions at the time of opening the closed-end transaction.
	//
	//
	TwoLegTransactionDetails *iso20022.TwoLegTransactionDetails1 `xml:"TwoLegTxDtls,omitempty"`

	// Specifies regulatory stipulations that financial institutions must be compliant with in the country, region, and/or area they conduct business.
	RegulatoryStipulations *iso20022.RegulatoryStipulations1 `xml:"RgltryStiptns,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeConfirmationV01) AddIdentification() *iso20022.TransactiontIdentification4 {
	s.Identification = new(iso20022.TransactiontIdentification4)
	return s.Identification
}

func (s *SecuritiesTradeConfirmationV01) AddNumberCount() *iso20022.NumberCount1Choice {
	s.NumberCount = new(iso20022.NumberCount1Choice)
	return s.NumberCount
}

func (s *SecuritiesTradeConfirmationV01) AddReferences() *iso20022.Linkages15 {
	newValue := new(iso20022.Linkages15)
	s.References = append(s.References, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddTradeDetails() *iso20022.Order14 {
	s.TradeDetails = new(iso20022.Order14)
	return s.TradeDetails
}

func (s *SecuritiesTradeConfirmationV01) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeConfirmationV01) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes31 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes31)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeConfirmationV01) AddUnderlyingFinancialInstrument() *iso20022.UnderlyingFinancialInstrument1 {
	newValue := new(iso20022.UnderlyingFinancialInstrument1)
	s.UnderlyingFinancialInstrument = append(s.UnderlyingFinancialInstrument, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddStipulations() *iso20022.FinancialInstrumentStipulations2 {
	s.Stipulations = new(iso20022.FinancialInstrumentStipulations2)
	return s.Stipulations
}

func (s *SecuritiesTradeConfirmationV01) AddConfirmationParties() *iso20022.ConfirmationParties2 {
	newValue := new(iso20022.ConfirmationParties2)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddSettlementParameters() *iso20022.SettlementDetails43 {
	s.SettlementParameters = new(iso20022.SettlementDetails43)
	return s.SettlementParameters
}

func (s *SecuritiesTradeConfirmationV01) AddStandingSettlementInstruction() *iso20022.StandingSettlementInstruction9 {
	s.StandingSettlementInstruction = new(iso20022.StandingSettlementInstruction9)
	return s.StandingSettlementInstruction
}

func (s *SecuritiesTradeConfirmationV01) AddDeliveringSettlementParties() *iso20022.SettlementParties23 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties23)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeConfirmationV01) AddReceivingSettlementParties() *iso20022.SettlementParties23 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties23)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeConfirmationV01) AddCashParties() *iso20022.CashParties6 {
	s.CashParties = new(iso20022.CashParties6)
	return s.CashParties
}

func (s *SecuritiesTradeConfirmationV01) AddClearingDetails() *iso20022.Clearing3 {
	s.ClearingDetails = new(iso20022.Clearing3)
	return s.ClearingDetails
}

func (s *SecuritiesTradeConfirmationV01) AddSettlementAmount() *iso20022.AmountAndDirection28 {
	s.SettlementAmount = new(iso20022.AmountAndDirection28)
	return s.SettlementAmount
}

func (s *SecuritiesTradeConfirmationV01) AddOtherAmounts() *iso20022.OtherAmounts16 {
	newValue := new(iso20022.OtherAmounts16)
	s.OtherAmounts = append(s.OtherAmounts, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddOtherPrices() *iso20022.OtherPrices1 {
	newValue := new(iso20022.OtherPrices1)
	s.OtherPrices = append(s.OtherPrices, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationV01) AddOtherBusinessParties() *iso20022.OtherParties18 {
	s.OtherBusinessParties = new(iso20022.OtherParties18)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeConfirmationV01) AddTwoLegTransactionDetails() *iso20022.TwoLegTransactionDetails1 {
	s.TwoLegTransactionDetails = new(iso20022.TwoLegTransactionDetails1)
	return s.TwoLegTransactionDetails
}

func (s *SecuritiesTradeConfirmationV01) AddRegulatoryStipulations() *iso20022.RegulatoryStipulations1 {
	s.RegulatoryStipulations = new(iso20022.RegulatoryStipulations1)
	return s.RegulatoryStipulations
}

func (s *SecuritiesTradeConfirmationV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
