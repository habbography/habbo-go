package group

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	habbo "github.com/habbography/habbo-go/pkg"
)

type Member struct {
	Online      bool      `json:"online"`
	Gender      string    `json:"gender"`
	Motto       string    `json:"motto"`
	HabboFigure string    `json:"habboFigure"`
	MemberSince time.Time `json:"memberSince"`
	UniqueId    string    `json:"uniqueId"`
	Name        string    `json:"name"`
	IsAdmin     bool      `json:"isAdmin"`
}

type Members struct {
	client   *habbo.BaseClient
	mutex    sync.Mutex
	isLoaded bool
	Id       string
	Members  []Member `json:"members"`
}

func NewMembers(groupId string, client *habbo.BaseClient) *Members {
	return &Members{
		Id:       groupId,
		client:   client,
		isLoaded: false,
	}
}

func (m *Members) Load(ctx context.Context) error {
	if m.isLoaded {
		return nil
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	res, err := m.client.Get(ctx, fmt.Sprintf("%s/groups/%s/members", m.client.BaseUrl, m.Id))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&m.Members); err != nil {
		return err
	}
	m.isLoaded = true
	return nil
}

func (m *Members) GetMembers(ctx context.Context) ([]Member, error) {
	if !m.isLoaded {
		if err := m.Load(ctx); err != nil {
			return nil, err
		}
	}
	return m.Members, nil
}
