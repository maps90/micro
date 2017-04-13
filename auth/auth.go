package auth

import (
	"fmt"

	"github.com/maps90/librarian"
	d "github.com/maps90/librarian/datasource"
)

type ApiUser struct {
	Id         uint8
	Token      *string
	ApiGroupId *string `db:"api_group_id"`
	Status     *string
	User       *string
	Apikey     *string
}

func (*ApiUser) PersistenceName() string {
	return "api_users"
}

func GetCustomAuth(bearer, token string) (bool, error) {
	read, ok := librarian.Get("mysql.slave").(d.DataAccessor)
	if !ok {
		return false, fmt.Errorf("cannot parse mysql into data accessor.")
	}

	results := ApiUser{}
	read.Find(&ApiUser{}, map[string]interface{}{"user": bearer, "apikey": token}, nil, &results)
	if results.Id == 0 {
		return false, fmt.Errorf("unauthorized access.")
	}

	return true, nil
}
