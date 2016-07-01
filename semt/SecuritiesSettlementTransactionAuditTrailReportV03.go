package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.03 Document"`
	Message *SecuritiesSettlementTransactionAuditTrailReportV03 `xml:"SctiesSttlmTxAudtTrlRpt"`
}

func (d *Document02200103) AddMessage() *SecuritiesSettlementTransactionAuditTrailReportV03 {
	d.Message = new(SecuritiesSettlementTransactionAuditTrailReportV03)
	return d.Message
}

// Scope
// This message is sent by the Market Infrastructure to the CSD to advise of the history of all the statuses, modifications, replacement and cancellation of a specific transaction during its whole life cycle when the instructing party is a direct participant to the Settlement Infrastructure.
// 
// Usage
// The message may also be used to: 
// - re-send a message sent by the market infrastructure to the direct participant,
// - provide a third party with a copy of a message being sent by the market infrastructure for information,
// - re-send to a third party a copy of a message being sent by the market infrastructure for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionAuditTrailReportV03 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Identification of the SecuritiesStatusQuery message sent to request this report.
	QueryReference *iso20022.Identification14 `xml:"QryRef,omitempty"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications29 `xml:"TxId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	//  Provides the history of status and reasons for a pending, posted or cancelled transaction.
	StatusTrail []*iso20022.StatusTrail6 `xml:"StsTrl,omitempty"`

}


func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddQueryReference() *iso20022.Identification14 {
	s.QueryReference = new(iso20022.Identification14)
	return s.QueryReference
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddTransactionIdentification() *iso20022.TransactionIdentifications29 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications29)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddSafekeepingAccount() *iso20022.SecuritiesAccount24 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddAccountOwner() *iso20022.PartyIdentification98 {
	s.AccountOwner = new(iso20022.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV03) AddStatusTrail() *iso20022.StatusTrail6 {
	newValue := new (iso20022.StatusTrail6)
	s.StatusTrail = append(s.StatusTrail, newValue)
	return newValue
}

