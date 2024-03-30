package token

import (
  "context"

  "github.com/gogf/gf/v2/errors/gerror"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/gogf/gf/v2/net/ghttp"
  "github.com/gogf/gf/v2/text/gstr"
  "github.com/golang-jwt/jwt/v5"
)

var (
  errorLogin      = gerror.New("登录身份已失效，请重新登录！")
  errorMultiLogin = gerror.New("账号已在其他地方登录，如非本人操作请及时修改登录密码！")
)

type JwtOption struct {
  SecretKey     string
  Key           string
  KeyValPrefix  string
  SigningMethod jwt.SigningMethod
}

type Jwt struct {
  *JwtOption
}

func New(funcs ...OptionFunc) (jwtUtil *Jwt) {
  option := &JwtOption{
    SecretKey:     "",
    Key:           "Authorization",
    KeyValPrefix:  "Bearer ",
    SigningMethod: jwt.SigningMethodHS256,
  }
  for _, fun := range funcs {
    fun(option)
  }
  return &Jwt{option}
}

func (j *Jwt) Sign(ctx context.Context, claims jwt.Claims) (string, error) {
  return jwt.NewWithClaims(j.SigningMethod, claims).SignedString([]byte(j.SecretKey))
}

func (j *Jwt) Parse(ctx context.Context, token string, claims jwt.Claims) (*jwt.Token, error) {
  parseWithClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte(j.SecretKey), nil
  })
  if err != nil {
    g.Log().Debugf(ctx, "parseToken err:%+v", err)
    return nil, err
  }
  if !parseWithClaims.Valid {
    return nil, errorLogin
  }
  return parseWithClaims, nil
}

func (j *Jwt) MustSign(ctx context.Context, claims jwt.Claims) string {
  sign, err := j.Sign(ctx, claims)
  if err != nil {
    g.Log().Warning(ctx, "签名错误")
  }
  return sign
}

// header > query form body
func (j *Jwt) GetToken(r *ghttp.Request) string {
  // 默认从请求头获取
  var authorization = r.Header.Get(j.Key)

  // 如果请求头不存在则从get参数获取
  if authorization == "" {
    return r.Get(gstr.ToLower(j.Key)).String()
  }
  if j.KeyValPrefix != "" {
    return gstr.Replace(authorization, "Bearer ", "")
  }
  return authorization
}

type OptionFunc func(*JwtOption)

func WithHeaderKey(v string) OptionFunc {
  return func(option *JwtOption) {
    option.Key = v
  }
}
func WithKeyValPrefix(v string) OptionFunc {
  return func(option *JwtOption) {
    option.KeyValPrefix = v
  }
}
func WithSecretKey(v string) OptionFunc {
  return func(option *JwtOption) {
    option.SecretKey = v
  }
}

func WithSigningMethod(v jwt.SigningMethod) OptionFunc {
  return func(option *JwtOption) {
    option.SigningMethod = v
  }
}
