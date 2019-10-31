package uri

import (
)

type URI interface {
	String() string
	Base() string
	Root() string
}
