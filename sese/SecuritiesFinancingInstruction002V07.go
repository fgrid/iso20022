package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03300207 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.033.002.07 Document"`
	Message *SecuritiesFinancingInstruction002V07 `xml:"SctiesFincgInstr"`
}

func (d *Document03300207) AddMessage() *SecuritiesFinancingInstruction002V07 {
	d.Message = new(SecuritiesFinancingInstruction002V07)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesFinancingInstruction to a securities financing transaction account servicer to notify the securities financing transaction account servicer of the details of a repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing transaction to allow the account servicer to manage the settlement and follow-up of the opening and closing leg of the transaction.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of securities financing transactions to a central securities depository or another settlement market infrastructure.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesFinancingInstruction002V07 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.RestrictedFINXMax16Text `xml:"TxId"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionTypeAndAdditionalParameters *iso20022.TransactionTypeAndAdditionalParameters18 `xml:"TxTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages43 `xml:"Lnkgs,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails59 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount56 `xml:"QtyAndAcctDtls"`

	// Details for the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails30 `xml:"SctiesFincgDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails105 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *iso20022.AmountAndDirection60 `xml:"OpngSttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingInstruction002V07) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingInstruction002V07) AddTransactionTypeAndAdditionalParameters() *iso20022.TransactionTypeAndAdditionalParameters18 {
	s.TransactionTypeAndAdditionalParameters = new(iso20022.TransactionTypeAndAdditionalParameters18)
	return s.TransactionTypeAndAdditionalParameters
}

func (s *SecuritiesFinancingInstruction002V07) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesFinancingInstruction002V07) AddLinkages() *iso20022.Linkages43 {
	newValue := new(iso20022.Linkages43)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesFinancingInstruction002V07) AddTradeDetails() *iso20022.SecuritiesTradeDetails59 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails59)
	return s.TradeDetails
}

func (s *SecuritiesFinancingInstruction002V07) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingInstruction002V07) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingInstruction002V07) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount56 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount56)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingInstruction002V07) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails30 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails30)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingInstruction002V07) AddSettlementParameters() *iso20022.SettlementDetails105 {
	s.SettlementParameters = new(iso20022.SettlementDetails105)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingInstruction002V07) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingInstruction002V07) AddDeliveringSettlementParties() *iso20022.SettlementParties44 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingInstruction002V07) AddReceivingSettlementParties() *iso20022.SettlementParties44 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingInstruction002V07) AddCashParties() *iso20022.CashParties30 {
	s.CashParties = new(iso20022.CashParties30)
	return s.CashParties
}

func (s *SecuritiesFinancingInstruction002V07) AddOpeningSettlementAmount() *iso20022.AmountAndDirection60 {
	s.OpeningSettlementAmount = new(iso20022.AmountAndDirection60)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingInstruction002V07) AddOtherAmounts() *iso20022.OtherAmounts35 {
	s.OtherAmounts = new(iso20022.OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingInstruction002V07) AddOtherBusinessParties() *iso20022.OtherParties29 {
	s.OtherBusinessParties = new(iso20022.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingInstruction002V07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
