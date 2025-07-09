package model

import (
	"database/sql"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		ulid        string
		uid         string
		displayName string
		deletedAt   sql.Null[time.Time]
		want        *User
	}{
		{
			name:        "正常なユーザー作成",
			ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			uid:         "firebase_uid_123",
			displayName: "テストユーザー",
			deletedAt:   sql.Null[time.Time]{},
			want: &User{
				Ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
				UID:         "firebase_uid_123",
				DisplayName: "テストユーザー",
				DeletedAt:   sql.Null[time.Time]{},
			},
		},
		{
			name:        "削除済みユーザー作成",
			ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			uid:         "firebase_uid_456",
			displayName: "削除済みユーザー",
			deletedAt:   sql.Null[time.Time]{Valid: true, V: time.Now()},
			want: &User{
				Ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
				UID:         "firebase_uid_456",
				DisplayName: "削除済みユーザー",
				DeletedAt:   sql.Null[time.Time]{Valid: true, V: time.Now()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.ulid, tt.uid, tt.displayName, tt.deletedAt)

			if got.Ulid != tt.want.Ulid {
				t.Errorf("NewUser().Ulid = %v, want %v", got.Ulid, tt.want.Ulid)
			}
			if got.UID != tt.want.UID {
				t.Errorf("NewUser().UID = %v, want %v", got.UID, tt.want.UID)
			}
			if got.DisplayName != tt.want.DisplayName {
				t.Errorf("NewUser().DisplayName = %v, want %v", got.DisplayName, tt.want.DisplayName)
			}
			if got.DeletedAt.Valid != tt.want.DeletedAt.Valid {
				t.Errorf("NewUser().DeletedAt.Valid = %v, want %v", got.DeletedAt.Valid, tt.want.DeletedAt.Valid)
			}
		})
	}
}

// ユーザーのビジネスロジックメソッドをテスト
func TestUser_IsDeleted(t *testing.T) {
	tests := []struct {
		name      string
		deletedAt sql.Null[time.Time]
		want      bool
	}{
		{
			name:      "削除されていないユーザー",
			deletedAt: sql.Null[time.Time]{},
			want:      false,
		},
		{
			name:      "削除されたユーザー",
			deletedAt: sql.Null[time.Time]{Valid: true, V: time.Now()},
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Ulid:        "test_ulid",
				UID:         "test_uid",
				DisplayName: "Test User",
				DeletedAt:   tt.deletedAt,
			}

			if got := user.IsDeleted(); got != tt.want {
				t.Errorf("User.IsDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
