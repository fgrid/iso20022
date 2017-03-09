package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02500105 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.025.001.05 Document"`
	Message *SecuritiesSettlementTransactionConfirmationV05 `xml:"SctiesSttlmTxConf"`
}

func (d *Document02500105) AddMessage() *SecuritiesSettlementTransactionConfirmationV05 {
	d.Message = new(SecuritiesSettlementTransactionConfirmationV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionConfirmation to an account owner to confirm the partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionConfirmationV05 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification15 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters17 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails31 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount28 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails70 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties17 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection36 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts18 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification15 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification15)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddAdditionalParameters() *iso20022.AdditionalParameters17 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters17)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddTradeDetails() *iso20022.SecuritiesTradeDetails31 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails31)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount28 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount28)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddSettlementParameters() *iso20022.SettlementDetails70 {
	s.SettlementParameters = new(iso20022.SettlementDetails70)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddDeliveringSettlementParties() *iso20022.SettlementParties11 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddReceivingSettlementParties() *iso20022.SettlementParties11 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddCashParties() *iso20022.CashParties17 {
	s.CashParties = new(iso20022.CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddSettledAmount() *iso20022.AmountAndDirection36 {
	s.SettledAmount = new(iso20022.AmountAndDirection36)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddOtherAmounts() *iso20022.OtherAmounts18 {
	s.OtherAmounts = new(iso20022.OtherAmounts18)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddOtherBusinessParties() *iso20022.OtherParties19 {
	s.OtherBusinessParties = new(iso20022.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
