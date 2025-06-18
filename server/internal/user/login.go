package user

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var signingKey = os.Getenv("JWT_KEY")

func checkUser(email string, password string) bool {
	user, err := GetUserByEmail(email)
	if err != nil {
		return false
	}

	return checkPassword(password, user.Password)
}

func setTokens(c *gin.Context, email string, id snowflake.ID) {
	verifTkn := generateToken(email, id, 60*time.Minute, false)
	refreshTkn := generateToken(email, id, 60*time.Hour, true)

	domain := getDomain(c)

	c.SetCookie("JWT_TOKEN", verifTkn, int(60*time.Minute), "/", domain, true, true)
	c.SetCookie("JWT_REFRESH", refreshTkn, int(60*time.Hour), "/", domain, true, true)
}

func generateToken(email string, id snowflake.ID, expiryTime time.Duration, isRefreshToken bool) string {
	jwtType := "validation"
	if isRefreshToken {
		jwtType = "refresh"
	}

	claims := Claims{
		Email: email,
		Type:  jwtType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryTime)),
			Subject:   id.String(),
			Issuer:    "server",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signingKey))

	if err != nil {
		log.Panic(err)
		return ""
	}

	return token
}

func getDomain(c *gin.Context) string {
	parsedUrl, err := url.Parse(c.Request.Header.Get("Origin"))
	if err != nil {
		log.Panic(err)
	}

	return parsedUrl.Hostname()
}

func checkPassword(password string, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	return err == nil
}
