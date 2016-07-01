package catp

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:catp.002.001.01 Document"`
	Message *ATMWithdrawalResponseV01 `xml:"ATMWdrwlRspn"`
}

func (d *Document00200101) AddMessage() *ATMWithdrawalResponseV01 {
	d.Message = new(ATMWithdrawalResponseV01)
	return d.Message
}

// The ATMWithdrawalResponse message is sent by an acquirer or its agent to an ATM in response to the ATMWithdrawalRequest to inform the ATM of the approval or decline of the withdrawal transaction.
type ATMWithdrawalResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMWdrwlRspn,omitempty"`

	// Information related to the response of an ATM withdrawal transaction from an ATM manager.
	ATMWithdrawalResponse *iso20022.ATMWithdrawalResponse1 `xml:"ATMWdrwlRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`

}


func (a *ATMWithdrawalResponseV01) AddHeader() *iso20022.Header20 {
	a.Header = new(iso20022.Header20)
	return a.Header
}

func (a *ATMWithdrawalResponseV01) AddProtectedATMWithdrawalResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMWithdrawalResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV01) AddATMWithdrawalResponse() *iso20022.ATMWithdrawalResponse1 {
	a.ATMWithdrawalResponse = new(iso20022.ATMWithdrawalResponse1)
	return a.ATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}

