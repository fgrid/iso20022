package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.01 Document"`
	Message *MoneyMarketUnsecuredMarketStatisticalReportV01 `xml:"MnyMktUscrdMktSttstclRpt"`
}

func (d *Document01300101) AddMessage() *MoneyMarketUnsecuredMarketStatisticalReportV01 {
	d.Message = new(MoneyMarketUnsecuredMarketStatisticalReportV01)
	return d.Message
}

// The MoneyMarketUnsecuredMarketStatisticalReport message is sent by the reporting agents to the relevant competent authority, to report all relevant unsecured money market transactions. 
type MoneyMarketUnsecuredMarketStatisticalReportV01 struct {

	// Provides the elements specific to the report.
	ReportHeader *iso20022.MoneyMarketReportHeader1 `xml:"RptHdr"`

	// Provides the reason why no activity is reported or the required list of transactions for the unsecured market segment.
	UnsecuredMarketReport *iso20022.UnsecuredMarketReport3Choice `xml:"UscrdMktRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MoneyMarketUnsecuredMarketStatisticalReportV01) AddReportHeader() *iso20022.MoneyMarketReportHeader1 {
	m.ReportHeader = new(iso20022.MoneyMarketReportHeader1)
	return m.ReportHeader
}

func (m *MoneyMarketUnsecuredMarketStatisticalReportV01) AddUnsecuredMarketReport() *iso20022.UnsecuredMarketReport3Choice {
	m.UnsecuredMarketReport = new(iso20022.UnsecuredMarketReport3Choice)
	return m.UnsecuredMarketReport
}

func (m *MoneyMarketUnsecuredMarketStatisticalReportV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

