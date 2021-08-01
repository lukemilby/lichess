package lichess

import (
	"fmt"
	"log"
	"os"
)

type UserRealTimeStatus struct {
}
type TopGamePrefs struct {
	Rating   int `json:"rating"`
	Progress int `json:"progress"`
}

type BulletPrefs struct {
	Bullet TopGamePrefs `json:"bullet"`
}

type BlitzPrefs struct {
	Blitz TopGamePrefs `json:"blitz"`
}
type RapidPrefs struct {
	Rapid TopGamePrefs `json:"rapid"`
}
type ClassicalPrefs struct {
	Classical TopGamePrefs `json:"classical"`
}
type UltraBulletPrefs struct {
	UltraBullet TopGamePrefs `json:"ultraBullet"`
}
type CrazyhousePrefs struct {
	Crazyhouse TopGamePrefs `json:"crazyhouse"`
}
type Chess960Prefs struct {
	Chess960 TopGamePrefs `json:"chess960"`
}
type KingOfTheHillPrefs struct {
	KingOfTheHill TopGamePrefs `json:"kingOfTheHill"`
}
type ThreeCheckPrefs struct {
	ThreeCheck TopGamePrefs `json:"threeCheck"`
}
type AntichessPrefs struct {
	Antichess TopGamePrefs `json:"antichess"`
}
type AtomicPrefs struct {
	Atomic TopGamePrefs `json:"atomic"`
}
type HordePrefs struct {
	Horde TopGamePrefs `json:"horde"`
}
type RacingKingsPrefs struct {
	RacingKings TopGamePrefs `json:"racingKings"`
}

type BulletPlayer struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Prefs    BulletPrefs `json:"prefs"`
	Title    string      `json:"title"`
	Patron   bool        `json:"patron"`
	Online   bool        `json:"online"`
}

type BlitzPlayer struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Prefs    BlitzPrefs `json:"prefs"`
	Title    string     `json:"title"`
	Patron   bool       `json:"patron"`
	Online   bool       `json:"online"`
}

type RapidPlayer struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Prefs    RapidPrefs `json:"prefs"`
	Title    string     `json:"title"`
	Patron   bool       `json:"patron"`
	Online   bool       `json:"online"`
}

type ClassicalPlayer struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Prefs    ClassicalPrefs `json:"prefs"`
	Title    string         `json:"title"`
	Patron   bool           `json:"patron"`
	Online   bool           `json:"online"`
}

type UltraBulletPlayer struct {
	ID       string           `json:"id"`
	Username string           `json:"username"`
	Prefs    UltraBulletPrefs `json:"prefs"`
	Title    string           `json:"title"`
	Patron   bool             `json:"patron"`
	Online   bool             `json:"online"`
}

type CrazyHousePlayer struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Prefs    CrazyhousePrefs `json:"prefs"`
	Title    string          `json:"title"`
	Patron   bool            `json:"patron"`
	Online   bool            `json:"online"`
}

type Chess960Player struct {
	ID       string        `json:"id"`
	Username string        `json:"username"`
	Prefs    Chess960Prefs `json:"prefs"`
	Title    string        `json:"title"`
	Patron   bool          `json:"patron"`
	Online   bool          `json:"online"`
}

type KingOfTheHillPlayer struct {
	ID       string             `json:"id"`
	Username string             `json:"username"`
	Prefs    KingOfTheHillPrefs `json:"prefs"`
	Title    string             `json:"title"`
	Patron   bool               `json:"patron"`
	Online   bool               `json:"online"`
}

type ThreeCheckPlayer struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Prefs    ThreeCheckPrefs `json:"prefs"`
	Title    string          `json:"title"`
	Patron   bool            `json:"patron"`
	Online   bool            `json:"online"`
}

type AntichessPlayer struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Prefs    AntichessPrefs `json:"prefs"`
	Title    string         `json:"title"`
	Patron   bool           `json:"patron"`
	Online   bool           `json:"online"`
}

type AtomicPlayer struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Prefs    AtomicPrefs `json:"prefs"`
	Title    string      `json:"title"`
	Patron   bool        `json:"patron"`
	Online   bool        `json:"online"`
}
type HordePlayer struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Prefs    HordePrefs `json:"prefs"`
	Title    string     `json:"title"`
	Patron   bool       `json:"patron"`
	Online   bool       `json:"online"`
}
type RacingKingsPlayer struct {
	ID       string           `json:"id"`
	Username string           `json:"username"`
	Prefs    RacingKingsPrefs `json:"prefs"`
	Title    string           `json:"title"`
	Patron   bool             `json:"patron"`
	Online   bool             `json:"online"`
}

type TopPlayers struct {
	Bullet        []BulletPlayer        `json:"bullet"`
	Blitz         []BlitzPlayer         `json:"blitz"`
	Rapid         []RapidPlayer         `json:"rapid"`
	Classical     []ClassicalPlayer     `json:"classical"`
	UltraBullet   []UltraBulletPlayer   `json:"ultraBullet"`
	CrazyHouse    []CrazyHousePlayer    `json:"crazyhouse"`
	Chess960      []Chess960Player      `json:"chess960"`
	KingOfTheHill []KingOfTheHillPlayer `json:"kingOfTheHill"`
	ThreeCheck    []ThreeCheckPlayer    `json:"threeCheck"`
	Antichess     []AntichessPlayer     `json:"antichess"`
	Atomic        []AtomicPlayer        `json:"atomic"`
	Horde         []HordePlayer         `json:"horde"`
	RacingKings   []RacingKingsPlayer   `json:"racingKings"`
}

type LeaderStats struct {
	Rating   int `json:"rating"`
	Progress int `json:"progress"`
}

type LeaderUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Prefs    struct {
		Bullet        LeaderStats `json:"bullet"`
		Blitz         LeaderStats `json:"blitz"`
		Rapid         LeaderStats `json:"rapid"`
		Classical     LeaderStats `json:"classical"`
		UltraBullet   LeaderStats `json:"ultraBullet"`
		CrazyHouse    LeaderStats `json:"crazyHouse"`
		Chess960      LeaderStats `json:"chess960"`
		KingOfTheHill LeaderStats `json:"kingOfTheHill"`
		ThreeCheck    LeaderStats `json:"threeCheck"`
		Antichess     LeaderStats `json:"antichess"`
		Atomic        LeaderStats `json:"atomic"`
		Horde         LeaderStats `json:"horde"`
		RacingKings   LeaderStats `json:"racingKings"`
	}
	Online bool   `json:"online"`
	Title  string `json:"title"`
	Patron bool   `json:"patron"`
}

type LeaderBoard struct {
	Users []LeaderUser `json:"users"`
}

//Player history struct
type HistoryDate struct {
}

type PrefType struct {
	Name   string `json:"name"`
	Points []int
}

type PlayerHistory struct {
	History []PrefType
}

func (c *Client) GetUserPublicData(user_name string) (*User, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s", user_name), nil)
	req.Header.Set("Accept", "application/vnd.lichess.v3+json")
	user := new(User)
	_, err = c.do(req, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

type UserStatus struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Online    bool   `json:"online"`
	Playing   bool   `json:"playing"`
	Streaming bool   `json:"streaming"`
	Patron    bool   `json:"patron"`
}

func (c *Client) GetRLUsersStatus(ids string) (*[]UserStatus, error) {
	req, err := c.newRequest("GET", "/api/users/status", nil)

	q := req.URL.Query()
	q.Add("ids", ids)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/vnd.lichess.v3+json")
	users_status := new([]UserStatus)
	_, err = c.do(req, &users_status)
	if err != nil {
		return nil, err
	}
	return users_status, err
}

//Gets top 10 players for each game type
func (c *Client) GetTopPlayers() (*TopPlayers, error) {
	req, err := c.newRequest("GET", "/player", nil)

	players := new(TopPlayers)
	_, err = c.do(req, &players)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return players, err
}

// Gets Leaderboard for game type as prefType and a max number of results
func (c *Client) GetLeaderboard(nb int, prefType string) (*LeaderBoard, error) {
	endpoint := fmt.Sprintf("/player/top/%d/%s", nb, prefType)
	req, err := c.newRequest("GET", endpoint, nil)
	req.Header.Set("Accept", "application/vnd.lichess.v3+json")
	leaderBoard := new(LeaderBoard)
	_, err = c.do(req, &leaderBoard)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return leaderBoard, err
}

// Gets Player information
func (c *Client) GetPlayer(username string) (*User, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s", username), nil)
	req.Header.Set("Accept", "application/vnd.lichess.v3+json")
	user := new(User)
	_, err = c.do(req, &user)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return user, err
}

// Gets player history
func (c *Client) GetPlayerHistory(username string) (*PlayerHistory, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s/rating-history", username), nil)
	req.Header.Set("Accept", "application/vnd.lichess.v3+json")
	history := new(PlayerHistory)
	_, err = c.do(req, &history)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return history, err
}
