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

func (l *Lxc) CreateLXC(lxc Lxc) (*int, error) {
	err := Db.QueryRow("INSERT INTO lxcs(name, image) VALUES($1, $2) RETURNING id", lxc.Name, lxc.Image).Scan(&lxc.Id)

	if err != nil {
		return nil, err
	}

	return &lxc.Id, nil
}

func (l *Lxc) DeleteLXC(id int) error {
	_, err := Db.Exec("delete from lxcs where id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func (l *Lxc) UpdateIpAddress(id int, ipAddress string) error {
	_, err := Db.Exec("UPDATE lxcs SET ip_address=$1 WHERE id=$2", ipAddress, id)

	if err != nil {
		return err
	}

	return nil
}

func (l *Lxc) UpdateState(id int, status string) error {
	_, err := Db.Exec("UPDATE lxcs SET status=$1 WHERE id=$2", status, id)

	if err != nil {
		return err
	}

	return nil
}

func (l *Lxc) UpdateLxdId(id, lxdId int) error {
	_, err := Db.Exec("UPDATE lxcs SET id_lxd=$1 WHERE id=$2", lxdId, id)

	if err != nil {
		return err
	}

	return nil
}
