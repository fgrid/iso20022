package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:catp.009.001.01 Document"`
	Message *ATMCompletionAcknowledgementV01 `xml:"ATMCmpltnAck"`
}

func (d *Document00900101) AddMessage() *ATMCompletionAcknowledgementV01 {
	d.Message = new(ATMCompletionAcknowledgementV01)
	return d.Message
}

// The ATMCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMCompletionAdvice message.
type ATMCompletionAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMCompletionAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdATMCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM completion on the ATM. manager.
	ATMCompletionAcknowledgement *iso20022.ATMCompletionAcknowledgement1 `xml:"ATMCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (a *ATMCompletionAcknowledgementV01) AddHeader() *iso20022.Header21 {
	a.Header = new(iso20022.Header21)
	return a.Header
}

func (a *ATMCompletionAcknowledgementV01) AddProtectedATMCompletionAcknowledgement() *iso20022.ContentInformationType10 {
	a.ProtectedATMCompletionAcknowledgement = new(iso20022.ContentInformationType10)
	return a.ProtectedATMCompletionAcknowledgement
}

func (a *ATMCompletionAcknowledgementV01) AddATMCompletionAcknowledgement() *iso20022.ATMCompletionAcknowledgement1 {
	a.ATMCompletionAcknowledgement = new(iso20022.ATMCompletionAcknowledgement1)
	return a.ATMCompletionAcknowledgement
}

func (a *ATMCompletionAcknowledgementV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}

