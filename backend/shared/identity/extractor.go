package identity

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc/metadata"
)

const (
	// UserIDHeader ...
	UserIDHeader = "x-wetalk-userid"
	// SecWebSocketProtocol this headers seems not for authentication
	// https://tools.ietf.org/html/rfc6455#section-4
	// But we don't to put token in query path or first message
	SecWebSocketProtocol = "Sec-WebSocket-Protocol"
)

var (
	// ErrNoMetadata ...
	ErrNoMetadata = errors.New("no metadata from context")
	// ErrEmptyMetadata ...
	ErrEmptyMetadata = errors.New("empty metadata")
)

// Extractor put it to an interface for easir when we want to mock
type Extractor interface {
	GetUserID(ctx context.Context) (string, error)
	GetWebSocketCredentials(r *http.Request) (string, string)
}

// NewExtractor ...
func NewExtractor() Extractor {
	return &extractorImpl{}
}

type extractorImpl struct {
}

// GetUserID ...
func (i *extractorImpl) GetUserID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrNoMetadata
	}

	mdUserID := md.Get(UserIDHeader)
	if len(mdUserID) < 1 {
		return "", ErrEmptyMetadata
	}
	return mdUserID[0], nil
}

// GetWebSocketCredentials ...
func (i *extractorImpl) GetWebSocketCredentials(r *http.Request) (string, string) {
	return r.Header.Get(UserIDHeader), r.Header.Get(SecWebSocketProtocol)
}
