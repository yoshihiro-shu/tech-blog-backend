package request

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/auth"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/model"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type Context struct {
	db     *model.DBContext
	cache  cache.RedisClient
	Conf   config.Configs
	Logger *log.Logger
}

func NewContext(conf config.Configs) *Context {
	return &Context{
		db:     model.New(conf),
		cache:  cache.New(conf.CacheRedis),
		Conf:   conf,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (c Context) MasterDB() *pg.DB {
	return c.db.Master()
}

func (c Context) RepricaDB() *pg.DB {
	return c.db.Reprica()
}

func (c Context) Cache() cache.RedisClient {
	return c.cache
}

func (c Context) UnmarshalFromRequest(r *http.Request, i interface{}) error {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, i)
}

func (c Context) GetAuthUserID(ctx context.Context) interface{} {
	return ctx.Value(auth.UserKey)
}
