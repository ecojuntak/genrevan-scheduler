package model

import (
	"fmt"

	"github.com/guregu/null"
)

type Lxc struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	IpAddress     null.String `json:"ip_address"`
	Image         string      `json:"image"`
	Status        string      `json:"status"`
	LxdId         null.Int    `json:"lxd_id"`
	HostPort      int         `json:"host_port"`
	ContainerPort int         `json:"container_port"`
	ErrorMessage  null.String `json:"error_message"`
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
		if err := rows.Scan(&lxc.Id, &lxc.Name, &lxc.IpAddress, &lxc.Image, &lxc.Status, &lxc.LxdId, &lxc.HostPort, &lxc.ContainerPort, &lxc.ErrorMessage); err != nil {
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

	if err := row.Scan(&lxc.Id, &lxc.Name, &lxc.IpAddress, &lxc.Image, &lxc.Status, &lxc.LxdId, &lxc.HostPort, &lxc.ContainerPort, &lxc.ErrorMessage); err != nil {
		return nil, err
	}

	return &lxc, nil
}

func (l *Lxc) GetLXCsByLXDId(id int) ([]Lxc, error) {
	queryString := fmt.Sprintf("select * from lxcs where id_lxd=%d", id)

	rows, err := Db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lxcs := []Lxc{}

	for rows.Next() {
		var lxc Lxc
		if err := rows.Scan(&lxc.Id, &lxc.Name, &lxc.IpAddress, &lxc.Image, &lxc.Status, &lxc.LxdId, &lxc.HostPort, &lxc.ContainerPort, &lxc.ErrorMessage); err != nil {
			return nil, err
		}
		lxcs = append(lxcs, lxc)
	}

	return lxcs, nil
}

func (l *Lxc) IsLXCsExist(id int, hostPort int) (bool, error) {
	queryString := fmt.Sprintf("select id_lxd, host_port from lxcs where id_lxd=%d and host_port=%d", id, hostPort)

	rows, err := Db.Query(queryString)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (l *Lxc) CreateLXC(lxc Lxc) (*int, error) {
	err := ValidateLXCName(lxc.Name)

	if err != nil {
		return nil, err
	}

	err = Db.QueryRow("INSERT INTO lxcs(name, image, id_lxd, host_port, container_port) VALUES($1, $2, $3, $4, $5) RETURNING id", lxc.Name, lxc.Image, lxc.LxdId, lxc.HostPort, lxc.ContainerPort).Scan(&lxc.Id)

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

func (l *Lxc) UpdateErrorMessage(id int, msg string) error {
	_, err := Db.Exec("UPDATE lxcs SET error_message=$1 WHERE id=$2", msg, id)

	if err != nil {
		return err
	}

	return nil
}
