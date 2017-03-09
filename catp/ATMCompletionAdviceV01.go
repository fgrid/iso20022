package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.008.001.01 Document"`
	Message *ATMCompletionAdviceV01 `xml:"ATMCmpltnAdvc"`
}

func (d *Document00800101) AddMessage() *ATMCompletionAdviceV01 {
	d.Message = new(ATMCompletionAdviceV01)
	return d.Message
}

// The ATMCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a transaction performed on the ATM.
type ATMCompletionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMCompletionAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMCmpltnAdvc,omitempty"`

	// Information related to the completion of an operation on the ATM.
	ATMCompletionAdvice *iso20022.ATMCompletionAdvice1 `xml:"ATMCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMCompletionAdviceV01) AddHeader() *iso20022.Header21 {
	a.Header = new(iso20022.Header21)
	return a.Header
}

func (a *ATMCompletionAdviceV01) AddProtectedATMCompletionAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMCompletionAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMCompletionAdvice
}

func (a *ATMCompletionAdviceV01) AddATMCompletionAdvice() *iso20022.ATMCompletionAdvice1 {
	a.ATMCompletionAdvice = new(iso20022.ATMCompletionAdvice1)
	return a.ATMCompletionAdvice
}

func (a *ATMCompletionAdviceV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
