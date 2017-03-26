# Golgang sample project

## Drivers API
-----------
#### Adding driver -
Issue POST request to http://ADDRESS:8080/api/driver/

#### Updating driver -
Issue PUT request to http://ADDRESS:8080/api/driver/

The following structure is used for those requests
type Driver {
	Id int
	Name string
	License_number string
}

example:
{""id":33, name": "test", "license_number":"test" }

#### Retrieving driver - 
issue GET request with the id parameter
http://ADDRESS:8080/api/driver/:id


#### Deleting driver-
issue DELETE request with the id parameter
http://ADDRESS:8080/api/driver/:id


## Metrics API
-----------

Issue api calls to http://ADDRESS:8080/api/metric/

#### Adding metric
Issur POST request using the following structure (json) -
```
Metric {
	metric_name string
	driver_id   string
	value       string
	lon, lat    float64
	timestamp   int64
}
```

Example:
{"metric_name": "test3", "driver_id":"3", "value":"5", "lon":0, "lat":5555,"timestamp":1352289170}


#### Deleting metrics - all matching metrics will be deleted.
#### Getting metrics (PUT) - all matching metrics will be retrieved
Note-  GET request body isn't read, please use PUT to issue actually GET (according to beego
documentation setting copyrequestbody = true in the conf should've fixed it. Content-Length in the request
is as expected but the content isn't read when accessing u.Ctx.Input.RequestBody) - Please use PUT in the meanwhile.

For the following requests use the following strcture (json)

```
MetricQueryInfo  {
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
```

Example queries -
{"metric_names": ["network.reception_strength", "gps.location_lost"], "driver_ids":["2","3"]}

all metrics with one of "network.reception_strength" and "gps.location_lost" names for drivers 2/3 will be retrieved.








