package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00400102 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.02 Document"`
	Message *ATMKeyDownloadResponseV02 `xml:"ATMKeyDwnldRspn"`
}

func (d *Document00400102) AddMessage() *ATMKeyDownloadResponseV02 {
	d.Message = new(ATMKeyDownloadResponseV02)
	return d.Message
}

// The ATMKeyDownloadResponse message is sent from an acquirer to an ATM in response to an ATMKeyDownloadRequest message, to download of one or several cryptographic keys.
type ATMKeyDownloadResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMKeyDownloadResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMKeyDwnldRspn,omitempty"`

	// Information related to the response of an ATM key download from an ATM manager.
	ATMKeyDownloadResponse *iso20022.ATMKeyDownloadResponse2 `xml:"ATMKeyDwnldRspn,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMKeyDownloadResponseV02) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMKeyDownloadResponseV02) AddProtectedATMKeyDownloadResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMKeyDownloadResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMKeyDownloadResponse
}

func (a *ATMKeyDownloadResponseV02) AddATMKeyDownloadResponse() *iso20022.ATMKeyDownloadResponse2 {
	a.ATMKeyDownloadResponse = new(iso20022.ATMKeyDownloadResponse2)
	return a.ATMKeyDownloadResponse
}

func (a *ATMKeyDownloadResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType13 {
	a.SecurityTrailer = new(iso20022.ContentInformationType13)
	return a.SecurityTrailer
}
