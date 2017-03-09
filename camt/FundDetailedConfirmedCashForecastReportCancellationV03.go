package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04500103 struct {
	XMLName xml.Name                                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.045.001.03 Document"`
	Message *FundDetailedConfirmedCashForecastReportCancellationV03 `xml:"FndDtldConfdCshFcstRptCxl"`
}

func (d *Document04500103) AddMessage() *FundDetailedConfirmedCashForecastReportCancellationV03 {
	d.Message = new(FundDetailedConfirmedCashForecastReportCancellationV03)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundDetailedConfirmedCashForecastReportCancellation messages to the report user, such as an investment manager, fund accountant or any other interested party, to cancel a previously sent FundDetailedConfirmedCashForecastReport.
// Usage
// The FundDetailedConfirmedCashForecastReportCancellation message is used to cancel an entire FundDetailedConfirmedCashForecastReport message that was previously sent. This message must contain the reference of the message to be cancelled.
// This message may also contain details of the message to be cancelled, but this is not recommended.
type FundDetailedConfirmedCashForecastReportCancellationV03 struct {

	// Identifies the message.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// The FundDetailedConfirmedCashForecastReport to be cancelled.
	CashForecastReportToBeCancelled *iso20022.FundDetailedConfirmedCashForecastReport3 `xml:"CshFcstRptToBeCanc,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	f.MessageIdentification = new(iso20022.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV03) AddPoolReference() *iso20022.AdditionalReference3 {
	f.PoolReference = new(iso20022.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	f.PreviousReference = new(iso20022.AdditionalReference3)
	return f.PreviousReference
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV03) AddRelatedReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV03) AddMessagePagination() *iso20022.Pagination {
	f.MessagePagination = new(iso20022.Pagination)
	return f.MessagePagination
}

func (f *FundDetailedConfirmedCashForecastReportCancellationV03) AddCashForecastReportToBeCancelled() *iso20022.FundDetailedConfirmedCashForecastReport3 {
	f.CashForecastReportToBeCancelled = new(iso20022.FundDetailedConfirmedCashForecastReport3)
	return f.CashForecastReportToBeCancelled
}
