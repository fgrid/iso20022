package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:setr.001.001.03 Document"`
	Message *RedemptionBulkOrderV03 `xml:"RedBlkOrdrV03"`
}

func (d *Document00100103) AddMessage() *RedemptionBulkOrderV03 {
	d.Message = new(RedemptionBulkOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative sends the RedemptionBulkOrder message to the executing party, for example, a transfer agent, to instruct a redemption from a financial instrument for two or more accounts.
// Usage
// The RedemptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are for the same financial instrument. The RedemptionBulkOrder can result in one single bulk cash settlement or several individual cash settlements.
// This message will be typically used by a party collecting orders, that is, a concentrator, bulking these individual orders into one bulk order before sending it to an executing party.
// For a single redemption order, the RedemptionOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for different financial instruments but for the same account, then the RedemptionOrder must be used.
type RedemptionBulkOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Common information related to all the orders.
	BulkOrderDetails *iso20022.RedemptionBulkOrder4 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (r *RedemptionBulkOrderV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderV03) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV03) AddBulkOrderDetails() *iso20022.RedemptionBulkOrder4 {
	r.BulkOrderDetails = new(iso20022.RedemptionBulkOrder4)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderV03) AddRelatedPartyDetails() *iso20022.Intermediary8 {
	newValue := new (iso20022.Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV03) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

