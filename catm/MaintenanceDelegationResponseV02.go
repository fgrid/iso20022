package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.006.001.02 Document"`
	Message *MaintenanceDelegationResponseV02 `xml:"MntncDlgtnRspn"`
}

func (d *Document00600102) AddMessage() *MaintenanceDelegationResponseV02 {
	d.Message = new(MaintenanceDelegationResponseV02)
	return d.Message
}

// The master terminal manager provides the outcome of a maintenance delegation request to a terminal manager.
type MaintenanceDelegationResponseV02 struct {

	// Maintenance delegation response message management information.
	Header *iso20022.Header29 `xml:"Hdr"`

	// Information related to the request of maintenance delegations.
	MaintenanceDelegationResponse *iso20022.MaintenanceDelegationResponse2 `xml:"MntncDlgtnRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (m *MaintenanceDelegationResponseV02) AddHeader() *iso20022.Header29 {
	m.Header = new(iso20022.Header29)
	return m.Header
}

func (m *MaintenanceDelegationResponseV02) AddMaintenanceDelegationResponse() *iso20022.MaintenanceDelegationResponse2 {
	m.MaintenanceDelegationResponse = new(iso20022.MaintenanceDelegationResponse2)
	return m.MaintenanceDelegationResponse
}

func (m *MaintenanceDelegationResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	m.SecurityTrailer = new(iso20022.ContentInformationType12)
	return m.SecurityTrailer
}
