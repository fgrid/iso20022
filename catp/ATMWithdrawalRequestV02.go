package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100102 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.001.001.02 Document"`
	Message *ATMWithdrawalRequestV02 `xml:"ATMWdrwlReq"`
}

func (d *Document00100102) AddMessage() *ATMWithdrawalRequestV02 {
	d.Message = new(ATMWithdrawalRequestV02)
	return d.Message
}

// The ATMWithdrawalRequest message is sent by an ATM to an acquirer or its agent to request the approval of a withdrawal transaction at an ATM.
type ATMWithdrawalRequestV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalRequest *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlReq,omitempty"`

	// Information related to the request of a withdrawal transaction from an ATM.
	ATMWithdrawalRequest *iso20022.ATMWithdrawalRequest2 `xml:"ATMWdrwlReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalRequestV02) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMWithdrawalRequestV02) AddProtectedATMWithdrawalRequest() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalRequest = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalRequest
}

func (a *ATMWithdrawalRequestV02) AddATMWithdrawalRequest() *iso20022.ATMWithdrawalRequest2 {
	a.ATMWithdrawalRequest = new(iso20022.ATMWithdrawalRequest2)
	return a.ATMWithdrawalRequest
}

func (a *ATMWithdrawalRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
