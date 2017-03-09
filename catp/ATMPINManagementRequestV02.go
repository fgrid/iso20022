package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:catp.010.001.02 Document"`
	Message *ATMPINManagementRequestV02 `xml:"ATMPINMgmtReq"`
}

func (d *Document01000102) AddMessage() *ATMPINManagementRequestV02 {
	d.Message = new(ATMPINManagementRequestV02)
	return d.Message
}

// The ATMPINManagementRequest message is sent by an ATM to an ATM manager to request an operation on the cardholder PIN.
type ATMPINManagementRequestV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMPINManagementRequest *iso20022.ContentInformationType10 `xml:"PrtctdATMPINMgmtReq,omitempty"`

	// Information related to the request of a PIN management from an ATM.
	ATMPINManagementRequest *iso20022.ATMPINManagementRequest2 `xml:"ATMPINMgmtReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMPINManagementRequestV02) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMPINManagementRequestV02) AddProtectedATMPINManagementRequest() *iso20022.ContentInformationType10 {
	a.ProtectedATMPINManagementRequest = new(iso20022.ContentInformationType10)
	return a.ProtectedATMPINManagementRequest
}

func (a *ATMPINManagementRequestV02) AddATMPINManagementRequest() *iso20022.ATMPINManagementRequest2 {
	a.ATMPINManagementRequest = new(iso20022.ATMPINManagementRequest2)
	return a.ATMPINManagementRequest
}

func (a *ATMPINManagementRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
