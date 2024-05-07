#this shell will tell who you are
#! /bin/bash

curl --location --request GET https://01.gritlab.ax/assets/superhero/all.json | jq '.[]| select ( .id ==70) .name'
