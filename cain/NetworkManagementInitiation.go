package cain

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.009.001.01 Document"`
	Message *NetworkManagementInitiation `xml:"NtwkMgmtInitn"`
}

func (d *Document00900101) AddMessage() *NetworkManagementInitiation {
	d.Message = new(NetworkManagementInitiation)
	return d.Message
}

// The NetworkManagementInitiation message covers the range of activities to control the operating condition of the network and may be initiated by any party to an acquirer, an issuer or an agent.
type NetworkManagementInitiation struct {

	// Information related to the protocol management.
	Header *iso20022.Header17 `xml:"Hdr"`

	// Information related to the network management.
	NetworkManagementInitiation *iso20022.AcquirerNetworkManagementInitiation1 `xml:"NtwkMgmtInitn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (n *NetworkManagementInitiation) AddHeader() *iso20022.Header17 {
	n.Header = new(iso20022.Header17)
	return n.Header
}

func (n *NetworkManagementInitiation) AddNetworkManagementInitiation() *iso20022.AcquirerNetworkManagementInitiation1 {
	n.NetworkManagementInitiation = new(iso20022.AcquirerNetworkManagementInitiation1)
	return n.NetworkManagementInitiation
}

func (n *NetworkManagementInitiation) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	n.SecurityTrailer = new(iso20022.ContentInformationType15)
	return n.SecurityTrailer
}
