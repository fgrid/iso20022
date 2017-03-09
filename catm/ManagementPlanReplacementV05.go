package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200105 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.05 Document"`
	Message *ManagementPlanReplacementV05 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200105) AddMessage() *ManagementPlanReplacementV05 {
	d.Message = new(ManagementPlanReplacementV05)
	return d.Message
}

// Terminal maintenance actions to be performed by a point of interaction (POI).
type ManagementPlanReplacementV05 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *iso20022.Header27 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *iso20022.ManagementPlan5 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (m *ManagementPlanReplacementV05) AddHeader() *iso20022.Header27 {
	m.Header = new(iso20022.Header27)
	return m.Header
}

func (m *ManagementPlanReplacementV05) AddManagementPlan() *iso20022.ManagementPlan5 {
	m.ManagementPlan = new(iso20022.ManagementPlan5)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV05) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	m.SecurityTrailer = new(iso20022.ContentInformationType12)
	return m.SecurityTrailer
}
