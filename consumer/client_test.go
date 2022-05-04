package consumer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dmitry-shidlovsky/TestPact/model"
)

func TestClient_GetUser_unit(t *testing.T) {
	userID := 10

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), fmt.Sprintf("/user/%d", userID))
		user, _ := json.Marshal(model.User{
			FirstName: "TestFirstName",
			LastName:  "TestLastName",
			ID:        userID,
			Type:      "admin",
			Username:  "testnickname",
		})
		rw.Write(user)
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	client := &Client{
		BaseURL: u,
	}
	user, err := client.GetUser(userID)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, userID)
}
