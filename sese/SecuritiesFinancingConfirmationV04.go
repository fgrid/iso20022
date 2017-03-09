package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500104 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.001.04 Document"`
	Message *SecuritiesFinancingConfirmationV04 `xml:"SctiesFincgConf"`
}

func (d *Document03500104) AddMessage() *SecuritiesFinancingConfirmationV04 {
	d.Message = new(SecuritiesFinancingConfirmationV04)
	return d.Message
}

// Scope
// A securities financing transaction account servicer sends a SecuritiesFinancingConfirmation to an account owner to confirm or advise of the partial or full settlement of the opening or closing leg of a securities financing transaction.
//
// The account servicer/owner relationship may be:
//
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
//
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
//
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
//
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
type SecuritiesFinancingConfirmationV04 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *iso20022.TransactionTypeAndAdditionalParameters3 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters18 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails6 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount18 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails11 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction3 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails60 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties17 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts17 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingConfirmationV04) AddTransactionIdentificationDetails() *iso20022.TransactionTypeAndAdditionalParameters3 {
	s.TransactionIdentificationDetails = new(iso20022.TransactionTypeAndAdditionalParameters3)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmationV04) AddAdditionalParameters() *iso20022.AdditionalParameters18 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters18)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmationV04) AddTradeDetails() *iso20022.SecuritiesTradeDetails6 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails6)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmationV04) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmationV04) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmationV04) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount18 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount18)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmationV04) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails11 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails11)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmationV04) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction3 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction3)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmationV04) AddSettlementParameters() *iso20022.SettlementDetails60 {
	s.SettlementParameters = new(iso20022.SettlementDetails60)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmationV04) AddDeliveringSettlementParties() *iso20022.SettlementParties10 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmationV04) AddReceivingSettlementParties() *iso20022.SettlementParties10 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmationV04) AddCashParties() *iso20022.CashParties17 {
	s.CashParties = new(iso20022.CashParties17)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmationV04) AddSettledAmount() *iso20022.AmountAndDirection2 {
	s.SettledAmount = new(iso20022.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmationV04) AddOtherAmounts() *iso20022.OtherAmounts17 {
	s.OtherAmounts = new(iso20022.OtherAmounts17)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmationV04) AddOtherBusinessParties() *iso20022.OtherParties19 {
	s.OtherBusinessParties = new(iso20022.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmationV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
