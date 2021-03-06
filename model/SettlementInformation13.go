package model

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation13 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount16 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount16 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount16 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount16 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation13) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInformation13) AddSettlementAccount() *CashAccount16 {
	s.SettlementAccount = new(CashAccount16)
	return s.SettlementAccount
}

func (s *SettlementInformation13) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}

func (s *SettlementInformation13) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation13) AddInstructingReimbursementAgentAccount() *CashAccount16 {
	s.InstructingReimbursementAgentAccount = new(CashAccount16)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation13) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation13) AddInstructedReimbursementAgentAccount() *CashAccount16 {
	s.InstructedReimbursementAgentAccount = new(CashAccount16)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInformation13) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInformation13) AddThirdReimbursementAgentAccount() *CashAccount16 {
	s.ThirdReimbursementAgentAccount = new(CashAccount16)
	return s.ThirdReimbursementAgentAccount
}
