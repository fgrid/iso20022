package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.031.001.01 Document"`
	Message *RejectCaseAssignment `xml:"camt.031.001.01"`
}

func (d *Document03100101) AddMessage() *RejectCaseAssignment {
	d.Message = new(RejectCaseAssignment)
	return d.Message
}

// Scope
// The Reject Case Assignment message is sent by a case assignee to a case creator or case assigner to reject a case given to him.
// Usage
// The Reject Case Assignment message is used to notify the case creator or case assigner the rejection of an assignment by the case assignee in a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// Rejecting a case assignment occurs when the case assignee is unable to trace the original payment instruction or when the case assignee is unable, or does not have authority, to process the assigned case.
// The Reject Case Assignment message covers one and only one case at a time. If the case assignee needs to reject several case assignments, then multiple Reject Case Assignment messages must be sent.
// The Reject Case Assignment message must be forwarded by all subsequent case assignee(s) until it reaches the case assigner.
// The Reject Case Assignment message must not be used in place of a Resolution Of Investigation or Case Status Report message.
type RejectCaseAssignment struct {

	// Identifies the assignment.
	Assignment *iso20022.CaseAssignment `xml:"Assgnmt"`

	// Identifies the case.
	Case *iso20022.Case `xml:"Case"`

	// Specifies the reason for not accepting a Case.
	Justification *iso20022.CaseAssignmentRejectionJustification `xml:"Justfn"`

}


func (r *RejectCaseAssignment) AddAssignment() *iso20022.CaseAssignment {
	r.Assignment = new(iso20022.CaseAssignment)
	return r.Assignment
}

func (r *RejectCaseAssignment) AddCase() *iso20022.Case {
	r.Case = new(iso20022.Case)
	return r.Case
}

func (r *RejectCaseAssignment) AddJustification() *iso20022.CaseAssignmentRejectionJustification {
	r.Justification = new(iso20022.CaseAssignmentRejectionJustification)
	return r.Justification
}

