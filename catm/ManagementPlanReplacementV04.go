package catm

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200104 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.04 Document"`
	Message *ManagementPlanReplacementV04 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200104) AddMessage() *ManagementPlanReplacementV04 {
	d.Message = new(ManagementPlanReplacementV04)
	return d.Message
}

// Terminal maintenance actions to be performed by a point of interaction (POI).
type ManagementPlanReplacementV04 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *iso20022.Header14 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *iso20022.ManagementPlan4 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`
}

func (m *ManagementPlanReplacementV04) AddHeader() *iso20022.Header14 {
	m.Header = new(iso20022.Header14)
	return m.Header
}

func (m *ManagementPlanReplacementV04) AddManagementPlan() *iso20022.ManagementPlan4 {
	m.ManagementPlan = new(iso20022.ManagementPlan4)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV04) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	m.SecurityTrailer = new(iso20022.ContentInformationType12)
	return m.SecurityTrailer
}
