package injection

// func InitializeAuthController() *userinterface.AuthController {
// 	wire.Build(
// 		infrastructure.GetDB,
// 		wire.Bind(
// 			new(application.Transaction),
// 			new(*infrastructure.DB),
// 		),
// 		infrastructure.NewAuthRepository,
// 		wire.Bind(
// 			new(domain.AuthRepository),
// 			new(*infrastructure.AuthRepository),
// 		),
// 		application.NewAuthService,
// 		userinterface.NewAuthController,
// 	)

// 	return nil
// }
