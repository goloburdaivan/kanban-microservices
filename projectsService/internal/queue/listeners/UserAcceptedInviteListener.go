package listeners

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"projectsService/internal/database/models"
	"projectsService/internal/enums"
	"projectsService/internal/repository"
)

type UserAcceptedInviteListener struct {
	repository *repository.ProjectUserRepository
}

func (u *UserAcceptedInviteListener) Listen(msg amqp.Delivery) {
	var invitation models.ProjectUser
	invitation.Role = enums.RoleMember

	if err := json.Unmarshal(msg.Body, &invitation); err != nil {
		log.Println("Unmarshal error:", err)
	}

	_, err := u.repository.Create(&invitation)

	if err != nil {
		log.Println("CreateRecord error:", err)
	}
}

func NewUserAcceptedInviteListener(repository *repository.ProjectUserRepository) *UserAcceptedInviteListener {
	return &UserAcceptedInviteListener{repository: repository}
}
