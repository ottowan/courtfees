*** Setting ***
Library     Selenium2Library
Library     BuiltIn
Library     String

*** Variable ***
${url_arbitration}      https://fees.coj.go.th/arbitration
${option_group}         inlineRadioOptions
${type_single_person}   personOne
${type_more_persons}    personMoreThanOne

${input_feecapital}     //*[@id="feeCapital"]
${btn_calculate}        //*[@id="calculate"]
${text_result}          //*[@id="result"]

*** Keywords ***
Select Feetype Option
    [Arguments]                 ${option_type}
    Select Radio Button         ${option_group}     ${option_type}

Insert Feecapital
    [Arguments]                 ${xpath_value}  ${value}
    Element Should Be Visible   ${xpath_value}
    Input Text                  ${xpath_value}  ${value}

Click Calculate
    [Arguments]                 ${xpath_btn}
    Element Should Be Visible   ${xpath_btn}
    Click Element               ${xpath_btn}

Verify Result
    [Arguments]                 ${xpath_result}     ${expected_value}
    Wait Until Element Is Visible           ${xpath_result}
    Element Should Be Visible               ${xpath_result}
    ${value}=   Get Text        ${xpath_result}
    Should Be Equal             ${value}            ${expected_value}

Open Court Fee On Browser
    Open Browser                about:blank             chrome
    Go To                       ${url_arbitration}

Test Case
    [Arguments]                 ${type}                 ${test_v}           ${expected_value}
    
    Select Feetype Option       ${type} 
    Insert Feecapital           ${input_feecapital}     ${test_v}
    Click Calculate             ${btn_calculate}
    Click Calculate             ${btn_calculate}
    Verify Result               ${text_result}          ${expected_value}  


*** Test Cases ***
Test arbitration check type single person input captital by 0 and expected value should be 6,000
    [Tags]  ${type_single_person}
    Open Court Fee On Browser
    Test Case   ${type_single_person}    0   6000

Test arbitration check type single person input captital by 2,000,000 and expected value should be 30,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    2000000   30000

Test arbitration check type single person input captital by 3,500,000 and expected value should be 45,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    3500000   45000

Test arbitration check type single person input captital by 5,000,000 and expected value should be 60,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    5000000   60000

Test arbitration check type single person input captital by 7,500,000 and expected value should be 80,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    7500000   80000

Test arbitration check type single person input captital by 10,000,000 and expected value should be 100,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    10000000   100000

Test arbitration check type single person input captital by 15,000,000 and expected value should be 130,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    15000000   130000

Test arbitration check type single person input captital by 20,000,000 and expected value should be 160,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    20000000   160000

Test arbitration check type single person input captital by 27,500,000 and expected value should be 190,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    27500000   190000

Test arbitration check type single person input captital by 35,000,000 and expected value should be 220,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    35000000   220000

Test arbitration check type single person input captital by 47,500,000 and expected value should be 245,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    47500000   245000

Test arbitration check type single person input captital by 50,000,000 and expected value should be 250,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    50000000   250000

Test arbitration check type single person input captital by 75,000,000 and expected value should be 275,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    75000000   275000

Test arbitration check type single person input captital by 100,000,000 and expected value should be 300,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    100000000   300000

Test arbitration check type single person input captital by 250,000,000 and expected value should be 375,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    250000000   375000

Test arbitration check type single person input captital by 500,000,000 and expected value should be 500,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    500000000   500000

Test arbitration check type single person input captital by 526,845,694 and expected value should be 510,738
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    526845694   510738  

Test arbitration check type single person input captital by 1,000,000,000 and expected value should be 700,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    1000000000   700000

Test arbitration check type single person input captital by 1,500,000,000 and expected value should be 850,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    1500000000   850000

Test arbitration check type single person input captital by 2,000,000,000 and expected value should be 1,000,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    2000000000   1000000

Test arbitration check type single person input captital by 10,000,000,000 and expected value should be 2,600,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    10000000000   2600000

Test arbitration check type single person input captital by 1,000,000 and expected value should be 30,000
    [Tags]  ${type_single_person}
    Test Case   ${type_single_person}    1000000   30000

Test arbitration check type more persons input captital by 0 and expected value should be 30,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    0   30000

Test arbitration check type more persons input captital by 2,000,000 and expected value should be 60,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    2000000   60000

Test arbitration check type more persons input captital by 3,500,000 and expected value should be 90,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    3500000   90000

Test arbitration check type more persons input captital by 5,000,000 and expected value should be 120,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    5000000   120000

Test arbitration check type more persons input captital by 7,500,000 and expected value should be 160,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    7500000   160000

Test arbitration check type more persons input captital by 10,000,000 and expected value should be 200,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    10000000   200000

Test arbitration check type more persons input captital by 15,000,000 and expected value should be 260,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    15000000   260000

Test arbitration check type more persons input captital by 20,000,000 and expected value should be 320,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    20000000   320000

Test arbitration check type more persons input captital by 27,500,000 and expected value should be 380,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    27500000   380000

Test arbitration check type more persons input captital by 35,000,000 and expected value should be 440,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    35000000   440000

Test arbitration check type more persons input captital by 47,500,000 and expected value should be 490,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    47500000   490000

Test arbitration check type more persons input captital by 50,000,000 and expected value should be 500,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    50000000   500000

Test arbitration check type more persons input captital by 75,000,000 and expected value should be 550,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    75000000   550000

Test arbitration check type more persons input captital by 100,000,000 and expected value should be 600,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    100000000   600000

Test arbitration check type more persons input captital by 250,000,000 and expected value should be 750,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    250000000   750000

Test arbitration check type more persons input captital by 500,000,000 and expected value should be 1,000,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    500000000   1000000

Test arbitration check type more persons input captital by 526,845,694 and expected value should be 1,021,476
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    526845694   1021476  

Test arbitration check type more persons input captital by 1,000,000,000 and expected value should be 1,400,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    1000000000   1400000

Test arbitration check type more persons input captital by 1,500,000,000 and expected value should be 1,700,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    1500000000   1700000

Test arbitration check type more persons input captital by 2,000,000,000 and expected value should be 2,000,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    2000000000   2000000

Test arbitration check type more persons input captital by 10,000,000,000 and expected value should be 5,200,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    10000000000   5200000

Test arbitration check type more persons input captital by 1,000,000 and expected value should be 60,000
    [Tags]  ${type_more_persons}
    Test Case   ${type_more_persons}    1000000   60000
    Close Browser