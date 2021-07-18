# gse
test for go symbol exports (pain)

## why
trying to see if i can get myself a shared object and cheat around instead of writing a whole
replacement for golang stuff in C

## some credits
uses like 2 files from mytls ([link](https://github.com/zedd3v/mytls/)).

## what we know
cgo is hard :(

### oh my god there is a specification (wow)
https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/edit?pli=1

## build requirements
* go
* clang

### build-env
current builds are being tested on `ghcr.io/alii/alibuntu:latest`, more can be read [here](https://github.com/alii/alibuntu).

## todo
we can create the shared library, but it's seemingly not recognized as a valid sofile by the system. we need this to work. :facepalm:

other then that, we need to make sure we can properly get an address, etc etc.
translating go objects into C would be nice, though might initally be a struggle.

however, go objects are structs, so should probably be fine? we will see

go's std would be useful in some use cases (cough) and if somehow this will decide
to link properly then life will be great 4 me

from some research (`readelf`), there is not a valid soname in the so file eghhhhhhh

okay another odd thing, it has a main method.. so it should be a program right?
but all articles say that c-shared libraries need a main method so it's quite misleading and idk
i'm committing this now in case i accidentally purge this image so lets see how it goes


