package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05700102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.057.001.02 Document"`
	Message *OrderConfirmationStatusReportV02 `xml:"OrdrConfStsRpt"`
}

func (d *Document05700102) AddMessage() *OrderConfirmationStatusReportV02 {
	d.Message = new(OrderConfirmationStatusReportV02)
	return d.Message
}

// Scope
// The OrderConfirmationStatusReport message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to report the status of an order confirmation or an order confirmation amendment.
// Usage
// The OrderConfirmationStatusReport message is used to report on the status of one or more individual:
// - subscription confirmations,
// - subscription confirmation amendments,
// - redemption confirmations,
// - redemption confirmation amendments,
// - switch order confirmations,
// - switch order confirmation amendments.
// One of the following statuses can be reported:
// - confirmation rejected, or,
// - amendment rejected, or,
// - sent to next party, or,
// - communication problem with next party, or,
// - confirmation accepted, or,
// - confirmation received.
// It is likely that the OrderConfirmationStatusReport is only sent by the order instructing party to the order executing party to reject an order confirmation or to reject an order confirmation amendment, although if an intermediary party is used, the statuses sent to next party and communication problem with next party are also likely be used. The statuses confirmation accepted and confirmation received would only be used in the event the order executing party sends a RequestForOrderConfirmationStatusReport message and one of the other statuses does not apply.
// If the status being reported is either confirmation rejected or amendment rejected, then a reason for the rejection must be given.
// The individual order confirmation or confirmation amendment for which the status is given is identified with its order reference. The message identification of the message in which the individual order confirmation or confirmation amendment was conveyed may also be quoted in RelatedReference, but this is not recommended.
type OrderConfirmationStatusReportV02 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the message or communication that was previously received.
	Reference *iso20022.References61Choice `xml:"Ref,omitempty"`

	// Status report details of an individual order confirmation.
	IndividualOrderConfirmationDetailsReport []*iso20022.IndividualOrderConfirmationStatusAndReason2 `xml:"IndvOrdrConfDtlsRpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (o *OrderConfirmationStatusReportV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	o.MessageIdentification = new(iso20022.MessageIdentification1)
	return o.MessageIdentification
}

func (o *OrderConfirmationStatusReportV02) AddReference() *iso20022.References61Choice {
	o.Reference = new(iso20022.References61Choice)
	return o.Reference
}

func (o *OrderConfirmationStatusReportV02) AddIndividualOrderConfirmationDetailsReport() *iso20022.IndividualOrderConfirmationStatusAndReason2 {
	newValue := new(iso20022.IndividualOrderConfirmationStatusAndReason2)
	o.IndividualOrderConfirmationDetailsReport = append(o.IndividualOrderConfirmationDetailsReport, newValue)
	return newValue
}

func (o *OrderConfirmationStatusReportV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	o.Extension = append(o.Extension, newValue)
	return newValue
}
