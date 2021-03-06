package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.04 Document"`
	Message *TransferCancellationStatusReportV04 `xml:"TrfCxlStsRpt"`
}

func (d *Document01000104) AddMessage() *TransferCancellationStatusReportV04 {
	d.Message = new(TransferCancellationStatusReportV04)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferCancellationStatusReport message to the instructing party, for example, an investment manager or one of its authorised representatives to provide the status of a previously received transfer cancellation instruction.
// Usage
// The TransferCancellationStatusReport message is used to report on the status of a transfer in or transfer out cancellation request.
// The reference of the transfer instruction for which the cancellation status is reported is identified in TransferReference. The message identification of the transfer cancellation request message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// The message identification of the transfer instruction request message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// One of the following statuses can be reported:
// - the transfer cancellation is accepted, or,
// - the transfer cancellation has been sent to the next party, or,
// - the transfer cancellation is complete and the reason for the status,
// - the transfer cancellation pending and the reason for the status,
// - the transfer cancellation is rejected and the reason for the status.
type TransferCancellationStatusReportV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *model.AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Reference to the message or communication that was previously received.
	Reference *model.References40Choice `xml:"Ref,omitempty"`

	// Status of the transfer cancellation instruction.
	StatusReport *model.CancellationStatusAndReason2 `xml:"StsRpt"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferCancellationStatusReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferCancellationStatusReportV04) AddCounterpartyReference() *model.AdditionalReference2 {
	t.CounterpartyReference = new(model.AdditionalReference2)
	return t.CounterpartyReference
}

func (t *TransferCancellationStatusReportV04) AddReference() *model.References40Choice {
	t.Reference = new(model.References40Choice)
	return t.Reference
}

func (t *TransferCancellationStatusReportV04) AddStatusReport() *model.CancellationStatusAndReason2 {
	t.StatusReport = new(model.CancellationStatusAndReason2)
	return t.StatusReport
}

func (t *TransferCancellationStatusReportV04) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferCancellationStatusReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
