package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.005.001.02 Document"`
	Message *ATMDiagnosticRequestV02 `xml:"ATMDgnstcReq"`
}

func (d *Document00500102) AddMessage() *ATMDiagnosticRequestV02 {
	d.Message = new(ATMDiagnosticRequestV02)
	return d.Message
}

// The ATMDiagnosticRequest message is sent from an ATM to an acquirer to verify the availability of the acquirer. The acquirer will also validate that this ATM is a valid ATM for its particular network.
type ATMDiagnosticRequestV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDiagnosticRequest *iso20022.ContentInformationType10 `xml:"PrtctdATMDgnstcReq,omitempty"`

	// Information related to the request of a diagnostic from an ATM.
	ATMDiagnosticRequest *iso20022.ATMDiagnosticRequest2 `xml:"ATMDgnstcReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDiagnosticRequestV02) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMDiagnosticRequestV02) AddProtectedATMDiagnosticRequest() *iso20022.ContentInformationType10 {
	a.ProtectedATMDiagnosticRequest = new(iso20022.ContentInformationType10)
	return a.ProtectedATMDiagnosticRequest
}

func (a *ATMDiagnosticRequestV02) AddATMDiagnosticRequest() *iso20022.ATMDiagnosticRequest2 {
	a.ATMDiagnosticRequest = new(iso20022.ATMDiagnosticRequest2)
	return a.ATMDiagnosticRequest
}

func (a *ATMDiagnosticRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
