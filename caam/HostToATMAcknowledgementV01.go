package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.008.001.01 Document"`
	Message *HostToATMAcknowledgementV01 `xml:"HstToATMAck"`
}

func (d *Document00800101) AddMessage() *HostToATMAcknowledgementV01 {
	d.Message = new(HostToATMAcknowledgementV01)
	return d.Message
}

// The HostToATMAcknowledgement message is sent by an ATM to a host to acknowledge the receipt of a HostToATMRequest message.
type HostToATMAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedHostToATMAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdHstToATMAck,omitempty"`

	// Information related to the acknowledgement from an ATM to contact the ATM manager.
	HostToATMAcknowledgement *iso20022.HostToATMAcknowledgement1 `xml:"HstToATMAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (h *HostToATMAcknowledgementV01) AddHeader() *iso20022.Header20 {
	h.Header = new(iso20022.Header20)
	return h.Header
}

func (h *HostToATMAcknowledgementV01) AddProtectedHostToATMAcknowledgement() *iso20022.ContentInformationType10 {
	h.ProtectedHostToATMAcknowledgement = new(iso20022.ContentInformationType10)
	return h.ProtectedHostToATMAcknowledgement
}

func (h *HostToATMAcknowledgementV01) AddHostToATMAcknowledgement() *iso20022.HostToATMAcknowledgement1 {
	h.HostToATMAcknowledgement = new(iso20022.HostToATMAcknowledgement1)
	return h.HostToATMAcknowledgement
}

func (h *HostToATMAcknowledgementV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	h.SecurityTrailer = new(iso20022.ContentInformationType15)
	return h.SecurityTrailer
}
