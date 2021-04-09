package app

import (
	"context"

	"go.temporal.io/sdk/activity"
)

func SendWelcomeEmail(ctx context.Context, customerId string) error {
  activity.GetLogger(ctx).Info("sending welcome email to customer", customerId)
  return nil
}

func SendEndOfTrialEmail(ctx context.Context, name string) error {
  activity.GetLogger(ctx).Info("sending end of trial email to customer", customerId)
  return nil
}

func SendMonthyChargeEmail(ctx context.Context, name string) error {
  activity.GetLogger(ctx).Info("sending monthly charge email to customer", customerId)
  return nil
}

func SendSorryToSeeYouGoEmail(ctx context.Context, name string) error {
  activity.GetLogger(ctx).Info("sending sorry to see you go email to customer", customerId)
  return nil
}

func ChargeMonthlyFee(ctx context.Context, name string) error {
  activity.GetLogger(ctx).Info("charging monthly fee for customer", customerId)
  return nil
}

func ProcessSubscriptionCancellation(ctx context.Context, name string) error {
  activity.GetLogger(ctx).Info("processing subscription cancellation for customer", customerId)
  return nil
}
