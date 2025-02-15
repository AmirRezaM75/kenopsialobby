package kenopsialobby

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LobbyRepository struct {
	baseUrl string
	token   string
}

func NewLobbyRepository(baseUrl, token string) LobbyRepository {
	return LobbyRepository{baseUrl: baseUrl, token: token}
}

func (lobbyService LobbyRepository) FindById(id string) (*Lobby, error) {
	url := fmt.Sprintf("%s/lobbies/%s", lobbyService.baseUrl, id)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("could not create new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Token", lobbyService.token)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("response status %d is not ok: %w", response.StatusCode, err)
	}

	var lobby Lobby

	err = json.NewDecoder(response.Body).Decode(&lobby)

	_ = response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("could not decode response body: %w", err)
	}

	return &lobby, nil
}
