package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500106 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.001.06 Document"`
	Message *SecuritiesFinancingConfirmationV06 `xml:"SctiesFincgConf"`
}

func (d *Document03500106) AddMessage() *SecuritiesFinancingConfirmationV06 {
	d.Message = new(SecuritiesFinancingConfirmationV06)
	return d.Message
}

// Scope
// A securities financing transaction account servicer sends a SecuritiesFinancingConfirmation to an account owner to confirm or advise of the partial or full settlement of the opening or closing leg of a securities financing transaction.
//
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
//
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesFinancingConfirmationV06 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *iso20022.TransactionTypeAndAdditionalParameters10 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters24 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails55 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount40 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails28 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails96 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties26 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection46 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts31 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingConfirmationV06) AddTransactionIdentificationDetails() *iso20022.TransactionTypeAndAdditionalParameters10 {
	s.TransactionIdentificationDetails = new(iso20022.TransactionTypeAndAdditionalParameters10)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmationV06) AddAdditionalParameters() *iso20022.AdditionalParameters24 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters24)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmationV06) AddTradeDetails() *iso20022.SecuritiesTradeDetails55 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails55)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmationV06) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmationV06) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmationV06) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount40 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount40)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmationV06) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails28 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails28)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmationV06) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmationV06) AddSettlementParameters() *iso20022.SettlementDetails96 {
	s.SettlementParameters = new(iso20022.SettlementDetails96)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmationV06) AddDeliveringSettlementParties() *iso20022.SettlementParties36 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmationV06) AddReceivingSettlementParties() *iso20022.SettlementParties36 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmationV06) AddCashParties() *iso20022.CashParties26 {
	s.CashParties = new(iso20022.CashParties26)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmationV06) AddSettledAmount() *iso20022.AmountAndDirection46 {
	s.SettledAmount = new(iso20022.AmountAndDirection46)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmationV06) AddOtherAmounts() *iso20022.OtherAmounts31 {
	s.OtherAmounts = new(iso20022.OtherAmounts31)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmationV06) AddOtherBusinessParties() *iso20022.OtherParties27 {
	s.OtherBusinessParties = new(iso20022.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmationV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
