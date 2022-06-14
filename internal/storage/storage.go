package storage

type IStorage interface {
	Products() IProductRepository
	TokenDetails() ITokenRepository
	Users() IUserRepository
}
