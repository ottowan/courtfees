*** Setting ***
Library     Selenium2Library
Library     BuiltIn
Library     String

*** Variable ***
${url_courtfees}        https://fees.coj.go.th/courtfees
${option_group}         inlineRadioOptions
${type_havecap}         havecapital
${type_noncap}          noncapital
${type_mortgage}        mortgage
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
    Element Should Be Visible   ${xpath_result}
    ${value}=   Get Value       ${xpath_result}
    Should Be Equal             ${value}            ${expected_value}

Open Court Fee On Browser
    Open Browser                about:blank             chrome
    Go To                       ${url_courtfees}

Test Case
    [Arguments]                 ${type}                 ${test_v}           ${expected_value}
    Click Calculate             ${btn_calculate}
    Select Feetype Option       ${type} 
    Insert Feecapital           ${input_feecapital}     ${test_v}
    Click Calculate             ${btn_calculate}
    Click Calculate             ${btn_calculate}
    Verify Result               ${text_result}          ${expected_value}  


*** Test Cases ***
Test type noncapital check capital is 0 and expected value should be 200
    [Tags]      noncapital
    Open Court Fee On Browser
    Test Case   ${type_noncap}   0   200.00 

Test type mortgage check capital is 500,000 and expected value should be 5,000
    [Tags]      mortgage
    Test Case   ${type_mortgage}   500000   5,000.00 

Test type mortgage check capital is 20,000,000 and expected value should be 100,000
    [Tags]      mortgage
    Test Case   ${type_mortgage}   20000000   100,000.00 

Test type mortgage check capital is 100,000,000 and expected value should be 150,000
    [Tags]      mortgage
    Test Case   ${type_mortgage}   100000000   150,000.00

Test type mortgage check capital is 150,000 and expected value should be 1,000
    [Tags]      mortgage
    Test Case   ${type_mortgage}   150000   1,000.00 

Test type mortgage check capital is 50,000 and expected value should be 500
    [Tags]      mortgage
    Test Case   ${type_mortgage}   50000   500.00 

Test type havecapital check capital is 1,000,000 and expected value should be 20,000
    [Tags]      havecapital
    Test Case   ${type_havecap}   1000000   20,000.00 

Test type havecapital check capital is 40,000,000 and expected value should be 200,000
    [Tags]      havecapital
    Test Case   ${type_havecap}   40000000   200,000.00 

Test type havecapital check capital is 50,000,000 and expected value should be 200,000
    [Tags]      havecapital
    Test Case   ${type_havecap}   50000000   200,000.00 

Test type havecapital check capital is 300,001 and expected value should be 6,000
    [Tags]      havecapital
    Test Case   ${type_havecap}   300001   6,000.00 

Test type havecapital check capital is 50,000,001 and expected value should be 200,000
    [Tags]      havecapital
    Test Case   ${type_havecap}   50000001   200,000.00 

Test type havecapital check capital is 526,845,694 and expected value should be 676,845
    [Tags]      havecapital
    Test Case   ${type_havecap}   526845694   676,845.00 

Test type havecapital check capital is 250,000 and expected value should be 1,000
    [Tags]      havecapital
    Test Case   ${type_havecap}   250000   1,000.00 

Test type havecapital check capital is 20,000 and expected value should be 400
    [Tags]      havecapital
    Test Case   ${type_havecap}   20000   400.00 

Test type mortgage check capital is 300,001 and expected value should be 3,000
    [Tags]      mortgage
    Test Case   ${type_mortgage}   300001   3,000.00 

Test type mortgage check fee capital is 50,000,000 and expected value should be 100,000
    [Tags]      mortgage
    Test Case   ${type_mortgage}   50000000   100,000.00
    Close Browser