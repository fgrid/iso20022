package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Document"`
	Message *HostToATMRequestV01 `xml:"HstToATMReq"`
}

func (d *Document00700101) AddMessage() *HostToATMRequestV01 {
	d.Message = new(HostToATMRequestV01)
	return d.Message
}

// The HostToATMRequest message is sent by a host to an ATM to request the ATM to contact a host by sending of a maintenance messages.
type HostToATMRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedHostToATMRequest *iso20022.ContentInformationType10 `xml:"PrtctdHstToATMReq,omitempty"`

	// Information related to the request to an ATM to contact the ATM manager.
	HostToATMRequest *iso20022.HostToATMRequest1 `xml:"HstToATMReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (h *HostToATMRequestV01) AddHeader() *iso20022.Header20 {
	h.Header = new(iso20022.Header20)
	return h.Header
}

func (h *HostToATMRequestV01) AddProtectedHostToATMRequest() *iso20022.ContentInformationType10 {
	h.ProtectedHostToATMRequest = new(iso20022.ContentInformationType10)
	return h.ProtectedHostToATMRequest
}

func (h *HostToATMRequestV01) AddHostToATMRequest() *iso20022.HostToATMRequest1 {
	h.HostToATMRequest = new(iso20022.HostToATMRequest1)
	return h.HostToATMRequest
}

func (h *HostToATMRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	h.SecurityTrailer = new(iso20022.ContentInformationType15)
	return h.SecurityTrailer
}

