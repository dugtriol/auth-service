package auth

import (
	`context`
	`errors`
	`log/slog`
	`time`

	`github.com/dugtriol/auth-service/internal/domain/models`
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver:    userSaver,
		usrProvider: userProvider,
		log:         log,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

func SaveUser(
	ctx context.Context,
	email string,
	passHash []byte,
) (uid int64, err error) {
	panic("implement me")
}

func User(ctx context.Context, email string) (models.User, error) { panic("implement me") }

func IsAdmin(ctx context.Context, userID int64) (bool, error) { panic("implement me") }

func App(ctx context.Context, appID int) (models.App, error) { panic("implement me") }
