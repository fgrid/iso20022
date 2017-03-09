package reda

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document06100101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.061.001.01 Document"`
	Message *NettingCutOffReferenceDataReportV01 `xml:"NetgCutOffRefDataRpt"`
}

func (d *Document06100101) AddMessage() *NettingCutOffReferenceDataReportV01 {
	d.Message = new(NettingCutOffReferenceDataReportV01)
	return d.Message
}

// The Netting Cut Off Reference Data Report message is sent to a participant by a central system to provide details of scheduled, changed, existing and previous netting cut off data held at a central system.
type NettingCutOffReferenceDataReportV01 struct {

	// Specifies the meta data for the netting cut off report including message pagination.
	ReportData *iso20022.NettingCutOffReportData1 `xml:"RptData"`

	// Provides the latest information related to the status of a netting cut off held at a central system.
	ParticipantNettingCutOffData []*iso20022.CutOffData1 `xml:"PtcptNetgCutOffData"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NettingCutOffReferenceDataReportV01) AddReportData() *iso20022.NettingCutOffReportData1 {
	n.ReportData = new(iso20022.NettingCutOffReportData1)
	return n.ReportData
}

func (n *NettingCutOffReferenceDataReportV01) AddParticipantNettingCutOffData() *iso20022.CutOffData1 {
	newValue := new(iso20022.CutOffData1)
	n.ParticipantNettingCutOffData = append(n.ParticipantNettingCutOffData, newValue)
	return newValue
}

func (n *NettingCutOffReferenceDataReportV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
