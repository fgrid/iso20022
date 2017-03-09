package admi

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00900102 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.009.001.02 Document"`
	Message *StaticDataRequestV02 `xml:"StatcDataReq"`
}

func (d *Document00900102) AddMessage() *StaticDataRequestV02 {
	d.Message = new(StaticDataRequestV02)
	return d.Message
}

// The StaticDataRequest message is sent by a participant of a central system to the central system to request a static data report.
//
type StaticDataRequestV02 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *iso20022.Max35Text `xml:"MsgId"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *iso20022.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Details of the request.
	DataRequestDetails *iso20022.RequestDetails3 `xml:"DataReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StaticDataRequestV02) SetMessageIdentification(value string) {
	s.MessageIdentification = (*iso20022.Max35Text)(&value)
}

func (s *StaticDataRequestV02) SetSettlementSessionIdentifier(value string) {
	s.SettlementSessionIdentifier = (*iso20022.Exact4AlphaNumericText)(&value)
}

func (s *StaticDataRequestV02) AddDataRequestDetails() *iso20022.RequestDetails3 {
	s.DataRequestDetails = new(iso20022.RequestDetails3)
	return s.DataRequestDetails
}

func (s *StaticDataRequestV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
