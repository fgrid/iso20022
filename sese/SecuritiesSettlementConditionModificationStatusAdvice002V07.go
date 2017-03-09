package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03100207 struct {
	XMLName xml.Name                                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.002.07 Document"`
	Message *SecuritiesSettlementConditionModificationStatusAdvice002V07 `xml:"SctiesSttlmCondModStsAdvc"`
}

func (d *Document03100207) AddMessage() *SecuritiesSettlementConditionModificationStatusAdvice002V07 {
	d.Message = new(SecuritiesSettlementConditionModificationStatusAdvice002V07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementConditionsModificationStatusAdvice to an account owner to advise the status of a modification request previously instructed by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// A SecuritiesSettlementConditionsModificationRequest may contain requests on multiple transactions. However, one SecuritiesSettlementConditionsModificationStatusAdvice must be sent per transaction modified unless the SecuritiesSettlementConditionsModificationRequest is rejected as a whole.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementConditionModificationStatusAdvice002V07 struct {

	// Identification of the SecuritiesSettlementConditionsModificationRequest.
	RequestReference *iso20022.Identification16 `xml:"ReqRef"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	RequestDetails *iso20022.RequestDetails16 `xml:"ReqDtls,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus58Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionModificationStatusAdvice002V07) AddRequestReference() *iso20022.Identification16 {
	s.RequestReference = new(iso20022.Identification16)
	return s.RequestReference
}

func (s *SecuritiesSettlementConditionModificationStatusAdvice002V07) AddAccountOwner() *iso20022.PartyIdentification109 {
	s.AccountOwner = new(iso20022.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionModificationStatusAdvice002V07) AddSafekeepingAccount() *iso20022.SecuritiesAccount30 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionModificationStatusAdvice002V07) AddRequestDetails() *iso20022.RequestDetails16 {
	s.RequestDetails = new(iso20022.RequestDetails16)
	return s.RequestDetails
}

func (s *SecuritiesSettlementConditionModificationStatusAdvice002V07) AddProcessingStatus() *iso20022.ProcessingStatus58Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus58Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementConditionModificationStatusAdvice002V07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
