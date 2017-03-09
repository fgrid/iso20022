package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02300207 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.023.002.07 Document"`
	Message *SecuritiesSettlementTransactionInstruction002V07 `xml:"SctiesSttlmTxInstr"`
}

func (d *Document02300207) AddMessage() *SecuritiesSettlementTransactionInstruction002V07 {
	d.Message = new(SecuritiesSettlementTransactionInstruction002V07)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementTransactionInstruction to an account servicer to instruct the receipt or delivery of financial instruments with or without payment, physically or by book-entry.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manages a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of transactions to a central securities depository or another settlement market infrastructure.
//
// Usage
// The instruction may be linked to other settlement instructions, for example, for a turnaround or back-to-back, or other transactions, for example, foreign exchange deal, using the linkage functionality.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionInstruction002V07 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.RestrictedFINXMax16Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *iso20022.SettlementTypeAndAdditionalParameters20 `xml:"SttlmTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages43 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails63 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount56 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails132 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *iso20022.AmountAndDirection60 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionInstruction002V07) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddSettlementTypeAndAdditionalParameters() *iso20022.SettlementTypeAndAdditionalParameters20 {
	s.SettlementTypeAndAdditionalParameters = new(iso20022.SettlementTypeAndAdditionalParameters20)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddLinkages() *iso20022.Linkages43 {
	newValue := new(iso20022.Linkages43)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddTradeDetails() *iso20022.SecuritiesTradeDetails63 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails63)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount56 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount56)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddSettlementParameters() *iso20022.SettlementDetails132 {
	s.SettlementParameters = new(iso20022.SettlementDetails132)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddDeliveringSettlementParties() *iso20022.SettlementParties44 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddReceivingSettlementParties() *iso20022.SettlementParties44 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddCashParties() *iso20022.CashParties30 {
	s.CashParties = new(iso20022.CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddSettlementAmount() *iso20022.AmountAndDirection60 {
	s.SettlementAmount = new(iso20022.AmountAndDirection60)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddOtherAmounts() *iso20022.OtherAmounts35 {
	s.OtherAmounts = new(iso20022.OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddOtherBusinessParties() *iso20022.OtherParties29 {
	s.OtherBusinessParties = new(iso20022.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionInstruction002V07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
