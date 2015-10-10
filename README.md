# Pagerduty report interpreter

## Requirements
 - `docker`

 ## Build
 `docker build -t dutyreader .`
 ## Run
 1. Copy the csv report from PagerDuty in working directory
 2. Run `docker run -it --rm dutyreader app ./incidents.csv "Alin Nica"`