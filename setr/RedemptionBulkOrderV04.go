package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100104 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.001.001.04 Document"`
	Message *RedemptionBulkOrderV04 `xml:"RedBlkOrdr"`
}

func (d *Document00100104) AddMessage() *RedemptionBulkOrderV04 {
	d.Message = new(RedemptionBulkOrderV04)
	return d.Message
}

// Scope
// The RedemptionBulkOrder message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to instruct a redemption from a financial instrument for two or more accounts.
// Usage
// The RedemptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are for the same financial instrument. The RedemptionBulkOrder can result in one single bulk cash settlement or several individual cash settlements.
// This message will be typically used by a party collecting orders, that is, a concentrator, bulking these individual orders into one bulk order before sending it to an executing party.
// For a single redemption order, the RedemptionOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for different financial instruments but for the same account, then the RedemptionOrder must be used.
type RedemptionBulkOrderV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// General information related to the orders.
	BulkOrderDetails *iso20022.RedemptionBulkOrder6 `xml:"BlkOrdrDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderV04) AddPoolReference() *iso20022.AdditionalReference9 {
	r.PoolReference = new(iso20022.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionBulkOrderV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV04) AddBulkOrderDetails() *iso20022.RedemptionBulkOrder6 {
	r.BulkOrderDetails = new(iso20022.RedemptionBulkOrder6)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderV04) AddCopyDetails() *iso20022.CopyInformation4 {
	r.CopyDetails = new(iso20022.CopyInformation4)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
