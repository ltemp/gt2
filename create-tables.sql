CREATE TABLE driver
(
    id INTEGER PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    license_number VARCHAR(10)
);

CREATE TABLE metric
(
    name VARCHAR(30) NOT NULL,
    value VARCHAR(30) NOT NULL,
    lon DECIMAL,
    lat DECIMAL,
    timestamp BIGINT,
    driver_id INTEGER REFERENCES driver(id) ON DELETE CASCADE
);

## can add multicolumn indexes
CREATE INDEX METRIC_NAME_IDX2 ON metric(NAME);
CREATE INDEX METRIC_NAME_IDX2 ON metric(value);
CREATE INDEX METRIC_NAME_IDX2 ON metric(lon);
CREATE INDEX METRIC_NAME_IDX2 ON metric(lat);
CREATE INDEX METRIC_NAME_IDX2 ON metric(timestamp);
CREATE INDEX METRIC_NAME_IDX2 ON metric(driver_id);