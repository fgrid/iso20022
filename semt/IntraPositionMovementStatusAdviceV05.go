package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400105 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.05 Document"`
	Message *IntraPositionMovementStatusAdviceV05 `xml:"IntraPosMvmntStsAdvc"`
}

func (d *Document01400105) AddMessage() *IntraPositionMovementStatusAdviceV05 {
	d.Message = new(IntraPositionMovementStatusAdviceV05)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementStatusAdvice to an account owner to advise the status of a intra-position movement instruction previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
type IntraPositionMovementStatusAdviceV05 struct {

	// Unambiguous identification of a transaction as per the account owner (or the instructing party managing the account).
	TransactionIdentification *iso20022.TransactionIdentifications29 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *iso20022.IntraPositionProcessingStatus5Choice `xml:"PrcgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus16Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.IntraPositionDetails39 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementStatusAdviceV05) AddTransactionIdentification() *iso20022.TransactionIdentifications29 {
	i.TransactionIdentification = new(iso20022.TransactionIdentifications29)
	return i.TransactionIdentification
}

func (i *IntraPositionMovementStatusAdviceV05) AddProcessingStatus() *iso20022.IntraPositionProcessingStatus5Choice {
	i.ProcessingStatus = new(iso20022.IntraPositionProcessingStatus5Choice)
	return i.ProcessingStatus
}

func (i *IntraPositionMovementStatusAdviceV05) AddSettlementStatus() *iso20022.SettlementStatus16Choice {
	i.SettlementStatus = new(iso20022.SettlementStatus16Choice)
	return i.SettlementStatus
}

func (i *IntraPositionMovementStatusAdviceV05) AddTransactionDetails() *iso20022.IntraPositionDetails39 {
	i.TransactionDetails = new(iso20022.IntraPositionDetails39)
	return i.TransactionDetails
}

func (i *IntraPositionMovementStatusAdviceV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
