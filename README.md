# Learning [Go](http://golang.org/)

This simple project acts as the basis for an introduction to the programming language Go. It is meant to be downloaded at the Codinglab by the participants.

## Setup
### Install Go
Head over to the [download page](http://golang.org/doc/install) and install the runtime for your system.

Be sure to write & run the simple example application as shown on the download page.

### Prepare a workspace
The recommended way to work with Go is by the use of a workspace as described [here](http://golang.org/doc/code.html). Setup a base directory containing the folders ```bin```, ```pkg``` and ```src``` and let the ```GOPATH``` environment variable point to this directory.

### Prepare IDE
Any text editor suffices to work with Go.

If no preference exists, use the following setup to have language support:
* Get a copy of [SublimeText](http://www.sublimetext.com/)
* Install the [Package Control](https://sublime.wbond.net/installation) plugin
* Use Package Control to install [GoSublime](https://github.com/DisposaBoy/GoSublime)

### Download this project
On the command line, navigate to the ```src``` directory of your workspace and enter the command
```
go get github.com/nightlyone/go-sample
```
to let Go download the project.

Then enter the directory ```src/github.com/nightlyone/go-sample``` and enter the commands
```
go get ./...
go install ./...
```
to fetch all dependencies and compile the application.

This should have placed an executable binary ```go-sample``` into the ```bin``` folder of your workspace.

## References

* [Go Homepage](http://golang.org/)
* [A Tour of Go](http://tour.golang.org/#1)
* Inspired by https://github.com/CodingDojoVie/learning-go
