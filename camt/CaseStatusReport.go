package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.039.001.01 Document"`
	Message *CaseStatusReport `xml:"camt.039.001.01"`
}

func (d *Document03900101) AddMessage() *CaseStatusReport {
	d.Message = new(CaseStatusReport)
	return d.Message
}

// Scope
// The Case Status Report message is sent by a case assignee to a case creator or case assigner.
// This message is used to report on the status of a case.
// Usage
// A Case Status Report message is sent in reply to a Case Status Report Request message. This message
// - covers one and only one case at a time. (If a case assignee needs to report on several cases, then multiple Case Status Report messages must be sent.)
// - may be forwarded to subsequent case assigner(s) until it reaches the end point
// - is able to indicate the fact that a case has been assigned to a party downstream in the payment processing chain
// - may not be used in place of a Resolution Of Investigation (except for the condition given in the next bullet point) or Notification Of Case Assignment message
// - may be skipped and replaced by a Resolution Of Investigation message if at the moment when the request for a investigation status arrives, the assignee has obtained a solution. (In this case a Resolution Of Investigation message can be sent in lieu of a Case Status Report and the case may be closed.)
type CaseStatusReport struct {

	// Specifies generic information about an investigation report.
	Header *iso20022.ReportHeader `xml:"Hdr"`

	// Identifies the case.
	Case *iso20022.Case `xml:"Case"`

	// Defines the status of the case.
	Status *iso20022.CaseStatus `xml:"Sts"`

	// Identifies the last assignment performed.
	NewAssignment *iso20022.CaseAssignment `xml:"NewAssgnmt,omitempty"`

}


func (c *CaseStatusReport) AddHeader() *iso20022.ReportHeader {
	c.Header = new(iso20022.ReportHeader)
	return c.Header
}

func (c *CaseStatusReport) AddCase() *iso20022.Case {
	c.Case = new(iso20022.Case)
	return c.Case
}

func (c *CaseStatusReport) AddStatus() *iso20022.CaseStatus {
	c.Status = new(iso20022.CaseStatus)
	return c.Status
}

func (c *CaseStatusReport) AddNewAssignment() *iso20022.CaseAssignment {
	c.NewAssignment = new(iso20022.CaseAssignment)
	return c.NewAssignment
}

