package group

import (
	"encoding/json"
	"fmt"
	"sync"

	habbo "github.com/habbography/habbo-go/pkg"
)

type GroupType string

const (
	GroupTypeNormal    GroupType = "NORMAL"
	GroupTypeClosed    GroupType = "CLOSED"
	GroupTypeExclusive GroupType = "EXCLUSIVE"
)

type Group struct {
	client          *habbo.BaseClient
	mutex           sync.Mutex
	isLoaded        bool
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Type            GroupType `json:"type"`
	RoomId          string    `json:"roomId"`
	BadgeCode       string    `json:"badgeCode"`
	PrimaryColour   string    `json:"primaryColour"`
	SecondaryColour string    `json:"secondaryColour"`
}

func NewGroup(groupId string, client *habbo.BaseClient) *Group {
	return &Group{
		client:   client,
		Id:       groupId,
		isLoaded: false,
	}
}

func (g *Group) Load() error {
	if g.isLoaded {
		return nil
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()
	res, err := g.client.HttpClient.Get(fmt.Sprintf("%s/groups/%s/", g.client.BaseUrl, g.Id))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(g); err != nil {
		return err
	}
	g.isLoaded = true
	return nil
}

func (g *Group) GetId() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.Id, nil
}

func (g *Group) GetName() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}

	}
	return g.Name, nil
}

func (g *Group) GetDescription() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.Description, nil
}

func (g *Group) GetType() (GroupType, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.Type, nil
}

func (g *Group) GetRoomId() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.RoomId, nil
}

func (g *Group) GetBadgeCode() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.BadgeCode, nil
}

func (g *Group) GetPrimaryColour() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.PrimaryColour, nil
}

func (g *Group) GetSecondaryColour() (string, error) {
	if !g.isLoaded {
		if err := g.Load(); err != nil {
			return "", err
		}
	}
	return g.SecondaryColour, nil
}
