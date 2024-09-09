package services

import (
	"fmt"
	"invitationService/internal/queue"
	"invitationService/internal/repository"
)

type InvitationAcceptService struct {
	repository *repository.InvitationRepository
	sender     *queue.RabbitMQSender
}

func (i *InvitationAcceptService) Accept(token string) error {
	invitation, err := i.repository.GetByToken(token)
	if err != nil {
		return fmt.Errorf("Token is not valid: %w", err)
	}

	i.repository.UpdateStatus(invitation, "active")
	i.sender.Send("user_accept", invitation)

	return nil
}

func NewInvitationAcceptService(
	repository *repository.InvitationRepository,
	sender *queue.RabbitMQSender,
) *InvitationAcceptService {
	return &InvitationAcceptService{repository: repository, sender: sender}
}
