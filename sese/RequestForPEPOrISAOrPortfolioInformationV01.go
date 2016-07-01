package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.019.001.01 Document"`
	Message *RequestForPEPOrISAOrPortfolioInformationV01 `xml:"ReqForPEPOrISAOrPrtflInfV01"`
}

func (d *Document01900101) AddMessage() *RequestForPEPOrISAOrPortfolioInformationV01 {
	d.Message = new(RequestForPEPOrISAOrPortfolioInformationV01)
	return d.Message
}

// Scope
// An instructing party, eg, a (new) plan manager sends the RequestForPEPorISAOrPortfolioInformation message to the executing party, eg, a (old) plan manager, on behalf of the initiating party, eg, an investor (client), to request information about financial instruments held on behalf of the client.
// Usage
// The RequestForPEPOrISAOrPortfolioInformation message is used to request information about one or more PEP or ISA or portfolio products held in a client's account for which it intends to instruct a transfer at a later time.
type RequestForPEPOrISAOrPortfolioInformationV01 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information identifying the primary individual investor, eg, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, eg, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying other individual investors, eg, name, address, social security number and date of birth.
	OtherIndividualInvestor []*iso20022.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, eg, name and address.
	PrimaryCorporateInvestor *iso20022.Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, eg, name and address.
	SecondaryCorporateInvestor *iso20022.Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, eg, name and address.
	OtherCorporateInvestor []*iso20022.Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	ClientAccount *iso20022.Account5 `xml:"ClntAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *iso20022.Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	NewPlanManager *iso20022.PartyIdentification2Choice `xml:"NewPlanMgr"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*iso20022.PEPISATransfer5 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddMessageReference() *iso20022.MessageIdentification1 {
	r.MessageReference = new(iso20022.MessageIdentification1)
	return r.MessageReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	r.PreviousReference = new(iso20022.AdditionalReference3)
	return r.PreviousReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	r.RelatedReference = new(iso20022.AdditionalReference3)
	return r.RelatedReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPrimaryIndividualInvestor() *iso20022.IndividualPerson8 {
	r.PrimaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return r.PrimaryIndividualInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddSecondaryIndividualInvestor() *iso20022.IndividualPerson8 {
	r.SecondaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return r.SecondaryIndividualInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddOtherIndividualInvestor() *iso20022.IndividualPerson8 {
	newValue := new (iso20022.IndividualPerson8)
	r.OtherIndividualInvestor = append(r.OtherIndividualInvestor, newValue)
	return newValue
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPrimaryCorporateInvestor() *iso20022.Organisation4 {
	r.PrimaryCorporateInvestor = new(iso20022.Organisation4)
	return r.PrimaryCorporateInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddSecondaryCorporateInvestor() *iso20022.Organisation4 {
	r.SecondaryCorporateInvestor = new(iso20022.Organisation4)
	return r.SecondaryCorporateInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddOtherCorporateInvestor() *iso20022.Organisation4 {
	newValue := new (iso20022.Organisation4)
	r.OtherCorporateInvestor = append(r.OtherCorporateInvestor, newValue)
	return newValue
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddClientAccount() *iso20022.Account5 {
	r.ClientAccount = new(iso20022.Account5)
	return r.ClientAccount
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddNomineeAccount() *iso20022.Account6 {
	r.NomineeAccount = new(iso20022.Account6)
	return r.NomineeAccount
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddNewPlanManager() *iso20022.PartyIdentification2Choice {
	r.NewPlanManager = new(iso20022.PartyIdentification2Choice)
	return r.NewPlanManager
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddProductTransfer() *iso20022.PEPISATransfer5 {
	newValue := new (iso20022.PEPISATransfer5)
	r.ProductTransfer = append(r.ProductTransfer, newValue)
	return newValue
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

