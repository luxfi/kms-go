package kms

import (
	api "github.com/luxfi/kms-go/packages/api/auth"
	"github.com/luxfi/kms-go/packages/errors"
	"github.com/luxfi/kms-go/packages/models"
)

type OciAuthLoginOptions struct {
	IdentityID  string
	PrivateKey  string
	Fingerprint string
	UserID      string
	TenancyID   string
	Region      string
	Passphrase  *string
}

type MachineIdentityCredential = api.MachineIdentityAuthLoginResponse

type Secret = models.Secret
type SecretImport = models.SecretImport

type APIError = errors.APIError
type RequestError = errors.RequestError
