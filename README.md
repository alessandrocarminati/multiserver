# Multiserver Tool
Multiserver is a simple, lightweight application providing both HTTP and TFTP
server functionality for read-only access to static files. 
Written in Go, it's a standalone, cross-platform binary that requires no 
dependencies... 
just place it in a directory, execute, and it's ready to go.

## Features
* HTTP & TFTP Server (Read-only): Serves files via both protocols.
* File Directory: Serves all files located in the `./file` directory.
* Cross-Platform Compatibility: Go-based for straightforward builds 
  on multiple architectures.
* No Daemon Mode: Suggested start command:
```
./multiserver &> ./log &
```

## Notes
* No Configuration Needed: Multiserver is ready to serve files immediately
  upon execution.
* Logging: Logs output to stdout.
