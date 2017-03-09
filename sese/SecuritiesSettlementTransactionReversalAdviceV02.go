package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02600102 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.026.001.02 Document"`
	Message *SecuritiesSettlementTransactionReversalAdviceV02 `xml:"SctiesSttlmTxRvslAdvc"`
}

func (d *Document02600102) AddMessage() *SecuritiesSettlementTransactionReversalAdviceV02 {
	d.Message = new(SecuritiesSettlementTransactionReversalAdviceV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionReversalAdvice to an account owner to reverse the confirmation of a partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
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
type SecuritiesSettlementTransactionReversalAdviceV02 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification9 `xml:"TxIdDtls"`

	// Reference to the unambiguous identification of the confirmation as per the account servicer.
	ConfirmationReference *iso20022.Identification1 `xml:"ConfRef"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters4 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails2 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount15 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails26 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction5 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties14 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties14 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties9 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts9 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties12 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification9 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification9)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddConfirmationReference() *iso20022.Identification1 {
	s.ConfirmationReference = new(iso20022.Identification1)
	return s.ConfirmationReference
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddAdditionalParameters() *iso20022.AdditionalParameters4 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters4)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddTradeDetails() *iso20022.SecuritiesTradeDetails2 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails2)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes20 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes20)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount15 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount15)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddSettlementParameters() *iso20022.SettlementDetails26 {
	s.SettlementParameters = new(iso20022.SettlementDetails26)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction5 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction5)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddDeliveringSettlementParties() *iso20022.SettlementParties14 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties14)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddReceivingSettlementParties() *iso20022.SettlementParties14 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties14)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddCashParties() *iso20022.CashParties9 {
	s.CashParties = new(iso20022.CashParties9)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddSettledAmount() *iso20022.AmountAndDirection2 {
	s.SettledAmount = new(iso20022.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddOtherAmounts() *iso20022.OtherAmounts9 {
	s.OtherAmounts = new(iso20022.OtherAmounts9)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddOtherBusinessParties() *iso20022.OtherParties12 {
	s.OtherBusinessParties = new(iso20022.OtherParties12)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
