package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.04 Document"`
	Message *AcceptorCancellationResponseV04 `xml:"AccptrCxlRspn"`
}

func (d *Document00600104) AddMessage() *AcceptorCancellationResponseV04 {
	d.Message = new(AcceptorCancellationResponseV04)
	return d.Message
}

// The AcceptorCancellationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the outcome of the cancellation request. If the response is positive, the acquirer has voided the financial data from the captured transaction.
type AcceptorCancellationResponseV04 struct {

	// Cancellation response message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the cancellation response.
	CancellationResponse *iso20022.AcceptorCancellationResponse4 `xml:"CxlRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationResponseV04) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorCancellationResponseV04) AddCancellationResponse() *iso20022.AcceptorCancellationResponse4 {
	a.CancellationResponse = new(iso20022.AcceptorCancellationResponse4)
	return a.CancellationResponse
}

func (a *AcceptorCancellationResponseV04) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}

