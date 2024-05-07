# This file will tell about my family
#! /bin/bash

echo "$opt" | tr -d '"'
curl --location --request GET https://01.gritlab.ax/assets/superhero/all.json | jq '.[]| select ( .id == 1) .relatives'
