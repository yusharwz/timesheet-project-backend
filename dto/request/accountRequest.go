package request

import "mime/multipart"

type (
	ActivateAccountRequest struct {
		Email    string
		Password string
	}

	UploadImagesRequest struct {
		UserID         string         `json:"userId"`
		SignatureImage multipart.File `json:"signature"`
	}

	LoginAccountRequest struct {
		Email    string `json:"email" binding:"email"`
		Password string `json:"password" binding:"required"`
	}

	EditAccountRequest struct {
		UserID string `json:"-"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Phone  string `json:"phone"`
	}

	ChangePasswordRequest struct {
		NewPassword string `json:"newPassword" binding:"password"`
	}

	UploadImagesRequest struct {
		UserID         string         `json:"user_id"`
		SignatureImage multipart.File `json:"signature"`
	}
)
