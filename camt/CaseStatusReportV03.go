package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03900103 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.039.001.03 Document"`
	Message *CaseStatusReportV03 `xml:"CaseStsRpt"`
}

func (d *Document03900103) AddMessage() *CaseStatusReportV03 {
	d.Message = new(CaseStatusReportV03)
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
// - may be skipped and replaced by a Resolution Of Investigation message when the request for a investigation status is received at the time the assigner has resolved the case. (In this case a Resolution Of Investigation message can be sent instead of a Case Status Report and the case may be closed.)
type CaseStatusReportV03 struct {

	// Specifies generic information about an investigation report.
	Header *iso20022.ReportHeader2 `xml:"Hdr"`

	// Identifies the investigation case.
	Case *iso20022.Case2 `xml:"Case"`

	// Defines the status of the case.
	Status *iso20022.CaseStatus2 `xml:"Sts"`

	// Identifies the change of an assignment for an investigation case from an assigner to a new assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	NewAssignment *iso20022.CaseAssignment2 `xml:"NewAssgnmt,omitempty"`
}

func (c *CaseStatusReportV03) AddHeader() *iso20022.ReportHeader2 {
	c.Header = new(iso20022.ReportHeader2)
	return c.Header
}

func (c *CaseStatusReportV03) AddCase() *iso20022.Case2 {
	c.Case = new(iso20022.Case2)
	return c.Case
}

func (c *CaseStatusReportV03) AddStatus() *iso20022.CaseStatus2 {
	c.Status = new(iso20022.CaseStatus2)
	return c.Status
}

func (c *CaseStatusReportV03) AddNewAssignment() *iso20022.CaseAssignment2 {
	c.NewAssignment = new(iso20022.CaseAssignment2)
	return c.NewAssignment
}
