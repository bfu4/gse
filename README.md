# gse
test for (**g**)o (**s**)ymbol (**e**)xports

proudly written at 2am in vim on a docker container (i dont hate myself) (lots go untested at 2am)

## why
trying to see if i can get myself a shared object and cheat around instead of writing a whole
replacement for golang stuff in C

## some credits
uses like 2 files from mytls ([link](https://github.com/zedd3v/mytls/)).

## running
for now this program can be ran like so:

```bash
$ GODEBUG=cgocheck=0 ./out/main
```

this is because we are passing back the go pointer at the moment.
this is ok! we are checking that it works. we deal with the rest later!
we will have to deal with the management of memory at some point, but the
calling convention for this program will be very different from how it is now.

## what we know
* go files are large
* this is okay, because our C is not
* we can shove these somewhere and our shared object probably wont be over 30Mb by the end. This will be okay, and (hopefully) gives a temporary fix to what is trying to be done :)

### oh my god there is a specification (wow)
* [vladimir vivien (THANK YOU)](https://github.com/vladimirvivien/go-cshared-examples)
* [google or someone idk](https://docs.google.com/document/d/1nr-TQHw_er6GOQRsF6T43GGhFDelrAP0NqSS_00RgZQ/edit?pli=1)

## build requirements
* go
* clang

### build-env
current builds are being tested on `ghcr.io/alii/alibuntu:latest`, more can be read [here](https://github.com/alii/alibuntu).

## progress
![man](https://cdn.discordapp.com/attachments/827656773296193597/866203029552365588/unknown.png)

## todo
translating go objects into C would be nice, though might initally be a struggle.

however, go objects are structs, so should probably be fine? we will see

go's std would be useful in some use cases (cough) and if somehow this will decide
to link properly then life will be great 4 me
