#!/bin/bash

# Wait for MySQL container to be ready
/usr/local/bin/wait-for-it.sh -t 0 mysql:3306 -- echo "MySQL is up!"

# Start your application
# Replace the command below with the command to start your application
exec ./output/bootstrap.sh