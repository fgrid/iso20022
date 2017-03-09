package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.003.001.02 Document"`
	Message *ATMWithdrawalCompletionAdviceV02 `xml:"ATMWdrwlCmpltnAdvc"`
}

func (d *Document00300102) AddMessage() *ATMWithdrawalCompletionAdviceV02 {
	d.Message = new(ATMWithdrawalCompletionAdviceV02)
	return d.Message
}

// The ATMWithdrawalCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a withdrawal transaction at an ATM.
// If the ATM is configured to only send negative completion, a generic completion message should be used instead of ATMCompletionAdvice.
type ATMWithdrawalCompletionAdviceV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAdvc,omitempty"`

	// Information related to the completion of a withdrawal transaction on the ATM.
	ATMWithdrawalCompletionAdvice *iso20022.ATMWithdrawalCompletionAdvice2 `xml:"ATMWdrwlCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAdviceV02) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMWithdrawalCompletionAdviceV02) AddProtectedATMWithdrawalCompletionAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV02) AddATMWithdrawalCompletionAdvice() *iso20022.ATMWithdrawalCompletionAdvice2 {
	a.ATMWithdrawalCompletionAdvice = new(iso20022.ATMWithdrawalCompletionAdvice2)
	return a.ATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
