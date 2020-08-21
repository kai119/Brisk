*** Settings ***
Documentation   Tests to verify that the main brisk program works correctly.
...             Including failing when it is supposed to fail and succeeding
...             when it is supposed to succeed
Metadata  Version 1.0.0
Library  Process

*** Variables ***



*** Test Cases ***
Brisk returns no error if help is parsed as an arguement
    [Tags]  011-test-brisk-with-arg
    ${output} =  Run Process  go run src/brisk/main.go help  shell=true
    Should Be Equal As Integers  ${output.rc}  0

Brisk returns error if help is called on a command that doesn't exist
    [Tags]  011-test-brisk-with-arg
    ${output} =  Run Process  go run src/brisk/main.go help fake  shell=true
    Should contain  ${output.stdout}  "fake" is not a command.
    Should Be Equal As Integers  ${output.rc}  1

Brisk returns error if no arguements are parsed to it
    [Tags]  012-test-brisk-without-arg
    ${output} =  Run Process  go run src/brisk/main.go  shell=true
    Should contain  ${output.stdout}  please enter a command.
    Should Be Equal As Integers  ${output.rc}  1

Brisk returns error if an arg that doesn't match a command is parsed to it
    [Tags]  013-test-brisk-bad-arg
    ${output} =  Run Process  go run src/brisk/main.go fake  shell=true
    Should contain  ${output.stdout}  "fake" is not a command.
    Should Be Equal As Integers  ${output.rc}  1

*** Keywords ***

    