# talker – Golang wrapper for Tolk library

[![Go Reference](https://pkg.go.dev/badge/github.com/SeanTolstoyevski/talker@master.svg)](https://pkg.go.dev/github.com/SeanTolstoyevski/talker@master?GOOS=windows)

## Welcome to Talker's Repo!

This Golang module is a wrapper for [Tolk library](https://github.com/dkager/tolk) (C).

No changes have been made to the use  API.
Only utf-8/ utf-16 conversions is automatically.
You do not need to do these conversions.

You can check the **examples folder** to see more stuff.

## İnstallation

**Note**: **GCC** is required. CGO was used. Before installing you must install GCC on your system.

`go get -u github.com/SeanTolstoyevski/talker`

This command gets the latest tagged version (git tag).

To install the latest commit:  
`go get github.com/SeanTolstoyevski/talker@master`

This command install  the last commit in the master branch.

