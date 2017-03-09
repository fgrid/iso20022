package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.011.001.01 Document"`
	Message *ATMExceptionAdviceV01 `xml:"ATMXcptnAdvc"`
}

func (d *Document01100101) AddMessage() *ATMExceptionAdviceV01 {
	d.Message = new(ATMExceptionAdviceV01)
	return d.Message
}

// The ATMExceptionAdvice message is sent by an ATM to an acquirer or its agent to inform of that an exception occurred outside a service.
type ATMExceptionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMExceptionAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMXcptnAdvc,omitempty"`

	// Information related to exceptions occurring on the ATM.
	ATMExceptionAdvice *iso20022.ATMExceptionAdvice1 `xml:"ATMXcptnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMExceptionAdviceV01) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMExceptionAdviceV01) AddProtectedATMExceptionAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMExceptionAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMExceptionAdvice
}

func (a *ATMExceptionAdviceV01) AddATMExceptionAdvice() *iso20022.ATMExceptionAdvice1 {
	a.ATMExceptionAdvice = new(iso20022.ATMExceptionAdvice1)
	return a.ATMExceptionAdvice
}

func (a *ATMExceptionAdviceV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
