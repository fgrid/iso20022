package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.03 Document"`
	Message *RedemptionBulkOrderCancellationRequestV03 `xml:"RedBlkOrdrCxlReqV03"`
}

func (d *Document00200103) AddMessage() *RedemptionBulkOrderCancellationRequestV03 {
	d.Message = new(RedemptionBulkOrderCancellationRequestV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the RedemptionBulkOrderCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent RedemptionBulkOrder instruction.
// Usage
// The RedemptionBulkOrderCancellationRequest message is used to either:
// - request the cancellation of an entire RedemptionBulkOrder message, that is, all the individual orders that it contained, or,
// - request the cancellation of one or more individual orders included in a previously sent RedemptionBulkOrder message. There is no amendment, but a cancellation and re-instruct policy.
// There are two ways to use the message.
// (1) When the RedemptionBulkOrderCancellationRequest message is used to request the cancellation of an entire RedemptionBulkOrder message, this can be done by either:
// - quoting the order references of all the individual orders listed in the RedemptionBulkOrder message, or,
// - quoting the details of all the individual orders (this includes the OrderReference) listed in RedemptionBulkOrder message, but this is not recommended.
// The message identification of the RedemptionBulkOrder message may also be quoted in PreviousReference.
// It is also possible to request the cancellation of an entire RedemptionBulkOrder message by quoting its message identification in PreviousReference, but this is not recommended.
// (2) When the RedemptionBulkOrderCancellationRequest message is used to request the cancellation of one or more individual orders, this can be done by either:
// - quoting the OrderReference of each individual order listed in the RedemptionOrder message, or,
// - quoting the details of each individual order (including the OrderReference) listed in RedemptionOrder message, but this is not recommended.
// The message identification of the RedemptionBulkOrder message in which the individual order was conveyed may also be quoted in PreviousReference.
// The deadline and acceptance of a cancellation request is subject to a service level agreement (SLA). This cancellation message is a cancellation request. There is no automatic acceptance of the cancellation.
type RedemptionBulkOrderCancellationRequestV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// References of the orders to be cancelled.
	CancellationByReference *iso20022.InvestmentFundOrder1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the orders to be cancelled.
	CancellationByOrderDetails *iso20022.RedemptionBulkOrderInstruction2 `xml:"CxlByOrdrDtls,omitempty"`

	// Message is a copy.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *RedemptionBulkOrderCancellationRequestV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderCancellationRequestV03) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderCancellationRequestV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	r.PreviousReference = new(iso20022.AdditionalReference3)
	return r.PreviousReference
}

func (r *RedemptionBulkOrderCancellationRequestV03) AddCancellationByReference() *iso20022.InvestmentFundOrder1 {
	r.CancellationByReference = new(iso20022.InvestmentFundOrder1)
	return r.CancellationByReference
}

func (r *RedemptionBulkOrderCancellationRequestV03) AddCancellationByOrderDetails() *iso20022.RedemptionBulkOrderInstruction2 {
	r.CancellationByOrderDetails = new(iso20022.RedemptionBulkOrderInstruction2)
	return r.CancellationByOrderDetails
}

func (r *RedemptionBulkOrderCancellationRequestV03) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}
