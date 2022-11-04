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
	"github.com/yoshihiro-shu/draft-backend/internal/cache"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/internal/model"
)

type Context struct {
	Db     *model.DBContext
	Cache  *cache.RedisContext
	Conf   config.Configs
	Logger *log.Logger
}

func NewContext(conf config.Configs) *Context {
	return &Context{
		Db:     model.New(conf),
		Cache:  cache.New(conf),
		Conf:   conf,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (c Context) DB() *pg.DB {
	return c.Db.PsqlDB
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
