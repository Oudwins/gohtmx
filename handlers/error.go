package handlers

import (
	"errors"

	"github.com/Oudwins/stackifyer/pkg/apierrors"
	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	print(err.Error())

	var appErr *apierrors.AppErr
	if errors.As(err, &appErr) {
		status := appErr.Status
		return c.Status(status).Render("partials/alerts/alertError", fiber.Map{
			"Msg": appErr.Error(),
		}, "partials/alerts/globalAlertBox")
	}
	// remember apierrors.InternalServerError, already instanciated
	return c.Status(401).Render("partials/alerts/alertError", fiber.Map{
		"Msg": err.Error(),
	}, "partials/alerts/globalAlertBox")
}
