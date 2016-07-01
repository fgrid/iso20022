package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.054.001.01 Document"`
	Message *InvoicePaymentReconciliationStatusV01 `xml:"InvcPmtRcncltnSts"`
}

func (d *Document05400101) AddMessage() *InvoicePaymentReconciliationStatusV01 {
	d.Message = new(InvoicePaymentReconciliationStatusV01)
	return d.Message
}

// The message InvoicePaymentReconciliationStatus is sent from a payee to a payer to acknowledge attribution of payments.
// A payee that has received payment reconciliation information uses this message to confirm or to question common understanding of payments and instalments.
// The payee may include references to the corresponding items of an InvoicePaymentReconciliationAdvice message or to other messages and may include the referenced data.
// The message can carry digital signatures if required by context.
type InvoicePaymentReconciliationStatusV01 struct {

	// Specifies a set of characteristics that unambiguously identify the status, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of payment reconciliation information.
	ReconciliationList []*iso20022.ReconciliationList1 `xml:"RcncltnList"`

	// Specifies the number of reconciliation lists.
	ReconciliationCount *iso20022.Max15NumericText `xml:"RcncltnCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`

}


func (i *InvoicePaymentReconciliationStatusV01) AddHeader() *iso20022.BusinessLetter1 {
	i.Header = new(iso20022.BusinessLetter1)
	return i.Header
}

func (i *InvoicePaymentReconciliationStatusV01) AddReconciliationList() *iso20022.ReconciliationList1 {
	newValue := new (iso20022.ReconciliationList1)
	i.ReconciliationList = append(i.ReconciliationList, newValue)
	return newValue
}

func (i *InvoicePaymentReconciliationStatusV01) SetReconciliationCount(value string) {
	i.ReconciliationCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoicePaymentReconciliationStatusV01) SetItemCount(value string) {
	i.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoicePaymentReconciliationStatusV01) SetControlSum(value string) {
	i.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (i *InvoicePaymentReconciliationStatusV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new (iso20022.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

