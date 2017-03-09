package tsin

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.010.001.01 Document"`
	Message *PartyRegistrationAndGuaranteeStatusV01 `xml:"PtyRegnAndGrntSts"`
}

func (d *Document01000101) AddMessage() *PartyRegistrationAndGuaranteeStatusV01 {
	d.Message = new(PartyRegistrationAndGuaranteeStatusV01)
	return d.Message
}

// The message PartyRegistrationAndGuaranteeStatus is either sent by a factoring service to a financing client to indicate the status of a factoring service agreement or from a guarantee issuer to a factoring client or a factoring service to indicate the guarantee covering a requested factoring service agreement. The message can also be sent to an interested party.
// The factoring service or guarantee issuer may include references to a corresponding PartyRegistrationAndGuaranteeRequest message or other related messages and may include referenced data.
// The message contains information about other parties to be notified about the financial service agreement or the guarantee and whether these parties are required to acknowledge the agreement.
// The message contains information returned by the financial institution indicating acceptance or rejection of the trade partner; a positive response is necessary before being able to assign financial items concerning the trade party.
// This message contains identifications of cash accounts to enable payer and payee to treat the transferred payment obligations.
// The message can carry digital signatures if required by context.
type PartyRegistrationAndGuaranteeStatusV01 struct {

	// Set of characteristics that unambiguously identify the status, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of agreements.
	AgreementList []*iso20022.FinancingAgreementList1 `xml:"AgrmtList"`

	// Number of agreement lists as control value.
	AgreementCount *iso20022.Max15NumericText `xml:"AgrmtCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (p *PartyRegistrationAndGuaranteeStatusV01) AddHeader() *iso20022.BusinessLetter1 {
	p.Header = new(iso20022.BusinessLetter1)
	return p.Header
}

func (p *PartyRegistrationAndGuaranteeStatusV01) AddAgreementList() *iso20022.FinancingAgreementList1 {
	newValue := new(iso20022.FinancingAgreementList1)
	p.AgreementList = append(p.AgreementList, newValue)
	return newValue
}

func (p *PartyRegistrationAndGuaranteeStatusV01) SetAgreementCount(value string) {
	p.AgreementCount = (*iso20022.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeStatusV01) SetItemCount(value string) {
	p.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeStatusV01) SetControlSum(value string) {
	p.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (p *PartyRegistrationAndGuaranteeStatusV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new(iso20022.EncapsulatedBusinessMessage1)
	p.AttachedMessage = append(p.AttachedMessage, newValue)
	return newValue
}
