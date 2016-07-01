package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:cain.012.001.01 Document"`
	Message *KeyExchangeResponse `xml:"KeyXchgRspn"`
}

func (d *Document01200101) AddMessage() *KeyExchangeResponse {
	d.Message = new(KeyExchangeResponse)
	return d.Message
}

// The KeyExchangeResponse message is sent by an acquirer, an issuer or an agent to answer to a KeyExchangeInitiation message and complete a cryptographic key exchange.
type KeyExchangeResponse struct {

	// Information related to the protocol management.
	Header *iso20022.Header17 `xml:"Hdr"`

	// Information related to the response to a key exchange.
	KeyExchangeResponse *iso20022.AcquirerKeyExchangeResponse1 `xml:"KeyXchgRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`

}


func (k *KeyExchangeResponse) AddHeader() *iso20022.Header17 {
	k.Header = new(iso20022.Header17)
	return k.Header
}

func (k *KeyExchangeResponse) AddKeyExchangeResponse() *iso20022.AcquirerKeyExchangeResponse1 {
	k.KeyExchangeResponse = new(iso20022.AcquirerKeyExchangeResponse1)
	return k.KeyExchangeResponse
}

func (k *KeyExchangeResponse) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	k.SecurityTrailer = new(iso20022.ContentInformationType12)
	return k.SecurityTrailer
}

