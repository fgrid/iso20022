package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03400206 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.034.002.06 Document"`
	Message *SecuritiesFinancingStatusAdvice002V06 `xml:"SctiesFincgStsAdvc"`
}

func (d *Document03400206) AddMessage() *SecuritiesFinancingStatusAdvice002V06 {
	d.Message = new(SecuritiesFinancingStatusAdvice002V06)
	return d.Message
}

// Scope
// An securities financing transaction account servicer sends a SecuritiesFinancingStatusAdvice to an account owner to advise the status of a securities financing transaction previously instructed by the account owner.
// The status advice may be sent as a response to the request of the account owner or not.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
// 
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesFinancingStatusAdvice002V06 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications35 `xml:"TxId"`

	// Processing status of the transaction.
	ProcessingStatus *iso20022.ProcessingStatus57Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *iso20022.MatchingStatus29Choice `xml:"MtchgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *iso20022.MatchingStatus29Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus21Choice `xml:"SttlmSts,omitempty"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *iso20022.RepoCallRequestStatus9Choice `xml:"RepoCallReqSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.SecuritiesFinancingTransactionDetails33 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesFinancingStatusAdvice002V06) AddTransactionIdentification() *iso20022.TransactionIdentifications35 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications35)
	return s.TransactionIdentification
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddProcessingStatus() *iso20022.ProcessingStatus57Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus57Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddMatchingStatus() *iso20022.MatchingStatus29Choice {
	s.MatchingStatus = new(iso20022.MatchingStatus29Choice)
	return s.MatchingStatus
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddInferredMatchingStatus() *iso20022.MatchingStatus29Choice {
	s.InferredMatchingStatus = new(iso20022.MatchingStatus29Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddSettlementStatus() *iso20022.SettlementStatus21Choice {
	s.SettlementStatus = new(iso20022.SettlementStatus21Choice)
	return s.SettlementStatus
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddRepoCallRequestStatus() *iso20022.RepoCallRequestStatus9Choice {
	s.RepoCallRequestStatus = new(iso20022.RepoCallRequestStatus9Choice)
	return s.RepoCallRequestStatus
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddTransactionDetails() *iso20022.SecuritiesFinancingTransactionDetails33 {
	s.TransactionDetails = new(iso20022.SecuritiesFinancingTransactionDetails33)
	return s.TransactionDetails
}

func (s *SecuritiesFinancingStatusAdvice002V06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

