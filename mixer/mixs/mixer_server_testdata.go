package mixs

var checkResponse01 = `
Precondition:
  Status: 
    Code: "0" #对应http状态码200
    Message: "code is 200, ok"
  Duration: 2 #2s
  ReferencedAttributes: 
    AttributeMatches:
    - Name: "xxx"
      Condition: 0 #EXACT,找到匹配
Quotas: 
- Name: "quotas01"
  Status: 
    Code: "0" #对应http状态码200
    Message: "code is 200, ok"
  GrantedAmount: 2
  Duration: 2 #2s
  ReferencedAttributes:
    AttributeMatches:
    - Name: "test-mixer-ref-name"
      Condition: 2 #EXACT,找到匹配
`
