package model

import "fmt"

type Metric struct {
	Id          int     `json:"id"`
	IdLxd       int     `json:"id_lxd"`
	CpuUsage    float64 `json:"cpu_usage"`
	MemoryUsage int     `json:"memory_usage"`
}

func (m *Metric) CreateMetric(ldxId int) (*int, error) {
	err := Db.QueryRow("INSERT INTO metrics(id_lxd) VALUES($1) RETURNING id", ldxId).Scan(&m.Id)

	if err != nil {
		return nil, err
	}

	return &m.Id, nil
}

func (m *Metric) GetMetric(id int) (*Metric, error) {
	queryString := fmt.Sprintf("select * from metrics where id=%d", id)

	row := Db.QueryRow(queryString)

	metric := Metric{}

	if err := row.Scan(&metric.Id, &metric.IdLxd, &metric.CpuUsage, &metric.MemoryUsage); err != nil {
		return nil, err
	}

	return &metric, nil
}

func (m *Metric) GetMetricByLXDId(lxdId int) (*Metric, error) {
	queryString := fmt.Sprintf("select * from metrics where id_lxd=%d", lxdId)

	row := Db.QueryRow(queryString)

	metric := Metric{}

	if err := row.Scan(&metric.Id, &metric.IdLxd, &metric.CpuUsage, &metric.MemoryUsage); err != nil {
		return nil, err
	}

	return &metric, nil
}
