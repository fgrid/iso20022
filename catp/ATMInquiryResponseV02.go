package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.007.001.02 Document"`
	Message *ATMInquiryResponseV02 `xml:"ATMNqryRspn"`
}

func (d *Document00700102) AddMessage() *ATMInquiryResponseV02 {
	d.Message = new(ATMInquiryResponseV02)
	return d.Message
}

// The ATMInquiryResponse message is sent by an ATM manager or its agent to the ATM to provide the information and the outcome of the verifications requested in the ATMInquiryRequest.
type ATMInquiryResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMInquiryResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMNqryRspn,omitempty"`

	// Information related to the response of an ATM inquiry from an ATM manager.
	ATMInquiryResponse *iso20022.ATMInquiryResponse2 `xml:"ATMNqryRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMInquiryResponseV02) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMInquiryResponseV02) AddProtectedATMInquiryResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMInquiryResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMInquiryResponse
}

func (a *ATMInquiryResponseV02) AddATMInquiryResponse() *iso20022.ATMInquiryResponse2 {
	a.ATMInquiryResponse = new(iso20022.ATMInquiryResponse2)
	return a.ATMInquiryResponse
}

func (a *ATMInquiryResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
