# Demo webserver written in GO

Requirements:

- go >=v1.21 (in theory should work >=v1.16 but untested)
- linux (necessary for systemd operation)
- superuser access

## Installation

The `install.sh` script is idempodent and can be safely reran. This installer will need super user permissions and should only be ran inside of a linux based machine.

```
sudo ./install.sh
```

## Local Developments

There is a `generate-data.sh` script that will write to a file specfied or by default write to `/data/metrics_from_special_app.txt` once every second with randomly generated data. Run this in a separate terminal if you want to test the webserver polling for updates.

```
go run server.go
```

```
./generate-data.sh
```

## Testing

```
go test
```

## Build

```
go build
```
