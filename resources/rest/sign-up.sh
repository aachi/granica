#!/bin/sh

# Vars
HOST="localhost"
PORT="8080"

post () {
  echo "POST $1"
  curl -X POST $1 --header 'Content-Type: application/json' -d @resources/rest/sign-up.json
}

post "http://$HOST:$PORT/sign-up"