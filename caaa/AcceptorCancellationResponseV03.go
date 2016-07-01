package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.03 Document"`
	Message *AcceptorCancellationResponseV03 `xml:"AccptrCxlRspn"`
}

func (d *Document00600103) AddMessage() *AcceptorCancellationResponseV03 {
	d.Message = new(AcceptorCancellationResponseV03)
	return d.Message
}

// The AcceptorCancellationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the outcome of the cancellation request. If the response is positive, the acquirer has voided the financial data from the captured transaction.
type AcceptorCancellationResponseV03 struct {

	// Cancellation response message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the cancellation response.
	CancellationResponse *iso20022.AcceptorCancellationResponse3 `xml:"CxlRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`

}


func (a *AcceptorCancellationResponseV03) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorCancellationResponseV03) AddCancellationResponse() *iso20022.AcceptorCancellationResponse3 {
	a.CancellationResponse = new(iso20022.AcceptorCancellationResponse3)
	return a.CancellationResponse
}

func (a *AcceptorCancellationResponseV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}

