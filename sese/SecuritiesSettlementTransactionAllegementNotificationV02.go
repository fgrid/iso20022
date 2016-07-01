package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02800102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.028.001.02 Document"`
	Message *SecuritiesSettlementTransactionAllegementNotificationV02 `xml:"SctiesSttlmTxAllgmtNtfctn"`
}

func (d *Document02800102) AddMessage() *SecuritiesSettlementTransactionAllegementNotificationV02 {
	d.Message = new(SecuritiesSettlementTransactionAllegementNotificationV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementNotification to an account owner to advise the account owner that a counterparty has alleged an instruction against the account owner's account at the account servicer and that the account servicer could not find the corresponding instruction of the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// using the relevant elements in the Business Application Header.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionAllegementNotificationV02 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *iso20022.SettlementTypeAndAdditionalParameters2 `xml:"SttlmTpAndAddtlParams"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *iso20022.Identification1 `xml:"MktInfrstrctrTxId,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails18 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes15 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount13 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails7 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails25 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies cash parties in the framework of a corporate action event.
	CashParties *iso20022.CashParties11 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *iso20022.AmountAndDirection22 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts8 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties11 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesSettlementTransactionAllegementNotificationV02) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddSettlementTypeAndAdditionalParameters() *iso20022.SettlementTypeAndAdditionalParameters2 {
	s.SettlementTypeAndAdditionalParameters = new(iso20022.SettlementTypeAndAdditionalParameters2)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddMarketInfrastructureTransactionIdentification() *iso20022.Identification1 {
	s.MarketInfrastructureTransactionIdentification = new(iso20022.Identification1)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddTradeDetails() *iso20022.SecuritiesTradeDetails18 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails18)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes15 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes15)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount13 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount13)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails7 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails7)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddSettlementParameters() *iso20022.SettlementDetails25 {
	s.SettlementParameters = new(iso20022.SettlementDetails25)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddDeliveringSettlementParties() *iso20022.SettlementParties11 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddReceivingSettlementParties() *iso20022.SettlementParties11 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddCashParties() *iso20022.CashParties11 {
	s.CashParties = new(iso20022.CashParties11)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddSettlementAmount() *iso20022.AmountAndDirection22 {
	s.SettlementAmount = new(iso20022.AmountAndDirection22)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddOtherAmounts() *iso20022.OtherAmounts8 {
	s.OtherAmounts = new(iso20022.OtherAmounts8)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddOtherBusinessParties() *iso20022.OtherParties11 {
	s.OtherBusinessParties = new(iso20022.OtherParties11)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionAllegementNotificationV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

