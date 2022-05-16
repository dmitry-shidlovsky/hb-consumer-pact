package consumer

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"

	"github.com/dmitry-shidlovsky/TestPact/model"
)

var commonHeaders = dsl.MapMatcher{
	"Content-Type":         term("application/json; charset=utf-8", `application\/json`),
	"X-Api-Correlation-Id": dsl.Like("100"),
}

func TestClient_GetUser_pact(t *testing.T) {
	t.Run("the user exists", func(t *testing.T) {
		id := 1

		pact.
			AddInteraction().
			Given("User dmitry exists").
			UponReceiving("A request to login with user 'dmitry'").
			WithRequest(request{
				Method: "GET",
				Path:   term("/user/1", "/user/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status:  200,
				Body:    dsl.Match(model.User{}),
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			user, err := client.GetUser(id)

			if err != nil {
				return err
			}

			if user.ID != 1 {
				t.Fatal("Oops")
			}

			return nil
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
}

func TestClient_GetUsers_pact(t *testing.T) {
	t.Run("the user exists", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("Users exists").
			UponReceiving("A request to login with users ").
			WithRequest(request{
				Method: "GET",
				Path:   term("/users/", "/users"),
			}).
			WillRespondWith(dsl.Response{
				Status:  200,
				Body:    dsl.Match([]model.User{}),
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			user, err := client.GetUsers()

			if err != nil {
				return err
			}

			if len(user) != 1 {
				t.Fatal("Oops")
			}
			if user[0].ID != 1 {
				t.Fatal("Oops")
			}

			return nil
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
}
