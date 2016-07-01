package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03300106 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.033.001.06 Document"`
	Message *SecuritiesFinancingInstructionV06 `xml:"SctiesFincgInstr"`
}

func (d *Document03300106) AddMessage() *SecuritiesFinancingInstructionV06 {
	d.Message = new(SecuritiesFinancingInstructionV06)
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
type SecuritiesFinancingInstructionV06 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionTypeAndAdditionalParameters *iso20022.TransactionTypeAndAdditionalParameters9 `xml:"TxTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages37 `xml:"Lnkgs,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails56 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount39 `xml:"QtyAndAcctDtls"`

	// Details for the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails28 `xml:"SctiesFincgDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails97 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *iso20022.AmountAndDirection46 `xml:"OpngSttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesFinancingInstructionV06) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesFinancingInstructionV06) AddTransactionTypeAndAdditionalParameters() *iso20022.TransactionTypeAndAdditionalParameters9 {
	s.TransactionTypeAndAdditionalParameters = new(iso20022.TransactionTypeAndAdditionalParameters9)
	return s.TransactionTypeAndAdditionalParameters
}

func (s *SecuritiesFinancingInstructionV06) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesFinancingInstructionV06) AddLinkages() *iso20022.Linkages37 {
	newValue := new (iso20022.Linkages37)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesFinancingInstructionV06) AddTradeDetails() *iso20022.SecuritiesTradeDetails56 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails56)
	return s.TradeDetails
}

func (s *SecuritiesFinancingInstructionV06) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingInstructionV06) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingInstructionV06) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount39 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount39)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingInstructionV06) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails28 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails28)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingInstructionV06) AddSettlementParameters() *iso20022.SettlementDetails97 {
	s.SettlementParameters = new(iso20022.SettlementDetails97)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingInstructionV06) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingInstructionV06) AddDeliveringSettlementParties() *iso20022.SettlementParties36 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingInstructionV06) AddReceivingSettlementParties() *iso20022.SettlementParties36 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingInstructionV06) AddCashParties() *iso20022.CashParties26 {
	s.CashParties = new(iso20022.CashParties26)
	return s.CashParties
}

func (s *SecuritiesFinancingInstructionV06) AddOpeningSettlementAmount() *iso20022.AmountAndDirection46 {
	s.OpeningSettlementAmount = new(iso20022.AmountAndDirection46)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingInstructionV06) AddOtherAmounts() *iso20022.OtherAmounts28 {
	s.OtherAmounts = new(iso20022.OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingInstructionV06) AddOtherBusinessParties() *iso20022.OtherParties27 {
	s.OtherBusinessParties = new(iso20022.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingInstructionV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

