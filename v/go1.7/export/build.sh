#!/bin/bash

rm __.PKGDEF _go_.o lib.a
go tool compile -newexport=1 -pack ./lib.go
ar -x lib.a
file __.PKGDEF

