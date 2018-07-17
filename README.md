# GENREVAN SCHEDULER

Genrevan Scheduler is an open source project to allocate new container efficiently based on load from LXD worker.
This scheduler use worker's metrics like CPU and memory usage to decide which LXD available to create new container.

## Limitation
- Cannot decide which LXD worker will allocate by disk usage.

## Depedencies
- Gorilla Mux
- Testify
- Postgresql

## Build
``` go install genrevan-scheduler ```

## Run
``` genrevan-scheduler ```

## Test
``` go test -v ```
