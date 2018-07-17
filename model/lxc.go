package model

import (
	"fmt"

	"github.com/guregu/null"
)

type Lxc struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	IpAddress null.String `json:"ip_address"`
	Image     string      `json:"image"`
	Status    string      `json:"status"`
	LxdId     null.Int    `json:"lxd_id"`
}

func (l *Lxc) GetLXCs() ([]Lxc, error) {
	rows, err := Db.Query("select * from lxcs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lxcs := []Lxc{}

	for rows.Next() {
		var lxc Lxc
		if err := rows.Scan(&lxc.Id, &lxc.Name, &lxc.IpAddress, &lxc.Image, &lxc.Status, &lxc.LxdId); err != nil {
			return nil, err
		}
		lxcs = append(lxcs, lxc)
	}

	return lxcs, nil
}

func (l *Lxc) GetLXC(id int) (*Lxc, error) {
	queryString := fmt.Sprintf("select * from lxcs where id=%d", id)

	row := Db.QueryRow(queryString)

	lxc := Lxc{}

	if err := row.Scan(&lxc.Id, &lxc.Name, &lxc.IpAddress, &lxc.Image, &lxc.Status, &lxc.LxdId); err != nil {
		return nil, err
	}

	return &lxc, nil
}
