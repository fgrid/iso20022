package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.03 Document"`
	Message *RequestForTransferStatusReportV03 `xml:"ReqForTrfStsRpt"`
}

func (d *Document00900103) AddMessage() *RequestForTransferStatusReportV03 {
	d.Message = new(RequestForTransferStatusReportV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the RequestForTransferStatusReport to the executing party, for example, a transfer agent to request the status of a previously instructed transfer.
// Usage
// The RequestForTransferStatusReport is used to request either:
// - the status of one or several transfer instructions or,
// - the status of one or several transfer cancellation instructions.
type RequestForTransferStatusReportV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Information to identify the transfer for which the status is requested.
	// 
	RequestDetails []*iso20022.MessageAndBusinessReference6 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (r *RequestForTransferStatusReportV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForTransferStatusReportV03) AddRequestDetails() *iso20022.MessageAndBusinessReference6 {
	newValue := new (iso20022.MessageAndBusinessReference6)
	r.RequestDetails = append(r.RequestDetails, newValue)
	return newValue
}

func (r *RequestForTransferStatusReportV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

