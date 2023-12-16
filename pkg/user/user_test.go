package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	habbo "github.com/habbography/habbo-go/pkg"
	"github.com/stretchr/testify/assert"
)

func TestUser_Load(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/users":
			w.Write([]byte(`{"uniqueId": "1234", "name": "test", "figureString": "test-figure"}`))
		case "/users/1234/":
			w.Write([]byte(`{"online": true, "profileVisible": true}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := &habbo.BaseClient{
		BaseUrl:    server.URL,
		HttpClient: server.Client(),
	}

	user := NewUser("test", client)

	err := user.Load()
	assert.NoError(t, err)
	assert.Equal(t, "1234", user.UniqueId)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test-figure", user.Figure)
	assert.True(t, user.ProfileVisible)
	assert.True(t, user.Online)
}
