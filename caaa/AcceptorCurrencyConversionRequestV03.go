package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600103 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.016.001.03 Document"`
	Message *AcceptorCurrencyConversionRequestV03 `xml:"AccptrCcyConvsReq"`
}

func (d *Document01600103) AddMessage() *AcceptorCurrencyConversionRequestV03 {
	d.Message = new(AcceptorCurrencyConversionRequestV03)
	return d.Message
}

// The AcceptorCurrencyConversionRequest message is sent by the card acceptor to the currency conversion service provider to request if the cardholder is able to pay in the currency of its card.
//
type AcceptorCurrencyConversionRequestV03 struct {

	// Currency Conversion request message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the currency conversion request.
	CurrencyConversionRequest *iso20022.AcceptorCurrencyConversionRequest3 `xml:"CcyConvsReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCurrencyConversionRequestV03) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorCurrencyConversionRequestV03) AddCurrencyConversionRequest() *iso20022.AcceptorCurrencyConversionRequest3 {
	a.CurrencyConversionRequest = new(iso20022.AcceptorCurrencyConversionRequest3)
	return a.CurrencyConversionRequest
}

func (a *AcceptorCurrencyConversionRequestV03) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
