package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.01 Document"`
	Message *IntraPositionMovementPostingReportV01 `xml:"IntraPosMvmntPstngRpt"`
}

func (d *Document01600101) AddMessage() *IntraPositionMovementPostingReportV01 {
	d.Message = new(IntraPositionMovementPostingReportV01)
	return d.Message
}

// Scope
// An account servicer sends an IntraPositionMovementPostingReport to an account owner to provide the details of increases and decreases in securities with a given status within a holding, ie, intra-position transfers, which occurred during a specified period, for all or selected securities in a specified safekeeping account which the account servicer holds for the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementPostingReportV01 struct {

	// Information that unambiguously identifies an IntraPositionMovementPostingReport message as known by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement15 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrument []*iso20022.FinancialInstrumentDetails1 `xml:"FinInstrm,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

}


func (i *IntraPositionMovementPostingReportV01) AddIdentification() *iso20022.DocumentIdentification11 {
	i.Identification = new(iso20022.DocumentIdentification11)
	return i.Identification
}

func (i *IntraPositionMovementPostingReportV01) AddPagination() *iso20022.Pagination {
	i.Pagination = new(iso20022.Pagination)
	return i.Pagination
}

func (i *IntraPositionMovementPostingReportV01) AddStatementGeneralDetails() *iso20022.Statement15 {
	i.StatementGeneralDetails = new(iso20022.Statement15)
	return i.StatementGeneralDetails
}

func (i *IntraPositionMovementPostingReportV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	i.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementPostingReportV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementPostingReportV01) AddFinancialInstrument() *iso20022.FinancialInstrumentDetails1 {
	newValue := new (iso20022.FinancialInstrumentDetails1)
	i.FinancialInstrument = append(i.FinancialInstrument, newValue)
	return newValue
}

func (i *IntraPositionMovementPostingReportV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	i.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return i.MessageOriginator
}

func (i *IntraPositionMovementPostingReportV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	i.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return i.MessageRecipient
}

