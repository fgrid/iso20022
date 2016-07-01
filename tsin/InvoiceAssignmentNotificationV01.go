package tsin

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00800101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.008.001.01 Document"`
	Message *InvoiceAssignmentNotificationV01 `xml:"InvcAssgnmtNtfctn"`
}

func (d *Document00800101) AddMessage() *InvoiceAssignmentNotificationV01 {
	d.Message = new(InvoiceAssignmentNotificationV01)
	return d.Message
}

// This message is sent from a factoring service provider or a factoring client to a trade partner to inform about assignments of financing items and, optionally, to an interested party.
// The information given to the trade party indicates that property of the payment obligation has been or is being transferred to the financial institution and that payments have to be done between the trade partner and the factoring service provider.
// The message indicates whether the notified party is required to acknowledge the notified assignment and to which party an acknowledgement has to be sent.
// This message can also be used outside a factoring context directly between a payer and a payee for example as a reminder about a payment obligation or to make an adjustment.
// If applicable, the message may reference corresponding items of an InvoiceFinancingRequest or InvoiceFinancingStatus or other related messages and may contain referenced data.
// The message can carry digital signatures if required by context.
type InvoiceAssignmentNotificationV01 struct {

	// Set of characteristics that unambiguously identify the assignment notification, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of assignment notifications.
	NotificationList []*iso20022.FinancingItemList1 `xml:"NtfctnList"`

	// Number of assignment notification lists.
	NotificationCount *iso20022.Max15NumericText `xml:"NtfctnCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`

}


func (i *InvoiceAssignmentNotificationV01) AddHeader() *iso20022.BusinessLetter1 {
	i.Header = new(iso20022.BusinessLetter1)
	return i.Header
}

func (i *InvoiceAssignmentNotificationV01) AddNotificationList() *iso20022.FinancingItemList1 {
	newValue := new (iso20022.FinancingItemList1)
	i.NotificationList = append(i.NotificationList, newValue)
	return newValue
}

func (i *InvoiceAssignmentNotificationV01) SetNotificationCount(value string) {
	i.NotificationCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentNotificationV01) SetItemCount(value string) {
	i.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentNotificationV01) SetControlSum(value string) {
	i.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (i *InvoiceAssignmentNotificationV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new (iso20022.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

