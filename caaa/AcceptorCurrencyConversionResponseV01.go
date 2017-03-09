package caaa

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.017.001.01 Document"`
	Message *AcceptorCurrencyConversionResponseV01 `xml:"AccptrCcyConvsRspn"`
}

func (d *Document01700101) AddMessage() *AcceptorCurrencyConversionResponseV01 {
	d.Message = new(AcceptorCurrencyConversionResponseV01)
	return d.Message
}

// The AcceptorCurrencyConversionResponse message is sent by currency conversion service provider to the card acceptor to return the result of a potential currency conversion for the cardholder.
//
type AcceptorCurrencyConversionResponseV01 struct {

	// Currency conversion response message management information.
	Header *iso20022.Header7 `xml:"Hdr"`

	// Information related to the outcome of the currency conversion.
	CurrencyConversionResponse *iso20022.AcceptorCurrencyConversionResponse1 `xml:"CcyConvsRspn"`

	// Trailer of the message containing a MAC (message authentication code).
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCurrencyConversionResponseV01) AddHeader() *iso20022.Header7 {
	a.Header = new(iso20022.Header7)
	return a.Header
}

func (a *AcceptorCurrencyConversionResponseV01) AddCurrencyConversionResponse() *iso20022.AcceptorCurrencyConversionResponse1 {
	a.CurrencyConversionResponse = new(iso20022.AcceptorCurrencyConversionResponse1)
	return a.CurrencyConversionResponse
}

func (a *AcceptorCurrencyConversionResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}
