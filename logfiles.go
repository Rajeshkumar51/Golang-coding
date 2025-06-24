package main

import (
	"os"
)

func createFile(name string, content string) {
	os.WriteFile(name, []byte(content), 0644)
}

func main() {
	createFile("server1.log", `INFO: Server started successfully
ERROR: Failed to connect to database
INFO: Request processed
ERROR: Timeout occurred while processing request`)

	createFile("server2.log", `INFO: Backup started
INFO: Backup completed
ERROR: Disk space low
INFO: Cleanup done`)

	createFile("server3.log", `INFO: User login detected
ERROR: Authentication failure
INFO: Session expired
ERROR: Invalid token received`)
}
