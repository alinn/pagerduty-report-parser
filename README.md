# Pagerduty report interpreter


## Setup Without docker
### Dependencies
- `golang`

### Run
1. Copy the csv report from PagerDuty in working directory
2. Run `go run main.go ./incidents.csv "Alin Nica"`

## Setup With docker
### Dependencies
- `docker`

### Build
`docker build -t dutyreader .`
### Run
1. Copy the csv report from PagerDuty in working directory
2. Run `docker run -it --rm dutyreader app ./incidents.csv "Alin Nica"`