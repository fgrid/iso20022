package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:catm.005.001.01 Document"`
	Message *MaintenanceDelegationRequestV01 `xml:"MntncDlgtnReq"`
}

func (d *Document00500101) AddMessage() *MaintenanceDelegationRequestV01 {
	d.Message = new(MaintenanceDelegationRequestV01)
	return d.Message
}

// A terminal manager requests to the master terminal manager the delegation of maintenance functions or maintenance operation on the terminal estate managed by the master terminal manager.
type MaintenanceDelegationRequestV01 struct {

	// Information related to the protocol management.
	Header *iso20022.Header16 `xml:"Hdr"`

	// Information related to the request of maintenance delegations.
	MaintenanceDelegationRequest *iso20022.MaintenanceDelegationRequest1 `xml:"MntncDlgtnReq"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`

}


func (m *MaintenanceDelegationRequestV01) AddHeader() *iso20022.Header16 {
	m.Header = new(iso20022.Header16)
	return m.Header
}

func (m *MaintenanceDelegationRequestV01) AddMaintenanceDelegationRequest() *iso20022.MaintenanceDelegationRequest1 {
	m.MaintenanceDelegationRequest = new(iso20022.MaintenanceDelegationRequest1)
	return m.MaintenanceDelegationRequest
}

func (m *MaintenanceDelegationRequestV01) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	m.SecurityTrailer = new(iso20022.ContentInformationType12)
	return m.SecurityTrailer
}

