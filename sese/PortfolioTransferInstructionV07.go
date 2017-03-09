package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01200107 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.012.001.07 Document"`
	Message *PortfolioTransferInstructionV07 `xml:"PrtflTrfInstr"`
}

func (d *Document01200107) AddMessage() *PortfolioTransferInstructionV07 {
	d.Message = new(PortfolioTransferInstructionV07)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee), sends the PortfolioTransferInstruction message to the executing party, for example, a (old) plan manager (Transferor), on behalf of the initiating party, for example, an investor (client), to instruct the transfer of financial instruments from the clients account at the old plan manager (Transferor) to the clients account at the new plan manager (Transferee) through a nominee account.
// Usage
// The PortfolioTransferInstruction message is used to instruct the withdrawal of one or more ISA or portfolio products from one account and deliver them to another account.
// The PortfolioTransferInstruction message is used to instruct one or more transfers for one client. Each transfer is for delivery to the same account. The account may be owned by one or more individual investors or one or more corporate investors. Each transfer is identified in TransferIdentification.
// If the instructing party does not have enough information to instruct the transfer, then it must first send a AccountHoldingInformationRequest message to the executing party in order to receive a AccountHoldingInformation message.
type PortfolioTransferInstructionV07 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *iso20022.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
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

	// Identification of a related party or intermediary.
	IntermediaryInformation []*iso20022.Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *iso20022.CashAccount34 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*iso20022.ISATransfer22 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PortfolioTransferInstructionV07) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferInstructionV07) AddPoolReference() *iso20022.AdditionalReference6 {
	p.PoolReference = new(iso20022.AdditionalReference6)
	return p.PoolReference
}

func (p *PortfolioTransferInstructionV07) AddPreviousReference() *iso20022.AdditionalReference6 {
	p.PreviousReference = new(iso20022.AdditionalReference6)
	return p.PreviousReference
}

func (p *PortfolioTransferInstructionV07) AddRelatedReference() *iso20022.AdditionalReference6 {
	p.RelatedReference = new(iso20022.AdditionalReference6)
	return p.RelatedReference
}

func (p *PortfolioTransferInstructionV07) AddPrimaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PortfolioTransferInstructionV07) AddSecondaryIndividualInvestor() *iso20022.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(iso20022.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PortfolioTransferInstructionV07) AddOtherIndividualInvestor() *iso20022.IndividualPerson8 {
	newValue := new(iso20022.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddPrimaryCorporateInvestor() *iso20022.Organisation21 {
	p.PrimaryCorporateInvestor = new(iso20022.Organisation21)
	return p.PrimaryCorporateInvestor
}

func (p *PortfolioTransferInstructionV07) AddSecondaryCorporateInvestor() *iso20022.Organisation21 {
	p.SecondaryCorporateInvestor = new(iso20022.Organisation21)
	return p.SecondaryCorporateInvestor
}

func (p *PortfolioTransferInstructionV07) AddOtherCorporateInvestor() *iso20022.Organisation21 {
	newValue := new(iso20022.Organisation21)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddTransferorAccount() *iso20022.Account19 {
	p.TransferorAccount = new(iso20022.Account19)
	return p.TransferorAccount
}

func (p *PortfolioTransferInstructionV07) AddNomineeAccount() *iso20022.Account19 {
	p.NomineeAccount = new(iso20022.Account19)
	return p.NomineeAccount
}

func (p *PortfolioTransferInstructionV07) AddTransferee() *iso20022.PartyIdentification70Choice {
	p.Transferee = new(iso20022.PartyIdentification70Choice)
	return p.Transferee
}

func (p *PortfolioTransferInstructionV07) AddIntermediaryInformation() *iso20022.Intermediary34 {
	newValue := new(iso20022.Intermediary34)
	p.IntermediaryInformation = append(p.IntermediaryInformation, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddCashAccount() *iso20022.CashAccount34 {
	p.CashAccount = new(iso20022.CashAccount34)
	return p.CashAccount
}

func (p *PortfolioTransferInstructionV07) AddProductTransfer() *iso20022.ISATransfer22 {
	newValue := new(iso20022.ISATransfer22)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}

func (p *PortfolioTransferInstructionV07) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
