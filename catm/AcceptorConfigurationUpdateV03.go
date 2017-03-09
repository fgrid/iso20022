package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.03 Document"`
	Message *AcceptorConfigurationUpdateV03 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300103) AddMessage() *AcceptorConfigurationUpdateV03 {
	d.Message = new(AcceptorConfigurationUpdateV03)
	return d.Message
}

// Update of the acceptor configuration to be dowloaded by the terminal management system.
type AcceptorConfigurationUpdateV03 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *iso20022.AcceptorConfiguration3 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType9 `xml:"SctyTrlr"`
}

func (a *AcceptorConfigurationUpdateV03) AddHeader() *iso20022.Header4 {
	a.Header = new(iso20022.Header4)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV03) AddAcceptorConfiguration() *iso20022.AcceptorConfiguration3 {
	a.AcceptorConfiguration = new(iso20022.AcceptorConfiguration3)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV03) AddSecurityTrailer() *iso20022.ContentInformationType9 {
	a.SecurityTrailer = new(iso20022.ContentInformationType9)
	return a.SecurityTrailer
}
