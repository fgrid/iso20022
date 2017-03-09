package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03600206 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.002.06 Document"`
	Message *SecuritiesFinancingModificationInstruction002V06 `xml:"SctiesFincgModInstr"`
}

func (d *Document03600206) AddMessage() *SecuritiesFinancingModificationInstruction002V06 {
	d.Message = new(SecuritiesFinancingModificationInstruction002V06)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesFinancingModificationInstruction to a securities financing transaction account servicer to notify the securities financing transaction account servicer of an update in the details of a repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing transaction that does not impact the original transaction securities quantity.
// Such a change may be:
// - the providing of closing details not available at the time of the sending of the Securities Financing Instruction, for example, termination date for an open repo,
// - the providing of a new rate, for example, a repo rate,
// - the rollover of a position extending the closing or maturity date.
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
type SecuritiesFinancingModificationInstruction002V06 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing), modification information and other parameters.
	TransactionTypeAndModificationAdditionalParameters *iso20022.TransactionTypeAndAdditionalParameters20 `xml:"TxTpAndModAddtlParams"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails12 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification20 `xml:"FinInstrmId"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount61 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingAdditionalDetails *iso20022.SecuritiesFinancingTransactionDetails32 `xml:"SctiesFincgAddtlDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails106 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *iso20022.AmountAndDirection66 `xml:"OpngSttlmAmt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddTransactionTypeAndModificationAdditionalParameters() *iso20022.TransactionTypeAndAdditionalParameters20 {
	s.TransactionTypeAndModificationAdditionalParameters = new(iso20022.TransactionTypeAndAdditionalParameters20)
	return s.TransactionTypeAndModificationAdditionalParameters
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddTradeDetails() *iso20022.SecuritiesTradeDetails12 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails12)
	return s.TradeDetails
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount61 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount61)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddSecuritiesFinancingAdditionalDetails() *iso20022.SecuritiesFinancingTransactionDetails32 {
	s.SecuritiesFinancingAdditionalDetails = new(iso20022.SecuritiesFinancingTransactionDetails32)
	return s.SecuritiesFinancingAdditionalDetails
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddSettlementParameters() *iso20022.SettlementDetails106 {
	s.SettlementParameters = new(iso20022.SettlementDetails106)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddDeliveringSettlementParties() *iso20022.SettlementParties44 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddReceivingSettlementParties() *iso20022.SettlementParties44 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddOpeningSettlementAmount() *iso20022.AmountAndDirection66 {
	s.OpeningSettlementAmount = new(iso20022.AmountAndDirection66)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingModificationInstruction002V06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
