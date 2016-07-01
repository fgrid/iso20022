package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.035.001.01 Document"`
	Message *SecuritiesFinancingConfirmationV01 `xml:"SctiesFincgConf"`
}

func (d *Document03500101) AddMessage() *SecuritiesFinancingConfirmationV01 {
	d.Message = new(SecuritiesFinancingConfirmationV01)
	return d.Message
}

// SCOPE
// A securities financing transaction account servicer sends a SecuritiesFinancingConfirmation to an account owner to confirm or advise of the partial or full settlement of the opening or closing leg of a securities financing transaction. 
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or 
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
// 
// USAGE
// The message may also be used to: 
// - re-send a message previously sent (the sub-function of the message is Duplicate) 
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy) 
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// 
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingConfirmationV01 struct {

	// Information that unambiguously identifies a SecuritiesFinancingConfirmation message as known by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
	TransactionIdentificationDetails *iso20022.TransactionTypeAndAdditionalParameters3 `xml:"TxIdDtls"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters2 `xml:"AddtlParams,omitempty"`

	// Details of the securities financing deal.
	TradeDetails *iso20022.SecuritiesTradeDetails6 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount6 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *iso20022.SecuritiesFinancingTransactionDetails3 `xml:"SctiesFincgDtls,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction1 `xml:"StgSttlmInstrDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails9 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties3 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection2 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts4 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties2 `xml:"OthrBizPties,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`

}


func (s *SecuritiesFinancingConfirmationV01) AddIdentification() *iso20022.DocumentIdentification11 {
	s.Identification = new(iso20022.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesFinancingConfirmationV01) AddTransactionIdentificationDetails() *iso20022.TransactionTypeAndAdditionalParameters3 {
	s.TransactionIdentificationDetails = new(iso20022.TransactionTypeAndAdditionalParameters3)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesFinancingConfirmationV01) AddAdditionalParameters() *iso20022.AdditionalParameters2 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters2)
	return s.AdditionalParameters
}

func (s *SecuritiesFinancingConfirmationV01) AddTradeDetails() *iso20022.SecuritiesTradeDetails6 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails6)
	return s.TradeDetails
}

func (s *SecuritiesFinancingConfirmationV01) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingConfirmationV01) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesFinancingConfirmationV01) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount6 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount6)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesFinancingConfirmationV01) AddSecuritiesFinancingDetails() *iso20022.SecuritiesFinancingTransactionDetails3 {
	s.SecuritiesFinancingDetails = new(iso20022.SecuritiesFinancingTransactionDetails3)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesFinancingConfirmationV01) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction1 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction1)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesFinancingConfirmationV01) AddSettlementParameters() *iso20022.SettlementDetails9 {
	s.SettlementParameters = new(iso20022.SettlementDetails9)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingConfirmationV01) AddDeliveringSettlementParties() *iso20022.SettlementParties5 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingConfirmationV01) AddReceivingSettlementParties() *iso20022.SettlementParties5 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingConfirmationV01) AddCashParties() *iso20022.CashParties3 {
	s.CashParties = new(iso20022.CashParties3)
	return s.CashParties
}

func (s *SecuritiesFinancingConfirmationV01) AddSettledAmount() *iso20022.AmountAndDirection2 {
	s.SettledAmount = new(iso20022.AmountAndDirection2)
	return s.SettledAmount
}

func (s *SecuritiesFinancingConfirmationV01) AddOtherAmounts() *iso20022.OtherAmounts4 {
	s.OtherAmounts = new(iso20022.OtherAmounts4)
	return s.OtherAmounts
}

func (s *SecuritiesFinancingConfirmationV01) AddOtherBusinessParties() *iso20022.OtherParties2 {
	s.OtherBusinessParties = new(iso20022.OtherParties2)
	return s.OtherBusinessParties
}

func (s *SecuritiesFinancingConfirmationV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	s.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesFinancingConfirmationV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	s.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesFinancingConfirmationV01) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

