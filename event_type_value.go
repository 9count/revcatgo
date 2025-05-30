package revcatgo

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/guregu/null.v4"
)

const (
	EventTypeTest                 = "TEST"
	EventTypeInitialPurchase      = "INITIAL_PURCHASE"
	EventTypeNonRenewingPurchase  = "NON_RENEWING_PURCHASE"
	EventTypeRenewal              = "RENEWAL"
	EventTypeProductChange        = "PRODUCT_CHANGE"
	EventTypeCancellation         = "CANCELLATION"
	EventTypeUnCancellation       = "UNCANCELLATION"
	EventTypeBillingIssue         = "BILLING_ISSUE"
	EventTypeSubscriberAlias      = "SUBSCRIBER_ALIAS"
	EventTypeSubscriptionPaused   = "SUBSCRIPTION_PAUSED"
	EventTypeTransfer             = "TRANSFER"
	EventTypeExpiration           = "EXPIRATION"
	EventTypeSubscriptionExtended = "SUBSCRIPTION_EXTENDED"
)

var validEventTypeValues = []string{
	EventTypeTest,
	EventTypeInitialPurchase,
	EventTypeNonRenewingPurchase,
	EventTypeRenewal,
	EventTypeProductChange,
	EventTypeCancellation,
	EventTypeUnCancellation,
	EventTypeBillingIssue,
	EventTypeSubscriberAlias,
	EventTypeSubscriptionPaused,
	EventTypeTransfer,
	EventTypeExpiration,
	EventTypeSubscriptionExtended,
}

type eventType struct {
	value null.String
}

func newEventType(s string) (*eventType, error) {
	if !contains(validEventTypeValues, s) {
		return &eventType{}, fmt.Errorf("eventType value should be one of the following: %v, got %v", strings.Join(validEventTypeValues, ", "), s)
	}
	return &eventType{value: null.StringFrom(s)}, nil
}

func (e eventType) String() string {
	return e.value.ValueOrZero()
}

// MarshalJSON serializes a store to JSON.
func (e eventType) MarshalJSON() ([]byte, error) {
	return e.value.MarshalJSON()
}

// UnmarshalJSON deserializes a store from JSON
func (e *eventType) UnmarshalJSON(b []byte) error {
	v := &eventType{}
	err := v.value.UnmarshalJSON(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of type: %w", err)
	}
	if !v.value.Valid {
		return errors.New("type is a required field")
	}
	_e, err := newEventType(v.value.ValueOrZero())
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of type: %w", err)
	}
	e.value = _e.value

	return nil
}
