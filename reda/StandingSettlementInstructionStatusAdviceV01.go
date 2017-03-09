package reda

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05800101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.058.001.01 Document"`
	Message *StandingSettlementInstructionStatusAdviceV01 `xml:"StgSttlmInstrStsAdvc"`
}

func (d *Document05800101) AddMessage() *StandingSettlementInstructionStatusAdviceV01 {
	d.Message = new(StandingSettlementInstructionStatusAdviceV01)
	return d.Message
}

// Scope
// The receiver of a StandingSettlementInstruction message sends the StandingSettlementInstructionStatusAdvice message to the instructing party (sender of the StandingSettlementInstruction message) to provide the status of a previously received StandingSettlementInstruction, StandingSettlementInstructionCancellation or StandingSettlementInstructionDeletion message.
//
// Usage
// The StandingSettlementInstructionStatusAdvice message is used to report one of the following statuses:
// -	a received status, or,
// -	an accepted status, or,
// -	a rejected status, or,
// -	a pending processing status, or,
// -	a proprietary status.
type StandingSettlementInstructionStatusAdviceV01 struct {

	// Date on which the SSI is effective.
	EffectiveDateDetails *iso20022.EffectiveDate1 `xml:"FctvDtDtls,omitempty"`

	// Unique and unambiguous master identification known to the sender (or its authorised agent) and receiver (or its authorised agent), below which the SSI will be lodged. This may be an account number or reference to a fund.
	// If no account or reference is available then “NONREF” must be specified.
	AccountIdentification []*iso20022.AccountIdentification26 `xml:"AcctId"`

	// Identifies the market for the standing settlement instruction.
	MarketIdentification *iso20022.MarketIdentificationOrCashPurpose1Choice `xml:"MktId"`

	// Settlement information that helps to identify the standing settlement  instruction, cancellation or deletion for which the status is sent.
	SettlementDetails *iso20022.PartyOrCurrency1Choice `xml:"SttlmDtls"`

	// Reference to a linked message that was previously received.
	RelatedMessageReference *iso20022.Max35Text `xml:"RltdMsgRef"`

	// Status of the standing settlement instruction, deletion or cancellation.
	ProcessingStatus *iso20022.ProcessingStatus43Choice `xml:"PrcgSts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddEffectiveDateDetails() *iso20022.EffectiveDate1 {
	s.EffectiveDateDetails = new(iso20022.EffectiveDate1)
	return s.EffectiveDateDetails
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddAccountIdentification() *iso20022.AccountIdentification26 {
	newValue := new(iso20022.AccountIdentification26)
	s.AccountIdentification = append(s.AccountIdentification, newValue)
	return newValue
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddMarketIdentification() *iso20022.MarketIdentificationOrCashPurpose1Choice {
	s.MarketIdentification = new(iso20022.MarketIdentificationOrCashPurpose1Choice)
	return s.MarketIdentification
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddSettlementDetails() *iso20022.PartyOrCurrency1Choice {
	s.SettlementDetails = new(iso20022.PartyOrCurrency1Choice)
	return s.SettlementDetails
}

func (s *StandingSettlementInstructionStatusAdviceV01) SetRelatedMessageReference(value string) {
	s.RelatedMessageReference = (*iso20022.Max35Text)(&value)
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddProcessingStatus() *iso20022.ProcessingStatus43Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus43Choice)
	return s.ProcessingStatus
}

func (s *StandingSettlementInstructionStatusAdviceV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
