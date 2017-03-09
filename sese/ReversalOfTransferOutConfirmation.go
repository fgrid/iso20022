package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.004.001.01 Document"`
	Message *ReversalOfTransferOutConfirmation `xml:"sese.004.001.01"`
}

func (d *Document00400101) AddMessage() *ReversalOfTransferOutConfirmation {
	d.Message = new(ReversalOfTransferOutConfirmation)
	return d.Message
}

// Scope
// The ReversalOfTransferOutConfirmation message is sent by an executing party to the instructing party or the instructing party's designated agent. This message is used to reverse a TransferOutConfirmation that was previously sent by the instructing party.
// Usage
// The ReversalOfTransferOutConfirmation message is sent by an executing party to reverse a previously sent TransferOutConfirmation.
// This message must contain the reference of the message to be reversed. The message may also contain all the details of the reversed message, but this is not recommended.
type ReversalOfTransferOutConfirmation struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Copy of the transfer out confirmation to reverse.
	TransferOutConfirmationToBeReversed *iso20022.TransferOut1 `xml:"TrfOutConfToBeRvsd,omitempty"`
}

func (r *ReversalOfTransferOutConfirmation) AddPreviousReference() *iso20022.AdditionalReference2 {
	r.PreviousReference = new(iso20022.AdditionalReference2)
	return r.PreviousReference
}

func (r *ReversalOfTransferOutConfirmation) AddPoolReference() *iso20022.AdditionalReference2 {
	r.PoolReference = new(iso20022.AdditionalReference2)
	return r.PoolReference
}

func (r *ReversalOfTransferOutConfirmation) AddRelatedReference() *iso20022.AdditionalReference2 {
	r.RelatedReference = new(iso20022.AdditionalReference2)
	return r.RelatedReference
}

func (r *ReversalOfTransferOutConfirmation) AddTransferOutConfirmationToBeReversed() *iso20022.TransferOut1 {
	r.TransferOutConfirmationToBeReversed = new(iso20022.TransferOut1)
	return r.TransferOutConfirmationToBeReversed
}
