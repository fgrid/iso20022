package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01900104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.019.001.04 Document"`
	Message *AccountHoldingInformationRequestV04 `xml:"AcctHldgInfReq"`
}

func (d *Document01900104) AddMessage() *AccountHoldingInformationRequestV04 {
	d.Message = new(AccountHoldingInformationRequestV04)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee) sends the AccountHoldingInformationRequest message to the executing party, for example, a (old) plan manager (Transferor), on behalf of the initiating party, for example, an investor (client), to request information about financial instruments held on behalf of the client.
// Usage
// The AccountHoldingInformationRequest message is used to request information about one or more ISA or portfolio products held in a client's account for which it intends to instruct a transfer at a later time.
type AccountHoldingInformationRequestV04 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Identifies the business flow type (assets to be delivered or received).
	BusinessFlowDirectionType *iso20022.BusinessFlowDirectionType1Code `xml:"BizFlowDrctnTp,omitempty"`

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*iso20022.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *iso20022.Organisation21 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *iso20022.Organisation21 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*iso20022.Organisation21 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *iso20022.Account19 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *iso20022.Account19 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *iso20022.PartyIdentification70Choice `xml:"Trfee"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*iso20022.ISATransfer27 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (a *AccountHoldingInformationRequestV04) AddMessageReference() *iso20022.MessageIdentification1 {
	a.MessageReference = new(iso20022.MessageIdentification1)
	return a.MessageReference
}

func (a *AccountHoldingInformationRequestV04) AddPoolReference() *iso20022.AdditionalReference6 {
	a.PoolReference = new(iso20022.AdditionalReference6)
	return a.PoolReference
}

func (a *AccountHoldingInformationRequestV04) AddPreviousReference() *iso20022.AdditionalReference6 {
	a.PreviousReference = new(iso20022.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountHoldingInformationRequestV04) AddRelatedReference() *iso20022.AdditionalReference6 {
	a.RelatedReference = new(iso20022.AdditionalReference6)
	return a.RelatedReference
}

func (a *AccountHoldingInformationRequestV04) SetBusinessFlowDirectionType(value string) {
	a.BusinessFlowDirectionType = (*iso20022.BusinessFlowDirectionType1Code)(&value)
}

func (a *AccountHoldingInformationRequestV04) AddPrimaryIndividualInvestor() *iso20022.IndividualPerson8 {
	a.PrimaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return a.PrimaryIndividualInvestor
}

func (a *AccountHoldingInformationRequestV04) AddSecondaryIndividualInvestor() *iso20022.IndividualPerson8 {
	a.SecondaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return a.SecondaryIndividualInvestor
}

func (a *AccountHoldingInformationRequestV04) AddOtherIndividualInvestor() *iso20022.IndividualPerson8 {
	newValue := new (iso20022.IndividualPerson8)
	a.OtherIndividualInvestor = append(a.OtherIndividualInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV04) AddPrimaryCorporateInvestor() *iso20022.Organisation21 {
	a.PrimaryCorporateInvestor = new(iso20022.Organisation21)
	return a.PrimaryCorporateInvestor
}

func (a *AccountHoldingInformationRequestV04) AddSecondaryCorporateInvestor() *iso20022.Organisation21 {
	a.SecondaryCorporateInvestor = new(iso20022.Organisation21)
	return a.SecondaryCorporateInvestor
}

func (a *AccountHoldingInformationRequestV04) AddOtherCorporateInvestor() *iso20022.Organisation21 {
	newValue := new (iso20022.Organisation21)
	a.OtherCorporateInvestor = append(a.OtherCorporateInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV04) AddTransferorAccount() *iso20022.Account19 {
	a.TransferorAccount = new(iso20022.Account19)
	return a.TransferorAccount
}

func (a *AccountHoldingInformationRequestV04) AddNomineeAccount() *iso20022.Account19 {
	a.NomineeAccount = new(iso20022.Account19)
	return a.NomineeAccount
}

func (a *AccountHoldingInformationRequestV04) AddTransferee() *iso20022.PartyIdentification70Choice {
	a.Transferee = new(iso20022.PartyIdentification70Choice)
	return a.Transferee
}

func (a *AccountHoldingInformationRequestV04) AddProductTransfer() *iso20022.ISATransfer27 {
	newValue := new (iso20022.ISATransfer27)
	a.ProductTransfer = append(a.ProductTransfer, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV04) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountHoldingInformationRequestV04) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

