*** Settings ***
Documentation   Tests to verify that the repl program works correctly.
...             It should take in any text and tokenize it, outputting each
...             token on its own line
Metadata  Version 1.0.0
Library  Process
Library  OperatingSystem
Resource  ../resource/resource.robot

Test Teardown  Remove File  tests/testdata/output.txt

*** Variables ***



*** Test Cases ***
REPL can start correctly
    [Tags]  021-test-repl-start
    ${command} =  Convert To String  test
    The REPL is run with the command ${command}

REPL correctly converts input to list of tokens
    [Tags]  022-test-repl-output
    ${command} =  Convert To String  go run src/brisk/main.go repl -c "if (5 < 10) { return true }"
    The output of command ${command} matches tests/testdata/022-expected-output.txt

REPL exit keyword works as expected
    [Tags]  023-test-repl-exit
    ${command} =  Convert To String  go run src/brisk/main.go repl -c exit
    The output of command ${command} matches tests/testdata/023-expected-output.txt

REPL shows help page
    [Tags]  024-test-repl-help
    ${command} =  Convert To String  go run src/brisk/main.go help repl
    The output of command ${command} matches tests/testdata/024-expected-output.txt



*** Keywords ***
The REPL is run with the command ${command}
    ${output} =  Run Process  go run src/brisk/main.go repl -c ${command}  alias=repl    shell=True
    Should Contain  ${output.stdout}  Hello! This is the BRISK programming language!
    