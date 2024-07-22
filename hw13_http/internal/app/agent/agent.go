package agent

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/agent"
	"github.com/Stern-Ritter/go/hw13_http/internal/model"
	service "github.com/Stern-Ritter/go/hw13_http/internal/service/agent"
)

func Run(cfg *agent.Config, lg *slog.Logger) error {
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
		agent.Logger.Error("Error marshalling request body", "url", url, "payload", string(payload), "error", err)
		return
	}

	resp, err := agent.SendRequest(url, http.MethodPost, map[string]string{"Content-Type": "application/json"}, payload)
	if err != nil {
		agent.Logger.Error("Error sending request to agent", "url", url, "payload", string(payload), "error", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		agent.Logger.Error("Error reading response body", "url", url, "payload", string(payload), "error", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		agent.Logger.Info("Successfully sent request to server", "url", url, "payload", string(payload), "statusCode",
			resp.StatusCode)
		return
	}

	savedUser := model.User{}
	err = json.Unmarshal(data, &savedUser)
	if err != nil {
		agent.Logger.Error("Error unmarshalling response body", "url", url, "payload", string(payload), "error", err)
		return
	}

	agent.Logger.Info("Successfully sent request to server", "url", url, "payload", user, "statusCode",
		resp.StatusCode, "responseBody", savedUser)
}

func sendGetUserRequest(agent *service.Agent, url string) {
	resp, err := agent.SendRequest(url, http.MethodGet, make(map[string]string), make([]byte, 0))
	if err != nil {
		agent.Logger.Error("Error sending request to agent", "url", url, "error", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		agent.Logger.Error("Error reading response body", "url", url, "error", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		agent.Logger.Info("Successfully sent request to server", "url", url, "statusCode", resp.StatusCode)
		return
	}

	savedUser := model.User{}
	err = json.Unmarshal(data, &savedUser)
	if err != nil {
		agent.Logger.Error("Error unmarshalling response body", "url", url, "error", err)
		return
	}

	agent.Logger.Info("Successfully sent request to server", "url", url, "statusCode", resp.StatusCode,
		"responseBody", savedUser)
}
