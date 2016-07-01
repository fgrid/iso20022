package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.044.001.01 Document"`
	Message *FundConfirmedCashForecastReportCancellationV01 `xml:"camt.044.001.01"`
}

func (d *Document04400101) AddMessage() *FundConfirmedCashForecastReportCancellationV01 {
	d.Message = new(FundConfirmedCashForecastReportCancellationV01)
	return d.Message
}

// Scope
// The FundConfirmedCashForecastReportCancellation message is sent by a report provider, such as a transfer agent or a designated agent of the fund, to a report user, such as an investment manager, a fund accountant or any other interested party.
// This message is used to cancel a previously sent FundConfirmedCashForecastReport message.
// Usage
// The FundConfirmedCashForecastReportCancellation message is used to cancel an entire FundConfirmedCashForecastReport message that was previously sent by the report provider.
// This message must contain the reference of the message to be cancelled. This message may also contain details of the message to be cancelled, but this is not recommended.
type FundConfirmedCashForecastReportCancellationV01 struct {

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// The FundDetailedConfirmedCashForecastReport to be cancelled.
	CashForecastReportToBeCancelled *iso20022.FundConfirmedCashForecastReport1 `xml:"CshFcstRptToBeCanc,omitempty"`

}


func (f *FundConfirmedCashForecastReportCancellationV01) AddPoolReference() *iso20022.AdditionalReference3 {
	f.PoolReference = new(iso20022.AdditionalReference3)
	return f.PoolReference
}

func (f *FundConfirmedCashForecastReportCancellationV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	f.PreviousReference = new(iso20022.AdditionalReference3)
	return f.PreviousReference
}

func (f *FundConfirmedCashForecastReportCancellationV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportCancellationV01) AddCashForecastReportToBeCancelled() *iso20022.FundConfirmedCashForecastReport1 {
	f.CashForecastReportToBeCancelled = new(iso20022.FundConfirmedCashForecastReport1)
	return f.CashForecastReportToBeCancelled
}

