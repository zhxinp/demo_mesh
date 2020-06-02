// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
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
