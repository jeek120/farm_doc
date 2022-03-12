
# base

---
## Req

> ### 公共请求头

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |no|uint32|请求id
 服务器偶数递增，客户端奇数递增||
## Resp

> ### 公共响应头

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |code|int32|||
> |msg|string|||
## Notice

> ### 公共通知头
# login

---
## LoginAccPasswdReq

> ### 登录请求

> ### 协议号
> `2241165165`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |account|string|||
> |passwd|string|||
## LoginAccPasswdResp

> ### 登录返回

> ### 协议号
> `2241165166`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |status|[Resp](#Resp)|||
> |token|string|||
> |addr|string|||
## RegisterMailReq

> ### 注册请求

> ### 协议号
> `2241137433`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |name|string|||
> |nickname|string|||
> |age|int32|||
> |sex|int32|||
> |invertCode|string|||
> |mail|string|||
> |phone|string|||
> |avatar|string|||
> |passwd|string|密码||
> |vcode|string|验证码||
## RegisterMailResp

> ### 注册响应

> ### 协议号
> `2241137434`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |status|[Resp](#Resp)|||
## GetWebTokenReq

> ### 换取Web的token请求

> ### 协议号
> `2241191325`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
## GetWebTokenResp

> ### 换取Web的token响应

> ### 协议号
> `2241191326`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |token|string|||
# system

---
## MessageResp

> ### 弹出框

> ### 协议号
> `2972717106`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |code|int32|错误码
 1-显示ok，2-显示cancel||
> |x|int32|倒计时
 当大于0时，会出现倒计时x秒后才可以点||
> |msg|string|||
## ConfirmResp

> ### 确认框

> ### 协议号
> `2972728842`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |ok|string|ok 按钮的文字显示||
> |x|int32|倒计时
 当大于0时，会出现倒计时x秒后才可以点||
> |msg|string|||
## ConfirmOkBtnReq

> ### 确认点击确认框的ok按钮

> ### 协议号
> `2972766165`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |req_no|uint32|请求id||