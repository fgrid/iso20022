package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.001.001.02 Document"`
	Message *RedemptionBulkOrderV02 `xml:"setr.001.001.02"`
}

func (d *Document00100102) AddMessage() *RedemptionBulkOrderV02 {
	d.Message = new(RedemptionBulkOrderV02)
	return d.Message
}

// Scope
// The RedemptionBulkOrder message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to instruct the executing party to redeem a financial instrument for two or more accounts.
// Usage
// The RedemptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, ie, account owners, but are the same financial instrument. The RedemptionBulkOrder can result in one single bulk cash settlement or several individual cash settlements.
// This message will be typically used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single redemption order, the RedemptionMultipleOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for different financial instruments but for the same account, then the RedemptionMultipleOrder must be used.
type RedemptionBulkOrderV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Common information related to all the orders.
	BulkOrderDetails *iso20022.RedemptionBulkOrder2 `xml:"BlkOrdrDtls"`

	// The information related to an intermediary.
	IntermediaryDetails []*iso20022.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderV02) AddMasterReference() *iso20022.AdditionalReference3 {
	r.MasterReference = new(iso20022.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionBulkOrderV02) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV02) AddBulkOrderDetails() *iso20022.RedemptionBulkOrder2 {
	r.BulkOrderDetails = new(iso20022.RedemptionBulkOrder2)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderV02) AddIntermediaryDetails() *iso20022.Intermediary4 {
	newValue := new(iso20022.Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV02) AddCopyDetails() *iso20022.CopyInformation1 {
	r.CopyDetails = new(iso20022.CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
