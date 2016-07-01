package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.02 Document"`
	Message *ProprietaryFormatInvestigationV02 `xml:"PrtryFrmtInvstgtn"`
}

func (d *Document03500102) AddMessage() *ProprietaryFormatInvestigationV02 {
	d.Message = new(ProprietaryFormatInvestigationV02)
	return d.Message
}

// Scope
// The Proprietary Format Investigation message type is used by financial institutions, with their own offices, and/or with other financial institutions with which they have established bilateral agreements.
// Usage
// The user should ensure that an existing standard message cannot be used before using the proprietary message.
// As defined in the scope, this ProprietaryFormatInvestigation message may only be used when bilaterally agreed.
// It is used as an envelope for a non standard message and provides means to manage an exception or investigation which falls outside the scope or capability of any other formatted message.
type ProprietaryFormatInvestigationV02 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage Rule: the Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case2 `xml:"Case"`

	// Proprietary information.
	ProprietaryData *iso20022.ProprietaryData4 `xml:"PrtryData"`

}


func (p *ProprietaryFormatInvestigationV02) AddAssignment() *iso20022.CaseAssignment2 {
	p.Assignment = new(iso20022.CaseAssignment2)
	return p.Assignment
}

func (p *ProprietaryFormatInvestigationV02) AddCase() *iso20022.Case2 {
	p.Case = new(iso20022.Case2)
	return p.Case
}

func (p *ProprietaryFormatInvestigationV02) AddProprietaryData() *iso20022.ProprietaryData4 {
	p.ProprietaryData = new(iso20022.ProprietaryData4)
	return p.ProprietaryData
}

