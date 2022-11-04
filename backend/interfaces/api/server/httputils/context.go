package httputils

// type Context struct {
// 	writer  http.ResponseWriter
// 	request *http.Request
// 	db      *model.DBContext
// 	cache   *cache.RedisContext
// 	conf    config.Configs
// 	logger  *log.Logger
// 	h       HandlerFunc
// }

// // HandlerFunc defines a function to serve HTTP requests.
// type HandlerFunc func(c Context) error

// func NewContext(conf config.Configs) *Context {
// 	return &Context{
// 		db:     model.New(conf),
// 		cache:  cache.New(conf),
// 		conf:   conf,
// 		logger: log.New(os.Stdout, "", log.LstdFlags),
// 	}
// }

// func (c *Context) Reset(w http.ResponseWriter, r *http.Request) {
// 	c.writer = w
// 	c.request = r
// }

// func WrapHandler(h http.Handler) HandlerFunc {
// 	return func(c Context) error {
// 		h.ServeHTTP(c.writer, c.request)
// 		return nil
// 	}
// }
