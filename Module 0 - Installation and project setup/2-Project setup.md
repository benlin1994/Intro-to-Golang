### Setting up your project 

It is important to have an organized workspace otherwise your code will be hard to manage in the future. 

To start, create a new directory, such as `/Users/username/go`, that will be used as the workspace for all of your Go projects. 

Inside your workspace folder, add the following three directories:

* bin - this folder contains all your executables
* src - this folder contains all your Go source code
* pkg - this folder contains all your imported packages

### Setting up your Go environment variables

We now need to configure some Go environment variables in order for our project structure to work correctly.

Run the following commands using your workspace path to configure your Go environment variables:

```shell
export GOPATH=/Users/username/go
export GOBIN=$GOPATH/bin
export PATH=$GOPATH/bin:$PATH

```
Here's a bit of explanation of what each command did:

* GOPATH specifies where your Go source files are located. 

* GOBIN specifies where your executables will live once you run a `go install` command. 

* PATH isn't an environment variable, but we are configuring it so we can run executables from our bin file easier.