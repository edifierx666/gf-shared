package cache

import (
  "github.com/gogf/gf/v2/database/gredis"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/gogf/gf/v2/os/gcache"
  "github.com/gogf/gf/v2/os/gctx"
  "github.com/gogf/gf/v2/os/gfile"
  "shared/library/cache/file"
)

type CacheType string

const (
  Redis  CacheType = "redis"
  File   CacheType = "file"
  Memory CacheType = "memory"
)

var (
  ctx = gctx.New()
)

func NewRedisCache(redis ...*gredis.Redis) gcache.Adapter {
  var r *gredis.Redis
  if len(redis) > 0 {
    r = redis[0]
  } else {
    r = g.Redis()
  }
  c := gcache.New()
  c.SetAdapter(gcache.NewAdapterRedis(r))
  return c
}

func NewFileCache(fileDir string) *gcache.Cache {
  if fileDir == "" {
    g.Log().Fatal(ctx, "file path must be configured for file caching.")
    return nil
  }
  if !gfile.Exists(fileDir) {
    if err := gfile.Mkdir(fileDir); err != nil {
      g.Log().Fatalf(ctx, "failed to create the cache directory. procedure, err:%+v", err)
      return nil
    }
  }
  c := gcache.New()

  c.SetAdapter(file.NewAdapterFile(fileDir))
  return c
}

func NewMemoryCache() *gcache.Cache {
  c := gcache.New()
  c.SetAdapter(gcache.NewAdapterMemory())
  return c
}
