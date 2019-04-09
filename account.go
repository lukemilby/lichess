package lichess

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type Profile struct {
	Bio string `json:"bio"`
	Country string `json:"country"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Links string `json:"links"`
	Location string `json:"location"`
}

type Count struct {
	AI int `json:"ai"`
	All int `json:"all"`
	Bookmark int `json:"bookmark"`
	Draw int `json:"draw"`
	Import int `json:"import"`
	Loss int `json:"loss"`
	LossH int `json:"lossH"`
	Me int `json:"me"`
	Playing int `json:"playing"`
	Rated int `json:"rated"`
	Win int `json:"win"`
	WinH int `json:"winH"`
}
type PrefsStats struct {
	Games int `json:"games"`
	Prog int `json:"prog"`
	Rating int `json:"rating"`
	RD int `json:"rd"`
}

type Perfs struct {
	Blitz PrefsStats `json:"blitz"`
	Bullet PrefsStats `json:"bullet"`
	Chess960 PrefsStats `json:"chess960"`
	Puzzle PrefsStats `json:"puzzle"`
}

type PlayTime struct {
	Total int `json:"total"`
	Tv int `json:"tv"`
}

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Title string `json:"title"`
	Online bool `json:"online"`
	Playing string `json:"playing"`
	CreatedAt int `json:"createdAt"`
	SeenAt int `json:"seenAt"`
	Profile Profile `json:"profile"`
	NbFollowers int `json:"nbFollowers"`
	NbFollowing int `json:"nbFollowing"`
	CompletionRate int `json:"completionRate"`
	Language string `json:"language"`
	Count Count `json:"count"`
	Perfs Perfs `json:"perfs"`
	Patron bool `json:"patron"`
	Disabled bool `json:"disabled"`
	Engine bool `json:"engine"`
	Booster bool `json:"booster"`
	PlayTime PlayTime `json:"playTime"`
}

type Email struct {
	Email string `json:"email"`
}

type ProfilePrefs struct {
	Dark bool `json:"dark"`
	Transp bool `json:"transp"`
	BackgroundImage string `json:"bgImg"`
	Is3D bool `json:"is3d"`
	Theme string `json:"theme"`
	PieceSet string `json:"pieceSet"`
	Theme3D string `json:"theme3d"`
	PieceSet3D string `json:"pieceSet3d"`
	SoundSet string `json:"soundSet"`
	BlindFold int `json:"blindfold"`
	AutoQueen int `json:"autoQueen"`
	AutoThreefold int `json:"autoThreefold"`
	Takeback int `json:"takeback"`
	ClockTenths int `json:"clockTenths"`
	ClockBar bool `json:"clockBar"`
	ClockSound bool `json:"clockSound"`
	Premove bool `json:"premove"`
	Animation int `json:"animation"`
	Captured bool `json:"captured"`
	Follow bool `json:"follow"`
	Highlight bool `json:"highlight"`
	Destination bool `json:"destination"`
	Coords int `json:"coords"`
	Replay int `json:"replay"`
	Callenge int `json:"callenge"`
	Message int `json:"message"`
	CoordColor int `json:"coordColor"`
	SubmitMove int `json:"submitMove"`
	ConfirmResign int `json:"confirmResign"`
	InsightShare int `json:"insightShare"`
	KeyboardMove int `json:"keyboardMove"`
	Zen int `json:"zen"`
	MoveEvent int `json:"moveEvent"`
}

type Preferences struct {
	Prefs ProfilePrefs `json:"prefs"`
}

type KidModeStatus struct {
	Kid bool `json:"kid"`
}

type Ok struct {
	Ok bool `json:"ok"`
}

// Will return profile information
func (c *Client) GetProfile() (*User, error) {
	req, err := c.newRequest("GET", "/api/account", nil)

	user := new(User)
	_, err = c.do(req, &user)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return user, err
}
// Will return the email for the account
func (c *Client) GetEmail()(*Email, error ){
	req, err := c.newRequest("GET", "/api/account/email", nil)

	email := new(Email)
	_, err = c.do(req, &email)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return email, err
}

// Will return preferences of the account
func (c * Client) GetPreferences()(*Preferences, error){
	req, err := c.newRequest("GET", "/api/account/preferences", nil)

	pref := new(Preferences)
	_, err = c.do(req, &pref)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return pref, err
}

// Will return the kid mode status of the account
func (c * Client) GetKidModeStatus()(*KidModeStatus, error){
	req, err := c.newRequest("GET", "/api/account/kid", nil)

	kid := new(KidModeStatus)
	_, err = c.do(req, &kid)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return kid, err
}

// Set Kid mode on account
func (c * Client) SetKidModeStatus(KidStaus bool)(*Ok, error){
	kidStaus := map[string]bool{
		"kid": KidStaus,
	}
	jsonPayload, _ := json.Marshal(kidStaus)
	req, err := c.newRequest("POST", "/api/account/kid", bytes.NewBuffer(jsonPayload))
	okStatus := new(Ok)
	_, err = c.do(req, &okStatus)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return okStatus, err
}
