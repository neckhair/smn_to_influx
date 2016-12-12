# SMN to Influxdb

Stores [SwissMetNet data](http://opendata.netcetera.com/smn/swagger) into an InfluxDB database.

    export INFLUXDB_URL=localhost:8086
    export INFLUXDB_DATABASE=smn
    smn_to_influx BUS

## Development

    docker-compose up
    docker-compose exec influx influx
    create database climate;

    export INFLUXDB_URL=localhost:8086
    export INFLUXDB_DATABASE=smn
    go build .
    smn_to_influx BUS

Then go to http://localhost:3000, login with `admin/admin` and create the influxdb datasource to http://localhost:8086 (direct connection). In grafana you can import the dashboard at `meta/dashboard.json` to view the data.
