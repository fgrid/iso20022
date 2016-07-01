package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.02 Document"`
	Message *PEPOrISAOrPortfolioTransferConfirmationV02 `xml:"PEPOrISAOrPrtflTrfConfV02"`
}

func (d *Document01300102) AddMessage() *PEPOrISAOrPortfolioTransferConfirmationV02 {
	d.Message = new(PEPOrISAOrPortfolioTransferConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a (old) plan manager, sends the PEPOrISAOrPortfolioTransferConfirmation message to the instructing party, eg, a (new) plan manager, to confirm the transfer of one or more PEP or ISA or portfolio products from the client's account at the old plan manager to the client's account at the new plan manager through a nominee account.
// Usage
// The PEPOrISAOrPortfolioTransferConfirmation message is used to confirm the transfer of one or more PEP or ISA or portfolio products.
// The reference of each product transfer confirmation is identified in TransferConfirmationIdentification. The reference of the original product transfer is specified in TransferInstructionReference. The message identification of the PEPOrISAPOrPortfolioTransferInstruction message in which the product transfers were conveyed may also be quoted in RelatedReference.
type PEPOrISAOrPortfolioTransferConfirmationV02 struct {

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

	// Information identifying the other individual investors, eg, name, address, social security number and date of birth.
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

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *iso20022.CashAccount11 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*iso20022.PEPISATransfer4 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPrimaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddSecondaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddOtherIndividualInvestor() *iso20022.IndividualPerson8 {
	newValue := new (iso20022.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPrimaryCorporateInvestor() *iso20022.Organisation4 {
	p.PrimaryCorporateInvestor = new(iso20022.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddSecondaryCorporateInvestor() *iso20022.Organisation4 {
	p.SecondaryCorporateInvestor = new(iso20022.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddOtherCorporateInvestor() *iso20022.Organisation4 {
	newValue := new (iso20022.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddClientAccount() *iso20022.Account5 {
	p.ClientAccount = new(iso20022.Account5)
	return p.ClientAccount
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddNomineeAccount() *iso20022.Account6 {
	p.NomineeAccount = new(iso20022.Account6)
	return p.NomineeAccount
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddNewPlanManager() *iso20022.PartyIdentification2Choice {
	p.NewPlanManager = new(iso20022.PartyIdentification2Choice)
	return p.NewPlanManager
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddCashAccount() *iso20022.CashAccount11 {
	p.CashAccount = new(iso20022.CashAccount11)
	return p.CashAccount
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddProductTransfer() *iso20022.PEPISATransfer4 {
	newValue := new (iso20022.PEPISATransfer4)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

