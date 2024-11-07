package chachrepositorys

import (
	usercachemodels "auth_api/infrastructure/dto_models/user_cache_models"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type UserCacheRepository struct {
	rdm *redis.Client
	ctx *context.Context
}

func NewUserCacheRepository(rmd *redis.Client, ctx *context.Context) *UserCacheRepository {

	return &UserCacheRepository{
		rdm: rmd,
		ctx: ctx,
	}
}

func (ucr UserCacheRepository) AddUser(users usercachemodels.UserCacheModel) error {
	json, _ := json.Marshal(users)
	return ucr.rdm.Set(*ucr.ctx, "users", json, 0).Err()
}
