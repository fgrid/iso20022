package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.012.001.01 Document"`
	Message *ATMExceptionAcknowledgementV01 `xml:"ATMXcptnAck"`
}

func (d *Document01200101) AddMessage() *ATMExceptionAcknowledgementV01 {
	d.Message = new(ATMExceptionAcknowledgementV01)
	return d.Message
}

// The ATMExceptionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMExceptionAdvice message.
type ATMExceptionAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMExceptionAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdATMXcptnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM exception.
	ATMExceptionAcknowledgement *iso20022.ATMExceptionAcknowledgement1 `xml:"ATMXcptnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMExceptionAcknowledgementV01) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMExceptionAcknowledgementV01) AddProtectedATMExceptionAcknowledgement() *iso20022.ContentInformationType10 {
	a.ProtectedATMExceptionAcknowledgement = new(iso20022.ContentInformationType10)
	return a.ProtectedATMExceptionAcknowledgement
}

func (a *ATMExceptionAcknowledgementV01) AddATMExceptionAcknowledgement() *iso20022.ATMExceptionAcknowledgement1 {
	a.ATMExceptionAcknowledgement = new(iso20022.ATMExceptionAcknowledgement1)
	return a.ATMExceptionAcknowledgement
}

func (a *ATMExceptionAcknowledgementV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
