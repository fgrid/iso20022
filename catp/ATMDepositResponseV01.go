package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.013.001.01 Document"`
	Message *ATMDepositResponseV01 `xml:"ATMDpstRspn"`
}

func (d *Document01300101) AddMessage() *ATMDepositResponseV01 {
	d.Message = new(ATMDepositResponseV01)
	return d.Message
}

// The ATMDepositResponse message is sent by an ATM manager or its agent to inform the ATM of the approval or decline of the deposit transaction.
type ATMDepositResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMDpstRspn,omitempty"`

	// Response to a deposit request.
	ATMDepositResponse *iso20022.ATMDepositResponse1 `xml:"ATMDpstRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositResponseV01) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMDepositResponseV01) AddProtectedATMDepositResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMDepositResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMDepositResponse
}

func (a *ATMDepositResponseV01) AddATMDepositResponse() *iso20022.ATMDepositResponse1 {
	a.ATMDepositResponse = new(iso20022.ATMDepositResponse1)
	return a.ATMDepositResponse
}

func (a *ATMDepositResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
