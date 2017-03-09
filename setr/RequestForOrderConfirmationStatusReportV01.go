package setr

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05800101 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.058.001.01 Document"`
	Message *RequestForOrderConfirmationStatusReportV01 `xml:"ReqForOrdrConfStsRptV01"`
}

func (d *Document05800101) AddMessage() *RequestForOrderConfirmationStatusReportV01 {
	d.Message = new(RequestForOrderConfirmationStatusReportV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, send the RequestForOrderConfirmationStatusReport message to the instructing party, for example, an investment manager or its authorised representative, to request the status of one or several order confirmations.
// Usage
// The RequestForOrderConfirmationStatusReport message is used to request the status of either:
// - one or several individual order confirmations, or,
// - one or several order confirmation messages.
// The response to a RequestForOrderConfirmationStatusReport message is the OrderConfirmationStatusReport message.
// When the RequestForOrderConfirmationStatusReport message is used to request the status of several individual order confirmations or one or more order confirmation messages, the executing party may receive several OrderConfirmationStatusReport messages from the instructing party.
// When the RequestForOrderConfirmationStatusReport is used to request the status of one or more individual order confirmations, each individual order confirmation is identified with its order reference. The message identification of the message in which the individual order confirmation was conveyed may also be quoted in PreviousReference.
// When the RequestForOrderConfirmationStatusReport is used to request the status of an order confirmation message, then the message identification of the order confirmation message is identified in PreviousReference.
type RequestForOrderConfirmationStatusReportV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Information to identify the order confirmation for which the status is requested.
	RequestDetails []*iso20022.MessageAndBusinessReference5 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RequestForOrderConfirmationStatusReportV01) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForOrderConfirmationStatusReportV01) AddRequestDetails() *iso20022.MessageAndBusinessReference5 {
	newValue := new(iso20022.MessageAndBusinessReference5)
	r.RequestDetails = append(r.RequestDetails, newValue)
	return newValue
}

func (r *RequestForOrderConfirmationStatusReportV01) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
