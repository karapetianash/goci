## Description

This `goci` is a simple, but useful, implementation of a `Continuous Integration (CI)` 
tool for your Go programs. 

A typical `CI pipeline` consists of several automated steps that
continuously ensure a code base or an application is ready to be merged
with some other developer’s code, usually in a shared version control
repository.

### For this example, the CI pipeline consists of:

* Building the program using `go build` to verify if the program structure is
valid.
* Executing tests using `go test` to ensure the program does what it’s
intended to do.
* Executing `gofmt` to ensure the program’s format conforms to the
standards.
* Executing `git push` to push the code to the remote shared Git repository
that hosts the program code.

The `goci` tool accepts only one flag `-p` of type string,
which represents the Go project directory on which to execute the `CI`
pipeline steps.

## Usage
To build the app, run the following command in the root folder:

```
> go build .
```
Above command will generate `goci` file. This name is defined in the `go.mod` file, and it will be the initialized module name.

After that you can run the program using the cmd.\
Example of using the pipeline for a project:

```
> .\goci.exe -p project/directory/
```
