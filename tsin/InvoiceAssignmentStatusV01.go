package tsin

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.007.001.01 Document"`
	Message *InvoiceAssignmentStatusV01 `xml:"InvcAssgnmtSts"`
}

func (d *Document00700101) AddMessage() *InvoiceAssignmentStatusV01 {
	d.Message = new(InvoiceAssignmentStatusV01)
	return d.Message
}

// The message InvoiceAssignmentStatus is sent by a factoring service provider to a factoring client and, optionally, to an interested party as a response to assignments requests.
// The factoring service provider returns a copy of items of corresponding requests together with an information about the status of treatment, for example acceptance, rejection or treatment not yet finished. A rejection can be the result of bad message syntax, but also for other motives such as risk, compliance or covenants.
// For each reported financial item, the factoring service provider includes a reference to the corresponding item of the InvoiceFinancingRequest message and may include the referenced item as well as data from other related and referenced messages.
// The message contains information about other parties to be notified and whether these parties are required to acknowledge the assignment.
// The message can carry digital signatures if required by context.
type InvoiceAssignmentStatusV01 struct {

	// Set of characteristics that unambiguously identify the assignment status, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of assignments of financial items.
	AssignmentList []*iso20022.FinancingItemList1 `xml:"AssgnmtList"`

	// Number of assignments.
	AssignmentCount *iso20022.Max15NumericText `xml:"AssgnmtCnt,omitempty"`

	// Total number of individual items in all assignments.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`

}


func (i *InvoiceAssignmentStatusV01) AddHeader() *iso20022.BusinessLetter1 {
	i.Header = new(iso20022.BusinessLetter1)
	return i.Header
}

func (i *InvoiceAssignmentStatusV01) AddAssignmentList() *iso20022.FinancingItemList1 {
	newValue := new (iso20022.FinancingItemList1)
	i.AssignmentList = append(i.AssignmentList, newValue)
	return newValue
}

func (i *InvoiceAssignmentStatusV01) SetAssignmentCount(value string) {
	i.AssignmentCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentStatusV01) SetItemCount(value string) {
	i.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentStatusV01) SetControlSum(value string) {
	i.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (i *InvoiceAssignmentStatusV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new (iso20022.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

