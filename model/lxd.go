package model

import (
	"fmt"

	"github.com/guregu/null"
)

type Lxd struct {
	Id        int         `json:"id"`
	Name      null.String `json:"name"`
	IpAddress string      `json:"ip_address"`
}

func (l *Lxd) CreateLXD(ip string) error {
	queryString := fmt.Sprintf("insert into lxds (ip_address) values('%s') on conflict do nothing", ip)
	_, err := Db.Exec(queryString)

	if err != nil {
		return err
	}

	return nil
}

func (l *Lxd) GetLXD(ip string) (*Lxd, error) {
	queryString := fmt.Sprintf("select * from lxds where ip_address='%s'", ip)
	row := Db.QueryRow(queryString)

	lxd := Lxd{}

	if err := row.Scan(&lxd.Id, &lxd.Name, &lxd.IpAddress); err != nil {
		return nil, err
	}

	return &lxd, nil
}
