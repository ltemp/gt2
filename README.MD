Golgang sample project

Drivers API
-----------
Issue api calls to http://ADDRESS:8080/api/driver/

INSERT (POST) -
The following structure is used
type Driver {
	Id int
	Name string
	License_number string
}

example:
{""id":33, name": "test", "license_number":"test" }

GET- issue GET request to
http://ADDRESS:8080/api/driver/:id


DELETE - issue request to
http://ADDRESS:8080/api/driver/:id


Metrics API
-----------

Issue api calls to http://ADDRESS:8080/api/metric/

INSERT - POST
Issur POST request using the following structure (json) -

Metric {
	metric_name string
	driver_id   string
	value       string
	lon, lat    float64
	timestamp   int64
}

Example:
{"metric_name": "test3", "driver_id":"3", "value":"5", "lon":0, "lat":5555,"timestamp":1352289170}


For the following requests use the following object (json)


MetricQueryInfo struct {
	metric_names  []string
	driver_ids    []string
	value         string
	min_lon       string
	max_lon       string
	min_lat       string
	max_lat       string
	min_timestamp string
	max_timestamp string
}


Note-  GET request body isn't read, please use PUT to issue actually GET (according to
documentation setting copyrequestbody = true in the conf should fixed it. Content-Length in the request
is as expected but the content isn't read when accessing u.Ctx.Input.RequestBody) - Please use PUT in the meanwhile.

PUT(GET) - all matching metrics will be retrieved
DELETE - all matching metrics will be deleted.

Example queries -
{"metric_names": ["network.reception_strength", "gps.location_lost"], "driver_ids":["2","3"]}







