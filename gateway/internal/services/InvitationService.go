package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/internal/http/requests"
	"net/http"
)

type InvitationService struct{}

func (i *InvitationService) Invite(invite *requests.InviteRequest) error {
	jsonData, err := json.Marshal(invite)
	if err != nil {
		return fmt.Errorf("Invalid request: %v", err)
	}

	resp, err := http.Post(
		"http://localhost:8083/invite",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return fmt.Errorf("Invite service is not available. Try later: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send the invite to user: %v", resp.Status)
	}

	return nil
}

func NewInvitationService() *InvitationService {
	return &InvitationService{}
}
