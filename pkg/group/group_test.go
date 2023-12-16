package group

import (
	"net/http"
	"net/http/httptest"
	"testing"

	habbo "github.com/habbography/habbo-go/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGroup_Load(t *testing.T) {
	responseData := `{
		"id": "1234",
		"name": "TestGroup",
		"description": "A test group",
		"type": "CLOSED",
		"roomId": "12345",
		"badgeCode": "b23174s281045b6bc499cf6df518e29e77a89fc174a2",
		"primaryColour": "FFFFFF",
		"secondaryColour": "FFFFFF"
	}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/groups/1234/":
			w.Write([]byte(responseData))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	client := &habbo.BaseClient{
		BaseUrl:    server.URL,
		HttpClient: server.Client(),
	}
	group := NewGroup("1234", client)
	err := group.Load()
	assert.NoError(t, err)
	assert.Equal(t, "TestGroup", group.Name)
	assert.Equal(t, GroupTypeClosed, group.Type)
	assert.Equal(t, "A test group", group.Description)
	assert.Equal(t, "b23174s281045b6bc499cf6df518e29e77a89fc174a2", group.BadgeCode)

}
