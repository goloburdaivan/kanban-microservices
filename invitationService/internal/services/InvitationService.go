package services

import (
	"fmt"
	"invitationService/internal/dto"
	"invitationService/internal/http/requests"
	"invitationService/internal/models"
	"invitationService/internal/queue"
	"invitationService/internal/repository"
)

type InvitationService struct {
	repository     *repository.InvitationRepository
	userRepository repository.UserGetter
	checker        ProjectExistenceChecker
	sender         *queue.RabbitMQSender
	emailService   *InvitationEmailService
}

func (i *InvitationService) Invite(request *requests.InvitationRequest) error {
	user, err := i.userRepository.GetUser(uint(request.UserID))

	if err != nil {
		return fmt.Errorf("Invalid user ID: %v", err)
	}

	err = i.checker.Get(request.ProjectID)

	if err != nil {
		return fmt.Errorf("Invalid project ID: %v", err)
	}

	invite := &models.Invitation{
		UserID:    request.UserID,
		ProjectID: uint(request.ProjectID),
	}

	record, err := i.repository.Create(invite)

	if err != nil {
		return fmt.Errorf("Failed to send invite to user: %v", err)
	}

	go i.emailService.Send(&dto.InvitationDTO{
		User:  user,
		Token: record.Token,
	})

	return nil
}

func (i *InvitationService) Accept(token string) error {
	invitation, err := i.repository.GetByToken(token)
	if err != nil {
		return fmt.Errorf("Token is not valid: %w", err)
	}

	i.repository.UpdateStatus(invitation, "active")
	i.sender.Send("user_accept", invitation)

	return nil
}

func NewInvitationService(
	repository *repository.InvitationRepository,
	sender *queue.RabbitMQSender,
	userRepository repository.UserGetter,
	service *InvitationEmailService,
	checker ProjectExistenceChecker,
) *InvitationService {
	return &InvitationService{
		repository:     repository,
		sender:         sender,
		userRepository: userRepository,
		emailService:   service,
		checker:        checker,
	}
}
