package repositorys

import (
	chachrepositorys "auth_api/infrastructure/repositorys/cache_repositorys"
	databaserepositorys "auth_api/infrastructure/repositorys/database_repositorys"
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	UserRepository      databaserepositorys.IUserRepository
	UserCacheRepository chachrepositorys.UserCacheRepository

	//diğer tablalrın repoları eklendekçe buraya gelecek
}

func NewUnitOfWork(rmd *redis.Client, db *gorm.DB) *UnitOfWork {
	ctx := context.Background()
	return &UnitOfWork{
		UserRepository:      databaserepositorys.NewUserRepository(db),
		UserCacheRepository: *chachrepositorys.NewUserCacheRepository(rmd, &ctx),
	}
}
