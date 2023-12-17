package group

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	habbo "github.com/habbography/habbo-go/pkg"
	"github.com/stretchr/testify/assert"
)

// Your existing code...

func TestMembersLoad(t *testing.T) {
	// Mock server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[{"online":true,"gender":"m","motto":"Test Motto","habboFigure":"test-figure","memberSince":"1970-01-01T00:00:00Z","uniqueId":"12345","name":"Test","isAdmin":false}]`))
	}))
	defer server.Close()

	client := &habbo.BaseClient{
		BaseUrl:    server.URL,
		HttpClient: server.Client(),
	}
	testMembers := NewMembers("1234", client)
	ctx := context.Background()

	err := testMembers.Load(ctx)

	expectedMember := Member{
		Online:      true,
		Gender:      "m",
		Motto:       "Test Motto",
		HabboFigure: "test-figure",
		MemberSince: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		UniqueId:    "12345",
		Name:        "Test",
		IsAdmin:     false,
	}

	assert.NoError(t, err)
	assert.Greater(t, len(testMembers.Members), 0)
	assert.EqualValues(t, expectedMember, testMembers.Members[0])
}
