package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document06200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Document"`
	Message *PayInScheduleV03 `xml:"PayInSchdl"`
}

func (d *Document06200103) AddMessage() *PayInScheduleV03 {
	d.Message = new(PayInScheduleV03)
	return d.Message
}

// The PayInSchedule message is sent by a central settlement system to the participant to provide notification of a series of timed payments scheduled for each currency at the time and date of the schedule generation. The central settlement system may send information about how the timed payments have been calculated.
type PayInScheduleV03 struct {

	// Party for which the pay-in schedule is generated.
	PartyIdentification *iso20022.PartyIdentification73Choice `xml:"PtyId"`

	// General information applicable to the report.
	ReportData *iso20022.ReportData4 `xml:"RptData"`

	// Projected net position for all currencies, projected long for the value date.
	PayInScheduleLongBalance []*iso20022.BalanceStatus2 `xml:"PayInSchdlLngBal,omitempty"`

	// Currency and total amount to be paid in by the corresponding deadline.
	PayInScheduleItem []*iso20022.PayInScheduleItems1 `xml:"PayInSchdlItm,omitempty"`

	// Factors used in the calculation of the pay-in schedule.
	PayInFactors *iso20022.PayInFactors1 `xml:"PayInFctrs,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (p *PayInScheduleV03) AddPartyIdentification() *iso20022.PartyIdentification73Choice {
	p.PartyIdentification = new(iso20022.PartyIdentification73Choice)
	return p.PartyIdentification
}

func (p *PayInScheduleV03) AddReportData() *iso20022.ReportData4 {
	p.ReportData = new(iso20022.ReportData4)
	return p.ReportData
}

func (p *PayInScheduleV03) AddPayInScheduleLongBalance() *iso20022.BalanceStatus2 {
	newValue := new (iso20022.BalanceStatus2)
	p.PayInScheduleLongBalance = append(p.PayInScheduleLongBalance, newValue)
	return newValue
}

func (p *PayInScheduleV03) AddPayInScheduleItem() *iso20022.PayInScheduleItems1 {
	newValue := new (iso20022.PayInScheduleItems1)
	p.PayInScheduleItem = append(p.PayInScheduleItem, newValue)
	return newValue
}

func (p *PayInScheduleV03) AddPayInFactors() *iso20022.PayInFactors1 {
	p.PayInFactors = new(iso20022.PayInFactors1)
	return p.PayInFactors
}

func (p *PayInScheduleV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

