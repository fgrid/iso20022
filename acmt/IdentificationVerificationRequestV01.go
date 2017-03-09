package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02300101 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.023.001.01 Document"`
	Message *IdentificationVerificationRequestV01 `xml:"IdVrfctnReq"`
}

func (d *Document02300101) AddMessage() *IdentificationVerificationRequestV01 {
	d.Message = new(IdentificationVerificationRequestV01)
	return d.Message
}

// Scope
// The IdentificationVerificationRequest message is sent by an assigner to an assignee. It is used to request the verification of party and/or account identification information.
// Usage
// The IdentificationVerificationRequest message is sent before the sending of one or several transactions messages.
// The IdentificationVerificationRequest message can contain one or more verification requests.
type IdentificationVerificationRequestV01 struct {

	// Identifies the identification assignment.
	Assignment *iso20022.IdentificationAssignment1 `xml:"Assgnmt"`

	// Information concerning the identification data that is requested to be verified.
	Verification []*iso20022.IdentificationVerification1 `xml:"Vrfctn"`
}

func (i *IdentificationVerificationRequestV01) AddAssignment() *iso20022.IdentificationAssignment1 {
	i.Assignment = new(iso20022.IdentificationAssignment1)
	return i.Assignment
}

func (i *IdentificationVerificationRequestV01) AddVerification() *iso20022.IdentificationVerification1 {
	newValue := new(iso20022.IdentificationVerification1)
	i.Verification = append(i.Verification, newValue)
	return newValue
}
