package firebase

import (
	"context"

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
	GetCredFromJWT(idToken string) (cred *Credential, err error)
	DeleteUser(ctx context.Context, uid string) error
}

// Goでは新しく生成されたAppのClientを *auth.Client として扱う
type firebaseGlue struct {
	client *auth.Client
}

// NewFirebaseGlue :
// https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja#verify_id_tokens_using_the_firebase_admin_sdk
// 戻り値は interface メソッドの実装
func NewFirebaseGlue() Glue {
	// Firebase Admin Appを初期化
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	// Firebase Admin AppからAuthサービスを取得
	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return firebaseGlue{firebaseAuth}
}

// tokenの検証
// 検証後にFirebaseのJWTトークンから必要な情報を抽出し、アプリケーション用のCredential構造体に変換
// NOTE: https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja#verify_id_tokens_using_the_firebase_admin_sdk
func (f firebaseGlue) GetCredFromJWT(idToken string) (cred *Credential, err error) {

	token, err := f.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	// FirebaseのJWTトークンのClaimsから必要な情報を抽出
	var email string
	if v, ok := token.Claims["email"]; ok {
		// NOTE: v.(string) は型アサーション
		email = v.(string)
	}

	var picture, name null.String
	if v, ok := token.Claims["picture"]; ok {
		picture = null.StringFrom(v.(string))
	}
	if v, ok := token.Claims["name"]; ok {
		name = null.StringFrom(v.(string))
	}

	if len(name.String) == 0 {
		name = null.StringFrom(email)
	}

	cred = &Credential{
		UID:           token.UID,
		Email:         email,
		EmailVerified: true,
		PictureURL:    picture,
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
