# This file looks for all files ending with .sh
#! /bin/bash
find -name '*.sh'| rev | cut -d '/' -f1 | rev | cut -d '.' -f1