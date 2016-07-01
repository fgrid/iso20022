package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.022.001.03 Document"`
	Message *SecuritiesStatusOrStatementQueryStatusAdviceV03 `xml:"SctiesStsOrStmtQryStsAdvc"`
}

func (d *Document02200103) AddMessage() *SecuritiesStatusOrStatementQueryStatusAdviceV03 {
	d.Message = new(SecuritiesStatusOrStatementQueryStatusAdviceV03)
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
// 
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesStatusOrStatementQueryStatusAdviceV03 struct {

	// Unambiguous identification of the query as per the account owner.
	QueryDetails *iso20022.DocumentIdentification24 `xml:"QryDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	StatusOrStatementRequested *iso20022.StatusOrStatement5Choice `xml:"StsOrStmtReqd,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus22Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesStatusOrStatementQueryStatusAdviceV03) AddQueryDetails() *iso20022.DocumentIdentification24 {
	s.QueryDetails = new(iso20022.DocumentIdentification24)
	return s.QueryDetails
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV03) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV03) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV03) AddStatusOrStatementRequested() *iso20022.StatusOrStatement5Choice {
	s.StatusOrStatementRequested = new(iso20022.StatusOrStatement5Choice)
	return s.StatusOrStatementRequested
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV03) AddProcessingStatus() *iso20022.ProcessingStatus22Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus22Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesStatusOrStatementQueryStatusAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

