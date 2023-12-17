package user

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	habbo "github.com/habbography/habbo-go/pkg"
)

type Badge struct {
	BadgeIndex  int    `json:"badgeIndex"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	client                      *habbo.BaseClient
	mutex                       sync.Mutex
	isLoaded                    bool
	UniqueId                    string    `json:"uniqueId"`
	Name                        string    `json:"name"`
	Figure                      string    `json:"figureString"`
	Motto                       string    `json:"motto"`
	Online                      bool      `json:"online"`
	LastAccessTime              time.Time `json:"lastAccessTime"`
	MemberSince                 time.Time `json:"memberSince"`
	ProfileVisible              bool      `json:"profileVisible"`
	CurrentLevel                int       `json:"currentLevel"`
	CurrentLevelCompletePercent int       `json:"currentLevelCompletePercent"`
	TotalExperience             int       `json:"totalExperience"`
	StarGemCount                int       `json:"starGemCount"`
	SelectedBadges              []Badge   `json:"selectedBadges"`
}

func NewUser(name string, client *habbo.BaseClient) *User {
	return &User{
		Name:     name,
		client:   client,
		isLoaded: false,
	}
}

func (u *User) Load(ctx context.Context) error {
	if u.isLoaded {
		return nil
	}
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.UniqueId == "" {
		res, err := u.client.Get(ctx, fmt.Sprintf("%s/users?name=%s", u.client.BaseUrl, u.Name))
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(u); err != nil {
			return err
		}
	}
	res, err := u.client.Get(ctx, fmt.Sprintf("%s/users/%s/", u.client.BaseUrl, u.UniqueId))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(u); err != nil {
		return err
	}
	u.isLoaded = true
	return nil
}

func (u *User) GetUniqueId(ctx context.Context) (string, error) {
	if err := u.Load(ctx); err != nil {
		return "", err
	}
	return u.UniqueId, nil
}

func (u *User) GetName(ctx context.Context) (string, error) {
	if err := u.Load(ctx); err != nil {
		return "", err
	}
	return u.Name, nil
}

func (u *User) GetFigure(ctx context.Context) (string, error) {
	if err := u.Load(ctx); err != nil {
		return "", err
	}
	return u.Figure, nil
}

func (u *User) GetMotto(ctx context.Context) (string, error) {
	if err := u.Load(ctx); err != nil {
		return "", err
	}
	return u.Motto, nil
}

func (u *User) GetOnline(ctx context.Context) (bool, error) {
	if err := u.Load(ctx); err != nil {
		return false, err
	}
	return u.Online, nil
}

func (u *User) GetLastAccessTime(ctx context.Context) (time.Time, error) {
	if err := u.Load(ctx); err != nil {
		return time.Now(), err
	}
	return u.LastAccessTime, nil
}

func (u *User) GetMemberSince(ctx context.Context) (time.Time, error) {
	if err := u.Load(ctx); err != nil {
		return time.Now(), err
	}
	return u.MemberSince, nil
}

func (u *User) GetProfileVisible(ctx context.Context) (bool, error) {
	if err := u.Load(ctx); err != nil {
		return false, err
	}
	return u.ProfileVisible, nil
}

func (u *User) GetCurrentLevel(ctx context.Context) (int, error) {
	if err := u.Load(ctx); err != nil {
		return 0, err
	}
	return u.CurrentLevel, nil
}

func (u *User) GetCurrentLevelCompletePercent(ctx context.Context) (int, error) {
	if err := u.Load(ctx); err != nil {
		return 0, err
	}
	return u.CurrentLevelCompletePercent, nil
}

func (u *User) GetTotalExperience(ctx context.Context) (int, error) {
	if err := u.Load(ctx); err != nil {
		return 0, err
	}
	return u.TotalExperience, nil
}

func (u *User) GetStarGemCount(ctx context.Context) (int, error) {
	if err := u.Load(ctx); err != nil {
		return 0, err
	}
	return u.StarGemCount, nil
}

func (u *User) GetSelectedBadges(ctx context.Context) ([]Badge, error) {
	if err := u.Load(ctx); err != nil {
		return nil, err
	}
	return u.SelectedBadges, nil
}
