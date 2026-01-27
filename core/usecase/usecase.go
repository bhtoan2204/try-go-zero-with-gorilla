package usecase

type Usecase interface {
	AuthUsecase() AuthUsecase
	RoomUsecase() RoomUsecase
}

type usecase struct {
	authUsecase AuthUsecase
	roomUsecase RoomUsecase
}
