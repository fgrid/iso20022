package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.02 Document"`
	Message *AcceptorRejectionV02 `xml:"AccptrRjctn"`
}

func (d *Document01500102) AddMessage() *AcceptorRejectionV02 {
	d.Message = new(AcceptorRejectionV02)
	return d.Message
}

// The AcceptorRejection message is sent by the acquirer (or its agent) to reject a message request or advice sent by an acceptor (or its agent), to indicate that the received message could not be processed.
type AcceptorRejectionV02 struct {

	// Rejection message management information.
	Header *iso20022.Header5 `xml:"Hdr"`

	// Information related to the reject.
	Reject *iso20022.AcceptorRejection1 `xml:"Rjct"`

}


func (a *AcceptorRejectionV02) AddHeader() *iso20022.Header5 {
	a.Header = new(iso20022.Header5)
	return a.Header
}

func (a *AcceptorRejectionV02) AddReject() *iso20022.AcceptorRejection1 {
	a.Reject = new(iso20022.AcceptorRejection1)
	return a.Reject
}

