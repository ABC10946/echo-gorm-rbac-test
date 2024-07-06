#!/bin/bash

sqlite3 ../test.db "DELETE FROM items"
sqlite3 ../test.db "DELETE FROM users"