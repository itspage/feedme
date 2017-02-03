package outlets

import (
	"github.com/itspage/feedme/order"
	"github.com/itspage/feedme/outlets/tossed"
	"github.com/itspage/feedme/user"
)

type Outlet interface {
	NewOrder(*user.Preferences, *user.Profile) (order.Order, error)
	Login(*user.Preferences) (*user.Profile, error)
}

// Select returns an Outlet based on the user's preferences
func Select(preferences *user.Preferences) (Outlet, error) {
	return tossed.Outlet(), nil
}
