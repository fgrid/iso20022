package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.008.001.01 Document"`
	Message *ReversalOfTransferInConfirmation `xml:"sese.008.001.01"`
}

func (d *Document00800101) AddMessage() *ReversalOfTransferInConfirmation {
	d.Message = new(ReversalOfTransferInConfirmation)
	return d.Message
}

// Scope
// The ReversalOfTransferInConfirmation message is sent by an executing party to the instructing party or the instructing party's designated agent.
// This message is used to reverse a TransferInConfirmation that was previously sent by the instructing party.
// Usage
// The ReversalOfTransferInConfirmation message is sent by an executing party to reverse a previously sent TransferInConfirmation.
// This message must contain the reference of the message to be reversed. The message may also contain all the details of the reversed message, but this is not recommended.
type ReversalOfTransferInConfirmation struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Copy of the transfer in confirmation to reverse.
	TransferInConfirmationToBeReversed *iso20022.TransferIn1 `xml:"TrfInConfToBeRvsd,omitempty"`
}

func (r *ReversalOfTransferInConfirmation) AddPreviousReference() *iso20022.AdditionalReference2 {
	r.PreviousReference = new(iso20022.AdditionalReference2)
	return r.PreviousReference
}

func (r *ReversalOfTransferInConfirmation) AddPoolReference() *iso20022.AdditionalReference2 {
	r.PoolReference = new(iso20022.AdditionalReference2)
	return r.PoolReference
}

func (r *ReversalOfTransferInConfirmation) AddRelatedReference() *iso20022.AdditionalReference2 {
	r.RelatedReference = new(iso20022.AdditionalReference2)
	return r.RelatedReference
}

func (r *ReversalOfTransferInConfirmation) AddTransferInConfirmationToBeReversed() *iso20022.TransferIn1 {
	r.TransferInConfirmationToBeReversed = new(iso20022.TransferIn1)
	return r.TransferInConfirmationToBeReversed
}
