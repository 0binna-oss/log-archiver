# Log Archiver Tool

A modular log archiver tool written in Go. This tool archives logs, compresses them and rotates old logs based on size and age.

# features

* Archive Logs: Save logs with a timestamp in the filename.

* Compress logs: Optionally compress archive logs using gzip.

* Log Rotation: Automatically delete old logs based on size and age.

* Configuration: Customize settings like archive directory, compression, and rotation rules via `config.json`

* CLI Support: Easy-to-use command-line interface

# Installation

Prerequisites

* Go(version 1.20 or higher)installed on your system.

Steps

1. Clone the repository:

   ```git clone https://github.com/0binna-oss/log-archiver.git```

      ```cd log-archiver```

2. Build the tool:

   ```go build -o log-archiver```

3. Run the tool:

   ```./log-archiver```
