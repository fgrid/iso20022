package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02300101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.023.001.01 Document"`
	Message *SecuritiesEndOfProcessReportV01 `xml:"SctiesEndOfPrcRpt"`
}

func (d *Document02300101) AddMessage() *SecuritiesEndOfProcessReportV01 {
	d.Message = new(SecuritiesEndOfProcessReportV01)
	return d.Message
}

// Scope
// Sent by an executing party or an instructing party to the custodian or an affirming party to notify that all the necessary SecuritiesTradeConfirmation message or any other notification of the process have been sent.
// It may also be sent through Central Matching Utility (CMU).
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager.
// The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The custodian or an affirming party is typically the custodian, trustee, financial institution, intermediary system/vendor communicating on behalf of them, or their agent.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: Executing Party, CMU or Instructing Party
// Respondent: Custodian or an affirming party does not need to respond.
type SecuritiesEndOfProcessReportV01 struct {

	// Number used to sequence pages when it is not possible for data to be conveyed in a single message and the data has to be split across several pages (messages).
	Pagination []*iso20022.Pagination `xml:"Pgntn,omitempty"`

	// Notifies the type of report transmitted.
	ReportGeneralDetails *iso20022.Report3 `xml:"RptGnlDtls"`

	// Parties involved in the confirmation of the details of a trade.
	ConfirmationParties []*iso20022.ConfirmationParties2 `xml:"ConfPties,omitempty"`

	// Party that identifies the underlying investor.
	Investor []*iso20022.PartyIdentificationAndAccount79 `xml:"Invstr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesEndOfProcessReportV01) AddPagination() *iso20022.Pagination {
	newValue := new(iso20022.Pagination)
	s.Pagination = append(s.Pagination, newValue)
	return newValue
}

func (s *SecuritiesEndOfProcessReportV01) AddReportGeneralDetails() *iso20022.Report3 {
	s.ReportGeneralDetails = new(iso20022.Report3)
	return s.ReportGeneralDetails
}

func (s *SecuritiesEndOfProcessReportV01) AddConfirmationParties() *iso20022.ConfirmationParties2 {
	newValue := new(iso20022.ConfirmationParties2)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesEndOfProcessReportV01) AddInvestor() *iso20022.PartyIdentificationAndAccount79 {
	newValue := new(iso20022.PartyIdentificationAndAccount79)
	s.Investor = append(s.Investor, newValue)
	return newValue
}

func (s *SecuritiesEndOfProcessReportV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
