//go:build wireinject
// +build wireinject

package injection

func InitializeAuthController() *userinterface.AuthController {
	wire.Build(
		infrastructure.GetDB,
		wire.Bind(
			new(application.Transaction),
			new(*infrastructure.DB),
		),
		infrastructure.NewAuthRepository,
		wire.Bind(
			new(domain.AuthRepository),
			new(*infrastructure.AuthRepository),
		),
		application.NewAuthService,
		userinterface.NewAuthController,
	)

	return nil
}

func InitializeUserController() *userinterface.UserController {
	wire.Build(
		infrastructure.GetDB,
		wire.Bind(
			new(application.Transaction),
			new(*infrastructure.DB),
		),
		infrastructure.NewUserRepository,
		wire.Bind(
			new(domain.UserRepository),
			new(*infrastructure.UserRepository),
		),
		application.NewUserService,
		userinterface.NewUserController,
	)

	return nil
}

func InitializeBoardController() *userinterface.BoardController {
	wire.Build(
		infrastructure.GetDB,
		wire.Bind(
			new(application.Transaction),
			new(*infrastructure.DB),
		),
		infrastructure.NewBoardRepository,
		wire.Bind(
			new(domain.BoardRepository),
			new(*infrastructure.BoardRepository),
		),
		application.NewBoardService,
		userinterface.NewBoardController,
	)

	return nil
}
