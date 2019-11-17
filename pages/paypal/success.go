package paypal

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
)

const adminID = "4J6qpK1ve"

// Success is called once the payment has been confirmed by the user on the PayPal website.
// However, the actual payment still needs to be executed and can fail.
func Success(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
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

	// Get access token
	_, err = c.GetAccessToken()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not get PayPal access token", err)
	}

	// Execute payment
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

	stringutils.PrettyPrint(sdkPayment)
	transaction := sdkPayment.Transactions[0]
	payment := arn.NewPayPalPayment(paymentID, payerID, user.ID, sdkPayment.Payer.PaymentMethod, transaction.Amount.Total, transaction.Amount.Currency)
	payment.Save()

	// Increase user's balance
	user.Balance += payment.Gems()

	// Save in DB
	user.Save()

	// Notify admin
	go func() {
		admin, _ := arn.GetUser(adminID)
		admin.SendNotification(&arn.PushNotification{
			Title:   user.Nick + " bought " + strconv.Itoa(payment.Gems()) + " gems",
			Message: user.Nick + " paid " + payment.Amount + " " + payment.Currency + " making his new balance " + strconv.Itoa(user.Balance),
			Icon:    user.AvatarLink("large"),
			Link:    "https://" + assets.Domain + "/api/paypalpayment/" + payment.ID,
			Type:    arn.NotificationTypePurchase,
		})
	}()

	return ctx.HTML(components.PayPalSuccess(payment))
}
