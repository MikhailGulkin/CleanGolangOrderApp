package api

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/application/commands"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewCustomerHandler(service *application.CustomerServices) *Handler {
	return &Handler{
		cs: service,
	}
}

type Handler struct {
	cs *application.CustomerServices
}

// UploadNewAvatar godoc
// @Summary Upload new avatar for customer
// @Description Upload new avatar for customer
// @Tags Customer
// @Accept			multipart/form-data
// @Produce json
// @Param id path string true "Customer ID"
// @Param  file	formData	file			true "Avatar"
// @Success 200 {object} commands.UploadAvatarDTO
// @Failure 400 {object} Error
// @Router /upload-avatar/{id}/ [post]
func (h *Handler) UploadNewAvatar(context *fiber.Ctx) error {
	id := context.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(context.Body()) == 0 {
		return context.Status(fiber.StatusBadRequest).JSON(
			&Error{Error: "Body can't be empty"},
		)
	}
	data, err := h.cs.Commands.UploadCustomerAvatar.Handle(commands.UploadAvatarCommand{
		CustomerID: uid,
		Avatar:     context.Body(),
	})
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(
			&Error{Error: err.Error()},
		)
	}
	return context.Status(fiber.StatusOK).JSON(data)
}
