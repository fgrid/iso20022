package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500104 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Document"`
	Message *AcceptorRejectionV04 `xml:"AccptrRjctn"`
}

func (d *Document01500104) AddMessage() *AcceptorRejectionV04 {
	d.Message = new(AcceptorRejectionV04)
	return d.Message
}

// The AcceptorRejection message is sent by the acquirer (or its agent) to reject a message request or advice sent by an acceptor (or its agent), to indicate that the received message could not be processed.
type AcceptorRejectionV04 struct {

	// Rejection message management information.
	Header *iso20022.Header13 `xml:"Hdr"`

	// Information related to the reject.
	Reject *iso20022.AcceptorRejection2 `xml:"Rjct"`
}

func (a *AcceptorRejectionV04) AddHeader() *iso20022.Header13 {
	a.Header = new(iso20022.Header13)
	return a.Header
}

func (a *AcceptorRejectionV04) AddReject() *iso20022.AcceptorRejection2 {
	a.Reject = new(iso20022.AcceptorRejection2)
	return a.Reject
}
