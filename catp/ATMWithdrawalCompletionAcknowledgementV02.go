package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400102 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.004.001.02 Document"`
	Message *ATMWithdrawalCompletionAcknowledgementV02 `xml:"ATMWdrwlCmpltnAck"`
}

func (d *Document00400102) AddMessage() *ATMWithdrawalCompletionAcknowledgementV02 {
	d.Message = new(ATMWithdrawalCompletionAcknowledgementV02)
	return d.Message
}

// The ATMWithdrawalCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMWithdrawalCompletionAdvice message.
type ATMWithdrawalCompletionAcknowledgementV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM withdrawal transaction from the ATM manager.
	ATMWithdrawalCompletionAcknowledgement *iso20022.ATMWithdrawalCompletionAcknowledgement2 `xml:"ATMWdrwlCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddProtectedATMWithdrawalCompletionAcknowledgement() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAcknowledgement = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAcknowledgement
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddATMWithdrawalCompletionAcknowledgement() *iso20022.ATMWithdrawalCompletionAcknowledgement2 {
	a.ATMWithdrawalCompletionAcknowledgement = new(iso20022.ATMWithdrawalCompletionAcknowledgement2)
	return a.ATMWithdrawalCompletionAcknowledgement
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
