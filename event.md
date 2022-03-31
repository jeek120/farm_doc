
# 一、base

## 1、Req

> ###  公共请求头

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |no|uint32| 请求id|服务器偶数递增，客户端奇数递增|
## 2、Resp

> ###  公共响应头

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |code|int32| 错误码|0-成功   >100000失败|
> |msg|string|||
## 3、Notice

> ###  公共通知头
# 二、blindbox

## 1、LandOpenBlindReq

> ###  开启盲盒请求

> ### 协议号
> `2271763913`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |id|int64| 盲盒id||
## 2、LandOpenBlindResp

> ###  开启盲盒结果

> ### 协议号
> `2271763914`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |t|int64| 盲盒类型||
> |land_id|int64| 土地id||
# 三、land

## 1、UnlockLandReq

> ###  解锁土地请求

> ### 协议号
> `539913625`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |land_id|int64|||
## 2、UnlockLandResp

> ###  解锁土地返回

> ### 协议号
> `539913625`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |land_id|int64|||
## 3、OperateLandReq

> ###  操作土地请求

> ### 协议号
> `554839481`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |land_id|int64|||
> |operate_type|enum|||
## 4、OperateLandResp

> ###  操作土地返回

> ### 协议号
> `554839482`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |land_id|int64|||
> |operate_type|enum|||
# 四、livestock

# 五、login

## 1、PlayerInfo

> ###  玩家信息实体

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |player_id|int64|||
> |name|string|||
> |nickname|string|||
> |age|int32|||
> |sex|int32|||
> |invertCode|string|||
> |mail|string|||
> |phone|string|||
> |avatar|string|||
> |money|int64|||
> |coin|int64|||
## 2、LoginAccPasswdReq

> ###  登录请求

> ### 协议号
> `3298260953`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |account|string|||
> |passwd|string|||
> |platform|int32|||
## 3、LoginAccPasswdResp

> ###  登录返回

> ### 协议号
> `3298260954`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |status|[Resp](#Resp)|||
> |player_id|int64|||
> |token|string|||
> |addr|string|||
## 4、RegisterMailReq

> ###  注册请求

> ### 协议号
> `4282977073`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |player_id|int64|||
> |name|string|||
> |nickname|string|||
> |age|int32|||
> |sex|int32|||
> |invertCode|string|||
> |mail|string|||
> |phone|string|||
> |avatar|string|||
> |passwd|string| 密码||
> |vcode|string| 验证码||
## 5、LoginTokenReq

> ###  token授权登录请求

> ### 协议号
> `1467768201`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |token|string|||
## 6、LoginTokenResp

> ###  token授权登录返回

> ### 协议号
> `1467768202`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |info|[PlayerInfo](#PlayerInfo)|||
## 7、LogoutResp

> ###  退出游戏

> ### 协议号
> `765576026`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |platform|int32| 1 - ANDROID PLATFORM = 1 << 0|2 - IOS<br/>4 - PC<br/>5 - H5<br/>6 - WEB|
## 8、RegisterMailResp

> ###  注册响应

> ### 协议号
> `4282977074`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |status|[Resp](#Resp)|||
## 9、GetWebTokenReq

> ###  换取Web的token请求

> ### 协议号
> `1348918585`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
## 10、GetWebTokenResp

> ###  换取Web的token响应

> ### 协议号
> `1348918586`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |token|string|||
## 11、PlayerInfoReq

> ###  根据玩家id请求玩家信息

> ### 协议号
> `2047325297`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |player_id|int64|||
## 12、PlayerInfoResp

> ###  玩家信息返回

> ### 协议号
> `2047325298`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |list|[PlayerInfo](#PlayerInfo)|||
# 六、market

## 1、GoodItem

> ###  商品列表

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |type_no|int32|||
> |name|int32| 如果字段为空，则从配置文件中获取名字||
> |price|int64| 售卖价格||
> |img|string| 预览图||
## 2、GoodListReq

> ###  市场列表请求

> ### 协议号
> `802074425`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |last_id|int64|||
## 3、GoodListResp

> ### 

> ### 协议号
> `802074426`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |item|[GoodItem](#GoodItem)|||
# 七、steal

## 1、StealReq

> ###  请求偷取列表

> ### 协议号
> `2191930257`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
## 2、Steal

> ###  偷取实体

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |ds|string|||
> |player_id|int64|||
## 3、StealResp

> ###  偷取列表返回

> ### 协议号
> `2191930257`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |list|[Steal](#Steal)|||
# 八、system

## 1、ResValue

> ###  配置值

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |s_val|string|||
> |i_val|int64|||
## 2、ResRecord

> ###  配置信息记录

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |record|[RecordEntry](#RecordEntry)|||
## 3、ResRecords

> ###  配置信息表

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |records|[ResRecord](#ResRecord)|||
## 4、ResConfigReq

> ###  获取资源配置

> ### 协议号
> `1056715489`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
## 5、ResConfigResp

> ###  获取配置文件返回信息

> ### 协议号
> `1056715490`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |data|[DataEntry](#DataEntry)|||
## 6、NoticeResp

> ###  消息提示
> 一般出现再屏幕的最上面，没有任何按钮

> ### 协议号
> `2611204202`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |position|int32| 位置|0-根据用户或客户端配置 1-top 2-bottom|
> |icon|int32| 0-没有icon||
> |msg|string|||
## 7、MessageResp

> ###  弹出框
> 有一个按钮的弹框，一般不会有倒计时

> ### 协议号
> `1269895266`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |code|int32| 错误码|1-显示ok，2-显示cancel|
> |icon|int32| 0-没有icon||
> |s|int32| 倒计时|当大于0时，会出现倒计时s秒后才可以点|
> |msg|string|||
## 8、ConfirmResp

> ###  确认框
> 有两个按钮的弹框，并且如果选择了Ok按钮，需要传回服务器

> ### 协议号
> `429080082`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |resp|[Resp](#Resp)|||
> |ok|string| ok 按钮的文字显示||
> |icon|int32| 0-没有icon||
> |x|int32| 倒计时|当大于0时，会出现倒计时x秒后才可以点|
> |msg|string|||
## 9、ConfirmOkBtnReq

> ###  确认点击确认框的ok按钮

> ### 协议号
> `1501745065`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
> |req_no|uint32| 请求id||
## 10、PingReq

> ###  ping

> ### 协议号
> `68964473`

> ### 协议详情
> |字段|类型|中文|描述|
> |---|---|---|---|
> |req|[Req](#Req)|||
# 协议号映射
```javascript
{
  "ConfirmOkBtnReq": 1501745065,
  "ConfirmResp": 429080082,
  "GetWebTokenReq": 1348918585,
  "GetWebTokenResp": 1348918586,
  "GoodListReq": 802074425,
  "GoodListResp": 802074426,
  "LandOpenBlindReq": 2271763913,
  "LandOpenBlindResp": 2271763914,
  "LoginAccPasswdReq": 3298260953,
  "LoginAccPasswdResp": 3298260954,
  "LoginTokenReq": 1467768201,
  "LoginTokenResp": 1467768202,
  "LogoutResp": 765576026,
  "MessageResp": 1269895266,
  "NoticeResp": 2611204202,
  "OperateLandReq": 554839481,
  "OperateLandResp": 554839482,
  "PingReq": 68964473,
  "PlayerInfoReq": 2047325297,
  "PlayerInfoResp": 2047325298,
  "RegisterMailReq": 4282977073,
  "RegisterMailResp": 4282977074,
  "ResConfigReq": 1056715489,
  "ResConfigResp": 1056715490,
  "StealReq": 2191930257,
  "StealResp": 2191930257,
  "UnlockLandReq": 539913625,
  "UnlockLandResp": 539913625
}
```