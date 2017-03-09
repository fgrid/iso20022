package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02200204 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.002.04 Document"`
	Message *SecuritiesStatusOrStatementQueryStatusAdvice002V04 `xml:"SctiesStsOrStmtQryStsAdvc"`
}

func (d *Document02200204) AddMessage() *SecuritiesStatusOrStatementQueryStatusAdvice002V04 {
	d.Message = new(SecuritiesStatusOrStatementQueryStatusAdvice002V04)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesStatusOrStatementQueryStatusAdvice to an account owner to advise the status of a status query or statement query previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesStatusOrStatementQueryStatusAdvice002V04 struct {

	// Unambiguous identification of the query as per the account owner.
	QueryDetails *iso20022.DocumentIdentification48 `xml:"QryDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	StatusOrStatementRequested *iso20022.StatusOrStatement8Choice `xml:"StsOrStmtReqd,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus64Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddQueryDetails() *iso20022.DocumentIdentification48 {
	s.QueryDetails = new(iso20022.DocumentIdentification48)
	return s.QueryDetails
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddSafekeepingAccount() *iso20022.SecuritiesAccount30 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddStatusOrStatementRequested() *iso20022.StatusOrStatement8Choice {
	s.StatusOrStatementRequested = new(iso20022.StatusOrStatement8Choice)
	return s.StatusOrStatementRequested
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddProcessingStatus() *iso20022.ProcessingStatus64Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus64Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesStatusOrStatementQueryStatusAdvice002V04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
