package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02800105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.05 Document"`
	Message *AdditionalPaymentInformationV05 `xml:"AddtlPmtInf"`
}

func (d *Document02800105) AddMessage() *AdditionalPaymentInformationV05 {
	d.Message = new(AdditionalPaymentInformationV05)
	return d.Message
}

// Scope
// The Additional Payment Information message is sent by an account servicing institution to an account owner.
// This message is used to provide additional or corrected information on a payment instruction or statement entry, in order to allow reconciliation.
// Usage
// The Additional Payment Information message provides elements which are usually not reported in a statement or advice (for example full remittance information or identification of parties other than the account servicing institution and the account owner). It complements information about a payment instruction that has already been received, in the form of one or several entries of the original payment instruction.
// The Additional Payment Information message covers one and only one original payment instruction. If several payment instructions need further details, multiple Additional Payment Information messages must be used, one for each of the payment instructions.
// The AdditionalPaymentInformation message may be used as a result of two investigation processes and in a way outlined below.
// - A Claim Non Receipt workflow raised by the creditor or recipient of the payment: This means that the payment instruction has reached the creditor or beneficiary. The account owner needs further details or correct information for its reconciliation processes. The Additional Payment Information can be used to provide the missing information.
// - A Request To Modify Payment workflow raised by the debtor or one of the intermediate agents upstream: When the payment instruction has reached its intended recipient and the modification does not affect the accounting at the account servicing institution, this Additional Payment Information message allows the account owner to receive further particulars or correct information about a payment instruction or an entry passed to its account.
// The Additional Payment Information message cannot be used to trigger a request for modification of a payment instruction activity. A Request To Modify Payment message must be used. In other words, if a debtor or one of intermediate agent (excluding the account servicing institution of the creditor) realises the some information was missing in the original payment instruction, he should not use an Additional Payment Information but instead a Request To Modify Payment message.
// It is assumed that when an account servicing institution sends out an Additional Payment Information message, the institution is fairly confident that this will resolve the case. Therefore it does not need to wait for a Resolution Of Investigation message. Neither does the account owner, or whoever receives the additional information, need to send back a Resolution Of Investigation message. Positive resolution in this case is implicit. Both parties are expected to close the case. In the event that the problem does not go away, a party can re-open the case.
type AdditionalPaymentInformationV05 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case3 `xml:"Case"`

	// Identifies the underlying payment instruction.
	Underlying *iso20022.UnderlyingTransaction2Choice `xml:"Undrlyg"`

	// Additional information to the underlying payment instruction.
	Information *iso20022.PaymentComplementaryInformation4 `xml:"Inf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (a *AdditionalPaymentInformationV05) AddAssignment() *iso20022.CaseAssignment3 {
	a.Assignment = new(iso20022.CaseAssignment3)
	return a.Assignment
}

func (a *AdditionalPaymentInformationV05) AddCase() *iso20022.Case3 {
	a.Case = new(iso20022.Case3)
	return a.Case
}

func (a *AdditionalPaymentInformationV05) AddUnderlying() *iso20022.UnderlyingTransaction2Choice {
	a.Underlying = new(iso20022.UnderlyingTransaction2Choice)
	return a.Underlying
}

func (a *AdditionalPaymentInformationV05) AddInformation() *iso20022.PaymentComplementaryInformation4 {
	a.Information = new(iso20022.PaymentComplementaryInformation4)
	return a.Information
}

func (a *AdditionalPaymentInformationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

