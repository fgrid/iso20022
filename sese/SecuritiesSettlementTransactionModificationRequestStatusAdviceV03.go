package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03900103 struct {
	XMLName xml.Name                                                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 `xml:"SctiesSttlmTxModReqStsAdvc"`
}

func (d *Document03900103) AddMessage() *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestStatusAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionModificationRequestStatusAdvice to an account owner to advise the status of a SecuritiesSettlementModificationRequest message previously sent by the account owner.
// The account servicer may be:
// - a central securities depository or another settlement market infrastructure managing securities settlement transactions on behalf of their participants
// - an custodian acting as an accounting and/or settlement agent.
//
// Usage
// The message may also be used to:
// - re-send a message sent by the account owner to the account servicer,
// - provide a third party with a copy of a message being sent by the account owner for information,
// - re-send to a third party a copy of a message being sent by the account owner for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	ModificationRequestReference *iso20022.Identification1 `xml:"ModReqRef"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications25 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ModificationProcessingStatus *iso20022.ModificationProcessingStatus4Choice `xml:"ModPrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails45 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddModificationRequestReference() *iso20022.Identification1 {
	s.ModificationRequestReference = new(iso20022.Identification1)
	return s.ModificationRequestReference
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddTransactionIdentification() *iso20022.TransactionIdentifications25 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications25)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddModificationProcessingStatus() *iso20022.ModificationProcessingStatus4Choice {
	s.ModificationProcessingStatus = new(iso20022.ModificationProcessingStatus4Choice)
	return s.ModificationProcessingStatus
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddTransactionDetails() *iso20022.TransactionDetails45 {
	s.TransactionDetails = new(iso20022.TransactionDetails45)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
