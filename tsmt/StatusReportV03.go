package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03700103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.037.001.03 Document"`
	Message *StatusReportV03 `xml:"StsRpt"`
}

func (d *Document03700103) AddMessage() *StatusReportV03 {
	d.Message = new(StatusReportV03)
	return d.Message
}

// Scope
// The StatusReport message is sent by the matching application to the requester of a status report.
// This message is used to report on the status of transactions registered in the matching application.
// Usage
// The StatusReport message can be sent by the matching application to report on the status and sub-status of all transactions that the requester of the report is involved in. The message can be sent in response to a StatusReportRequest message.
type StatusReportV03 struct {

	// Identifies the report. 
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Reference to the previous message requesting the report.
	RelatedMessageReference *iso20022.MessageIdentification1 `xml:"RltdMsgRef"`

	// Describes a transaction and its status.
	ReportedItems []*iso20022.StatusReportItems2 `xml:"RptdItms,omitempty"`

}


func (s *StatusReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	s.ReportIdentification = new(iso20022.MessageIdentification1)
	return s.ReportIdentification
}

func (s *StatusReportV03) AddRelatedMessageReference() *iso20022.MessageIdentification1 {
	s.RelatedMessageReference = new(iso20022.MessageIdentification1)
	return s.RelatedMessageReference
}

func (s *StatusReportV03) AddReportedItems() *iso20022.StatusReportItems2 {
	newValue := new (iso20022.StatusReportItems2)
	s.ReportedItems = append(s.ReportedItems, newValue)
	return newValue
}

