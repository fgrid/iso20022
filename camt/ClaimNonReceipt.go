package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02700101 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.01 Document"`
	Message *ClaimNonReceipt `xml:"camt.027.001.01"`
}

func (d *Document02700101) AddMessage() *ClaimNonReceipt {
	d.Message = new(ClaimNonReceipt)
	return d.Message
}

// Scope
// The Claim Non Receipt message is sent by a case creator/case assigner to a case assignee.
// This message allows to initiate an investigation in case the beneficiary of a payment has not received an expected payment.
// Usage
// Note 1: Although there are cases where a creditor would contact the creditor's bank when claiming non-receipt, the activity only retained the scenario where the beneficiary claims non-receipt with its debtor, the debtor in its turn contacting the first agent. Therefore the investigation follows the same route as the original instruction.
// Note 2: This message is also used to report a missing cover. The rationale behind this is that the beneficiary of the cover (receiver of the payment instruction) missing the cover would be in the very same position as a beneficiary expecting a credit to its account and would therefore trigger the same processes.
// In case of a Missing cover, the case will be assigned to the sender of the payment instruction, before following the route of the payment instruction.
type ClaimNonReceipt struct {

	// Identifies an assignment.
	Assignment *iso20022.CaseAssignment `xml:"Assgnmt"`

	// Identifies a case.
	Case *iso20022.Case `xml:"Case"`

	// Identifies the payment instruction for which the Creditor has not received the funds.
	// Note:
	// In case of a missing cover, it must be the Field 20 of the received MT103.
	// In case of a claim non receipt initiated by the Debtor, it must be the identification of the instruction (Field 20 of MT103, Instruction Identification of the Payment Initiation or the Bulk Credit Transfer).
	Underlying *iso20022.PaymentInstructionExtract `xml:"Undrlyg"`

	// Indicates if the claim non receipt is a missing cover. The absence of the component means that the message is not for a missing cover.
	MissingCover *iso20022.MissingCover `xml:"MssngCover,omitempty"`
}

func (c *ClaimNonReceipt) AddAssignment() *iso20022.CaseAssignment {
	c.Assignment = new(iso20022.CaseAssignment)
	return c.Assignment
}

func (c *ClaimNonReceipt) AddCase() *iso20022.Case {
	c.Case = new(iso20022.Case)
	return c.Case
}

func (c *ClaimNonReceipt) AddUnderlying() *iso20022.PaymentInstructionExtract {
	c.Underlying = new(iso20022.PaymentInstructionExtract)
	return c.Underlying
}

func (c *ClaimNonReceipt) AddMissingCover() *iso20022.MissingCover {
	c.MissingCover = new(iso20022.MissingCover)
	return c.MissingCover
}
