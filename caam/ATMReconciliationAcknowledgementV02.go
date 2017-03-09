package caam

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.010.001.02 Document"`
	Message *ATMReconciliationAcknowledgementV02 `xml:"ATMRcncltnAck"`
}

func (d *Document01000102) AddMessage() *ATMReconciliationAcknowledgementV02 {
	d.Message = new(ATMReconciliationAcknowledgementV02)
	return d.Message
}

// The ATMReconciliationAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMReconciliationAdvice message.
type ATMReconciliationAcknowledgementV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMReconciliationAcknowledgement *iso20022.ContentInformationType10 `xml:"PrtctdATMRcncltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM reconciliation from the ATM manager.
	ATMReconciliationAcknowledgement *iso20022.ATMReconciliationAcknowledgement2 `xml:"ATMRcncltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMReconciliationAcknowledgementV02) AddHeader() *iso20022.Header32 {
	a.Header = new(iso20022.Header32)
	return a.Header
}

func (a *ATMReconciliationAcknowledgementV02) AddProtectedATMReconciliationAcknowledgement() *iso20022.ContentInformationType10 {
	a.ProtectedATMReconciliationAcknowledgement = new(iso20022.ContentInformationType10)
	return a.ProtectedATMReconciliationAcknowledgement
}

func (a *ATMReconciliationAcknowledgementV02) AddATMReconciliationAcknowledgement() *iso20022.ATMReconciliationAcknowledgement2 {
	a.ATMReconciliationAcknowledgement = new(iso20022.ATMReconciliationAcknowledgement2)
	return a.ATMReconciliationAcknowledgement
}

func (a *ATMReconciliationAcknowledgementV02) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
