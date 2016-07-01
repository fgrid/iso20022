package tsin

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.013.001.01 Document"`
	Message *InvoiceAssignmentAcknowledgementV01 `xml:"InvcAssgnmtAck"`
}

func (d *Document01300101) AddMessage() *InvoiceAssignmentAcknowledgementV01 {
	d.Message = new(InvoiceAssignmentAcknowledgementV01)
	return d.Message
}

// The InvoiceAssignmentAcknowledgement message is sent from a trade partner to communicate the status of payment obligations related to financial items.  The message can be sent independently or as a response to an InvoiceAssignmentNotification message.
// Depending on legal contexts the message may be required as a response to an InvoiceAssignmentNotification message in order for the assignment to become effective.
// The trade party may include references to the corresponding items of an InvoiceAssignmentRequest, InvoiceAssignmentStatus or InvoiceAssignmentNotification or other messages and may include referenced data.
// The message can carry digital signatures if required by context.
type InvoiceAssignmentAcknowledgementV01 struct {

	// Set of characteristics that unambiguously identify the status, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of payment status information.
	PaymentStatusList []*iso20022.FinancingItemList1 `xml:"PmtStsList"`

	// Number of payment information lists as control value.
	PaymentStatusCount *iso20022.Max15NumericText `xml:"PmtStsCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`

}


func (i *InvoiceAssignmentAcknowledgementV01) AddHeader() *iso20022.BusinessLetter1 {
	i.Header = new(iso20022.BusinessLetter1)
	return i.Header
}

func (i *InvoiceAssignmentAcknowledgementV01) AddPaymentStatusList() *iso20022.FinancingItemList1 {
	newValue := new (iso20022.FinancingItemList1)
	i.PaymentStatusList = append(i.PaymentStatusList, newValue)
	return newValue
}

func (i *InvoiceAssignmentAcknowledgementV01) SetPaymentStatusCount(value string) {
	i.PaymentStatusCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentAcknowledgementV01) SetItemCount(value string) {
	i.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentAcknowledgementV01) SetControlSum(value string) {
	i.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (i *InvoiceAssignmentAcknowledgementV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new (iso20022.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

