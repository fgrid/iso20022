package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:catp.009.001.02 Document"`
	Message *ATMCompletionAcknowledgementV02 `xml:"ATMCmpltnAck"`
}

func (d *Document00900102) AddMessage() *ATMCompletionAcknowledgementV02 {
	d.Message = new(ATMCompletionAcknowledgementV02)
	return d.Message
}

// The ATMCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMCompletionAdvice message.
type ATMCompletionAcknowledgementV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMCompletionAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdATMCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM completion on the ATM. manager.
	ATMCompletionAcknowledgement *iso20022.ATMCompletionAcknowledgement2 `xml:"ATMCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMCompletionAcknowledgementV02) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMCompletionAcknowledgementV02) AddProtectedATMCompletionAcknowledgement() *iso20022.ContentInformationType10 {
	a.ProtectedATMCompletionAcknowledgement = new(iso20022.ContentInformationType10)
	return a.ProtectedATMCompletionAcknowledgement
}

func (a *ATMCompletionAcknowledgementV02) AddATMCompletionAcknowledgement() *iso20022.ATMCompletionAcknowledgement2 {
	a.ATMCompletionAcknowledgement = new(iso20022.ATMCompletionAcknowledgement2)
	return a.ATMCompletionAcknowledgement
}

func (a *ATMCompletionAcknowledgementV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
