package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.003.001.01 Document"`
	Message *ATMWithdrawalCompletionAdviceV01 `xml:"ATMWdrwlCmpltnAdvc"`
}

func (d *Document00300101) AddMessage() *ATMWithdrawalCompletionAdviceV01 {
	d.Message = new(ATMWithdrawalCompletionAdviceV01)
	return d.Message
}

// The ATMWithdrawalCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a withdrawal transaction at an ATM.
// If the ATM is configured to only send negative completion, a generic completion message should be used instead of ATMCompletionAdvice.
type ATMWithdrawalCompletionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAdvc,omitempty"`

	// Information related to the completion of a withdrawal transaction on the ATM.
	ATMWithdrawalCompletionAdvice *iso20022.ATMWithdrawalCompletionAdvice1 `xml:"ATMWdrwlCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAdviceV01) AddHeader() *iso20022.Header21 {
	a.Header = new(iso20022.Header21)
	return a.Header
}

func (a *ATMWithdrawalCompletionAdviceV01) AddProtectedATMWithdrawalCompletionAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV01) AddATMWithdrawalCompletionAdvice() *iso20022.ATMWithdrawalCompletionAdvice1 {
	a.ATMWithdrawalCompletionAdvice = new(iso20022.ATMWithdrawalCompletionAdvice1)
	return a.ATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
