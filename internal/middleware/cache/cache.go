package cache

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Option interface {
	apply(*option)
}

type option struct {
	maxAge  int
	sMaxAge int
	noCache bool
	noStore bool
	public  bool
	private bool
}

type Cache struct {
	opt *option
}

type MaxAge struct {
	value int
}

func (m *MaxAge) apply(opt *option) {
	opt.maxAge = m.value
}

func NewMaxAge(value int) Option {
	return &MaxAge{value: value}
}

type SMaxAge struct {
	value int
}

func (s *SMaxAge) apply(opt *option) {
	opt.sMaxAge = s.value
}

func NewSMaxAge(value int) Option {
	return &SMaxAge{value: value}
}

type NoCache struct {
}

func (n *NoCache) apply(opt *option) {
	opt.noCache = true
}

func NewNoCache() Option {
	return &NoCache{}
}

type NoStore struct {
}

func (n *NoStore) apply(opt *option) {
	opt.noStore = true
}

func NewNoStore() Option {
	return &NoStore{}
}

type Public struct {
}

func (p *Public) apply(opt *option) {
	opt.public = true
}

func NewPublic() Option {
	return &Public{}
}

type Private struct {
}

func (p *Private) apply(opt *option) {
	opt.private = true
}

func NewPrivate() Option {
	return &Private{}
}

func (c *Cache) CacheMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var headers []string
		if c.opt.maxAge > 0 {
			headers = append(headers, fmt.Sprintf("max-age=%d", c.opt.maxAge))
		}
		if c.opt.sMaxAge > 0 {
			headers = append(headers, fmt.Sprintf("s-maxage=%d", c.opt.sMaxAge))
		}
		if c.opt.public {
			headers = append(headers, "public")
		} else if c.opt.private {
			headers = append(headers, "private")
		}
		if c.opt.noCache {
			headers = append(headers, "no-cache")
		}
		if c.opt.noStore {
			headers = append(headers, "no-store")
		}
		ctx.Header("Cache-Control", strings.Join(headers, ","))
	}
}

func NewCache(o ...Option) *Cache {
	var opt = &option{}
	for _, f := range o {
		f.apply(opt)
	}
	return &Cache{
		opt: opt,
	}
}
