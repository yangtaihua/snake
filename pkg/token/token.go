// 主要是配合jwt来生成用户登录token

package token

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/1024casts/snake/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
)

// Payload is the data of the JSON web token.
type Payload struct {
	UserID   uint64
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the payloads if the token was valid.
func Parse(tokenString string, secret string) (*Payload, error) {
	payloads := &Payload{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return payloads, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloads.UserID = uint64(claims["user_id"].(float64))
		payloads.Username = claims["username"].(string)
		return payloads, nil

		// Other errors.
	} else {
		return payloads, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Payload, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := config.Conf.App.JwtSecret

	if len(header) == 0 {
		return &Payload{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	_, err := fmt.Sscanf(header, "Bearer %s", &t)
	if err != nil {
		fmt.Printf("fmt.Sscanf err: %+v", err)
	}
	return Parse(t, secret)
}

// Sign signs the payload with the specified secret.
func Sign(ctx context.Context, payload map[string]interface{}, secret string, timeout int64) (tokenString string, err error) {
	// The token content.
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["nbf"] = now
	claims["iat"] = now
	claims["exp"] = now + timeout

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
