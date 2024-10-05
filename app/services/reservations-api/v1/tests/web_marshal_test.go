package tests

import (
	"time"

	"github.com/ameghdadian/service/app/services/reservations-api/v1/handlers/businessgrp"
	"github.com/ameghdadian/service/app/services/reservations-api/v1/handlers/usergrp"
	"github.com/ameghdadian/service/business/core/business"
	"github.com/ameghdadian/service/business/core/user"
)

func toAppUser(usr user.User) usergrp.AppUser {
	roles := make([]string, len(usr.Roles))
	for i, role := range usr.Roles {
		roles[i] = role.Name()
	}

	return usergrp.AppUser{
		ID:           usr.ID.String(),
		Name:         usr.Name,
		Email:        usr.Email.Address,
		Roles:        roles,
		PasswordHash: nil, // This field is not marshalled.
		PhoneNo:      usr.PhoneNo.Number(),
		Enabled:      usr.Enabled,
		DateCreated:  usr.DateCreated.Format(time.RFC3339),
		DateUpdated:  usr.DateUpdated.Format(time.RFC3339),
	}
}

func toAppUsers(users []user.User) []usergrp.AppUser {
	items := make([]usergrp.AppUser, len(users))
	for i, usr := range users {
		items[i] = toAppUser(usr)
	}

	return items
}

func toAppUserPtr(usr user.User) *usergrp.AppUser {
	appUsr := toAppUser(usr)
	return &appUsr
}

// ----------------------------------------------------------

func toAppBusiness(b business.Business) businessgrp.AppBusiness {
	return businessgrp.AppBusiness{
		ID:          b.ID.String(),
		OwnerID:     b.OwnerID.String(),
		Name:        b.Name,
		Description: b.Desc,
		DateCreated: b.DateCreated.Format(time.RFC3339),
		DateUpdated: b.DateUpdated.Format(time.RFC3339),
	}
}

func toAppBusinesses(bsns []business.Business) []businessgrp.AppBusiness {
	items := make([]businessgrp.AppBusiness, len(bsns))
	for i, b := range bsns {
		items[i] = toAppBusiness(b)
	}

	return items
}

func toAppBusinessPtr(b business.Business) *businessgrp.AppBusiness {
	appBsn := toAppBusiness(b)
	return &appBsn
}
