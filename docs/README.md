# Documentation

So, to state the obvious, the files in this folder are various bits of documentation for the project. 

## ficsiteconf.json
This file is an example JSON file that has been scrubbed of potentially risky information about my file system, passwords, etc. You will want to put this file someplace safe on your system, and set an appropriate env variable to point to it. For the development version of this code, I am simply calling that variable $FICSITECONF as seen in app/config.go.