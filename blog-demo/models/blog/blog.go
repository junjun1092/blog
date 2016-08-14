package blog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/g"
)

func OneById(id int64) *blog {
	if id < 0 {
		return nil
	}
	key := fmt.Sprintf("%d", id)
	val := g.BlogCacheGet(key)
	if val == nil {
		if p := OneByIdInDB(); p != nil {
			g.BlogCachePut(key, *p)
		}
		return nil
	}
	ret := val.(blog)
	return &ret
}

func OneByIdInDB(id int64) *blog {
	if id < 0 {
		return nil
	}
	o := blog{Id: id}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}
func main() {

}
