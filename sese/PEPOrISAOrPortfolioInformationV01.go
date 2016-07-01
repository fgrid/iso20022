package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01800101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.01 Document"`
	Message *PEPOrISAOrPortfolioInformationV01 `xml:"PEPOrISAOrPrtflInfV01"`
}

func (d *Document01800101) AddMessage() *PEPOrISAOrPortfolioInformationV01 {
	d.Message = new(PEPOrISAOrPortfolioInformationV01)
	return d.Message
}

// Scope
// An executing party, eg, a (old) plan manager, sends the PEPOrISAOrPortfolioInformation message to the instructing party, eg, a (new) plan manager, to provide information about financial instruments held on behalf of a client.
// Usage
// The PEPOrISAOrPortfolioInformation message is used to provide information about one or more PEP or ISA or portfolio products held in a client's account.
type PEPOrISAOrPortfolioInformationV01 struct {

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
	ProductTransfer []*iso20022.PEPISATransfer6 `xml:"PdctTrf"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (p *PEPOrISAOrPortfolioInformationV01) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PEPOrISAOrPortfolioInformationV01) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PEPOrISAOrPortfolioInformationV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PEPOrISAOrPortfolioInformationV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PEPOrISAOrPortfolioInformationV01) AddPrimaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioInformationV01) AddSecondaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioInformationV01) AddOtherIndividualInvestor() *iso20022.IndividualPerson8 {
	newValue := new (iso20022.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioInformationV01) AddPrimaryCorporateInvestor() *iso20022.Organisation4 {
	p.PrimaryCorporateInvestor = new(iso20022.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioInformationV01) AddSecondaryCorporateInvestor() *iso20022.Organisation4 {
	p.SecondaryCorporateInvestor = new(iso20022.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioInformationV01) AddOtherCorporateInvestor() *iso20022.Organisation4 {
	newValue := new (iso20022.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioInformationV01) AddClientAccount() *iso20022.Account5 {
	p.ClientAccount = new(iso20022.Account5)
	return p.ClientAccount
}

func (p *PEPOrISAOrPortfolioInformationV01) AddNomineeAccount() *iso20022.Account6 {
	p.NomineeAccount = new(iso20022.Account6)
	return p.NomineeAccount
}

func (p *PEPOrISAOrPortfolioInformationV01) AddNewPlanManager() *iso20022.PartyIdentification2Choice {
	p.NewPlanManager = new(iso20022.PartyIdentification2Choice)
	return p.NewPlanManager
}

func (p *PEPOrISAOrPortfolioInformationV01) AddProductTransfer() *iso20022.PEPISATransfer6 {
	newValue := new (iso20022.PEPISATransfer6)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioInformationV01) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

