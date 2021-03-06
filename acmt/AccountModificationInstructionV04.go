package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.003.001.04 Document"`
	Message *AccountModificationInstructionV04 `xml:"AcctModInstr"`
}

func (d *Document00300104) AddMessage() *AccountModificationInstructionV04 {
	d.Message = new(AccountModificationInstructionV04)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent, sends the AccountModificationInstruction message to the account servicer, for example, a registrar, transfer agent or custodian bank to modify, that is, create, update or delete specific details of an existing investment fund account.
// Usage
// The AccountModificationInstruction message is used to modify the details of an existing account.
// The AccountModificationInstruction message has three specific uses:
// - to maintain/update any of the existing account details, for example, to update the address of the beneficiary or modify the preference to income from distribution to capitalisation, or,
// - to add/create specific details to the existing account when these details were not yet recorded at the time of account creation, for example, to add a second address or to establish new cash settlement standing instructions, or,
// - to delete specific account details, for example, delete cash standing instructions.
// This message cannot be used to delete an entire account, as institution specific and regulatory rules pertaining to account deletion are diverse.
// The usage of this message may be subject to service level agreement (SLA) between the counterparties.
// Execution of the AccountModificationInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountModificationInstructionV04 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Provides detailed information about the application modification instruction.
	InstructionDetails *model.InvestmentAccountModificationDetails `xml:"InstrDtls,omitempty"`

	// Investment account selection information used to identify the account for which the information is modified.
	InvestmentAccountSelection *model.InvestmentAccountSelection2 `xml:"InvstmtAcctSelctn"`

	// Information related to general characteristics of an investment account to be inserted, updated or deleted.
	ModifiedInvestmentAccount *model.InvestmentAccount39 `xml:"ModfdInvstmtAcct,omitempty"`

	// Information related to the account related parties, for example, primary owner to be inserted, updated or deleted.
	ModifiedAccountParties []*model.AccountParties8 `xml:"ModfdAcctPties,omitempty"`

	// Information related to intermediaries to be inserted, updated or deleted.
	ModifiedIntermediaries []*model.ModificationScope20 `xml:"ModfdIntrmies,omitempty"`

	// Information related to referred placement agent in the hedge fund industry to be inserted, updated or deleted.
	ModifiedPlacement *model.ReferredAgent1 `xml:"ModfdPlcmnt,omitempty"`

	// Eligibility conditions information related to new issues allocation to be inserted, updated or deleted.
	ModifiedIssueAllocation *model.ModificationScope9 `xml:"ModfdIsseAllcn,omitempty"`

	// Information related to a savings plan to be either inserted, updated or deleted.
	ModifiedSavingsInvestmentPlan []*model.ModificationScope18 `xml:"ModfdSvgsInvstmtPlan,omitempty"`

	// Information related to a withdrawal plan to be either inserted, updated or deleted.
	ModifiedWithdrawalInvestmentPlan []*model.ModificationScope18 `xml:"ModfdWdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction associated to the investment fund transaction and to be either inserted or deleted.
	ModifiedCashSettlement []*model.InvestmentFundCashSettlementInformation8 `xml:"ModfdCshSttlm,omitempty"`

	// Information related to documents to be added, deleted or updated.
	//
	ModifiedServiceLevelAgreement []*model.ModificationScope10 `xml:"ModfdSvcLvlAgrmt,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountModificationInstructionV04) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountModificationInstructionV04) AddPreviousReference() *model.AdditionalReference3 {
	a.PreviousReference = new(model.AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountModificationInstructionV04) AddInstructionDetails() *model.InvestmentAccountModificationDetails {
	a.InstructionDetails = new(model.InvestmentAccountModificationDetails)
	return a.InstructionDetails
}

func (a *AccountModificationInstructionV04) AddInvestmentAccountSelection() *model.InvestmentAccountSelection2 {
	a.InvestmentAccountSelection = new(model.InvestmentAccountSelection2)
	return a.InvestmentAccountSelection
}

func (a *AccountModificationInstructionV04) AddModifiedInvestmentAccount() *model.InvestmentAccount39 {
	a.ModifiedInvestmentAccount = new(model.InvestmentAccount39)
	return a.ModifiedInvestmentAccount
}

func (a *AccountModificationInstructionV04) AddModifiedAccountParties() *model.AccountParties8 {
	newValue := new(model.AccountParties8)
	a.ModifiedAccountParties = append(a.ModifiedAccountParties, newValue)
	return newValue
}

func (a *AccountModificationInstructionV04) AddModifiedIntermediaries() *model.ModificationScope20 {
	newValue := new(model.ModificationScope20)
	a.ModifiedIntermediaries = append(a.ModifiedIntermediaries, newValue)
	return newValue
}

func (a *AccountModificationInstructionV04) AddModifiedPlacement() *model.ReferredAgent1 {
	a.ModifiedPlacement = new(model.ReferredAgent1)
	return a.ModifiedPlacement
}

func (a *AccountModificationInstructionV04) AddModifiedIssueAllocation() *model.ModificationScope9 {
	a.ModifiedIssueAllocation = new(model.ModificationScope9)
	return a.ModifiedIssueAllocation
}

func (a *AccountModificationInstructionV04) AddModifiedSavingsInvestmentPlan() *model.ModificationScope18 {
	newValue := new(model.ModificationScope18)
	a.ModifiedSavingsInvestmentPlan = append(a.ModifiedSavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV04) AddModifiedWithdrawalInvestmentPlan() *model.ModificationScope18 {
	newValue := new(model.ModificationScope18)
	a.ModifiedWithdrawalInvestmentPlan = append(a.ModifiedWithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV04) AddModifiedCashSettlement() *model.InvestmentFundCashSettlementInformation8 {
	newValue := new(model.InvestmentFundCashSettlementInformation8)
	a.ModifiedCashSettlement = append(a.ModifiedCashSettlement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV04) AddModifiedServiceLevelAgreement() *model.ModificationScope10 {
	newValue := new(model.ModificationScope10)
	a.ModifiedServiceLevelAgreement = append(a.ModifiedServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV04) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountModificationInstructionV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
