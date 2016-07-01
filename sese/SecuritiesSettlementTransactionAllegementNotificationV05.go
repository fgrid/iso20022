package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02800105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.028.001.05 Document"`
	Message *SecuritiesSettlementTransactionAllegementNotificationV05 `xml:"SctiesSttlmTxAllgmtNtfctn"`
}

func (d *Document02800105) AddMessage() *SecuritiesSettlementTransactionAllegementNotificationV05 {
	d.Message = new(SecuritiesSettlementTransactionAllegementNotificationV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementNotification to an account owner to advise the account owner that a counterparty has alleged an instruction against the account owner's account at the account servicer and that the account servicer could not find the corresponding instruction of the account owner.
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
type SecuritiesSettlementTransactionAllegementNotificationV05 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *iso20022.SettlementTypeAndAdditionalParameters12 `xml:"SttlmTpAndAddtlParams"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *iso20022.Identification14 `xml:"MktInfrstrctrTxId,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails54 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount42 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails29 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails99 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies cash parties in the framework of a corporate action event.
	CashParties *iso20022.CashParties25 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *iso20022.AmountAndDirection48 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts32 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties28 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesSettlementTransactionAllegementNotificationV05) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddSettlementTypeAndAdditionalParameters() *iso20022.SettlementTypeAndAdditionalParameters12 {
	s.SettlementTypeAndAdditionalParameters = new(iso20022.SettlementTypeAndAdditionalParameters12)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddMarketInfrastructureTransactionIdentification() *iso20022.Identification14 {
	s.MarketInfrastructureTransactionIdentification = new(iso20022.Identification14)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddTradeDetails() *iso20022.SecuritiesTradeDetails54 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails54)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount42 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount42)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails29 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails29)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddSettlementParameters() *iso20022.SettlementDetails99 {
	s.SettlementParameters = new(iso20022.SettlementDetails99)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddDeliveringSettlementParties() *iso20022.SettlementParties36 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddReceivingSettlementParties() *iso20022.SettlementParties36 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddCashParties() *iso20022.CashParties25 {
	s.CashParties = new(iso20022.CashParties25)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddSettlementAmount() *iso20022.AmountAndDirection48 {
	s.SettlementAmount = new(iso20022.AmountAndDirection48)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddOtherAmounts() *iso20022.OtherAmounts32 {
	s.OtherAmounts = new(iso20022.OtherAmounts32)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddOtherBusinessParties() *iso20022.OtherParties28 {
	s.OtherBusinessParties = new(iso20022.OtherParties28)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

