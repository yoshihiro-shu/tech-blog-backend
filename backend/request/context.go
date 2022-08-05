package request

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/yoshi429/test/auth"
	"github.com/yoshi429/test/cache"
	"github.com/yoshi429/test/config"
	"github.com/yoshi429/test/model"
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
