package reda

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:reda.001.001.03 Document"`
	Message *PriceReportV03 `xml:"PricRptV03"`
}

func (d *Document00100103) AddMessage() *PriceReportV03 {
	d.Message = new(PriceReportV03)
	return d.Message
}

// Scope
// A report provider, eg, a transfer agent, fund accountant or market data provider, sends the PriceReport message to the report recipient, eg, a fund management company, transfer agent, market data provider, regulator or other interested party to provide the net asset value and price information for financial instruments on specific trade dates and, optionally, to quote price variation information.
// Usage
// The PriceReport message is used to:
// - report prices for one or several different financial instruments for one or several different trade dates,
// - report statistical information about the valuation of a financial instrument,
// - inform another party that the quotation of a financial instrument is suspended,
// - report prices that are used for purposes other than the execution of investment funds orders.
type PriceReportV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*iso20022.PriceValuation3 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (p *PriceReportV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	p.MessageIdentification = new(iso20022.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportV03) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	p.PreviousReference = append(p.PreviousReference, newValue)
	return newValue
}

func (p *PriceReportV03) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PriceReportV03) AddMessagePagination() *iso20022.Pagination {
	p.MessagePagination = new(iso20022.Pagination)
	return p.MessagePagination
}

func (p *PriceReportV03) AddPriceValuationDetails() *iso20022.PriceValuation3 {
	newValue := new (iso20022.PriceValuation3)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReportV03) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

