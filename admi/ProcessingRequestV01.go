package admi

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.017.001.01 Document"`
	Message *ProcessingRequestV01 `xml:"PrcgReq"`
}

func (d *Document01700101) AddMessage() *ProcessingRequestV01 {
	d.Message = new(ProcessingRequestV01)
	return d.Message
}

// The Processing Request message is sent by a participant to a central system to request the initiation of a system process suported by a central system.
type ProcessingRequestV01 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *iso20022.Max35Text `xml:"MsgId"`

	// Indicates the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *iso20022.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Contains the details of the processing request.
	Request *iso20022.RequestDetails19 `xml:"Req"`
}

func (p *ProcessingRequestV01) SetMessageIdentification(value string) {
	p.MessageIdentification = (*iso20022.Max35Text)(&value)
}

func (p *ProcessingRequestV01) SetSettlementSessionIdentifier(value string) {
	p.SettlementSessionIdentifier = (*iso20022.Exact4AlphaNumericText)(&value)
}

func (p *ProcessingRequestV01) AddRequest() *iso20022.RequestDetails19 {
	p.Request = new(iso20022.RequestDetails19)
	return p.Request
}
