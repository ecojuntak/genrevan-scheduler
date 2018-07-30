# GENREVAN SCHEDULER
[![Build Status](https://travis-ci.org/go-squads/genrevan-scheduler.svg?branch=master)](https://travis-ci.org/go-squads/genrevan-scheduler)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-squads/genrevan-scheduler)](https://goreportcard.com/report/github.com/go-squads/genrevan-scheduler)

Genrevan Scheduler is an open source project builded by GO-JEK intern team. The main purpose of this project is to allocate new container efficiently based on load from LXD worker.
This scheduler use LXD worker's metrics like CPU and memory usage to decide which LXD available to create new container.

## Depedencies
- Gorilla Mux
- Testify
- Postgresql

##Configuration
Copy and rename ``` development.example.yaml ``` to ``` development.yaml ``` in folder ``` /config ```. Do the same thing for testing configuration file. Set your environment variable on the configuration file.


## Build
``` go install genrevan-scheduler ```

## Migration
``` genrevan-scheduler migrate ```

## Seeding
``` genrevan-scheduler seed ```
Please define you own seeder first.

## Run
``` genrevan-scheduler start ```

## Test
``` go test -v -race ./... ```
