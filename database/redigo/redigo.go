package redigo

import (
	"github.com/gomodule/redigo/redis"
)

// SubscribeFunc 订阅回调函数
type SubscribeFunc func(c redis.PubSubConn) (err error)

// ExecuteFunc 普通回调函数
type ExecuteFunc func(c redis.Conn) (res interface{}, err error)

type Client struct{ mode ModeInterface }

func New(mode ModeInterface) *Client {
	return &Client{mode: mode}
}

// Mode 当前客户端使用模式
// alone 单机或者代理入口模式
// cluster Redis-Cluster集群模式
// sentinel Redis-Sentinel哨兵模式
func (c *Client) Mode() string {
	return c.mode.String()
}

func (c *Client) Close() error {
	return c.mode.Close()
}

func (c *Client) Execute(fn ExecuteFunc) (res interface{}, err error) {
	conn := c.mode.GetConn()
	defer conn.Close()
	if res, err = fn(conn); err != nil {
		if _, ok := err.(redis.Error); ok {
			return
		} else if newConn, newErr := c.mode.NewConn(); newErr != nil {
			return
		} else {
			defer newConn.Close()
			res, err = fn(newConn)
		}
	}
	return
}

func (c *Client) MustExec(fun ExecuteFunc) interface{} {
	res, err := c.Execute(fun)
	if err != nil {
		panic(err)
	}
	return res
}

func (c *Client) Subscribe(fn SubscribeFunc) error {
	conn, err := c.mode.NewConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(redis.PubSubConn{Conn: conn})
}

func (c *Client) Int(fn ExecuteFunc) (int, error) {
	return redis.Int(c.Execute(fn))
}

func (c *Client) Ints(fn ExecuteFunc) ([]int, error) {
	return redis.Ints(c.Execute(fn))
}

func (c *Client) IntMap(fn ExecuteFunc) (map[string]int, error) {
	return redis.IntMap(c.Execute(fn))
}

func (c *Client) Int64(fn ExecuteFunc) (int64, error) {
	return redis.Int64(c.Execute(fn))
}

func (c *Client) Int64s(fn ExecuteFunc) ([]int64, error) {
	return redis.Int64s(c.Execute(fn))
}
func (c *Client) Int64Map(fn ExecuteFunc) (map[string]int64, error) {
	return redis.Int64Map(c.Execute(fn))
}

func (c *Client) Uint64(fn ExecuteFunc) (uint64, error) {
	return redis.Uint64(c.Execute(fn))
}

func (c *Client) Bool(fn ExecuteFunc) (bool, error) {
	return redis.Bool(c.Execute(fn))
}

func (c *Client) String(fn ExecuteFunc) (string, error) {
	return redis.String(c.Execute(fn))
}

func (c *Client) StringMap(fn ExecuteFunc) (map[string]string, error) {
	return redis.StringMap(c.Execute(fn))
}

func (c *Client) Strings(fn ExecuteFunc) ([]string, error) {
	return redis.Strings(c.Execute(fn))
}

func (c *Client) Bytes(fn ExecuteFunc) ([]byte, error) {
	return redis.Bytes(c.Execute(fn))
}

func (c *Client) ByteSlices(fn ExecuteFunc) ([][]byte, error) {
	return redis.ByteSlices(c.Execute(fn))
}

func (c *Client) Positions(fn ExecuteFunc) ([]*[2]float64, error) {
	return redis.Positions(c.Execute(fn))
}

func (c *Client) Float64(fn ExecuteFunc) (float64, error) {
	return redis.Float64(c.Execute(fn))
}

func (c *Client) Float64s(fn ExecuteFunc) ([]float64, error) {
	return redis.Float64s(c.Execute(fn))
}

func (c *Client) Values(fn ExecuteFunc) ([]interface{}, error) {
	return redis.Values(c.Execute(fn))
}
