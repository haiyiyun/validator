package validator

import (
	"strconv"
	"time"

	"github.com/haiyiyun/utils/help"
)

type Required struct{}

func (r Required) IsSatisfied(obj interface{}) bool {
	if obj == nil {
		return false
	}

	if str, ok := obj.(string); ok {
		return len(str) > 0
	}
	if list, ok := obj.([]interface{}); ok {
		return len(list) > 0
	}
	if list, ok := obj.([]string); ok {
		return len(list) > 0
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	if i, ok := obj.(int); ok {
		return i != 0
	}
	if t, ok := obj.(time.Time); ok {
		return !t.IsZero()
	}

	return true
}

func (r Required) DefaultMessage() string {
	return "Required"
}

type Or struct {
	Want     int
	Included int
}

func (o Or) IsSatisfied(obj interface{}) bool {
	ss := obj.([]string)
	if len(ss) < o.Want {
		return false
	}

	o.Included = 0
	for _, s := range ss {
		if (Required{}).IsSatisfied(s) {
			o.Included++
		}
	}

	if o.Included >= o.Want {
		return true
	}

	return false
}

func (o Or) DefaultMessage() string {
	return "Contains at least " + strconv.Itoa(o.Want) + ", currently contains only " + strconv.Itoa(o.Included)
}

type Have struct {
	Item string
}

func (h Have) IsSatisfied(obj interface{}) bool {
	return help.NewSlice(obj).CheckItem(h.Item)
}

func (h Have) DefaultMessage() string {
	return "Not found " + h.Item
}
