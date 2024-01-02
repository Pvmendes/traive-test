#!/bin/bash

# MongoDB connection details
MONGODB_HOST=localhost
MONGODB_PORT=27017
MONGODB_ADMIN_USER=root
MONGODB_ADMIN_PASS=123456
DB_NAME=mytransactionsdb
DB_USER=appUser
DB_PASS=123456

# Create MongoDB database and user
mongo admin --host $MONGODB_HOST --port $MONGODB_PORT -u $MONGODB_ADMIN_USER -p $MONGODB_ADMIN_PASS --eval "db.getSiblingDB('$DB_NAME').createUser({user: '$DB_USER', pwd: '$DB_PASS', roles: ['readWrite']})"

echo "MongoDB setup completed:"
echo "Database: $DB_NAME"
echo "User: $DB_USER"
echo "Password: $DB_PASS"
