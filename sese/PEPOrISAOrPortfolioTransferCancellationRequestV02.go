package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.014.001.02 Document"`
	Message *PEPOrISAOrPortfolioTransferCancellationRequestV02 `xml:"PEPOrISAOrPrtflTrfCxlReqV02"`
}

func (d *Document01400102) AddMessage() *PEPOrISAOrPortfolioTransferCancellationRequestV02 {
	d.Message = new(PEPOrISAOrPortfolioTransferCancellationRequestV02)
	return d.Message
}

// Scope
// An instructing party, eg, a (new) plan manager, sends the PEPOrISAOrPortfolioTransferCancellationRequest message to the executing party, eg, a (old) plan manager, to request the cancellation of a previously sent PEPOrISAOrPortfolioTransferInstruction.
// Usage
// The PEPOrISAOrPortfolioTransferCancellationRequest message is used to request the cancellation of an entire PEPOrISAOrPortfolioTransferInstruction message, ie, all the product transfers that it contained. The cancellation request can be specified either by:
// - quoting the transfer references of all the product transfers listed in the PEPOrISAOrPortfolioTransferInstruction message, or,
// - quoting the details of all the product transfers (this includes TransferReference) listed in PEPOrISAOrPortfolioTransferInstruction message.
// The message identification of the PEPOrISAOrPortfolioTransferInstruction may also be quoted in PreviousReference. It is also possible to request the cancellation of PEPOrISAOrPortfolioTransferInstruction by just quoting its message identification in PreviousReference.
type PEPOrISAOrPortfolioTransferCancellationRequestV02 struct {

	// Identifies the message.
	MessageReference *iso20022.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *iso20022.PEPISATransfer7 `xml:"CxlByTrfInstrDtls,omitempty"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *iso20022.TransferReference3 `xml:"CxlByRef,omitempty"`

}


func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddMessageReference() *iso20022.MessageIdentification1 {
	p.MessageReference = new(iso20022.MessageIdentification1)
	return p.MessageReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddPoolReference() *iso20022.AdditionalReference3 {
	p.PoolReference = new(iso20022.AdditionalReference3)
	return p.PoolReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	p.PreviousReference = new(iso20022.AdditionalReference3)
	return p.PreviousReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	p.RelatedReference = new(iso20022.AdditionalReference3)
	return p.RelatedReference
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddCancellationByTransferInstructionDetails() *iso20022.PEPISATransfer7 {
	p.CancellationByTransferInstructionDetails = new(iso20022.PEPISATransfer7)
	return p.CancellationByTransferInstructionDetails
}

func (p *PEPOrISAOrPortfolioTransferCancellationRequestV02) AddCancellationByReference() *iso20022.TransferReference3 {
	p.CancellationByReference = new(iso20022.TransferReference3)
	return p.CancellationByReference
}

