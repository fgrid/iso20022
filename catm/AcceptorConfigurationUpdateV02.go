package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.02 Document"`
	Message *AcceptorConfigurationUpdateV02 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300102) AddMessage() *AcceptorConfigurationUpdateV02 {
	d.Message = new(AcceptorConfigurationUpdateV02)
	return d.Message
}

// Update of the acceptor configuration to be dowloaded by the terminal management system.
type AcceptorConfigurationUpdateV02 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *iso20022.AcceptorConfiguration2 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType4 `xml:"SctyTrlr"`
}

func (a *AcceptorConfigurationUpdateV02) AddHeader() *iso20022.Header4 {
	a.Header = new(iso20022.Header4)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV02) AddAcceptorConfiguration() *iso20022.AcceptorConfiguration2 {
	a.AcceptorConfiguration = new(iso20022.AcceptorConfiguration2)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV02) AddSecurityTrailer() *iso20022.ContentInformationType4 {
	a.SecurityTrailer = new(iso20022.ContentInformationType4)
	return a.SecurityTrailer
}
