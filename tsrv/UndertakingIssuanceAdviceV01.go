package tsrv

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.002.001.01 Document"`
	Message *UndertakingIssuanceAdviceV01 `xml:"UdrtkgIssncAdvc"`
}

func (d *Document00200101) AddMessage() *UndertakingIssuanceAdviceV01 {
	d.Message = new(UndertakingIssuanceAdviceV01)
	return d.Message
}

// The UndertakingIssuanceAdvice message is sent by an advising party to the beneficiary, either directly or via one or more other advising parties in the transaction chain, to advise the issuance of an undertaking. Other interested parties may also be informed of the advice. The undertaking advised could be a demand guarantee, standby letter of credit, or counter-undertaking (counter-guarantee or counter-standby). In addition to providing details on the applicable rules, expiry date, the amount, required documents, and terms and conditions of the undertaking, the advice may provide information from the sender such as confirmation details.
type UndertakingIssuanceAdviceV01 struct {

	// Party advising the undertaking to the beneficiary or to another party.
	AdvisingParty *iso20022.PartyIdentification43 `xml:"AdvsgPty"`

	// Additional party that advises the undertaking.
	SecondAdvisingParty *iso20022.PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Date on which the undertaking is advised.
	DateOfAdvice *iso20022.DateAndDateTimeChoice `xml:"DtOfAdvc"`

	// Details related to the advice of the issued undertaking.
	UndertakingIssuanceAdviceDetails *iso20022.UndertakingAdvice1 `xml:"UdrtkgIssncAdvcDtls"`

	// Additional information specific to the bank-to-bank communication.
	BankToBankInformation []*iso20022.Max2000Text `xml:"BkToBkInf,omitempty"`

	// Digital signature of the undertaking advice.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingIssuanceAdviceV01) AddAdvisingParty() *iso20022.PartyIdentification43 {
	u.AdvisingParty = new(iso20022.PartyIdentification43)
	return u.AdvisingParty
}

func (u *UndertakingIssuanceAdviceV01) AddSecondAdvisingParty() *iso20022.PartyIdentification43 {
	u.SecondAdvisingParty = new(iso20022.PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *UndertakingIssuanceAdviceV01) AddDateOfAdvice() *iso20022.DateAndDateTimeChoice {
	u.DateOfAdvice = new(iso20022.DateAndDateTimeChoice)
	return u.DateOfAdvice
}

func (u *UndertakingIssuanceAdviceV01) AddUndertakingIssuanceAdviceDetails() *iso20022.UndertakingAdvice1 {
	u.UndertakingIssuanceAdviceDetails = new(iso20022.UndertakingAdvice1)
	return u.UndertakingIssuanceAdviceDetails
}

func (u *UndertakingIssuanceAdviceV01) AddBankToBankInformation(value string) {
	u.BankToBankInformation = append(u.BankToBankInformation, (*iso20022.Max2000Text)(&value))
}

func (u *UndertakingIssuanceAdviceV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
