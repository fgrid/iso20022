package reda

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:reda.003.001.03 Document"`
	Message *PriceReportCorrectionV03 `xml:"PricRptCrrctnV03"`
}

func (d *Document00300103) AddMessage() *PriceReportCorrectionV03 {
	d.Message = new(PriceReportCorrectionV03)
	return d.Message
}

// Scope
// A report provider, eg, a transfer agent, fund accountant or market data provider, sends the PriceReportCorrection message to the report recipient, eg, a fund management company, transfer agent, market data provider, regulator or any other interested party to correct at least one of the prices for a financial instrument that was reported in a previously sent PriceReport message.
// Usage
// The PriceReportCorrection message is used to correct information reported in a previously sent PriceReport.
// If an entire PriceReport message must be corrected, eg, due to an incorrect trade date, it is recommended that a PriceReportCancellation message is used to cancel the entire PriceReport message and a new PriceReport sent.
type PriceReportCorrectionV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint. 
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// Information related to the correction of a price of a financial instrument.
	PriceCorrectionDetails []*iso20022.PriceCorrection3 `xml:"PricCrrctnDtls"`

}


func (p *PriceReportCorrectionV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	p.MessageIdentification = new(iso20022.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportCorrectionV03) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportCorrectionV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PriceReportCorrectionV03) AddMessagePagination() *iso20022.Pagination {
	p.MessagePagination = new(iso20022.Pagination)
	return p.MessagePagination
}

func (p *PriceReportCorrectionV03) AddPriceCorrectionDetails() *iso20022.PriceCorrection3 {
	newValue := new (iso20022.PriceCorrection3)
	p.PriceCorrectionDetails = append(p.PriceCorrectionDetails, newValue)
	return newValue
}

