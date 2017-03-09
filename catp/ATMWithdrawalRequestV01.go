package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:catp.001.001.01 Document"`
	Message *ATMWithdrawalRequestV01 `xml:"ATMWdrwlReq"`
}

func (d *Document00100101) AddMessage() *ATMWithdrawalRequestV01 {
	d.Message = new(ATMWithdrawalRequestV01)
	return d.Message
}

// The ATMWithdrawalRequest message is sent by an ATM to an acquirer or its agent to request the approval of a withdrawal transaction at an ATM.
type ATMWithdrawalRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalRequest *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlReq,omitempty"`

	// Information related to the request of a withdrawal transaction from an ATM.
	ATMWithdrawalRequest *iso20022.ATMWithdrawalRequest1 `xml:"ATMWdrwlReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalRequestV01) AddHeader() *iso20022.Header20 {
	a.Header = new(iso20022.Header20)
	return a.Header
}

func (a *ATMWithdrawalRequestV01) AddProtectedATMWithdrawalRequest() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalRequest = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalRequest
}

func (a *ATMWithdrawalRequestV01) AddATMWithdrawalRequest() *iso20022.ATMWithdrawalRequest1 {
	a.ATMWithdrawalRequest = new(iso20022.ATMWithdrawalRequest1)
	return a.ATMWithdrawalRequest
}

func (a *ATMWithdrawalRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
