package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500206 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.002.06 Document"`
	Message *SecuritiesFinancingConfirmation002V06 `xml:"SctiesFincgConf"`
}

func (d *Document03500206) AddMessage() *SecuritiesFinancingConfirmation002V06 {
	d.Message = new(SecuritiesFinancingConfirmation002V06)
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
type SecuritiesFinancingConfirmation002V06 struct {

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *iso20022.TransactionTypeAndAdditionalParameters12 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters26 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails58 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount46 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails30 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails104 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties30 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection60 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts34 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesFinancingConfirmation002V06) AddTransactionIdentificationDetails() *iso20022.TransactionTypeAndAdditionalParameters12 {
	s.TransactionIdentificationDetails = new(iso20022.TransactionTypeAndAdditionalParameters12)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmation002V06) AddAdditionalParameters() *iso20022.AdditionalParameters26 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters26)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmation002V06) AddTradeDetails() *iso20022.SecuritiesTradeDetails58 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails58)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmation002V06) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmation002V06) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmation002V06) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount46 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount46)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmation002V06) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails30 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails30)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmation002V06) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmation002V06) AddSettlementParameters() *iso20022.SettlementDetails104 {
	s.SettlementParameters = new(iso20022.SettlementDetails104)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmation002V06) AddDeliveringSettlementParties() *iso20022.SettlementParties44 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmation002V06) AddReceivingSettlementParties() *iso20022.SettlementParties44 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmation002V06) AddCashParties() *iso20022.CashParties30 {
	s.CashParties = new(iso20022.CashParties30)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmation002V06) AddSettledAmount() *iso20022.AmountAndDirection60 {
	s.SettledAmount = new(iso20022.AmountAndDirection60)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmation002V06) AddOtherAmounts() *iso20022.OtherAmounts34 {
	s.OtherAmounts = new(iso20022.OtherAmounts34)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmation002V06) AddOtherBusinessParties() *iso20022.OtherParties29 {
	s.OtherBusinessParties = new(iso20022.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmation002V06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

