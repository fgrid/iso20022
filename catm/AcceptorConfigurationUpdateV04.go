package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300104 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.04 Document"`
	Message *AcceptorConfigurationUpdateV04 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300104) AddMessage() *AcceptorConfigurationUpdateV04 {
	d.Message = new(AcceptorConfigurationUpdateV04)
	return d.Message
}

// Update of the acceptor configuration to be downloaded by the terminal management system.
type AcceptorConfigurationUpdateV04 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *iso20022.Header14 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *iso20022.AcceptorConfiguration4 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`
}

func (a *AcceptorConfigurationUpdateV04) AddHeader() *iso20022.Header14 {
	a.Header = new(iso20022.Header14)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV04) AddAcceptorConfiguration() *iso20022.AcceptorConfiguration4 {
	a.AcceptorConfiguration = new(iso20022.AcceptorConfiguration4)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV04) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	a.SecurityTrailer = new(iso20022.ContentInformationType12)
	return a.SecurityTrailer
}
