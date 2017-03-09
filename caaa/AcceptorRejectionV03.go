package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.03 Document"`
	Message *AcceptorRejectionV03 `xml:"AccptrRjctn"`
}

func (d *Document01500103) AddMessage() *AcceptorRejectionV03 {
	d.Message = new(AcceptorRejectionV03)
	return d.Message
}

// The AcceptorRejection message is sent by the acquirer (or its agent) to reject a message request or advice sent by an acceptor (or its agent), to indicate that the received message could not be processed.
type AcceptorRejectionV03 struct {

	// Rejection message management information.
	Header *iso20022.Header9 `xml:"Hdr"`

	// Information related to the reject.
	Reject *iso20022.AcceptorRejection2 `xml:"Rjct"`
}

func (a *AcceptorRejectionV03) AddHeader() *iso20022.Header9 {
	a.Header = new(iso20022.Header9)
	return a.Header
}

func (a *AcceptorRejectionV03) AddReject() *iso20022.AcceptorRejection2 {
	a.Reject = new(iso20022.AcceptorRejection2)
	return a.Reject
}
