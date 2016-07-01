package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.044.001.01 Document"`
	Message *SecuritiesTradeConfirmationStatusAdviceV01 `xml:"SctiesTradConfStsAdvc"`
}

func (d *Document04400101) AddMessage() *SecuritiesTradeConfirmationStatusAdviceV01 {
	d.Message = new(SecuritiesTradeConfirmationStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent from Central Matching Utility (CMU) to an executing party or an instructing party to advise the status of the SecuritiesTradeConfirmation message previously sent by the party. The status may be a processing, pending processing, affirmed or disaffirmed, cancel or replacement by an instructing party, a custodian or an affirming party, internal matching, and/or matching status.
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager or of other categories of investors. The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: In central matching the Initiator is the Central Matching Utility.
// Respondent: no response is needed by the recipient of the message.
type SecuritiesTradeConfirmationStatusAdviceV01 struct {

	// Information that unambiguously identifies an SecuritiesTradeConfirmationStatusAdvice message as known by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.TransactiontIdentification4 `xml:"Id"`

	// Link to another transaction that must be processed after, before or at the same time.
	References []*iso20022.Linkages18 `xml:"Refs"`

	// Provides details on the affitrmation status of a trade.
	AffirmationStatus *iso20022.AffirmationStatus6Choice `xml:"AffirmSts,omitempty"`

	// Provides the processing status of a trade.
	ProcessingStatus *iso20022.ProcessingStatus17Choice `xml:"PrcgSts,omitempty"`

	// Provides details on the matching status of a trade.
	MatchingStatus *iso20022.MatchingStatus9Choice `xml:"MtchgSts,omitempty"`

	// Provides the replacement processing status of a trade.
	ReplacementProcessingStatus *iso20022.ReplacementProcessingStatus7Choice `xml:"RplcmntPrcgSts,omitempty"`

	// Provides details on the cancellation status of a trade.
	CancellationProcessingStatus *iso20022.CancellationProcessingStatus6Choice `xml:"CxlPrcgSts,omitempty"`

	// Details of the trading party.
	PartyTradingDetails *iso20022.Order16 `xml:"PtyTradgDtls,omitempty"`

	// Details of the trading counterparty.
	CounterpartyTradingDetails *iso20022.Order16 `xml:"CtrPtyTradgDtls,omitempty"`

	// Parties used for acting parties that applies either to the whole message or to individual sides.
	ConfirmationParties []*iso20022.ConfirmationParties4 `xml:"ConfPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties23 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties23 `xml:"RcvgSttlmPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddIdentification() *iso20022.TransactiontIdentification4 {
	s.Identification = new(iso20022.TransactiontIdentification4)
	return s.Identification
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddReferences() *iso20022.Linkages18 {
	newValue := new (iso20022.Linkages18)
	s.References = append(s.References, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddAffirmationStatus() *iso20022.AffirmationStatus6Choice {
	s.AffirmationStatus = new(iso20022.AffirmationStatus6Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddProcessingStatus() *iso20022.ProcessingStatus17Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus17Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddMatchingStatus() *iso20022.MatchingStatus9Choice {
	s.MatchingStatus = new(iso20022.MatchingStatus9Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddReplacementProcessingStatus() *iso20022.ReplacementProcessingStatus7Choice {
	s.ReplacementProcessingStatus = new(iso20022.ReplacementProcessingStatus7Choice)
	return s.ReplacementProcessingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddCancellationProcessingStatus() *iso20022.CancellationProcessingStatus6Choice {
	s.CancellationProcessingStatus = new(iso20022.CancellationProcessingStatus6Choice)
	return s.CancellationProcessingStatus
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddPartyTradingDetails() *iso20022.Order16 {
	s.PartyTradingDetails = new(iso20022.Order16)
	return s.PartyTradingDetails
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddCounterpartyTradingDetails() *iso20022.Order16 {
	s.CounterpartyTradingDetails = new(iso20022.Order16)
	return s.CounterpartyTradingDetails
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddConfirmationParties() *iso20022.ConfirmationParties4 {
	newValue := new (iso20022.ConfirmationParties4)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddDeliveringSettlementParties() *iso20022.SettlementParties23 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties23)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddReceivingSettlementParties() *iso20022.SettlementParties23 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties23)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeConfirmationStatusAdviceV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

