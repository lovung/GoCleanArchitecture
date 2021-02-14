package jwtutil

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateToken(t *testing.T) {
	InitJWTSession("test_secret", time.Hour*12, time.Hour*24)
	dummyIssueTime, _ := time.Parse(time.RFC3339, "2020-01-01T01:00:05Z")
	tests := []struct {
		name       string
		arg        AuthClaims
		wantAToken string
		wantRToken string
	}{
		{
			name: "SuccsssCase: Correctly put claims",
			arg: AuthClaims{
				UserID:    1,
				IssueTime: dummyIssueTime,
			},
			wantAToken: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.` +
				`eyJleHAiOjE1Nzc4ODM2MDUsImlhdCI6MTU3Nzg0MDQwNSwidXNlcl9pZCI6MX0.` +
				`_OuEn9wfc91EkxpMBu0OGwdwKBaP4jGNdqt0-TTwgbU`,
			wantRToken: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.` +
				`eyJleHAiOjE1Nzc5MjY4MDUsImlhdCI6MTU3Nzg0MDQwNSwidXNlcl9pZCI6MX0.` +
				`eFZh2LXnYE5E4nWrmb7wLFyrY32I-U7r9ysIN4NSDIg`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAToken, gotRToken, _ := GenerateToken(tt.arg)
			if gotAToken != tt.wantAToken {
				t.Errorf("GenerateToken aToken = %v, want = %v", gotAToken, tt.wantAToken)
			}
			if gotRToken != tt.wantRToken {
				t.Errorf("GenerateToken rToken = %v, want = %v", gotRToken, tt.wantRToken)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	InitJWTSession("test_secret", time.Hour*12, time.Hour*24)
	validToken, _, _ := GenerateToken(
		AuthClaims{
			UserID:    1,
			IssueTime: time.Now(),
		},
	)

	dummyIssueTime, _ := time.Parse(time.RFC3339, "2020-01-01T01:00:05Z")
	invalidToken, _, _ := GenerateToken(
		AuthClaims{
			UserID:    1,
			IssueTime: dummyIssueTime,
		},
	)

	tests := []struct {
		name    string
		arg     string
		want    AuthClaims
		wantErr bool
	}{
		{
			name: "SuccessCase: put the token issued right now",
			arg:  validToken,
			want: AuthClaims{
				UserID: 1,
			},
			wantErr: false,
		},
		{
			name:    "FailureCase: put the token issued in the past",
			arg:     invalidToken,
			want:    AuthClaims{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := VerifyToken(tt.arg)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("VerifyToken: clamis = %v, want = %v", got, tt.want)
			}
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("VerifyToken: err = %v, want = %v", gotErr, tt.wantErr)
			}
		})
	}
}
