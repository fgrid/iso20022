package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.002.001.02 Document"`
	Message *ATMWithdrawalResponseV02 `xml:"ATMWdrwlRspn"`
}

func (d *Document00200102) AddMessage() *ATMWithdrawalResponseV02 {
	d.Message = new(ATMWithdrawalResponseV02)
	return d.Message
}

// The ATMWithdrawalResponse message is sent by an acquirer or its agent to an ATM in response to the ATMWithdrawalRequest to inform the ATM of the approval or decline of the withdrawal transaction.
type ATMWithdrawalResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlRspn,omitempty"`

	// Information related to the response of an ATM withdrawal transaction from an ATM manager.
	ATMWithdrawalResponse *iso20022.ATMWithdrawalResponse2 `xml:"ATMWdrwlRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalResponseV02) AddHeader() *iso20022.Header31 {
	a.Header = new(iso20022.Header31)
	return a.Header
}

func (a *ATMWithdrawalResponseV02) AddProtectedATMWithdrawalResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV02) AddATMWithdrawalResponse() *iso20022.ATMWithdrawalResponse2 {
	a.ATMWithdrawalResponse = new(iso20022.ATMWithdrawalResponse2)
	return a.ATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
