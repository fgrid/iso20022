package tsin

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.009.001.01 Document"`
	Message *PartyRegistrationAndGuaranteeRequestV01 `xml:"PtyRegnAndGrntReq"`
}

func (d *Document00900101) AddMessage() *PartyRegistrationAndGuaranteeRequestV01 {
	d.Message = new(PartyRegistrationAndGuaranteeRequestV01)
	return d.Message
}

// The message PartyRegistrationAndGuaranteeRequest is sent by a factoring client either to a financial service or a guarantee issuer. The message can also be sent from a financial service to a guarantee issuer. Furthermore, the message can be sent to an interested party for example a fiscal authority. When the message is sent to a guarantee issuer, the factoring client or financial service provider requests a guarantee for the factoring agreement concerning the indicated trade party. When the message is sent to a financial service, the financial client requests an agreement to execute assignments of financial items. The financial client may request the guarantee amount to be obtained in case of insolvency of the trade partner for a corresponding account receivable directly from the financial service. Alternatively and depending on the contractual and product definition, the financial client may be required to include a copy of a guarantee status received from a guarantee issuer.
// The message can carry digital signatures if required by context.
type PartyRegistrationAndGuaranteeRequestV01 struct {

	// Set of characteristics that unambiguously identify the request, common parameters, documents and identifications.
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


func (p *PartyRegistrationAndGuaranteeRequestV01) AddHeader() *iso20022.BusinessLetter1 {
	p.Header = new(iso20022.BusinessLetter1)
	return p.Header
}

func (p *PartyRegistrationAndGuaranteeRequestV01) AddAgreementList() *iso20022.FinancingAgreementList1 {
	newValue := new (iso20022.FinancingAgreementList1)
	p.AgreementList = append(p.AgreementList, newValue)
	return newValue
}

func (p *PartyRegistrationAndGuaranteeRequestV01) SetAgreementCount(value string) {
	p.AgreementCount = (*iso20022.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeRequestV01) SetItemCount(value string) {
	p.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeRequestV01) SetControlSum(value string) {
	p.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (p *PartyRegistrationAndGuaranteeRequestV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new (iso20022.EncapsulatedBusinessMessage1)
	p.AttachedMessage = append(p.AttachedMessage, newValue)
	return newValue
}

