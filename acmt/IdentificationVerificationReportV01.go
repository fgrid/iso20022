package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02400101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.024.001.01 Document"`
	Message *IdentificationVerificationReportV01 `xml:"IdVrfctnRpt"`
}

func (d *Document02400101) AddMessage() *IdentificationVerificationReportV01 {
	d.Message = new(IdentificationVerificationReportV01)
	return d.Message
}

// Scope
// The IdentificationVerificationReport message is sent by an assigner to an assignee. It is used to confirm whether or not the presented party and/or account identification information is correct.
// Usage
// The IdentificationVerificationReport message is sent as a response to an IdentificationVerificationRequest message.
// The IdentificationVerificationReport message can contain one or more reports.
// The IdentificationVerificationReport message may include a reason if the presented party and/or account identification information is confirmed to be incorrect.
// The IdentificationVerificationReport message may include the correct party and/or account identification information.
type IdentificationVerificationReportV01 struct {

	// Identifies the identification assignment.
	Assignment *iso20022.IdentificationAssignment1 `xml:"Assgnmt"`

	// Provides for the reference to the original identification assignment.
	OriginalAssignment *iso20022.MessageIdentification4 `xml:"OrgnlAssgnmt,omitempty"`

	// Information concerning the verification of the identification data for which verification was requested.
	Report []*iso20022.VerificationReport1 `xml:"Rpt"`
}

func (i *IdentificationVerificationReportV01) AddAssignment() *iso20022.IdentificationAssignment1 {
	i.Assignment = new(iso20022.IdentificationAssignment1)
	return i.Assignment
}

func (i *IdentificationVerificationReportV01) AddOriginalAssignment() *iso20022.MessageIdentification4 {
	i.OriginalAssignment = new(iso20022.MessageIdentification4)
	return i.OriginalAssignment
}

func (i *IdentificationVerificationReportV01) AddReport() *iso20022.VerificationReport1 {
	newValue := new(iso20022.VerificationReport1)
	i.Report = append(i.Report, newValue)
	return newValue
}
