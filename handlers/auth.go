package handlers

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/Oudwins/stackifyer/data"
	"github.com/Oudwins/stackifyer/db"
	"github.com/Oudwins/stackifyer/pkg/apierrors"
	"github.com/Oudwins/stackifyer/types"
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

func SignUp(c *fiber.Ctx) error {
	// credentials := struct {
	// Email                string
	// Password             string
	// PasswordConfirmation string
	// }{}

	newUser := supa.UserCredentials{}
	err := c.BodyParser(&newUser)
	if err != nil {
		print("error")
		print(err.Error())
	}

	fmt.Println(newUser)

	return c.Render("index", fiber.Map{
		"Title": newUser.Email,
	})
}

func SignIn(c *fiber.Ctx) error {
	credentials := supa.UserCredentials{}
	if err := c.BodyParser(&credentials); err != nil {
		return apierrors.New(fmt.Errorf("Invalid credentials: %w", err), http.StatusBadRequest)
	}

	// c.Append("HX-Redirect", "/users")
	ctx := context.Background()
	user, err := db.DbClient.Auth.SignIn(ctx, credentials)

	if err != nil {
		return apierrors.New(err, http.StatusUnauthorized)
	}

	token := new(fiber.Cookie)
	token.Name = "Authorization"
	token.Value = user.AccessToken

	c.Cookie(token)
	return c.SendString(user.User.Role)
}

func SignOut(c *fiber.Ctx) error {
	c.ClearCookie("Authorization")
	return c.Redirect("/")
}

func WithAuthenticated(c *fiber.Ctx) error {
	// TODO: Instead of DB call get the user info from the bloody token
	token := c.Cookies("Authorization")

	if token == "" {
		return c.Next()
	}

	ctx := context.Background()
	user, err := db.DbClient.Auth.User(ctx, token)
	if err == nil {
		c.Locals(types.LocalsKeyAuthenticatedDetails, &data.AuthenticatedDetails{User: user})
	}
	return c.Next()
}

func RestrictedTo(roles []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		details := getAuthenticatedDetails(c)
		if details == nil {
			return c.Redirect("/signin")
		}
		if roles != nil && !slices.Contains(roles, details.User.Role) {
			return c.Redirect("/signin")
		}

		return c.Next()
	}
}

var Restricted = RestrictedTo(nil)

func getAuthenticatedDetails(c *fiber.Ctx) *data.AuthenticatedDetails {
	val := c.Locals(types.LocalsKeyAuthenticatedDetails)
	if details, ok := val.(*data.AuthenticatedDetails); ok {
		return details
	}
	return nil
}
