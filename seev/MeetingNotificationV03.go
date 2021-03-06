package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.03 Document"`
	Message *MeetingNotificationV03 `xml:"MtgNtfctn"`
}

func (d *Document00100103) AddMessage() *MeetingNotificationV03 {
	d.Message = new(MeetingNotificationV03)
	return d.Message
}

// Scope
// A notifying party, eg, an issuer, its agent or an intermediary, sends the MeetingNotification message to a party holding the right to vote, to announce a shareholders meeting.
// Usage
// The MeetingNotification message is used to announce a shareholders meeting, for example, it provides information on the participation details and requirements for the meeting, the vote parameters and the resolutions. The MeetingNotification message may also be used to announce an update.
// To notify an update, the Amendment building block must be filled in. Any building block that is modified must be included in the amendment message. The information previously notified and not repeated in the amendment message remains valid.
// To update the resolutions of the agenda, the complete list of resolutions must be repeated in the amendment message. The resolutions that are deleted should be assigned the status Withdrawn.
type MeetingNotificationV03 struct {

	// Identifies the meeting notification message.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Information specific to an amendment.
	Amendment *model.AmendInformation1 `xml:"Amdmnt,omitempty"`

	// Defines the global status of the event contained in the notification.
	NotificationStatus *model.NotificationStatus1 `xml:"NtfctnSts"`

	// Specifies information about the meeting. This component contains meeting identifications, various deadlines, contact persons, electronic and postal locations for accessing information and proxy assignment parameters.
	Meeting *model.MeetingNotice3 `xml:"Mtg"`

	// Dates and details of the shareholders meeting.
	MeetingDetails []*model.Meeting3 `xml:"MtgDtls"`

	// Party notifying the meeting.
	NotifyingParty *model.PartyIdentification9Choice `xml:"NtifngPty"`

	// Specifies the institution that is the issuer of the security to which the meeting applies.
	Issuer *model.IssuerInformation1 `xml:"Issr"`

	// Agents of the issuer.
	IssuerAgent []*model.IssuerAgent1 `xml:"IssrAgt,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	Security []*model.SecurityPosition6 `xml:"Scty"`

	// Detailed information of a resolution proposed to the vote.
	Resolution []*model.Resolution2 `xml:"Rsltn,omitempty"`

	// Specifies the conditions to be allowed to vote, the different voting methods and options, the voting deadlines and the parameters of the incentive premium.
	Vote *model.VoteParameters2 `xml:"Vote,omitempty"`

	// Specifies the entitlement ratio and the different deadlines for calculating the entitlement.
	EntitlementSpecification *model.EntitlementAssessment2 `xml:"EntitlmntSpcfctn"`

	// Specifies requirements relative to the use of Power of Attorney.
	PowerOfAttorneyRequirements *model.PowerOfAttorneyRequirements2 `xml:"PwrOfAttnyRqrmnts,omitempty"`
}

func (m *MeetingNotificationV03) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingNotificationV03) AddAmendment() *model.AmendInformation1 {
	m.Amendment = new(model.AmendInformation1)
	return m.Amendment
}

func (m *MeetingNotificationV03) AddNotificationStatus() *model.NotificationStatus1 {
	m.NotificationStatus = new(model.NotificationStatus1)
	return m.NotificationStatus
}

func (m *MeetingNotificationV03) AddMeeting() *model.MeetingNotice3 {
	m.Meeting = new(model.MeetingNotice3)
	return m.Meeting
}

func (m *MeetingNotificationV03) AddMeetingDetails() *model.Meeting3 {
	newValue := new(model.Meeting3)
	m.MeetingDetails = append(m.MeetingDetails, newValue)
	return newValue
}

func (m *MeetingNotificationV03) AddNotifyingParty() *model.PartyIdentification9Choice {
	m.NotifyingParty = new(model.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingNotificationV03) AddIssuer() *model.IssuerInformation1 {
	m.Issuer = new(model.IssuerInformation1)
	return m.Issuer
}

func (m *MeetingNotificationV03) AddIssuerAgent() *model.IssuerAgent1 {
	newValue := new(model.IssuerAgent1)
	m.IssuerAgent = append(m.IssuerAgent, newValue)
	return newValue
}

func (m *MeetingNotificationV03) AddSecurity() *model.SecurityPosition6 {
	newValue := new(model.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingNotificationV03) AddResolution() *model.Resolution2 {
	newValue := new(model.Resolution2)
	m.Resolution = append(m.Resolution, newValue)
	return newValue
}

func (m *MeetingNotificationV03) AddVote() *model.VoteParameters2 {
	m.Vote = new(model.VoteParameters2)
	return m.Vote
}

func (m *MeetingNotificationV03) AddEntitlementSpecification() *model.EntitlementAssessment2 {
	m.EntitlementSpecification = new(model.EntitlementAssessment2)
	return m.EntitlementSpecification
}

func (m *MeetingNotificationV03) AddPowerOfAttorneyRequirements() *model.PowerOfAttorneyRequirements2 {
	m.PowerOfAttorneyRequirements = new(model.PowerOfAttorneyRequirements2)
	return m.PowerOfAttorneyRequirements
}
