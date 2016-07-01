package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.004.001.03 Document"`
	Message *RedemptionOrderV03 `xml:"RedOrdrV03"`
}

func (d *Document00400103) AddMessage() *RedemptionOrderV03 {
	d.Message = new(RedemptionOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the RedemptionOrder message to the executing party, for example, a transfer agent, to instruct the redemption of one or more financial instruments for one investment fund account.
// Usage
// The RedemptionOrder message is used to instruct single redemption orders, that is, a message containing one order for one financial instrument and related to one investment account. The RedemptionOrder message may also be used for multiple orders, that is, a message containing several orders related to the same investment account for different financial instruments.
// For a single redemption order, the RedemptionOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for the same financial instrument but for different accounts, then the RedemptionBulkOrder message must be used.
type RedemptionOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	MultipleOrderDetails *iso20022.RedemptionMultipleOrder4 `xml:"MltplOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (r *RedemptionOrderV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionOrderV03) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionOrderV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionOrderV03) AddMultipleOrderDetails() *iso20022.RedemptionMultipleOrder4 {
	r.MultipleOrderDetails = new(iso20022.RedemptionMultipleOrder4)
	return r.MultipleOrderDetails
}

func (r *RedemptionOrderV03) AddRelatedPartyDetails() *iso20022.Intermediary8 {
	newValue := new (iso20022.Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrderV03) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}

func (r *RedemptionOrderV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

