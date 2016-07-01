package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.003.001.02 Document"`
	Message *RedemptionBulkOrderConfirmationV02 `xml:"setr.003.001.02"`
}

func (d *Document00300102) AddMessage() *RedemptionBulkOrderConfirmationV02 {
	d.Message = new(RedemptionBulkOrderConfirmationV02)
	return d.Message
}

// Scope
// The RedemptionBulkOrderConfirmation message is sent by an executing party, eg, a transfer agent, to an instructing party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the executing party and the instructing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to confirm the details of the execution of a RedemptionBulkOrder message.
// Usage
// The RedemptionBulkOrderConfirmation message is sent, after the price has been determined, to confirm the execution of all individual orders.
// There is usually one bulk confirmation message for one bulk order message.
// A RedemptionBulkOrder must in all cases be responded to by a RedemptionBulkOrderConfirmation and in no circumstances by a RedemptionMultipleOrderConfirmation.
// If the executing party needs to confirm a RedemptionMultipleOrder message, then the RedemptionMultipleOrderConfirmation message must be used.
type RedemptionBulkOrderConfirmationV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef"`

	// General information related to the execution of investment fund orders.
	BulkExecutionDetails *iso20022.RedemptionBulkExecution2 `xml:"BlkExctnDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*iso20022.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (r *RedemptionBulkOrderConfirmationV02) AddMasterReference() *iso20022.AdditionalReference3 {
	r.MasterReference = new(iso20022.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionBulkOrderConfirmationV02) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderConfirmationV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	r.RelatedReference = new(iso20022.AdditionalReference3)
	return r.RelatedReference
}

func (r *RedemptionBulkOrderConfirmationV02) AddBulkExecutionDetails() *iso20022.RedemptionBulkExecution2 {
	r.BulkExecutionDetails = new(iso20022.RedemptionBulkExecution2)
	return r.BulkExecutionDetails
}

func (r *RedemptionBulkOrderConfirmationV02) AddIntermediaryDetails() *iso20022.Intermediary4 {
	newValue := new (iso20022.Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderConfirmationV02) AddCopyDetails() *iso20022.CopyInformation1 {
	r.CopyDetails = new(iso20022.CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderConfirmationV02) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

