package firebase

import (
	"context"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/auth/domain/model"
)

// FirebaseAuthGlue - Firebase認証の接着剤インターフェース
type FirebaseAuthGlue interface {
	CheckLoginJWT(ctx context.Context, idToken string) (cred *model.Credential, err error)
	DeleteAccount(ctx context.Context, uid string) error
}