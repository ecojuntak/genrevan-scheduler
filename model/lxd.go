package model

import (
	"fmt"
	"strings"

	"github.com/guregu/null"
)

type Lxd struct {
	Id        int         `json:"id"`
	Name      null.String `json:"name"`
	IpAddress string      `json:"ip_address"`
}

func (l *Lxd) CreateLXD(ip string) error {
	queryString := fmt.Sprintf("insert into lxds (ip_address) values('%s')", ip)
	_, err := Db.Exec(queryString)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil
		} else {
			return err
		}
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

func (l *Lxd) GetLXDs() ([]Lxd, error) {
	rows, err := Db.Query("select * from lxds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lxds := []Lxd{}

	for rows.Next() {
		var lxd Lxd
		if err := rows.Scan(&lxd.Id, &lxd.Name, &lxd.IpAddress); err != nil {
			return nil, err
		}
		lxds = append(lxds, lxd)
	}

	return lxds, nil
}
