package validator

import (
	"fmt"
	"regexp"
)

type Match struct {
	Regexp *regexp.Regexp
}

func (m Match) IsSatisfied(obj interface{}) bool {
	str := obj.(string)
	return m.Regexp.MatchString(str)
}

func (m Match) DefaultMessage() string {
	return fmt.Sprint("Must match", m.Regexp)
}

var EmailRegexp = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

type Email struct {
	Match
}

func (e Email) DefaultMessage() string {
	return fmt.Sprint("Must be a valid email address")
}

var ChinaMobileRegexp = regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")

type ChinaMobile struct {
	Match
}

func (cm ChinaMobile) DefaultMessage() string {
	return fmt.Sprint("Must be a valid china mobile")
}

type Confirm string

func (c Confirm) IsSatisfied(obj interface{}) bool {
	str, ok := obj.(string)
	if ok {
		return string(c) == str
	}
	return false
}

func (c Confirm) DefaultMessage() string {
	return "不相等"
}
