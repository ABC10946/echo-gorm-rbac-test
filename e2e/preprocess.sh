#!/bin/bash

sqlite3 ../test.db "INSERT INTO users (username, password) VALUES ('admin', 'admin')"