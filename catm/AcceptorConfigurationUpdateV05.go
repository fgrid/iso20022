package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300105 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.05 Document"`
	Message *AcceptorConfigurationUpdateV05 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300105) AddMessage() *AcceptorConfigurationUpdateV05 {
	d.Message = new(AcceptorConfigurationUpdateV05)
	return d.Message
}

// Update of the acceptor configuration to be downloaded by the terminal management system.
type AcceptorConfigurationUpdateV05 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *iso20022.Header27 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *iso20022.AcceptorConfiguration5 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorConfigurationUpdateV05) AddHeader() *iso20022.Header27 {
	a.Header = new(iso20022.Header27)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV05) AddAcceptorConfiguration() *iso20022.AcceptorConfiguration5 {
	a.AcceptorConfiguration = new(iso20022.AcceptorConfiguration5)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV05) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	a.SecurityTrailer = new(iso20022.ContentInformationType12)
	return a.SecurityTrailer
}
