package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/volatiletech/null"
)

// NOTE: アーキテクチャパターン
// https://zenn.dev/cloud_ace/articles/firebase-auth-guide

// NOTE: 認証の流れ
// https://firebase.google.com/docs/admin/setup?hl=ja

// Credential
type Credential struct {
	UID           string
	Email         string
	EmailVerified bool
	Disabled      bool
	PictureURL    null.String
	DisplayName   null.String
}

// FirebaseGlue
type Glue interface {
	GetCredFromJWT(ctx context.Context, idToken string) (cred *Credential, err error)
	DeleteUser(ctx context.Context, uid string) error
}

// Goでは新しく生成されたAppのClientを *auth.Client として扱う
type firebaseGlue struct {
	client *auth.Client
}

// NewFirebaseGlue :
// https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja#verify_id_tokens_using_the_firebase_admin_sdk
// 戻り値は interface メソッドの実装
func NewFirebaseGlue() (Glue, error) {
	// Firebase Admin Appを初期化
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}
	// Firebase Admin AppからAuthサービスを取得
	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get Firebase auth client: %w", err)
	}
	return &firebaseGlue{firebaseAuth}, nil
}

// tokenの検証
// 検証後にFirebaseのJWTトークンから必要な情報を抽出し、アプリケーション用のCredential構造体に変換
// NOTE: https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja#verify_id_tokens_using_the_firebase_admin_sdk
func (f firebaseGlue) GetCredFromJWT(ctx context.Context, idToken string) (cred *Credential, err error) {

	token, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	// FirebaseのJWTトークンのClaimsから必要な情報を抽出
	var email string
	if v, ok := token.Claims["email"]; ok {
		// NOTE: v.(string) は型アサーション
		if emailStr, ok := v.(string); ok {
			email = emailStr
		}
	}
	var emailVerified bool
	if v, ok := token.Claims["email_verified"]; ok {
		if verified, ok := v.(bool); ok {
			emailVerified = verified
		}
	}

	var picture, name null.String
	if v, ok := token.Claims["picture"]; ok {
		if pictureStr, ok := v.(string); ok {
			picture = null.StringFrom(pictureStr)
		}
	}
	if v, ok := token.Claims["name"]; ok {
		if nameStr, ok := v.(string); ok {
			name = null.StringFrom(nameStr)
		}
	}

	if len(name.String) == 0 {
		name = null.StringFrom(email)
	}

	cred = &Credential{
		UID:           token.UID,
		Email:         email,
		EmailVerified: emailVerified,
		PictureURL:    picture,
		Disabled:      false,
		DisplayName:   name,
	}
	return cred, nil
}

func (f firebaseGlue) DeleteUser(ctx context.Context, uid string) error {
	if _, err := f.client.GetUser(ctx, uid); err != nil {
		if auth.IsUserNotFound(err) {
			return nil
		} else {
			return err
		}
	}

	if err := f.client.DeleteUser(ctx, uid); err != nil {
		return err
	}
	return nil
}
