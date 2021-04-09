package app

import (
    "time"

    "go.temporal.io/sdk/workflow"
)

// SubscriptionWorkflow is a sample Temporal Workflow function
// that shows what a basic subscription Workflow might look like
func SubscriptionWorkflow(ctx workflow.Context, customerId string, email strin) error {

  trialPeriod := true;
  subscribed := true;

  // Define Activity options
	activityoptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
  // Add Activity options to context
	ctx = workflow.WithActivityOptions(ctx, activityoptions)

  // Send welcome email
	err := workflow.ExecuteActivity(ctx, SendWelcomeEmail, customerId).Get(ctx, nil)
	if err != nil {
		return err
	}

  // TODO
  // Add logic to capture Signal that unsubscribes customer

  for subscribed {
    if !trialPeriod {
      err := workflow.ExecuteActivity(ctx, ChargeMonthlyFee, customerId).Get(ctx, nil)
      if err != nil {
        return err
      }
      err := workflow.ExecuteActivity(ctx, SendMonthyChargeEmail, customerId).Get(ctx, nil)
      if err != nil {
        return err
      }
    } else {
      workflow.GetLogger(ctx).Info("Trial period active, no charge for customerId", customerId)
      trialPeriod = false
    }
    _ = workflow.Sleep(ctx, 30*time.Second)

  }

	workflow.GetLogger(ctx).Info("Done", zap.String("result", result))
	return nil
}
