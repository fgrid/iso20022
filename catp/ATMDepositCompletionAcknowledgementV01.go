package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.015.001.01 Document"`
	Message *ATMDepositCompletionAcknowledgementV01 `xml:"ATMDpstCmpltnAck"`
}

func (d *Document01500101) AddMessage() *ATMDepositCompletionAcknowledgementV01 {
	d.Message = new(ATMDepositCompletionAcknowledgementV01)
	return d.Message
}

// The ATMDepositCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMDepositCompletionAdvice message.
type ATMDepositCompletionAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositCompletionAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdATMDpstCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM deposit transaction from the ATM manager.
	ATMDepositCompletionAcknowledgement *iso20022.ATMDepositCompletionAcknowledgement1 `xml:"ATMDpstCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositCompletionAcknowledgementV01) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMDepositCompletionAcknowledgementV01) AddProtectedATMDepositCompletionAcknowledgement() *iso20022.ContentInformationType10 {
	a.ProtectedATMDepositCompletionAcknowledgement = new(iso20022.ContentInformationType10)
	return a.ProtectedATMDepositCompletionAcknowledgement
}

func (a *ATMDepositCompletionAcknowledgementV01) AddATMDepositCompletionAcknowledgement() *iso20022.ATMDepositCompletionAcknowledgement1 {
	a.ATMDepositCompletionAcknowledgement = new(iso20022.ATMDepositCompletionAcknowledgement1)
	return a.ATMDepositCompletionAcknowledgement
}

func (a *ATMDepositCompletionAcknowledgementV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
