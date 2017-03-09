package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03200207 struct {
	XMLName xml.Name                                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.032.002.07 Document"`
	Message *SecuritiesSettlementTransactionGenerationNotification002V07 `xml:"SctiesSttlmTxGnrtnNtfctn"`
}

func (d *Document03200207) AddMessage() *SecuritiesSettlementTransactionGenerationNotification002V07 {
	d.Message = new(SecuritiesSettlementTransactionGenerationNotification002V07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionGenerationNotification to an account owner to advise the account owner of a securities settlement transaction that has been generated/created by the account servicer for the account owner. The reason for creation can range from the automatic transformation of pending settlement instructions following a corporate event to the recovery of an account owner transactions' database initiated by its account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionGenerationNotification002V07 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification24 `xml:"TxIdDtls"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages43 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails70 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails []*iso20022.QuantityAndAccount56 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails137 `xml:"SttlmParams"`

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

	// Specifies the reason why the transaction was generated.
	GeneratedReason []*iso20022.GeneratedReason6 `xml:"GnrtdRsn,omitempty"`

	// Status and reason of the transaction.
	StatusAndReason *iso20022.StatusAndReason29 `xml:"StsAndRsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification24 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification24)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddLinkages() *iso20022.Linkages43 {
	newValue := new(iso20022.Linkages43)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddTradeDetails() *iso20022.SecuritiesTradeDetails70 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails70)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount56 {
	newValue := new(iso20022.QuantityAndAccount56)
	s.QuantityAndAccountDetails = append(s.QuantityAndAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddSettlementParameters() *iso20022.SettlementDetails137 {
	s.SettlementParameters = new(iso20022.SettlementDetails137)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddDeliveringSettlementParties() *iso20022.SettlementParties44 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddReceivingSettlementParties() *iso20022.SettlementParties44 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddCashParties() *iso20022.CashParties30 {
	s.CashParties = new(iso20022.CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddSettlementAmount() *iso20022.AmountAndDirection60 {
	s.SettlementAmount = new(iso20022.AmountAndDirection60)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddOtherAmounts() *iso20022.OtherAmounts35 {
	s.OtherAmounts = new(iso20022.OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddOtherBusinessParties() *iso20022.OtherParties29 {
	s.OtherBusinessParties = new(iso20022.OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddGeneratedReason() *iso20022.GeneratedReason6 {
	newValue := new(iso20022.GeneratedReason6)
	s.GeneratedReason = append(s.GeneratedReason, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddStatusAndReason() *iso20022.StatusAndReason29 {
	s.StatusAndReason = new(iso20022.StatusAndReason29)
	return s.StatusAndReason
}

func (s *SecuritiesSettlementTransactionGenerationNotification002V07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
