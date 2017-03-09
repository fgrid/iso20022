package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.006.001.01 Document"`
	Message *UndertakingAmendmentAdviceV01 `xml:"UdrtkgAmdmntAdvc"`
}

func (d *Document00600101) AddMessage() *UndertakingAmendmentAdviceV01 {
	d.Message = new(UndertakingAmendmentAdviceV01)
	return d.Message
}

// The UndertakingAmendmentAdvice message is sent by an advising party to the beneficiary, either directly or via one or more other advising parties in the transaction chain, to advise the content of a proposed amendment to an undertaking. Information about the message may also be sent to other interested parties. The proposed undertaking amendment could be to a demand guarantee, standby letter of credit, or counter-undertaking (counter-guarantee or counter-standby). In addition to providing the terms of the proposed amendment and relevant details on proposed changes to the undertaking, the message may provide information from the sender such as confirmation details. It may also be used to advise the proposed termination or cancellation of the undertaking.
type UndertakingAmendmentAdviceV01 struct {

	// Party advising the undertaking to the beneficiary or to another party.
	AdvisingParty *iso20022.PartyIdentification43 `xml:"AdvsgPty"`

	// Additional party that advises the undertaking.
	SecondAdvisingParty *iso20022.PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Date on which the undertaking is advised.
	DateOfAdvice *iso20022.ISODate `xml:"DtOfAdvc"`

	// Details related to the advice of the proposed amended undertaking.
	UndertakingAmendmentAdviceDetails *iso20022.Amendment2 `xml:"UdrtkgAmdmntAdvcDtls"`

	// Additional information specific to the bank-to-bank communication.
	BankToBankInformation []*iso20022.Max2000Text `xml:"BkToBkInf,omitempty"`

	// Digital signature of the proposed amendment advice.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentAdviceV01) AddAdvisingParty() *iso20022.PartyIdentification43 {
	u.AdvisingParty = new(iso20022.PartyIdentification43)
	return u.AdvisingParty
}

func (u *UndertakingAmendmentAdviceV01) AddSecondAdvisingParty() *iso20022.PartyIdentification43 {
	u.SecondAdvisingParty = new(iso20022.PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *UndertakingAmendmentAdviceV01) SetDateOfAdvice(value string) {
	u.DateOfAdvice = (*iso20022.ISODate)(&value)
}

func (u *UndertakingAmendmentAdviceV01) AddUndertakingAmendmentAdviceDetails() *iso20022.Amendment2 {
	u.UndertakingAmendmentAdviceDetails = new(iso20022.Amendment2)
	return u.UndertakingAmendmentAdviceDetails
}

func (u *UndertakingAmendmentAdviceV01) AddBankToBankInformation(value string) {
	u.BankToBankInformation = append(u.BankToBankInformation, (*iso20022.Max2000Text)(&value))
}

func (u *UndertakingAmendmentAdviceV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
