package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03700101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.037.001.01 Document"`
	Message *DebitAuthorisationRequest `xml:"camt.037.001.01"`
}

func (d *Document03700101) AddMessage() *DebitAuthorisationRequest {
	d.Message = new(DebitAuthorisationRequest)
	return d.Message
}

// Scope
// The Debit Authorisation Request message is sent by an account servicing institution to an account owner. This message is used to request authorisation to debit an account.
// Usage
// The Debit Authorisation Request message must be answered with a Debit Authorisation Response message.
// The Debit Authorisation Request message can be used to request debit authorisation in a:
// - request to modify payment case (in the case of a lower final amount or change of creditor)
// - request to cancel payment case (full amount)
// - unable to apply case (the creditor whose account has been credited is not the intended beneficiary)
// - claim non receipt case (the creditor whose account has been credited is not the intended beneficiary)
// The Debit Authorisation Request message covers one and only one payment instruction at a time. If an account servicing institution needs to request debit authorisation for several instructions, then multiple Debit Authorisation Request messages must be sent.
// The Debit Authorisation Request must be used exclusively between the account servicing institution and the account owner. It must not be used in place of a Request To Modify Payment or Request To Cancel Payment message between subsequent agents.
type DebitAuthorisationRequest struct {

	// Identifies the case assignment.
	Assignment *iso20022.CaseAssignment `xml:"Assgnmt"`

	// Identifies the case.
	Case *iso20022.Case `xml:"Case"`

	// Identifies the underlying payment instructrion.
	Underlying *iso20022.PaymentInstructionExtract `xml:"Undrlyg"`

	// Detailed information about the request.
	Detail *iso20022.DebitAuthorisationDetails `xml:"Dtl"`

}


func (d *DebitAuthorisationRequest) AddAssignment() *iso20022.CaseAssignment {
	d.Assignment = new(iso20022.CaseAssignment)
	return d.Assignment
}

func (d *DebitAuthorisationRequest) AddCase() *iso20022.Case {
	d.Case = new(iso20022.Case)
	return d.Case
}

func (d *DebitAuthorisationRequest) AddUnderlying() *iso20022.PaymentInstructionExtract {
	d.Underlying = new(iso20022.PaymentInstructionExtract)
	return d.Underlying
}

func (d *DebitAuthorisationRequest) AddDetail() *iso20022.DebitAuthorisationDetails {
	d.Detail = new(iso20022.DebitAuthorisationDetails)
	return d.Detail
}

