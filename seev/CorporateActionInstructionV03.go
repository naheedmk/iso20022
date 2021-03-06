package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300103 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.03 Document"`
	Message *CorporateActionInstructionV03 `xml:"CorpActnInstr"`
}

func (d *Document03300103) AddMessage() *CorporateActionInstructionV03 {
	d.Message = new(CorporateActionInstructionV03)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstruction message to an account servicer to instruct election on a corporate action event.
// This message is used to provide the custodian with instructions on how the account owner wishes to proceed with a corporate action event. Instructions include investment decisions regarding the exercise of rights issues, the election of stock or cash when the option is available, and decisions on the conversion or tendering of securities.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionV03 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *model.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent cancelled instruction document.
	CancelledInstructionIdentification *model.DocumentIdentification15 `xml:"CancInstrId,omitempty"`

	// Identification of a previously sent instruction cancellation request document.
	InstructionCancellationRequestIdentification *model.DocumentIdentification15 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation36 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance17 `xml:"AcctDtls"`

	// Provides information about the beneficial owner of the securities.
	BeneficialOwnerDetails []*model.PartyIdentification56 `xml:"BnfclOwnrDtls,omitempty"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *model.CorporateActionOption38 `xml:"CorpActnInstr"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative21 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionV03) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*model.YesNoIndicator)(&value)
}

func (c *CorporateActionInstructionV03) AddCancelledInstructionIdentification() *model.DocumentIdentification15 {
	c.CancelledInstructionIdentification = new(model.DocumentIdentification15)
	return c.CancelledInstructionIdentification
}

func (c *CorporateActionInstructionV03) AddInstructionCancellationRequestIdentification() *model.DocumentIdentification15 {
	c.InstructionCancellationRequestIdentification = new(model.DocumentIdentification15)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionV03) AddOtherDocumentIdentification() *model.DocumentIdentification13 {
	newValue := new(model.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionV03) AddEventsLinkage() *model.CorporateActionEventReference1 {
	newValue := new(model.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionInstructionV03) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation36 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation36)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionV03) AddAccountDetails() *model.AccountAndBalance17 {
	c.AccountDetails = new(model.AccountAndBalance17)
	return c.AccountDetails
}

func (c *CorporateActionInstructionV03) AddBeneficialOwnerDetails() *model.PartyIdentification56 {
	newValue := new(model.PartyIdentification56)
	c.BeneficialOwnerDetails = append(c.BeneficialOwnerDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionV03) AddCorporateActionInstruction() *model.CorporateActionOption38 {
	c.CorporateActionInstruction = new(model.CorporateActionOption38)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionV03) AddAdditionalInformation() *model.CorporateActionNarrative21 {
	c.AdditionalInformation = new(model.CorporateActionNarrative21)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
