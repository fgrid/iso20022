package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03600101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.036.001.01 Document"`
	Message *SecuritiesFinancingModificationInstructionV01 `xml:"SctiesFincgModInstr"`
}

func (d *Document03600101) AddMessage() *SecuritiesFinancingModificationInstructionV01 {
	d.Message = new(SecuritiesFinancingModificationInstructionV01)
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
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingModificationInstructionV01 struct {

	// Information that unambiguously identifies a SecuritiesFinancingModificationInstruction message as known by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing), modification information and other parameters.
	TransactionTypeAndModificationAdditionalParameters *iso20022.TransactionTypeAndAdditionalParameters2 `xml:"TxTpAndModAddtlParams"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails5 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification11 `xml:"FinInstrmId"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount7 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingAdditionalDetails *iso20022.SecuritiesFinancingTransactionDetails1 `xml:"SctiesFincgAddtlDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails3 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *iso20022.AmountAndDirection10 `xml:"OpngSttlmAmt,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`

}


func (s *SecuritiesFinancingModificationInstructionV01) AddIdentification() *iso20022.DocumentIdentification11 {
	s.Identification = new(iso20022.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesFinancingModificationInstructionV01) AddTransactionTypeAndModificationAdditionalParameters() *iso20022.TransactionTypeAndAdditionalParameters2 {
	s.TransactionTypeAndModificationAdditionalParameters = new(iso20022.TransactionTypeAndAdditionalParameters2)
	return s.TransactionTypeAndModificationAdditionalParameters
}

func (s *SecuritiesFinancingModificationInstructionV01) AddTradeDetails() *iso20022.SecuritiesTradeDetails5 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails5)
	return s.TradeDetails
}

func (s *SecuritiesFinancingModificationInstructionV01) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingModificationInstructionV01) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount7 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount7)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingModificationInstructionV01) AddSecuritiesFinancingAdditionalDetails() *iso20022.SecuritiesFinancingTransactionDetails1 {
	s.SecuritiesFinancingAdditionalDetails = new(iso20022.SecuritiesFinancingTransactionDetails1)
	return s.SecuritiesFinancingAdditionalDetails
}

func (s *SecuritiesFinancingModificationInstructionV01) AddSettlementParameters() *iso20022.SettlementDetails3 {
	s.SettlementParameters = new(iso20022.SettlementDetails3)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingModificationInstructionV01) AddDeliveringSettlementParties() *iso20022.SettlementParties5 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingModificationInstructionV01) AddReceivingSettlementParties() *iso20022.SettlementParties5 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingModificationInstructionV01) AddOpeningSettlementAmount() *iso20022.AmountAndDirection10 {
	s.OpeningSettlementAmount = new(iso20022.AmountAndDirection10)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingModificationInstructionV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	s.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesFinancingModificationInstructionV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	s.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesFinancingModificationInstructionV01) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

