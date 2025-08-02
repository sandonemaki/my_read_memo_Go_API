package model

import "github.com/volatiletech/null"

// Credential
type Credential struct {
	UID           string
	Email         string
	EmailVerified bool
	Disabled      bool
	PictureURL    null.String
	DisplayName   null.String
}

// 追加：GetSafeDisplayName - 安全に表示名を取得する
//
// 表示名が設定されていない場合は、メールアドレスを代わりに使う。
// これは「何かしらの名前を必ず返す」ための安全措置。
//
// 例：
// - DisplayNameがある場合: "山田太郎" を返す
// - DisplayNameがない場合: "yamada@example.com" を返す
func (c *Credential) GetSafeDisplayName() string {
	// null.Stringは「値があるかもしれないし、ないかもしれない」文字列型
	if c.DisplayName.Valid && c.DisplayName.String != "" {
		return c.DisplayName.String
	}
	// フォールバック: 表示名がない場合はメールアドレスを使用
	return c.Email
}
