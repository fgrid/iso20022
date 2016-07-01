package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.014.001.03 Document"`
	Message *PortfolioTransferCancellationRequestV03 `xml:"PrtflTrfCxlReq"`
}

func (d *Document01400103) AddMessage() *PortfolioTransferCancellationRequestV03 {
	d.Message = new(PortfolioTransferCancellationRequestV03)
	return d.Message
}

// Scope
// An instructing party, eg, a (new) plan manager (Transferee), sends the PortfolioTransferCancellationRequest message to the executing party, eg, a (old) plan manager (Transferor), to request the cancellation of a previously sent PortfolioTransferInstruction.
// Usage
// The PortfolioTransferCancellationRequest message is used to request the cancellation of an entire PortfolioTransferInstruction message, ie, all the product transfers that it contained. The cancellation request can be specified either by:
// - quoting the transfer references of all the product transfers listed in the PortfolioTransferInstruction message, or,
// - quoting the details of all the product transfers (this includes TransferReference) listed in PortfolioTransferInstruction message.
// The message identification of the PortfolioTransferInstruction may also be quoted in PreviousReference. It is also possible to request the cancellation of PortfolioTransferInstruction by just quoting its message identification in PreviousReference.
type PortfolioTransferCancellationRequestV03 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *iso20022.PEPISATransfer11 `xml:"CxlByTrfInstrDtls,omitempty"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *iso20022.TransferReference3 `xml:"CxlByRef,omitempty"`

}


func (p *PortfolioTransferCancellationRequestV03) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferCancellationRequestV03) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PortfolioTransferCancellationRequestV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PortfolioTransferCancellationRequestV03) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PortfolioTransferCancellationRequestV03) AddCancellationByTransferInstructionDetails() *iso20022.PEPISATransfer11 {
	p.CancellationByTransferInstructionDetails = new(iso20022.PEPISATransfer11)
	return p.CancellationByTransferInstructionDetails
}

func (p *PortfolioTransferCancellationRequestV03) AddCancellationByReference() *iso20022.TransferReference3 {
	p.CancellationByReference = new(iso20022.TransferReference3)
	return p.CancellationByReference
}

