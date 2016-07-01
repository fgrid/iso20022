package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.003.001.03 Document"`
	Message *ActivityReportRequestV03 `xml:"ActvtyReqRpt"`
}

func (d *Document00300103) AddMessage() *ActivityReportRequestV03 {
	d.Message = new(ActivityReportRequestV03)
	return d.Message
}

// Scope
// The ActivityReportRequest message is sent by any party involved in a transaction to the matching application.
// This message is used to request a report on all transactions for which an activity has taken place within a given time span.
// Usage
// The ActivityReportRequest message can be sent
// - at a pre-determined time. The message requests a report on all transactions that the requester is involved in and for which an activity has taken place within the last 24 hours.
// - on demand. The message requests a report on all transactions that the requester is involved in and for which an activity has taken place within a time span specified by the requester.
type ActivityReportRequestV03 struct {

	// Identifies the request message.
	RequestIdentification *iso20022.MessageIdentification1 `xml:"ReqId"`

	// Specifies the entities of the submitter for which the activities have to be reported.
	EntitiesToBeReported []*iso20022.BICIdentification1 `xml:"NttiesToBeRptd,omitempty"`

	// Specifies the period for which activities have to be reported.
	ReportPeriod *iso20022.DateTimePeriodDetails1 `xml:"RptPrd"`

}


func (a *ActivityReportRequestV03) AddRequestIdentification() *iso20022.MessageIdentification1 {
	a.RequestIdentification = new(iso20022.MessageIdentification1)
	return a.RequestIdentification
}

func (a *ActivityReportRequestV03) AddEntitiesToBeReported() *iso20022.BICIdentification1 {
	newValue := new (iso20022.BICIdentification1)
	a.EntitiesToBeReported = append(a.EntitiesToBeReported, newValue)
	return newValue
}

func (a *ActivityReportRequestV03) AddReportPeriod() *iso20022.DateTimePeriodDetails1 {
	a.ReportPeriod = new(iso20022.DateTimePeriodDetails1)
	return a.ReportPeriod
}

