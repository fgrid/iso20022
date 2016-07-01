package admi

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 Document"`
	Message *SystemEventAcknowledgementV01 `xml:"SysEvtAck"`
}

func (d *Document01100101) AddMessage() *SystemEventAcknowledgementV01 {
	d.Message = new(SystemEventAcknowledgementV01)
	return d.Message
}

// The SystemEventAcknowledgement message is sent by a participant of a central system to the central system to acknowledge the notification of an occurrence of an event in a central system.
// 
type SystemEventAcknowledgementV01 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *iso20022.Max35Text `xml:"MsgId"`

	// Represents the original reference of the system event notification for which the acknowledgement is given, as assigned by the central system.
	OriginatorReference *iso20022.Max35Text `xml:"OrgtrRef,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *iso20022.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Details of the system event being acknowledged.
	AcknowledgementDetails *iso20022.Event1 `xml:"AckDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SystemEventAcknowledgementV01) SetMessageIdentification(value string) {
	s.MessageIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SystemEventAcknowledgementV01) SetOriginatorReference(value string) {
	s.OriginatorReference = (*iso20022.Max35Text)(&value)
}

func (s *SystemEventAcknowledgementV01) SetSettlementSessionIdentifier(value string) {
	s.SettlementSessionIdentifier = (*iso20022.Exact4AlphaNumericText)(&value)
}

func (s *SystemEventAcknowledgementV01) AddAcknowledgementDetails() *iso20022.Event1 {
	s.AcknowledgementDetails = new(iso20022.Event1)
	return s.AcknowledgementDetails
}

func (s *SystemEventAcknowledgementV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

