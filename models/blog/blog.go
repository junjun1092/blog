package blog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/g"
)

func OneById(id int64) *Blog {
	if id <= 0 {
		return nil
	}
	key := fmt.Sprintf("%d", id)
	val := g.BlogCacheGet(key)
	if val == nil {
		if p := OneByIdInDB(id); p != nil {
			g.BlogCachePut(key, *p)
			return p
		}
		return nil
	}
	ret := val.(Blog)
	return ret
}

func OneByIdInDB(id int64) *Blog {
	if id <= 0 {
		return nil
	}
	o := Blog{Id: id}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}
