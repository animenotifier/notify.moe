package paypal

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const adminID = "4J6qpK1ve"

// Success ...
func Success(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	paymentID := ctx.Query("paymentId")
	token := ctx.Query("token")
	payerID := ctx.Query("PayerID")

	if paymentID == "" || payerID == "" || token == "" {
		return ctx.Error(http.StatusBadRequest, "Invalid parameters", errors.New("paymentId, token and PayerID are required"))
	}

	c, err := arn.PayPal()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not initiate PayPal client", err)
	}

	c.SetAccessToken(token)
	execute, err := c.ExecuteApprovedPayment(paymentID, payerID)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error executing PayPal payment", err)
	}

	if execute.State != "approved" {
		return ctx.Error(http.StatusInternalServerError, "PayPal payment has not been approved", err)
	}

	sdkPayment, err := c.GetPayment(paymentID)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not retrieve payment information", err)
	}

	arn.PrettyPrint(sdkPayment)

	transaction := sdkPayment.Transactions[0]

	payment := &arn.PayPalPayment{
		ID:       paymentID,
		PayerID:  payerID,
		UserID:   user.ID,
		Method:   sdkPayment.Payer.PaymentMethod,
		Amount:   transaction.Amount.Total,
		Currency: transaction.Amount.Currency,
		Created:  arn.DateTimeUTC(),
	}

	err = payment.Save()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not save payment in the database", err)
	}

	// Increase user's balance
	user.Balance += payment.Gems()

	// Save in DB
	err = user.Save()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not save new balance", err)
	}

	// Notify admin
	go func() {
		admin, _ := arn.GetUser(adminID)
		admin.SendNotification(&arn.Notification{
			Title:   user.Nick + " bought " + strconv.Itoa(payment.Gems()) + " AD",
			Message: user.Nick + " paid " + payment.Amount + " " + payment.Currency + " making his new balance " + strconv.Itoa(user.Balance),
			Icon:    user.LargeAvatar(),
			Link:    "https://" + ctx.App.Config.Domain + "/api/paypalpayment/" + payment.ID,
		})
	}()

	return ctx.HTML(components.PayPalSuccess(payment))
}
