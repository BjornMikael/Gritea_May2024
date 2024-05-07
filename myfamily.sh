# This file will tell about my family
#! /bin/bash

curl -s https://01.gritlab.ax/assets/superhero/all.json | jq --arg hero_id "$HERO_ID" '.[] | select(.id == ( $hero_id | tonumber)) | .connections.relatives' | tr -d '"'