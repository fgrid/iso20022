package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02700205 struct {
	XMLName xml.Name                                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.002.05 Document"`
	Message *SecuritiesTransactionCancellationRequestStatusAdvice002V05 `xml:"SctiesTxCxlReqStsAdvc"`
}

func (d *Document02700205) AddMessage() *SecuritiesTransactionCancellationRequestStatusAdvice002V05 {
	d.Message = new(SecuritiesTransactionCancellationRequestStatusAdvice002V05)
	return d.Message
}

// Scope
// An account servicer sends an SecuritiesTransactionCancellationRequestStatusAdvice to an account owner to advise the status of a securities transaction cancellation request previously sent by the account owner.
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
type SecuritiesTransactionCancellationRequestStatusAdvice002V05 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	CancellationRequestReference *iso20022.Identification16 `xml:"CxlReqRef"`

	// Unambiguous identification of the transaction as known by the account servicer.
	TransactionIdentification *iso20022.TransactionIdentifications39 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus65Choice `xml:"PrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails88 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequestStatusAdvice002V05) AddCancellationRequestReference() *iso20022.Identification16 {
	s.CancellationRequestReference = new(iso20022.Identification16)
	return s.CancellationRequestReference
}

func (s *SecuritiesTransactionCancellationRequestStatusAdvice002V05) AddTransactionIdentification() *iso20022.TransactionIdentifications39 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications39)
	return s.TransactionIdentification
}

func (s *SecuritiesTransactionCancellationRequestStatusAdvice002V05) AddProcessingStatus() *iso20022.ProcessingStatus65Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus65Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesTransactionCancellationRequestStatusAdvice002V05) AddTransactionDetails() *iso20022.TransactionDetails88 {
	s.TransactionDetails = new(iso20022.TransactionDetails88)
	return s.TransactionDetails
}

func (s *SecuritiesTransactionCancellationRequestStatusAdvice002V05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
