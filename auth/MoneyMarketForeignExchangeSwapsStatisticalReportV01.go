package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:auth.014.001.01 Document"`
	Message *MoneyMarketForeignExchangeSwapsStatisticalReportV01 `xml:"MnyMktFXSwpsSttstclRpt"`
}

func (d *Document01400101) AddMessage() *MoneyMarketForeignExchangeSwapsStatisticalReportV01 {
	d.Message = new(MoneyMarketForeignExchangeSwapsStatisticalReportV01)
	return d.Message
}

// The MoneyMarketSecuredMarketStatisticalReport message is sent by the reporting agents  to the relevant competent authority, to report all daily Foreign Exchange Swaps (FX Swaps) transactions.
type MoneyMarketForeignExchangeSwapsStatisticalReportV01 struct {

	// Provides the elements specific to the report.
	ReportHeader *iso20022.MoneyMarketReportHeader1 `xml:"RptHdr"`

	// Provides the reason why no activity is reported or the required list of transactions for the foreign exchange swaps segment.
	ForeignExchangeSwapsReport *iso20022.ForeignExchangeSwap2Choice `xml:"FXSwpsRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (m *MoneyMarketForeignExchangeSwapsStatisticalReportV01) AddReportHeader() *iso20022.MoneyMarketReportHeader1 {
	m.ReportHeader = new(iso20022.MoneyMarketReportHeader1)
	return m.ReportHeader
}

func (m *MoneyMarketForeignExchangeSwapsStatisticalReportV01) AddForeignExchangeSwapsReport() *iso20022.ForeignExchangeSwap2Choice {
	m.ForeignExchangeSwapsReport = new(iso20022.ForeignExchangeSwap2Choice)
	return m.ForeignExchangeSwapsReport
}

func (m *MoneyMarketForeignExchangeSwapsStatisticalReportV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

