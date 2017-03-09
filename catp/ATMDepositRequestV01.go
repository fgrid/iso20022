package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.012.001.01 Document"`
	Message *ATMDepositRequestV01 `xml:"ATMDpstReq"`
}

func (d *Document01200101) AddMessage() *ATMDepositRequestV01 {
	d.Message = new(ATMDepositRequestV01)
	return d.Message
}

// The ATMDepositRequest message is sent by an ATM to an acquirer or its agent to request the approval of a deposit transaction at an ATM, before or after deposit media inside the ATM.
type ATMDepositRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositRequest *iso20022.ContentInformationType10 `xml:"PrtctdATMDpstReq,omitempty"`

	// Information related to the request of a deposit transaction from an ATM.
	ATMDepositRequest *iso20022.ATMDepositRequest1 `xml:"ATMDpstReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositRequestV01) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMDepositRequestV01) AddProtectedATMDepositRequest() *iso20022.ContentInformationType10 {
	a.ProtectedATMDepositRequest = new(iso20022.ContentInformationType10)
	return a.ProtectedATMDepositRequest
}

func (a *ATMDepositRequestV01) AddATMDepositRequest() *iso20022.ATMDepositRequest1 {
	a.ATMDepositRequest = new(iso20022.ATMDepositRequest1)
	return a.ATMDepositRequest
}

func (a *ATMDepositRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
