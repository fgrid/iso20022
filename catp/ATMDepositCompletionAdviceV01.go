package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.014.001.01 Document"`
	Message *ATMDepositCompletionAdviceV01 `xml:"ATMDpstCmpltnAdvc"`
}

func (d *Document01400101) AddMessage() *ATMDepositCompletionAdviceV01 {
	d.Message = new(ATMDepositCompletionAdviceV01)
	return d.Message
}

// The ATMDepositCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a deposit transaction at an ATM.
// If the ATM is configured to only send negative completion, a generic completion message should be used instead of ATMCompletionAdvice.
// This message can be used each time a bundle is put in the ATM safe and/or at the end of a multi bundle deposit.
type ATMDepositCompletionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositCompletionAdvice *iso20022.ContentInformationType10 `xml:"PrtctdATMDpstCmpltnAdvc,omitempty"`

	// Information related to the completion of a deposit transaction on the ATM.
	ATMDepositCompletionAdvice *iso20022.ATMDepositCompletionAdvice1 `xml:"ATMDpstCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositCompletionAdviceV01) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMDepositCompletionAdviceV01) AddProtectedATMDepositCompletionAdvice() *iso20022.ContentInformationType10 {
	a.ProtectedATMDepositCompletionAdvice = new(iso20022.ContentInformationType10)
	return a.ProtectedATMDepositCompletionAdvice
}

func (a *ATMDepositCompletionAdviceV01) AddATMDepositCompletionAdvice() *iso20022.ATMDepositCompletionAdvice1 {
	a.ATMDepositCompletionAdvice = new(iso20022.ATMDepositCompletionAdvice1)
	return a.ATMDepositCompletionAdvice
}

func (a *ATMDepositCompletionAdviceV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
