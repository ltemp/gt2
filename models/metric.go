package models

import (
	"bytes"
	"strings"
	"fmt"
)

// Used to represent a single metric
type Metric struct {
	Name      string `json:"metric_name"`
	Driver_id int `json:"driver_id,string"`
	Value     string
	Lon, Lat  float64
	Timestamp int64
}

// Used to represent a metric query
type MetricQueryInfo struct {
	Metric_names  []string `json:"metric_names"`
	Driver_ids    []string `json:"driver_ids"`
	Value         string `json:"value"`
	Min_lon       string `json:"min_lon"`
	Max_lon       string `json:"max_lon"`
	Min_lat       string `json:"min_lat"`
	Max_lat       string `json:"max_lat"`
	Min_Timestamp string `json:"min_timestamp"`
	Max_Timestamp string `json:"max_timestamp"`
}

func InsertMetric(metric Metric) (error) {
	stmt, err := DBCon.Prepare("INSERT INTO metric(name,driver_id,value,lon,lat,timestamp) VALUES($1,$2,$3,$4,$5,$6)")
	_, err = stmt.Exec(metric.Name, metric.Driver_id, metric.Value, metric.Lon, metric.Lat, metric.Timestamp)
	return err
}

func DeleteMetrics(info MetricQueryInfo) (error, bool) {
	var query bytes.Buffer
	query.WriteString("DELETE FROM metric WHERE ")
	buildMetricsInfoConditions(info, &query)
	res, err := DBCon.Exec(query.String())
	if err != nil {
		return err, false
	}
	rows, _ := res.RowsAffected()
	return err, rows > 0
}

func LoadMetrics(info MetricQueryInfo) (error, []Metric) {
	var query bytes.Buffer
	query.WriteString("SELECT * FROM metric WHERE ")
	buildMetricsInfoConditions(info, &query)
	// TODO: Need to first select here the count, if the size of the result is too big
	// then return a error.
	rows, err := DBCon.Query(query.String())
	fmt.Print(query.String())
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	var metrics []Metric
	for rows.Next() {
		res := Metric{}
		err := rows.Scan(&res.Name, &res.Value, &res.Lon, &res.Lat,  &res.Timestamp, &res.Driver_id)
		if err != nil {
			return err,nil
		}
		metrics = append(metrics, res)
	}

	return nil, metrics
}


func buildMetricsInfoConditions(info MetricQueryInfo, query *bytes.Buffer) {
	if info.Metric_names != nil {
		query.WriteString("name in ('" + (strings.Join(info.Metric_names[:],"','")) + "') AND ")
	}

	if info.Driver_ids != nil {
		query.WriteString("driver_id in (" + (strings.Join(info.Driver_ids[:],",")) + ") AND ")
	}

	if info.Min_lat != "" {
		query.WriteString("lat >= ")
		query.WriteString(info.Min_lat)
		query.WriteString(" AND ")
	}

	if info.Max_lat != "" {
		query.WriteString("lat <= ")
		query.WriteString(info.Max_lat)
		query.WriteString(" AND ")
	}


	if info.Min_lon != "" {
		query.WriteString("lon >= ")
		query.WriteString(info.Min_lon)
		query.WriteString(" AND ")
	}

	if info.Max_lon != "" {
		query.WriteString("lon <= ")
		query.WriteString(info.Max_lon)
		query.WriteString(" AND ")
	}

	if info.Min_Timestamp != "" {
		query.WriteString("timestamp >= ")
		query.WriteString(info.Min_lon)
		query.WriteString(" AND ")
	}

	if info.Max_Timestamp != "" {
		query.WriteString("timestamp < ")
		query.WriteString(info.Max_Timestamp)
		query.WriteString(" AND ")
	}

	query.WriteString(" true;")
}
