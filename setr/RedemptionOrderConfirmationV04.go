package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600104 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.006.001.04 Document"`
	Message *RedemptionOrderConfirmationV04 `xml:"RedOrdrConf"`
}

func (d *Document00600104) AddMessage() *RedemptionOrderConfirmationV04 {
	d.Message = new(RedemptionOrderConfirmationV04)
	return d.Message
}

// Scope
// The RedemptionOrderConfirmation message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to confirm the details of execution for a previously received RedemptionOrder message.
// Usage
// The RedemptionOrderConfirmation message is used to confirm the execution of one or more individual orders.
// A RedemptionOrder message containing more than one individual order may be responded to by more than one RedemptionOrderConfirmation message, as the valuation cycle of the financial instruments in each individual order may be different. When a RedemptionOrderConfirmation message contains fewer confirmations that originally instructed in the original RedemptionOrder message, there is no specification indication in the confirmation for this. Reconciliation must be based on the references.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the RedemptionOrder message in which the individual order was conveyed may also be quoted in RelatedReference but this is not recommended.
// A RedemptionOrder must in all cases be responded to by a RedemptionOrderConfirmation message and in no circumstances by a RedemptionBulkOrderConfirmation message.
// If the executing party needs to confirm one or more redemption orders for the same financial instrument then a RedemptionBulkOrderConfirmation message must be used.
// When the message is used to convey a confirmation amendment/s, the AmendmentIndicator must be present with the value ‘true’ or ‘1’. When this is the case, the message must only contain a confirmation amendment/s and not contain both a confirmation amendment/s and a ‘new’ confirmation/s.
type RedemptionOrderConfirmationV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// General information related to the execution of investment fund orders.
	MultipleExecutionDetails *iso20022.RedemptionMultipleExecution5 `xml:"MltplExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionOrderConfirmationV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionOrderConfirmationV04) AddPoolReference() *iso20022.AdditionalReference9 {
	r.PoolReference = new(iso20022.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionOrderConfirmationV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionOrderConfirmationV04) AddRelatedReference() *iso20022.AdditionalReference8 {
	r.RelatedReference = new(iso20022.AdditionalReference8)
	return r.RelatedReference
}

func (r *RedemptionOrderConfirmationV04) AddMultipleExecutionDetails() *iso20022.RedemptionMultipleExecution5 {
	r.MultipleExecutionDetails = new(iso20022.RedemptionMultipleExecution5)
	return r.MultipleExecutionDetails
}

func (r *RedemptionOrderConfirmationV04) AddCopyDetails() *iso20022.CopyInformation4 {
	r.CopyDetails = new(iso20022.CopyInformation4)
	return r.CopyDetails
}

func (r *RedemptionOrderConfirmationV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
