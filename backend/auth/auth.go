package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"wetalk/monorepo/backend/api/identity"
	"wetalk/monorepo/backend/ent"
	"wetalk/monorepo/backend/ent/user"
)

type Server struct {
	identity.UnimplementedAuthServer
	Client *ent.Client
}

func (s *Server) Register(ctx context.Context, req *identity.RegisterRequest) (*identity.AuthReply, error) {

	if req.Password != req.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, "Passwords miss match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hashUser(req.Username, req.Password)), bcrypt.DefaultCost)

	u, err := s.Client.User.Create().
		SetUsername(req.Username).
		SetSafeID(uuid.New().String()).
		SetPassword(string(hashedPassword)).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return &identity.AuthReply{UserId: u.SafeID, Token: s.createToken(u.SafeID)}, nil
}
func (s *Server) Login(ctx context.Context, req *identity.LoginRequest) (resp *identity.AuthReply, err error) {

	u, err := s.Client.User.Query().
		Where(user.Username(req.Username), user.DeletedAtIsNil()).
		Only(ctx)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(hashUser(req.Username, req.Password)))

	if err != nil {
		return resp, status.Error(codes.Unauthenticated, "incorrect email or password")
	}

	return &identity.AuthReply{UserId: u.SafeID, Token: s.createToken(u.SafeID)}, nil
}
func (s *Server) VerifyToken(ctx context.Context, req *identity.VerifyTokenRequest) (*identity.VerifyTokenReply, error) {

	resp := &identity.VerifyTokenReply{Valid: false}

	token, _ := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		resp.UserId = claims["id"].(string)
		resp.Valid = true
		return resp, nil
	}

	return resp, nil
}

func (s *Server) createToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "auth-app",
		"id":  id,
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte("secret"))
	return jwtToken
}

func hashUser(email string, password string) string {
	return fmt.Sprintf("%s:%s", email, password)
}
