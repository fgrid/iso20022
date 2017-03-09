package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03600104 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.04 Document"`
	Message *SecuritiesFinancingModificationInstructionV04 `xml:"SctiesFincgModInstr"`
}

func (d *Document03600104) AddMessage() *SecuritiesFinancingModificationInstructionV04 {
	d.Message = new(SecuritiesFinancingModificationInstructionV04)
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
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingModificationInstructionV04 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing), modification information and other parameters.
	TransactionTypeAndModificationAdditionalParameters *iso20022.TransactionTypeAndAdditionalParameters7 `xml:"TxTpAndModAddtlParams"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails5 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount16 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingAdditionalDetails *iso20022.SecuritiesFinancingTransactionDetails19 `xml:"SctiesFincgAddtlDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails72 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *iso20022.AmountAndDirection10 `xml:"OpngSttlmAmt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingModificationInstructionV04) AddTransactionTypeAndModificationAdditionalParameters() *iso20022.TransactionTypeAndAdditionalParameters7 {
	s.TransactionTypeAndModificationAdditionalParameters = new(iso20022.TransactionTypeAndAdditionalParameters7)
	return s.TransactionTypeAndModificationAdditionalParameters
}

func (s *SecuritiesFinancingModificationInstructionV04) AddTradeDetails() *iso20022.SecuritiesTradeDetails5 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails5)
	return s.TradeDetails
}

func (s *SecuritiesFinancingModificationInstructionV04) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingModificationInstructionV04) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount16 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount16)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingModificationInstructionV04) AddSecuritiesFinancingAdditionalDetails() *iso20022.SecuritiesFinancingTransactionDetails19 {
	s.SecuritiesFinancingAdditionalDetails = new(iso20022.SecuritiesFinancingTransactionDetails19)
	return s.SecuritiesFinancingAdditionalDetails
}

func (s *SecuritiesFinancingModificationInstructionV04) AddSettlementParameters() *iso20022.SettlementDetails72 {
	s.SettlementParameters = new(iso20022.SettlementDetails72)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingModificationInstructionV04) AddDeliveringSettlementParties() *iso20022.SettlementParties10 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingModificationInstructionV04) AddReceivingSettlementParties() *iso20022.SettlementParties10 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingModificationInstructionV04) AddOpeningSettlementAmount() *iso20022.AmountAndDirection10 {
	s.OpeningSettlementAmount = new(iso20022.AmountAndDirection10)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingModificationInstructionV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
