package casbin

import (
	casbin "github.com/casbin/casbin/v2"
)

func Check(e *casbin.Enforcer, sub, obj, ac string) {
	ok, _ := e.Enforce(sub, obj, act)
}
