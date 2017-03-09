package pain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.01 Document"`
	Message *CreditorPaymentActivationRequestV01 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300101) AddMessage() *CreditorPaymentActivationRequestV01 {
	d.Message = new(CreditorPaymentActivationRequestV01)
	return d.Message
}

// Scope
// This message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents.
// It is used to initiate a creditor payment activation request.
type CreditorPaymentActivationRequestV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*iso20022.PaymentInstruction5 `xml:"PmtInf"`
}

func (c *CreditorPaymentActivationRequestV01) AddGroupHeader() *iso20022.GroupHeader45 {
	c.GroupHeader = new(iso20022.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV01) AddPaymentInformation() *iso20022.PaymentInstruction5 {
	newValue := new(iso20022.PaymentInstruction5)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}
