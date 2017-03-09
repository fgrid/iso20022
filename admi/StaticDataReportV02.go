package admi

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000102 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:admi.010.001.02 Document"`
	Message *StaticDataReportV02 `xml:"StatcDataRpt"`
}

func (d *Document01000102) AddMessage() *StaticDataReportV02 {
	d.Message = new(StaticDataReportV02)
	return d.Message
}

// The StaticDataReport message is sent by a central system to the participant to provide static data held in the system.
//
type StaticDataReportV02 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *iso20022.Max35Text `xml:"MsgId"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *iso20022.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Report type and returned data.
	ReportDetails *iso20022.RequestDetails5 `xml:"RptDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StaticDataReportV02) SetMessageIdentification(value string) {
	s.MessageIdentification = (*iso20022.Max35Text)(&value)
}

func (s *StaticDataReportV02) SetSettlementSessionIdentifier(value string) {
	s.SettlementSessionIdentifier = (*iso20022.Exact4AlphaNumericText)(&value)
}

func (s *StaticDataReportV02) AddReportDetails() *iso20022.RequestDetails5 {
	s.ReportDetails = new(iso20022.RequestDetails5)
	return s.ReportDetails
}

func (s *StaticDataReportV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
