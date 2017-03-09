package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05200101 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.052.001.01 Document"`
	Message *RedemptionOrderConfirmationAmendmentV01 `xml:"RedOrdrConfAmdmntV01"`
}

func (d *Document05200101) AddMessage() *RedemptionOrderConfirmationAmendmentV01 {
	d.Message = new(RedemptionOrderConfirmationAmendmentV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the RedemptionOrderConfirmationAmendment message to the instructing party, for example, an investment manager or its authorised representative to amend a previously sent RedemptionOrderConfirmation message.
// Usage
// The RedemptionOrderConfirmationAmendment message is used to amend one or more previously sent redemption order confirmations.
// Each individual order confirmation amendment specified is identified in DealReference. The reference of the original individual order is specified in OrderReference.
// The message identification of the RedemptionOrder message in which the individual orders were conveyed may also be quoted in RelatedReference. The message identification of the RedemptionOrderConfirmation message in which the original order confirmations were conveyed may also be quoted in PreviousReference.
type RedemptionOrderConfirmationAmendmentV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// General information related to the execution of investment fund orders.
	MultipleExecutionDetails *iso20022.RedemptionMultipleExecution3 `xml:"MltplExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	r.RelatedReference = new(iso20022.AdditionalReference3)
	return r.RelatedReference
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddMultipleExecutionDetails() *iso20022.RedemptionMultipleExecution3 {
	r.MultipleExecutionDetails = new(iso20022.RedemptionMultipleExecution3)
	return r.MultipleExecutionDetails
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddRelatedPartyDetails() *iso20022.Intermediary9 {
	newValue := new(iso20022.Intermediary9)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}

func (r *RedemptionOrderConfirmationAmendmentV01) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
