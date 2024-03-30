package token

import (
  "github.com/gogf/gf/v2/os/gtime"
  "github.com/golang-jwt/jwt/v5"
)

type PayLoad struct {
  Issuer    string           `json:"iss,omitempty"`
  Subject   string           `json:"sub,omitempty"`
  Audience  jwt.ClaimStrings `json:"aud,omitempty"`
  ExpiresAt *gtime.Time      `json:"exp,omitempty"`
  NotBefore *gtime.Time      `json:"nbf,omitempty"`
  IssuedAt  *gtime.Time      `json:"iat,omitempty"`
  ID        string           `json:"jti,omitempty"`
}

func (p PayLoad) GetExpirationTime() (*jwt.NumericDate, error) {
  if p.ExpiresAt == nil {
    return nil, nil
  }
  return jwt.NewNumericDate(p.ExpiresAt.Time), nil
}

func (p PayLoad) GetIssuedAt() (*jwt.NumericDate, error) {
  if p.IssuedAt == nil {
    return nil, nil
  }
  return jwt.NewNumericDate(p.IssuedAt.Time), nil

}

func (p PayLoad) GetNotBefore() (*jwt.NumericDate, error) {
  if p.NotBefore == nil {
    return nil, nil
  }
  return jwt.NewNumericDate(p.NotBefore.Time), nil
}

func (p PayLoad) GetIssuer() (string, error) {
  return p.Issuer, nil
}

func (p PayLoad) GetSubject() (string, error) {
  return p.Subject, nil
}

func (p PayLoad) GetAudience() (jwt.ClaimStrings, error) {
  return p.Audience, nil
}
