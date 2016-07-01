package reda

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:reda.002.001.04 Document"`
	Message *PriceReportCancellationV04 `xml:"PricRptCxl"`
}

func (d *Document00200104) AddMessage() *PriceReportCancellationV04 {
	d.Message = new(PriceReportCancellationV04)
	return d.Message
}

// SCOPE
// A report provider, for example, a transfer agent, fund accountant or market data provider, sends the PriceReportCancellation message to the report recipient, for example, a fund management company, transfer agent, market data provider, regulator or any other interested party to cancel previously sent prices.
// 
// USAGE
// The PriceReportCancellation message is used to either:
// - cancel an entire PriceReport that was previously sent (by quoting the business reference of the original price report in the PriceReportIdentification element), or,
// - cancel one or more individual prices from a previously sent price report (by using the PriceDetailsToBeCancelled sequence).
// Technically, it is possible to cancel all the prices individually by using the PriceDetailsToBeCancelled sequence, but this is not recommended.
// The cancellation should not contain the cancellation of prices for more than one NAV date. 
// 
// 
type PriceReportCancellationV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// Unique and unambiguous identifier for the price report, as assigned by the reporting party.
	PriceReportIdentification *iso20022.Max35Text `xml:"PricRptId"`

	// Unique and unambiguous identifier for the cancellation of the previous price report, as assigned by the reporting party.
	CancellationIdentification *iso20022.Max35Text `xml:"CxlId"`

	// Reason for the cancellation.
	CancellationReason *iso20022.Max350Text `xml:"CxlRsn,omitempty"`

	// Date or date and time the price will be corrected.
	ExpectedPriceCorrectionDate *iso20022.DateAndDateTime1Choice `xml:"XpctdPricCrrctnDt,omitempty"`

	// Indicates whether or not all the prices of the referenced price report are cancelled.
	CompletePriceCancellation *iso20022.YesNoIndicator `xml:"CmpltPricCxl"`

	// Details of prices to be cancelled.
	CancelledPriceValuationDetails []*iso20022.PriceReport3 `xml:"CancPricValtnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (p *PriceReportCancellationV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	p.MessageIdentification = new(iso20022.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportCancellationV04) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportCancellationV04) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PriceReportCancellationV04) AddMessagePagination() *iso20022.Pagination {
	p.MessagePagination = new(iso20022.Pagination)
	return p.MessagePagination
}

func (p *PriceReportCancellationV04) SetPriceReportIdentification(value string) {
	p.PriceReportIdentification = (*iso20022.Max35Text)(&value)
}

func (p *PriceReportCancellationV04) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*iso20022.Max35Text)(&value)
}

func (p *PriceReportCancellationV04) SetCancellationReason(value string) {
	p.CancellationReason = (*iso20022.Max350Text)(&value)
}

func (p *PriceReportCancellationV04) AddExpectedPriceCorrectionDate() *iso20022.DateAndDateTime1Choice {
	p.ExpectedPriceCorrectionDate = new(iso20022.DateAndDateTime1Choice)
	return p.ExpectedPriceCorrectionDate
}

func (p *PriceReportCancellationV04) SetCompletePriceCancellation(value string) {
	p.CompletePriceCancellation = (*iso20022.YesNoIndicator)(&value)
}

func (p *PriceReportCancellationV04) AddCancelledPriceValuationDetails() *iso20022.PriceReport3 {
	newValue := new (iso20022.PriceReport3)
	p.CancelledPriceValuationDetails = append(p.CancelledPriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReportCancellationV04) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

