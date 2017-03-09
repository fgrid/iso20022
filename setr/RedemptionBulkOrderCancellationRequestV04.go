package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200104 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.04 Document"`
	Message *RedemptionBulkOrderCancellationRequestV04 `xml:"RedBlkOrdrCxlReq"`
}

func (d *Document00200104) AddMessage() *RedemptionBulkOrderCancellationRequestV04 {
	d.Message = new(RedemptionBulkOrderCancellationRequestV04)
	return d.Message
}

// Scope
// The RedemptionBulkOrderCancellationRequest message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to request the cancellation of a previously sent RedemptionBulkOrder.
// Usage
// The RedemptionBulkOrderCancellationRequest message is used to request the cancellation of one or more individual orders.
// There is no amendment, but a cancellation and re-instruct policy.
// To request the cancellation of one or more individual orders, the order reference of each individual order listed in the original RedemptionBulkOrder message is specified in the order reference element. The message identification of the RedemptionBulkOrder message which contains the individual orders to be cancelled may also be quoted in PreviousReference but this is not recommended.
// The deadline and acceptance of a cancellation request is subject to a service level agreement (SLA). This cancellation message is a cancellation request. There is no automatic acceptance of the cancellation.
// The rejection or acceptance of a RedemptionBulkOrderCancellationRequest is made using an OrderCancellationStatusReport message.
type RedemptionBulkOrderCancellationRequestV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the individual order to be cancelled.
	OrderReferences []*iso20022.InvestmentFundOrder9 `xml:"OrdrRefs"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`
}

func (r *RedemptionBulkOrderCancellationRequestV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderCancellationRequestV04) AddPoolReference() *iso20022.AdditionalReference9 {
	r.PoolReference = new(iso20022.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionBulkOrderCancellationRequestV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	r.PreviousReference = new(iso20022.AdditionalReference8)
	return r.PreviousReference
}

func (r *RedemptionBulkOrderCancellationRequestV04) SetMasterReference(value string) {
	r.MasterReference = (*iso20022.Max35Text)(&value)
}

func (r *RedemptionBulkOrderCancellationRequestV04) AddOrderReferences() *iso20022.InvestmentFundOrder9 {
	newValue := new(iso20022.InvestmentFundOrder9)
	r.OrderReferences = append(r.OrderReferences, newValue)
	return newValue
}

func (r *RedemptionBulkOrderCancellationRequestV04) AddCopyDetails() *iso20022.CopyInformation4 {
	r.CopyDetails = new(iso20022.CopyInformation4)
	return r.CopyDetails
}
