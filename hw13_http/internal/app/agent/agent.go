package agent

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/agent"
	"github.com/Stern-Ritter/go/hw13_http/internal/model"
	service "github.com/Stern-Ritter/go/hw13_http/internal/service/agent"
	"github.com/sirupsen/logrus"
)

func Run(cfg *agent.Config, lg *logrus.Logger) error {
	c := &http.Client{}
	a := service.NewAgent(c, cfg, lg)

	sendPostUserRequest(a, strings.Join([]string{cfg.ServerURL, cfg.ResourceEndpoint}, "/"), model.User{
		Name:     "first",
		Email:    "first@example.com",
		Password: "password",
	})

	sendPostUserRequest(a, strings.Join([]string{cfg.ServerURL, cfg.ResourceEndpoint}, "/"), model.User{
		Name:     "second",
		Email:    "second@example.com",
		Password: "password",
	})

	sendGetUserRequest(a, strings.Join([]string{cfg.ServerURL, cfg.ResourceEndpoint, "1"}, "/"))
	sendGetUserRequest(a, strings.Join([]string{cfg.ServerURL, cfg.ResourceEndpoint, "2"}, "/"))
	sendGetUserRequest(a, strings.Join([]string{cfg.ServerURL, cfg.ResourceEndpoint, "3"}, "/"))

	return nil
}

func sendPostUserRequest(agent *service.Agent, url string, user model.User) {
	payload, err := json.Marshal(user)
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "payload": string(payload), "error": err}).
			Error("Error marshalling request body")
		return
	}

	resp, err := agent.SendRequest(url, http.MethodPost, map[string]string{"Content-Type": "application/json"}, payload)
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "payload": string(payload), "error": err}).
			Error("Error sending request to agent")
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "payload": string(payload), "error": err}).
			Error("Error reading response body")
		return
	}

	if resp.StatusCode != http.StatusOK {
		agent.Logger.WithFields(logrus.Fields{"url": url, "payload": string(payload), "statusCode": resp.StatusCode}).
			Info("Successfully sent request to server", "url")
		return
	}

	savedUser := model.User{}
	err = json.Unmarshal(data, &savedUser)
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "payload": string(payload), "error": err}).
			Error("Error unmarshalling response body")
		return
	}

	agent.Logger.WithFields(logrus.Fields{
		"url": url, "payload": user, "statusCode": resp.StatusCode,
		"responseBody": savedUser,
	}).
		Info("Successfully sent request to server")
}

func sendGetUserRequest(agent *service.Agent, url string) {
	resp, err := agent.SendRequest(url, http.MethodGet, make(map[string]string), make([]byte, 0))
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "error": err}).
			Error("Error sending request to agent")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "error": err}).
			Error("Error reading response body")
		return
	}

	if resp.StatusCode != http.StatusOK {
		agent.Logger.WithFields(logrus.Fields{"url": url, "statusCode": resp.StatusCode}).
			Info("Successfully sent request to server")
		return
	}

	savedUser := model.User{}
	err = json.Unmarshal(data, &savedUser)
	if err != nil {
		agent.Logger.WithFields(logrus.Fields{"url": url, "error": err}).
			Error("Error unmarshalling response body", "url")
		return
	}

	agent.Logger.WithFields(logrus.Fields{"url": url, "statusCode": resp.StatusCode, "responseBody": savedUser}).
		Info("Successfully sent request to server")
}
