package user

import (
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

func (u *User) Load() error {
	if u.isLoaded {
		return nil
	}
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.UniqueId == "" {
		res, err := u.client.HttpClient.Get(fmt.Sprintf("%s/users?name=%s", u.client.BaseUrl, u.Name))
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(u); err != nil {
			return err
		}
	}
	res, err := u.client.HttpClient.Get(fmt.Sprintf("%s/users/%s/", u.client.BaseUrl, u.UniqueId))
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

func (u *User) Refresh() error {
	u.isLoaded = false
	return u.Load()
}

func (u *User) GetUniqueId() (string, error) {
	if err := u.Load(); err != nil {
		return "", err
	}
	return u.UniqueId, nil
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetFigure() (string, error) {
	if err := u.Load(); err != nil {
		return "", err
	}
	return u.Figure, nil
}

func (u *User) GetMotto() (string, error) {
	if err := u.Load(); err != nil {
		return "", err
	}
	return u.Motto, nil
}

func (u *User) GetOnline() (bool, error) {
	if err := u.Load(); err != nil {
		return false, err
	}
	return u.Online, nil
}

func (u *User) GetLastAccessTime() (time.Time, error) {
	if err := u.Load(); err != nil {
		return time.Now(), err
	}
	return u.LastAccessTime, nil
}

func (u *User) GetMemberSince() (time.Time, error) {
	if err := u.Load(); err != nil {
		return time.Now(), err
	}
	return u.MemberSince, nil
}

func (u *User) GetProfileVisible() (bool, error) {
	if err := u.Load(); err != nil {
		return false, err
	}
	return u.ProfileVisible, nil
}

func (u *User) GetCurrentLevel() (int, error) {
	if err := u.Load(); err != nil {
		return 0, err
	}
	return u.CurrentLevel, nil
}

func (u *User) GetCurrentLevelCompletePercent() (int, error) {
	if err := u.Load(); err != nil {
		return 0, err
	}
	return u.CurrentLevelCompletePercent, nil
}

func (u *User) GetTotalExperience() (int, error) {
	if err := u.Load(); err != nil {
		return 0, err
	}
	return u.TotalExperience, nil
}

func (u *User) GetStarGemCount() (int, error) {
	if err := u.Load(); err != nil {
		return 0, err
	}
	return u.StarGemCount, nil
}

func (u *User) GetSelectedBadges() ([]Badge, error) {
	if err := u.Load(); err != nil {
		return nil, err
	}
	return u.SelectedBadges, nil
}
