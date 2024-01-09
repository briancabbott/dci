package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/oauth2"
	"k8s.io/client-go/util/retry"
)

var (
	// defaultScopes:
	// - cloud-platform is the base scope to authenticate to GCP.
	// - userinfo.email is used to authenticate to GKE APIs with gserviceaccount
	//   email instead of numeric uniqueID.
	defaultScopes = []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/userinfo.email"}
)

// defaultCredentialsTokenProvider provides default credential tokens.
type defaultCredentialsTokenProvider struct {
	googleDefaultTokenSource func(ctx context.Context, scope ...string) (oauth2.TokenSource, error)
}

func (p *defaultCredentialsTokenProvider) token() (string, *time.Time, error) {
	var tok *oauth2.Token

	// Retries (max 4 retries with approx delay 10*ms+jitter setup) help get around occasional network glitches
	err := retry.OnError(retry.DefaultBackoff, func(err error) bool { return true }, func() error {
		ts, err := p.googleDefaultTokenSource(context.Background(), defaultScopes...)
		if err != nil {
			return fmt.Errorf("cannot construct google default token source: %w", err)
		}

		tok, err = ts.Token()
		if err != nil {
			return fmt.Errorf("cannot retrieve default token from google default token source: %w", err)
		}

		return nil
	})
	if err != nil {
		return "", nil, fmt.Errorf("getting google default token failed after multiple retries: %w", err)
	}

	return tok.AccessToken, &tok.Expiry, nil
}

func (p *defaultCredentialsTokenProvider) useCache() bool { return false }

func (p *defaultCredentialsTokenProvider) getExtraArgs() []string { return []string{} }

func (p *defaultCredentialsTokenProvider) getGcloudArgs() []string { return []string{} }
