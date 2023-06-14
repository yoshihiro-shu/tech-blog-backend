package mock_test

// TODO mockテスト用のredis clientを使用したい
type MockRedisClient struct{}

func (m MockRedisClient) GET(key string, i interface{}) error { return nil }
func (m MockRedisClient) SET(key string, i interface{}) error { return nil }
