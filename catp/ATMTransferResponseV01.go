package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Document"`
	Message *ATMTransferResponseV01 `xml:"ATMTrfRspn"`
}

func (d *Document01700101) AddMessage() *ATMTransferResponseV01 {
	d.Message = new(ATMTransferResponseV01)
	return d.Message
}

// The ATMTransferResponse message is sent by an acquirer or its agent to inform the ATM of the approval or decline of the transfer transaction.
type ATMTransferResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMTransferResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMTrfRspn,omitempty"`

	// Information related to the response of an ATM transfer from an ATM manager.
	ATMTransferResponse *iso20022.ATMTransferResponse1 `xml:"ATMTrfRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMTransferResponseV01) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMTransferResponseV01) AddProtectedATMTransferResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMTransferResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMTransferResponse
}

func (a *ATMTransferResponseV01) AddATMTransferResponse() *iso20022.ATMTransferResponse1 {
	a.ATMTransferResponse = new(iso20022.ATMTransferResponse1)
	return a.ATMTransferResponse
}

func (a *ATMTransferResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
