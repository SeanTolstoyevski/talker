# talker
 Golang wrapper that provides APIs such as Speech for Screen readers

## Welcome to Talker's Repo!

This Golang module is a Golang wrapper for [Tolk library](https://github.com/dkager/tolk).

There is currently no active English documentation.
The **talker.go file** has been commented. You can read this file for English.

Available documentation:

* [Turkish - Türkçe](https://github.com/SeanTolstoyevski/talker/blob/master/docs/README_tr.rst)

## İnstallation

One thing you have to decide is what architecture you will use Talker on. Because Talker has different DLLs for two different architectures.

* `bit64`: suitable for 64 bit machines.
* `bit32`: Suitable for 32 bit compilations.

You must entry these parameters as `-tags xxx`.

Example:
`go get -tags bit64 github.com/SeanTolstoyevski/talker`

