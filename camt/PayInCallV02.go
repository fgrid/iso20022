package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document06100102 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.061.001.02 Document"`
	Message *PayInCallV02 `xml:"PayInCall"`
}

func (d *Document06100102) AddMessage() *PayInCallV02 {
	d.Message = new(PayInCallV02)
	return d.Message
}

// The PayInCall message is sent by a central settlement system to request additional funding from a settlement member impacted by a failure situation.
type PayInCallV02 struct {

	// Party for which the PayInCall is generated.
	PartyIdentification *iso20022.PartyIdentification73Choice `xml:"PtyId"`

	// Contains  the report generation information and the report items.
	ReportData *iso20022.ReportData5 `xml:"RptData"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *iso20022.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PayInCallV02) AddPartyIdentification() *iso20022.PartyIdentification73Choice {
	p.PartyIdentification = new(iso20022.PartyIdentification73Choice)
	return p.PartyIdentification
}

func (p *PayInCallV02) AddReportData() *iso20022.ReportData5 {
	p.ReportData = new(iso20022.ReportData5)
	return p.ReportData
}

func (p *PayInCallV02) SetSettlementSessionIdentifier(value string) {
	p.SettlementSessionIdentifier = (*iso20022.Exact4AlphaNumericText)(&value)
}

func (p *PayInCallV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
