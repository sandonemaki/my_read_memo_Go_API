package firebase

import (
	"context"

	// Firebase Admin SDK
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// FirebaseGlue
type Glue interface {
}

// NOTE: https://zenn.dev/cloud_ace/articles/firebase-auth-guide
type firebaseGlue struct {
	firebaseAuth *auth.Client
}

// NewFirebaseGlue :
// NewFirebaseGlueは、新しいFireBase Admin SDKインスタンスを初期化し、接着剤インターフェイスを返します。
// 初期化が失敗した場合、Firebase認証サービスとパニックを設定します。
// https://firebase.google.com/docs/admin/setup?hl=ja#initialize-multiple-apps
func NewFirebaseGlue() Glue {
	// Firebase Admin Appを初期化
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	// Firebase認証クライアントの取得
	firebaseAuth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return firebaseGlue{firebaseAuth}
}
