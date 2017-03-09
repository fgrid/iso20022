package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600105 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.05 Document"`
	Message *AcceptorCancellationResponseV05 `xml:"AccptrCxlRspn"`
}

func (d *Document00600105) AddMessage() *AcceptorCancellationResponseV05 {
	d.Message = new(AcceptorCancellationResponseV05)
	return d.Message
}

// The AcceptorCancellationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the outcome of the cancellation request. If the response is positive, the acquirer has voided the financial data from the captured transaction.
type AcceptorCancellationResponseV05 struct {

	// Cancellation response message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the cancellation response.
	CancellationResponse *iso20022.AcceptorCancellationResponse5 `xml:"CxlRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCancellationResponseV05) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorCancellationResponseV05) AddCancellationResponse() *iso20022.AcceptorCancellationResponse5 {
	a.CancellationResponse = new(iso20022.AcceptorCancellationResponse5)
	return a.CancellationResponse
}

func (a *AcceptorCancellationResponseV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
