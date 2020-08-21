*** Settings ***
Library  Process
Library  OperatingSystem



*** Variables ***



*** Keywords ***
The output of command ${command} matches ${file}
    ${output} =  Run Process  ${command}  alias=repl    shell=True
    Create File  tests/testdata/output.txt  ${output.stdout}
    ${fileOut} =  Get File  tests/testdata/output.txt
    ${fileExpected} =  Get File  ${file}
    Should Be Equal As Strings  ${fileOut}  ${fileExpected}
    Should Be Equal As Integers  ${output.rc}  0