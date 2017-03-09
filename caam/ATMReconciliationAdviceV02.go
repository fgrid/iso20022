package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Document"`
	Message *ATMReconciliationAdviceV02 `xml:"ATMRcncltnAdvc"`
}

func (d *Document00900102) AddMessage() *ATMReconciliationAdviceV02 {
	d.Message = new(ATMReconciliationAdviceV02)
	return d.Message
}

// The ATMReconciliationAdvice message is sent by an ATM to an acquirer or its agent to send all the counters of the ATM. It can be sent by an operator function or as a response of a command sent by an agent or a server.
type ATMReconciliationAdviceV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMReconciliationAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMRcncltnAdvc,omitempty"`

	// Information related to the reconciliation of an ATM.
	ATMReconciliationAdvice *iso20022.ATMReconciliationAdvice2 `xml:"ATMRcncltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMReconciliationAdviceV02) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMReconciliationAdviceV02) AddProtectedATMReconciliationAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMReconciliationAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMReconciliationAdvice
}

func (a *ATMReconciliationAdviceV02) AddATMReconciliationAdvice() *iso20022.ATMReconciliationAdvice2 {
	a.ATMReconciliationAdvice = new(iso20022.ATMReconciliationAdvice2)
	return a.ATMReconciliationAdvice
}

func (a *ATMReconciliationAdviceV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
