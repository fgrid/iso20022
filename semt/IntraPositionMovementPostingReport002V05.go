package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600205 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.05 Document"`
	Message *IntraPositionMovementPostingReport002V05 `xml:"IntraPosMvmntPstngRpt"`
}

func (d *Document01600205) AddMessage() *IntraPositionMovementPostingReport002V05 {
	d.Message = new(IntraPositionMovementPostingReport002V05)
	return d.Message
}

// Scope
// An account servicer sends an IntraPositionMovementPostingReport to an account owner to provide the details of increases and decreases in securities with a given status within a holding, that is, intra-position transfers, which occurred during a specified period, for all or selected securities in a specified safekeeping account which the account servicer holds for the account owner.
//
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// :
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type IntraPositionMovementPostingReport002V05 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement49 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrument []*iso20022.FinancialInstrumentDetails22 `xml:"FinInstrm,omitempty"`
}

func (i *IntraPositionMovementPostingReport002V05) AddPagination() *iso20022.Pagination {
	i.Pagination = new(iso20022.Pagination)
	return i.Pagination
}

func (i *IntraPositionMovementPostingReport002V05) AddStatementGeneralDetails() *iso20022.Statement49 {
	i.StatementGeneralDetails = new(iso20022.Statement49)
	return i.StatementGeneralDetails
}

func (i *IntraPositionMovementPostingReport002V05) AddAccountOwner() *iso20022.PartyIdentification103Choice {
	i.AccountOwner = new(iso20022.PartyIdentification103Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementPostingReport002V05) AddSafekeepingAccount() *iso20022.SecuritiesAccount27 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount27)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementPostingReport002V05) AddFinancialInstrument() *iso20022.FinancialInstrumentDetails22 {
	newValue := new(iso20022.FinancialInstrumentDetails22)
	i.FinancialInstrument = append(i.FinancialInstrument, newValue)
	return newValue
}
