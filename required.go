package validator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/haiyiyun/utils/help"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Required struct{}

func (r Required) IsSatisfied(obj interface{}) bool {
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

	return false
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
	Item interface{}
}

func (h Have) IsSatisfied(obj interface{}) bool {
	if str, ok := h.Item.(string); ok {
		strs := []string{}
		if objs, ok := obj.([]interface{}); ok {
			for _, v := range objs {
				if vstr, ok := v.(string); ok {
					strs = append(strs, vstr)
				}
			}

			return help.NewSlice(strs).CheckItem(str)
		}

	}

	if i, ok := h.Item.(int); ok {
		is := []int{}
		if objs, ok := obj.([]interface{}); ok {
			for _, v := range objs {
				if vint, ok := v.(int); ok {
					is = append(is, vint)
				}
			}

			return help.NewSlice(is).CheckDigital(i)
		}
	}

	return false
}

func (h Have) DefaultMessage() string {
	return fmt.Sprint("Not found ", h.Item)
}

type Digital struct {
}

func (d Digital) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		if _, err := strconv.Atoi(str); err == nil {
			return true
		}
	}

	return false
}

func (d Digital) DefaultMessage() string {
	return fmt.Sprint("not digital")
}

type Float struct {
}

func (d Float) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		if _, err := strconv.ParseFloat(str, 64); err == nil {
			return true
		}
	}

	return false
}

func (d Float) DefaultMessage() string {
	return fmt.Sprint("not float")
}

type BsonObjectID struct {
}

func (boid BsonObjectID) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		if _, err := primitive.ObjectIDFromHex(str); err == nil {
			return true
		}
	}

	return false
}

func (boid BsonObjectID) DefaultMessage() string {
	return fmt.Sprint("not bson ObjectID")
}
