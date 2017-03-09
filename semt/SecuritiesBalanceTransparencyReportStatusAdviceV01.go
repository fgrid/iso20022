package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04200101 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.042.001.01 Document"`
	Message *SecuritiesBalanceTransparencyReportStatusAdviceV01 `xml:"SctiesBalTrnsprncyRptStsAdvc"`
}

func (d *Document04200101) AddMessage() *SecuritiesBalanceTransparencyReportStatusAdviceV01 {
	d.Message = new(SecuritiesBalanceTransparencyReportStatusAdviceV01)
	return d.Message
}

// SCOPE
//
// An account owner, such as a custodian, central securities depository, international securities depository or transfer agent, sends the SecuritiesBalanceTransparencyReportStatusAdvice message in response to a SecuritiesBalanceTransparencyReport, to accept or reject the statement of holdings as sent in a SecuritiesBalanceTransparencyReport.
//
// USAGE
// The SecuritiesBalanceTransparencyReportStatusAdvice is used to accept (Accepted), partially accept (Accepted With Exception) or reject (Rejected) a previously received SecuritiesBalanceTransparencyReport.
type SecuritiesBalanceTransparencyReportStatusAdviceV01 struct {

	// Unique and unambiguous identification of the status advice message.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Identification of the party that is the sender of the status advice message.
	SenderIdentification *iso20022.PartyIdentification100 `xml:"SndrId"`

	// Identification of the party that is the receiver of the status advice message.
	ReceiverIdentification *iso20022.PartyIdentification100 `xml:"RcvrId,omitempty"`

	// Reference of the statement for which the status advice has been issued.
	RelatedStatement *iso20022.StatementReference1 `xml:"RltdStmt"`

	// Status of the referenced statement.
	Status *iso20022.ReportItemStatus1Choice `xml:"Sts"`

	// Number of items for each identical transaction status.
	NumberOfItemsPerStatus []*iso20022.NumberOfItemsPerStatus1 `xml:"NbOfItmsPerSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddSenderIdentification() *iso20022.PartyIdentification100 {
	s.SenderIdentification = new(iso20022.PartyIdentification100)
	return s.SenderIdentification
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddReceiverIdentification() *iso20022.PartyIdentification100 {
	s.ReceiverIdentification = new(iso20022.PartyIdentification100)
	return s.ReceiverIdentification
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddRelatedStatement() *iso20022.StatementReference1 {
	s.RelatedStatement = new(iso20022.StatementReference1)
	return s.RelatedStatement
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddStatus() *iso20022.ReportItemStatus1Choice {
	s.Status = new(iso20022.ReportItemStatus1Choice)
	return s.Status
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddNumberOfItemsPerStatus() *iso20022.NumberOfItemsPerStatus1 {
	newValue := new(iso20022.NumberOfItemsPerStatus1)
	s.NumberOfItemsPerStatus = append(s.NumberOfItemsPerStatus, newValue)
	return newValue
}

func (s *SecuritiesBalanceTransparencyReportStatusAdviceV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
