package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/glog"
)

// AlexaMessaging Alexa Messaging interface
type AlexaMessaging struct {
	ClientID     string
	clientSecret string
	Bearer       string
	BearerExpiry time.Time
}

// NewAlexaMessaging Creates a new initialized AlexaMessaging
func NewAlexaMessaging(clientID, clientSecret string) *AlexaMessaging {
	a := new(AlexaMessaging)
	a.ClientID = clientID
	a.clientSecret = clientSecret

	return a
}

func (a *AlexaMessaging) getBearer() error {
	var tokenURL string = "https://api.amazon.com/auth/O2/token"
	client := &http.Client{}

	data := url.Values{}
	data.Set("client_id", a.ClientID)
	data.Add("client_secret", a.clientSecret)
	data.Add("grant_type", "client_credentials")
	data.Add("scope", "alexa:skill_messaging")

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	if err != nil {
		glog.Errorln(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorln(err)
		return err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorln(err)
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		glog.Errorln(err)
		return err
	}
	if resp.StatusCode == 200 {
		token := make(map[string]interface{})
		err := json.Unmarshal(f, &token)
		if err != nil {
			glog.Errorln(err)
		}

		if val, ok := token["access_token"]; ok {
			a.Bearer = val.(string)
		} else {
			err := errors.New("error, no access_token")
			glog.Errorln(err)
			return err
		}

		if val, ok := token["expires_in"]; ok {
			a.BearerExpiry = time.Now().Add((time.Duration(int64(val.(float64))) * time.Second))

		} else {
			err := errors.New("error, no expires_in")
			glog.Errorln(err)
			return err
		}

	}

	return nil
}

// NewNotification Sends a new notification to a user
func (a *AlexaMessaging) NewNotification(userID, message string) error {
	if a.ClientID == "" || a.clientSecret == "" {
		err := errors.New("error, no client ID or client Secret")
		glog.Errorln(err)
		return err
	}

	if a.Bearer == "" || a.BearerExpiry.After(time.Now()) {
		glog.Infof("need to refresh bearer\n")
		a.getBearer()
	} else {
		glog.Infof("bearer looks good\n")
	}

	var apiURL string = fmt.Sprintf("https://api.amazonalexa.com/v1/skillmessages/users/%s", userID)
	client := &http.Client{}

	m := make(map[string]interface{})
	m["expiresAfterSeconds"] = 60
	d := make(map[string]interface{})
	d["notification"] = message
	m["data"] = d

	jsonValue, err := json.Marshal(m)
	if err != nil {
		glog.Errorln(err)
		return err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Bearer))
	if err != nil {
		glog.Errorln(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorln(err)
		return err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorln(err)
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		glog.Errorln(err)
		return err
	}

	spew.Dump(resp.StatusCode)
	spew.Dump(f)

	return nil
}
