# BRISK

BRISK is a simple and fast interpreted programming language.

## Running the source code

To run BRISK directly from source code, first go to https://golang.org/dl/ and follow the instructions to install golang on your operating system.

navigate to the root of the BRISK folder and run the command 
`go run src/brisk/main.go repl` to start the repl

## Download and Install

To install BRISK, download the desired release version from https://github.ibm.com/Kai-Mumford-CIC-UK/brisk/releases.

Once downloaded, extract it to */usr/local*, using a command like:

`tar -C /usr/local -xzf brisk-$VERSION-linux.tar.gz`

Where $VERSION is the version number you have downloaded (eg v1.0.0).

(This command may need to be run as root or using `sudo`)

Once extracted, add */usr/local/brisk/bin* to your PATH environment variable, do this with the command:

`export PATH=$PATH:/usr/local/brisk/bin`

## Useage

To start the repl use the command

`brisk repl`

This will start up the REPL and you should see a new BRISK shell start. You will know a new shell has started because you will see a `>>` in the terminal window. To exit the REPL type `exit`.

## Testing

BRISK is tested autonomously using both unit tests and system tests. These tests will be run automatically in Travis as part of a CI pipeline. **NOTE: code cannot be merged into the `dev` or `master` branches until the most recent Travis pipeline has passed**.

To run linting manually, navigate to the root of the repository and run `./run_static_analysis.sh`, this will build a docker image with the repository's code in it and run `golangci-lint` on the entire repository.

To run unit tests manually, navigate to the root of the repository and run `./run_unit_tests.sh`, this will build a docker image with the repository's code in it and run `go test` on the entire repository.

System tests are run using ROBOT framework, which in turn runs on python. in order to run ROBOT system tests, downlad python and install robotframework using pip. See https://robotframework.org/robotframework/latest/RobotFrameworkUserGuide.html#installation-instructions for instruction on downloading and installing both python and robot on your operating system. To run ROBOT system tests, navigate to the root of the repository and run `robot tests/scripts`. This will run the tests and produce html files that can be viewed in a browser for a detailed result of the tests.