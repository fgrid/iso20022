package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02400103 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.024.001.03 Document"`
	Message *SecuritiesSettlementTransactionStatusAdviceV03 `xml:"SctiesSttlmTxStsAdvc"`
}

func (d *Document02400103) AddMessage() *SecuritiesSettlementTransactionStatusAdviceV03 {
	d.Message = new(SecuritiesSettlementTransactionStatusAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionStatusAdvice to an account owner to advise the status of a securities settlement transaction instruction previously sent by the account owner or the status of a settlement transaction existing in the books of the servicer for the account of the owner. The status may be a processing, pending processing, internal matching, matching and/or settlement status.
// The status advice may be sent as a response to the request of the account owner or not.
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
type SecuritiesSettlementTransactionStatusAdviceV03 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications16 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *iso20022.ProcessingStatus19Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *iso20022.MatchingStatus7Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *iso20022.MatchingStatus7Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus7Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails36 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddTransactionIdentification() *iso20022.TransactionIdentifications16 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications16)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddProcessingStatus() *iso20022.ProcessingStatus19Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus19Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddInferredMatchingStatus() *iso20022.MatchingStatus7Choice {
	s.InferredMatchingStatus = new(iso20022.MatchingStatus7Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddMatchingStatus() *iso20022.MatchingStatus7Choice {
	s.MatchingStatus = new(iso20022.MatchingStatus7Choice)
	return s.MatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddSettlementStatus() *iso20022.SettlementStatus7Choice {
	s.SettlementStatus = new(iso20022.SettlementStatus7Choice)
	return s.SettlementStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddTransactionDetails() *iso20022.TransactionDetails36 {
	s.TransactionDetails = new(iso20022.TransactionDetails36)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionStatusAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
