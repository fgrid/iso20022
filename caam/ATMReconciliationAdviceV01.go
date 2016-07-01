package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.01 Document"`
	Message *ATMReconciliationAdviceV01 `xml:"ATMRcncltnAdvc"`
}

func (d *Document00900101) AddMessage() *ATMReconciliationAdviceV01 {
	d.Message = new(ATMReconciliationAdviceV01)
	return d.Message
}

// The ATMReconciliationAdvice message is sent by an ATM to an acquirer or its agent to send all the counters of the ATM. It can be sent by an operator function or as a response of a command sent by an agent or a server.
type ATMReconciliationAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMReconciliationAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMRcncltnAdvc,omitempty"`

	// Information related to the reconciliation of an ATM.
	ATMReconciliationAdvice *iso20022.ATMReconciliationAdvice1 `xml:"ATMRcncltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (a *ATMReconciliationAdviceV01) AddHeader() *iso20022.Header21 {
	a.Header = new(iso20022.Header21)
	return a.Header
}

func (a *ATMReconciliationAdviceV01) AddProtectedATMReconciliationAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMReconciliationAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMReconciliationAdvice
}

func (a *ATMReconciliationAdviceV01) AddATMReconciliationAdvice() *iso20022.ATMReconciliationAdvice1 {
	a.ATMReconciliationAdvice = new(iso20022.ATMReconciliationAdvice1)
	return a.ATMReconciliationAdvice
}

func (a *ATMReconciliationAdviceV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}

