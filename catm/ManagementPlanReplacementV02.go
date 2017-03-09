package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.02 Document"`
	Message *ManagementPlanReplacementV02 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200102) AddMessage() *ManagementPlanReplacementV02 {
	d.Message = new(ManagementPlanReplacementV02)
	return d.Message
}

// Terminal maintenance actions to be performed by a point of interaction (POI).
type ManagementPlanReplacementV02 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *iso20022.ManagementPlan2 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType4 `xml:"SctyTrlr"`
}

func (m *ManagementPlanReplacementV02) AddHeader() *iso20022.Header4 {
	m.Header = new(iso20022.Header4)
	return m.Header
}

func (m *ManagementPlanReplacementV02) AddManagementPlan() *iso20022.ManagementPlan2 {
	m.ManagementPlan = new(iso20022.ManagementPlan2)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV02) AddSecurityTrailer() *iso20022.ContentInformationType4 {
	m.SecurityTrailer = new(iso20022.ContentInformationType4)
	return m.SecurityTrailer
}
