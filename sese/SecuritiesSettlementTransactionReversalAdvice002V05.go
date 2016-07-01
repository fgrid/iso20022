package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02600205 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.026.002.05 Document"`
	Message *SecuritiesSettlementTransactionReversalAdvice002V05 `xml:"SctiesSttlmTxRvslAdvc"`
}

func (d *Document02600205) AddMessage() *SecuritiesSettlementTransactionReversalAdvice002V05 {
	d.Message = new(SecuritiesSettlementTransactionReversalAdvice002V05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionReversalAdvice to an account owner to reverse the confirmation of a partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// 
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionReversalAdvice002V05 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification24 `xml:"TxIdDtls"`

	// Reference to the unambiguous identification of the confirmation as per the account servicer.
	ConfirmationReference *iso20022.Identification16 `xml:"ConfRef"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters28 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails62 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount51 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails110 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties32 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection60 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts38 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification24 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification24)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddConfirmationReference() *iso20022.Identification16 {
	s.ConfirmationReference = new(iso20022.Identification16)
	return s.ConfirmationReference
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddAdditionalParameters() *iso20022.AdditionalParameters28 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters28)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddTradeDetails() *iso20022.SecuritiesTradeDetails62 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails62)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount51 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount51)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddSettlementParameters() *iso20022.SettlementDetails110 {
	s.SettlementParameters = new(iso20022.SettlementDetails110)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddDeliveringSettlementParties() *iso20022.SettlementParties44 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddReceivingSettlementParties() *iso20022.SettlementParties44 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddCashParties() *iso20022.CashParties32 {
	s.CashParties = new(iso20022.CashParties32)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddSettledAmount() *iso20022.AmountAndDirection60 {
	s.SettledAmount = new(iso20022.AmountAndDirection60)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddOtherAmounts() *iso20022.OtherAmounts38 {
	s.OtherAmounts = new(iso20022.OtherAmounts38)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddOtherBusinessParties() *iso20022.OtherParties29 {
	s.OtherBusinessParties = new(iso20022.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdvice002V05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

