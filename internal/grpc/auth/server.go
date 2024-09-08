package auth

import authv1 `github.com/dugtriol/auth-protos/gen/go/auth`

type serverAPI struct {
	authv1.UnimplementedAuthServer
}
