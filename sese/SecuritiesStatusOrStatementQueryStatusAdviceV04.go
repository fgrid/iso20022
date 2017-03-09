package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02200104 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.04 Document"`
	Message *SecuritiesStatusOrStatementQueryStatusAdviceV04 `xml:"SctiesStsOrStmtQryStsAdvc"`
}

func (d *Document02200104) AddMessage() *SecuritiesStatusOrStatementQueryStatusAdviceV04 {
	d.Message = new(SecuritiesStatusOrStatementQueryStatusAdviceV04)
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
type SecuritiesStatusOrStatementQueryStatusAdviceV04 struct {

	// Unambiguous identification of the query as per the account owner.
	QueryDetails *iso20022.DocumentIdentification30 `xml:"QryDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	StatusOrStatementRequested *iso20022.StatusOrStatement7Choice `xml:"StsOrStmtReqd,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus55Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV04) AddQueryDetails() *iso20022.DocumentIdentification30 {
	s.QueryDetails = new(iso20022.DocumentIdentification30)
	return s.QueryDetails
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV04) AddAccountOwner() *iso20022.PartyIdentification98 {
	s.AccountOwner = new(iso20022.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV04) AddSafekeepingAccount() *iso20022.SecuritiesAccount19 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV04) AddStatusOrStatementRequested() *iso20022.StatusOrStatement7Choice {
	s.StatusOrStatementRequested = new(iso20022.StatusOrStatement7Choice)
	return s.StatusOrStatementRequested
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV04) AddProcessingStatus() *iso20022.ProcessingStatus55Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus55Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
