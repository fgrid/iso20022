package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.016.001.02 Document"`
	Message *AcceptorCurrencyConversionRequestV02 `xml:"AccptrCcyConvsReq"`
}

func (d *Document01600102) AddMessage() *AcceptorCurrencyConversionRequestV02 {
	d.Message = new(AcceptorCurrencyConversionRequestV02)
	return d.Message
}

// The AcceptorCurrencyConversionRequest message is sent by the card acceptor to the currency conversion service provider to request if the cardholder is able to pay in the currency of its card.
//
type AcceptorCurrencyConversionRequestV02 struct {

	// Currency Conversion request message management information.
	Header *iso20022.Header10 `xml:"Hdr"`

	// Information related to the currency conversion request.
	CurrencyConversionRequest *iso20022.AcceptorCurrencyConversionRequest2 `xml:"CcyConvsReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCurrencyConversionRequestV02) AddHeader() *iso20022.Header10 {
	a.Header = new(iso20022.Header10)
	return a.Header
}

func (a *AcceptorCurrencyConversionRequestV02) AddCurrencyConversionRequest() *iso20022.AcceptorCurrencyConversionRequest2 {
	a.CurrencyConversionRequest = new(iso20022.AcceptorCurrencyConversionRequest2)
	return a.CurrencyConversionRequest
}

func (a *AcceptorCurrencyConversionRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType11 {
	a.SecurityTrailer = new(iso20022.ContentInformationType11)
	return a.SecurityTrailer
}
