package reda

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:reda.002.001.03 Document"`
	Message *PriceReportCancellationV03 `xml:"PricRptCxlV03"`
}

func (d *Document00200103) AddMessage() *PriceReportCancellationV03 {
	d.Message = new(PriceReportCancellationV03)
	return d.Message
}

// Scope
// A report provider, eg, a transfer agent, fund accountant or market data provider, sends the PriceReportCancellation message to the report recipient, eg, a fund management company, transfer agent, market data provider, regulator or any other interested party to cancel a previously sent PriceReport message.
// Usage
// The PriceReportCancellation is used to cancel an entire PriceReport message that was previously sent.
// If only a part of the information needs to be cancelled and replaced, the PriceReportCorrection message must be used.
// This message must contain the reference of the message to be cancelled.This message may also contain details of the message to be cancelled, but this is not recommended.
type PriceReportCancellationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// Common information related to all the price reports to be cancelled.
	PriceReportToBeCancelled *iso20022.PriceReport2 `xml:"PricRptToBeCanc,omitempty"`

}


func (p *PriceReportCancellationV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	p.MessageIdentification = new(iso20022.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportCancellationV03) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportCancellationV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PriceReportCancellationV03) AddMessagePagination() *iso20022.Pagination {
	p.MessagePagination = new(iso20022.Pagination)
	return p.MessagePagination
}

func (p *PriceReportCancellationV03) AddPriceReportToBeCancelled() *iso20022.PriceReport2 {
	p.PriceReportToBeCancelled = new(iso20022.PriceReport2)
	return p.PriceReportToBeCancelled
}

