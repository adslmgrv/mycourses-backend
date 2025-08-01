package model

import (
	"time"

	"github.com/google/uuid"
)

type Tier string

const (
	TierStudent   Tier = "student"
	TierTeam      Tier = "team"
	TierCommunity Tier = "community"
)

// Subscription represents a user's or team's membership level and its validity period.
//
// ID is generated by an external billing service.
//
// BackedBy specifies the entity responsible for the subscription.
// If the Tier is "team" or "community", this refers to the user (owner) who pays for or manages the team subscription.
// For other tiers ("student"), this simply indicates the user to whom the individual subscription is directly tied.
//
// InvitationCode is a code used for joining collective subscriptions (e.g., "team" tier).
// It is set to null if it's not possible to join them. It is nil if joining via code isn't possible
// for that specific subscription.
//
// Since is updated with each new billing cycle, reflecting the start of the
// current subscription period, not necessarily the initial payment date.
//
// Till denotes when the next recurring payment will be due, marking the end of the
// current subscription period.
type Subscription struct {
	ID             string
	BackedBy       uuid.UUID
	Tier           Tier
	InvitationCode *string
	Till           time.Time
	Since          time.Time
}

// SubscriptionMember links a specific user to a subscription, indicating their active membership.
//
// Since marks the exact time when this user became a member of the specified subscription.
type SubscriptionMember struct {
	SubscriptionID string
	UserID         uuid.UUID
	Since          time.Time
}
