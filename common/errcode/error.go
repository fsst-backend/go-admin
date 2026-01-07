package errcode

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// --------------------------这个结构体只有在生成的handler文件中使用--------------------------
type InvalidParams struct {
	Msg string
}

func (e *InvalidParams) Error() string {
	return e.Msg
}

func NewInvalidParams(msg string) *InvalidParams {
	return &InvalidParams{Msg: msg}
}

// --------------------------这个结构体只有在生成的handler文件中使用--------------------------

// ErrCode 错误代码
type ErrCode int

const (
	// OK OK=0 don not export
	OK ErrCode = iota
)

const RpcSuccess = 0

// 10000为请求错误
const (
	RespCodeBadRequest       ErrCode = 10000 // 请求错误
	RespCodeBadChannel       ErrCode = 10001 // 渠道错误
	RespCodeBadRepo          ErrCode = 10002 // 仓库错误
	RespCodeBadModule        ErrCode = 10003 // 模块错误
	RespCodeBadPostID        ErrCode = 10004 // 帖子ID错误
	RespCodeBadGroupID       ErrCode = 10005 // 群组ID错误
	RespCodeBadQuery         ErrCode = 10006 // 查询错误
	RespCodeBadParamValue    ErrCode = 10007 // 参数值错误
	RespCodeNotFoundSource   ErrCode = 10008 // 未找到数据源
	RespCodeBadRequestHeader ErrCode = 10009 // 请求头错误

	RespCodeTelAndEmailEmpty       ErrCode = 10010 // 手机号和邮箱都为空
	RespCodeInvalidEmailFormat     ErrCode = 10011 // 邮箱格式不正确
	RespCodeInvalidTelFormat       ErrCode = 10012 // 手机号格式不正确
	RespCodeInvalidPassword        ErrCode = 10014 // 密码不能为空
	RespCodeInvalidNameFormat      ErrCode = 10015 // 名字格式不正确
	RespCodeInvalidDocumentNumber  ErrCode = 10016 // 证件号码格式不正确
	RespCodePasswordMismatch       ErrCode = 10017 // 两次密码不一致
	RespCodeTelEmpty               ErrCode = 10018 // 手机号不能为空
	RespCodeEmailEmpty             ErrCode = 10019 // 手机号不能为空
	RespCodeInvalidValidDateFormat ErrCode = 10020 // 有效期格式错误

	// SMS 验证码相关错误码
	RespCodeSMSSuccess                   ErrCode = 10021 // SMS 操作成功
	RespCodeSMSOperationFailed           ErrCode = 10022 // SMS 操作失败，系统繁忙请稍后再试
	RespCodeSMSInternalError             ErrCode = 10023 // SMS 内部错误，系统繁忙请稍后再试
	RespCodeSMSConfigError               ErrCode = 10024 // SMS 配置错误，系统繁忙请稍后再试
	RespCodeSMSTemplateEmpty             ErrCode = 10025 // SMS 模板错误，系统繁忙请稍后再试
	RespCodeSMSBusinessError             ErrCode = 10026 // SMS 业务错误
	RespCodeSMSFrequencyLimit            ErrCode = 10027 // SMS 发送频率限制
	RespCodeSMSDataError                 ErrCode = 10028 // SMS 数据错误，系统繁忙请稍后再试
	RespCodeSMSVerifyCodeExpired         ErrCode = 10029 // 验证码已过期
	RespCodeSMSVerifyCodeError           ErrCode = 10030 // 验证码错误
	RespCodeSMSVerifyCodeUsed            ErrCode = 10031 // 验证码已使用
	RespCodeSMSVerifyCodeInvalid         ErrCode = 10032 // 验证码无效
	RespCodeSMSVerifyCodeNotUsed         ErrCode = 10033 // 验证码未使用
	RespCodeSMSGoogleJWTVerifyFailed     ErrCode = 10034 // 谷歌JWT验证失败
	RespCodeSMSGoogleJWTFailedToOpenAuth ErrCode = 10035 // 谷歌JWT打开授权json失败
)

// 20000为账号错误
const (
	RespCodeAccountTelError                 ErrCode = 20000 // 手机号错误
	RespCodeAccountServerError              ErrCode = 20001 // 服务器错误
	RespCodeAccountCodeError                ErrCode = 20002 // 验证码错误
	RespCodeAccountExistError               ErrCode = 20003 // 账号已存在
	RespCodeAccountNoExistError             ErrCode = 20004 // 账号不存在
	RespCodeAccountDelError                 ErrCode = 20005 // 删除账号错误
	RespCodeAccountQueryError               ErrCode = 20006 // 查询账号错误
	RespCodeAccountParseError               ErrCode = 20007 // 解析账号错误
	RespCodeAccountPasswordError            ErrCode = 20008 // 密码错误
	RespCodeAccountInvalidError             ErrCode = 20009 // 账号无效
	RespCodeAccountRatelimitError           ErrCode = 20010 // 请求频率限制
	RespCodeAccountDiffPasswordError        ErrCode = 20011 // 密码不一致
	RespCodeAccountUpdateUserError          ErrCode = 20012 // 更新用户信息错误
	RespCodeAccountUpdatePasswordError      ErrCode = 20013 // 更新密码错误
	RespCodeAccountUpdateAvatarError        ErrCode = 20014 // 更新头像错误
	RespCodeAccountThirdCodeError           ErrCode = 20015 // 第三方验证码错误
	RespCodeAccountThirdAccessTokenError    ErrCode = 20016 // 第三方AccessToken错误
	RespCodeAccountThirdUserInfoError       ErrCode = 20017 // 第三方用户信息错误
	RespCodeAccountThirdLoginError          ErrCode = 20018 // 第三方登录错误
	RespCodeAccountTelNoBindError           ErrCode = 20019 // 手机号未绑定
	RespCodeAccountTelBindError             ErrCode = 20020 // 手机号绑定错误
	RespCodeAccountTelBoundError            ErrCode = 20021 // 手机号已绑定
	RespCodeAccountTelUnBindError           ErrCode = 20022 // 手机号解绑错误
	RespCodeAccountMPApplicationCodeError   ErrCode = 20023 // 小程序验证码错误
	RespCodeAccountMPApplicationActiveError ErrCode = 20024 // 小程序激活错误
	RespCodeAccountMPApplicationLogoError   ErrCode = 20025 // 小程序Logo错误
	RespCodeAccountMPPublishContentError    ErrCode = 20026 // 小程序发布内容错误
	RespCodeAccountMPNoPermissionError      ErrCode = 20027 // 小程序无权限
	RespCodeAccountMPGetFollowError         ErrCode = 20028 // 获取小程序关注列表错误
	RespCodeAccountMPFollowedError          ErrCode = 20029 // 已关注小程序
	RespCodeAccountMPFollowError            ErrCode = 20030 // 关注小程序错误
	RespCodeAccountGetMPFollowListError     ErrCode = 20031 // 获取小程序关注列表错误
	RespCodeAccountGetMPTopicListError      ErrCode = 20032 // 获取小程序话题列表错误
	RespCodeAccountGetMPSVideoListError     ErrCode = 20033 // 获取小程序视频列表错误
	RespCodeAccountInvitationCodeError      ErrCode = 20034 // 邀请码错误
	RespCodeAccountGetTopicsError           ErrCode = 20035 // 获取话题错误
	RespCodeAccountForbiddenError           ErrCode = 20036 // 账号被禁用
	RespCodeAccountTelBindedError           ErrCode = 20037 // 手机号已绑定
	RespCodeAccountApprovalAvatarError      ErrCode = 20038 // 头像审核错误
	RespCodeAccountApprovalAliasError       ErrCode = 20039 // 昵称审核错误
	RespCodeAccountApprovalSignError        ErrCode = 20040 // 签名审核错误
	RespCodeAccountNoTelError               ErrCode = 20041 // 无手机号
	RespCodeAccountAuthThreeError           ErrCode = 20042 // 三方认证错误
	RespCodeAccountAuthFiveError            ErrCode = 20043 // 五方认证错误
	RespCodeAccountLogoutError              ErrCode = 20044 // 登出错误
	RespCodeAccountDelUserError             ErrCode = 20045 // 删除用户错误
	RespCodeAccountGetUserAddrError         ErrCode = 20046 // 获取用户地址错误
	RespCodeAccountAddUserAddrError         ErrCode = 20047 // 添加用户地址错误
	RespCodeAccountUpdateDefaultAddrError   ErrCode = 20048 // 更新默认地址错误
	RespCodeAccountGetAddrInfoError         ErrCode = 20049 // 获取地址信息错误
	RespCodeAccountDeleteAddrError          ErrCode = 20050 // 删除地址错误
	RespCodeAccountLoginOneStepError        ErrCode = 20051 // 一步登录错误
	RespCodeAccountVerifyLoginError         ErrCode = 20052 // 验证登录错误
	RespCodeAccountUpdateBindError          ErrCode = 20053 // 更新绑定错误
	RespCodeAccountInvideCodeError          ErrCode = 20054 // 邀请码错误
	RespCodeAccountAliasLengthError         ErrCode = 20055 // 昵称长度错误
	RespCodeRepeatedBindCardError           ErrCode = 20056 // 重复绑定卡片
	RespCodeRepeatedModifyAliasError        ErrCode = 20057 // 重复修改昵称
	RespCodeRepeatedModifyAvatarError       ErrCode = 20058 // 重复修改头像
	RespCodeRepeatedModifySignError         ErrCode = 20059 // 重复修改签名
	RespCodeReportApplyOrderSearchError     ErrCode = 20060 // 申请订单搜索错误
	RespCodeAccountFindPasswordNoTelError   ErrCode = 20061 // 找回密码无手机号
	RespCodeAccountFindPasswordTicketError  ErrCode = 20062 // 找回密码凭证错误
	RespCodeAccountGetCodeError             ErrCode = 20063 // 获取验证码错误
	RespCodeAccountPasswordLenError         ErrCode = 20064 // 密码长度错误
	RespCodeAccountNotLoginError            ErrCode = 20065 // 未登录错误
	RespCodeFailedSwitchAccounts            ErrCode = 20066 // 切换账号失败
	RespCodeRequestTradePlatformError       ErrCode = 20067 // 请求交易平台错误
	RespCodeTransactionAccountError         ErrCode = 20068 // 交易账号错误
	RespCodeActionTypeError                 ErrCode = 20069 // 交易类型错误
	RespCodeTransactionGetSymbolError       ErrCode = 20070 // 获取交易品种错误
	RespCodeUnsupportedVCodeType            ErrCode = 20071 // 不支持的验证码类型
	RespCodeVerifyCodeInvalid               ErrCode = 20072 // 验证码未验证或已过期
	RespCodeAccountEmailBindedError         ErrCode = 20073 // 邮箱已绑定
	RespCodeGetIpLocationFailed             ErrCode = 20074 // 获取IP归属地失败
	RespCodeOcrServiceError                 ErrCode = 20075 // OCR服务异常
	RespCodeAccountAgeNotAllowed            ErrCode = 20076 // 年龄不满足要求
	RespCodeAccountIdCardInvalid            ErrCode = 20077 // 身份证号码格式无效
	RespCodeAccountIdCardBirthInvalid       ErrCode = 20078 // 身份证出生日期无效
	RespCodeThreeFactorsAuthFailed          ErrCode = 20079 // 身份证三要素认证失败
	RespCodeFourFactorsAuthFailed           ErrCode = 20080 // 银行卡四要素认证失败
	RespCodeInviteCodeUsedError             ErrCode = 20081 // 邀请码已被使用
	RespCodeTransactionGetPositionError     ErrCode = 20082 // 获取持仓错误
	RespCodeAccountNoPermissionError        ErrCode = 20083 // 账号无权限
	RespCodeInvalidUserStatus               ErrCode = 20084 // 无效的用户状态
	RespCodeTransactionAccountNotActive     ErrCode = 20085 // 交易账号未激活
	RespCodeOcrServiceFirstError            ErrCode = 20086 // OCR首次识别失败 -无法读取您的证件中的信息，请确保证件清晰可见。
	RespCodeOcrServiceRetryError            ErrCode = 20087 // OCR重试识别失败 -无法读取您的证件中的信息，您可手动填写信息。
	RespCodeOcrServiceLimitError            ErrCode = 20088 // OCR调用次数超限
	RespCodeTwoFactorsAuthFailed            ErrCode = 20089 // 身份证二要素认证失败
	ErrCommissionWithdrawAmountTooLarge     ErrCode = 20090 //  提现金额过大
	RespCodeSymbolNotInSessionTradeError    ErrCode = 20091 //  未在交易时间段内
	RespCodeSymbolHasQuotesAvailableError   ErrCode = 20092 //  该品种没有可用的报价
	RespCodeSymbolInHolidayError            ErrCode = 20093 //  该品种在交易节假日内
	RespCodeFundsPasswordError              ErrCode = 20094 // 资金密码错误
	RespCodeSlideVerifyError                ErrCode = 20095 // 滑动校验错误
	RespCodeDuplicateIdentityNumber         ErrCode = 20096 // 证件已被使用
	RespCodeNewPasswordSameOldPassword      ErrCode = 20097 // 新密码与旧密码相同
	RespCodeGoogleEmailAlreadyUsed          ErrCode = 20098 // 该谷歌账号邮箱已被注册
	RespCodeAppleEmailAlreadyUsed           ErrCode = 20099 // 该苹果账号邮箱已被注册
	RespCodeGoogleEmailNotExist             ErrCode = 20100 // 该谷歌账号没有邮箱不能注册登录
	RespCodeAppleEmailNotExist              ErrCode = 20101 // 该苹果账号没有邮箱不能注册登录
)

const (
	RespCodeTemporaryAuthCodeInvalid ErrCode = 21000 // 临时授权码无效
	RespCodeTemporaryAuthCodeTimeout ErrCode = 21001 // 临时授权码超时
)

// 30000为互动类型错误
const (
	RespCodeCollectionExistError                   ErrCode = 30000 // 收藏已存在
	RespCodeCollectionNoExistError                 ErrCode = 30001 // 收藏不存在
	RespCodeCollectionServerError                  ErrCode = 30002 // 收藏服务错误
	RespCodeCollectionDelError                     ErrCode = 30003 // 删除收藏错误
	RespCodeGetCollectionError                     ErrCode = 30004 // 获取收藏错误
	RespCodeIntegrationCreateError                 ErrCode = 30005 // 积分创建错误
	RespCodeIntegrationTypeError                   ErrCode = 30006 // 积分类型错误
	RespCodeMessageGetFollowError                  ErrCode = 30007 // 获取关注消息错误
	RespCodeMessageGetStarError                    ErrCode = 30008 // 获取点赞消息错误
	RespCodeMessageGetCommentError                 ErrCode = 30009 // 获取评论消息错误
	RespCodeMessageGetNoticeError                  ErrCode = 30010 // 获取通知消息错误
	RespCodeMessageAppInfoNoExistError             ErrCode = 30011 // 应用信息不存在
	RespCodeOpenUserInfoFail                       ErrCode = 30012 // 获取用户信息失败
	RespCodeIntegrationRecordNotExistError         ErrCode = 30013 // 积分记录不存在
	RespCodeIntegrationNotEnough                   ErrCode = 30014 // 积分不足
	RespCodeIntegrationUpdateFail                  ErrCode = 30015 // 积分更新失败
	RespCodeIntegrationLogFail                     ErrCode = 30016 // 积分日志失败
	RespCodeLocalLiveUserMutedError                ErrCode = 30017 // 本地直播用户被禁言
	RespCodeNotMeetConditionError                  ErrCode = 30018 // 不满足条件
	RespCodeLevelNotFoundError                     ErrCode = 30019 // 等级不存在
	RespCodePositionNotClosedError                 ErrCode = 30020 // 仓位未平仓
	RespCodeFollowingError                         ErrCode = 30021 // 跟单失败
	RespCodeFollowerNotAllowedError                ErrCode = 30022 // 带单员不可以跟单
	RespCodeFollowingNotCanceledError              ErrCode = 30023 // 存在未取消的跟单
	RespCodeSignalOrderNoModificationsAllowedError ErrCode = 30024 // 带单订单不允许修改
	RespCodeFollowingTimeoutError                  ErrCode = 30025 // 跟单超时
)

// 40000为评论错误
const (
	RespCodeIllegalCommentError       ErrCode = 40000 // 非法评论
	RespCodeGetCommentError           ErrCode = 40001 // 获取评论错误
	RespCodeCommentMediaRelationError ErrCode = 40002 // 评论媒体关系错误
	RespCodeSubmitCommentError        ErrCode = 40003 // 提交评论错误
	RespCodeIllegalStarError          ErrCode = 40004 // 非法点赞
	RespCodeStaredError               ErrCode = 40005 // 已点赞
	RespCodeCommentClosedError        ErrCode = 40006 // 评论已关闭
	RespCodeCommentIntervalError      ErrCode = 40007 // 评论间隔错误
	RespCodeCommentBlackIpError       ErrCode = 40008 // 评论IP黑名单
)

// 50000为系统相关错误
const (
	RespCodeCreateFeedbackError  ErrCode = 50000 // 创建反馈错误
	RespCodeCreateReportError    ErrCode = 50001 // 创建报告错误
	RespCodeFeedbackContactError ErrCode = 50002 // 反馈联系方式错误
	RespCodeReportContactError   ErrCode = 50003 // 报告联系方式错误
	RespCodeReportPictureError   ErrCode = 50004 // 报告图片错误
	RespCodeReportRepeatError    ErrCode = 50005 // 重复报告错误
	RespCodeRateLimitError       ErrCode = 50006 // 请求频率限制错误
)

// 60000为oss video错误
const (
	RespCodeCreateOSSVideoError   ErrCode = 60000 // 创建OSS视频错误
	RespCodeQueryOSSVideoError    ErrCode = 60001 // 查询OSS视频错误
	RespCodeParamNilOSSError      ErrCode = 60002 // OSS参数为空
	RespCodeNilOSSVideoError      ErrCode = 60003 // OSS视频为空
	RespCodeCreateMPVideoError    ErrCode = 60004 // 创建小程序视频错误
	RespCodeOSSCallbackParseError ErrCode = 60005 // OSS回调解析错误
	RespCodeOSSUpdateStateError   ErrCode = 60006 // OSS状态更新错误
	RespCodeQueryMPVideoError     ErrCode = 60007 // 查询小程序视频错误
	RespCodeOSSCreateBucketError  ErrCode = 60008 // 创建OSS Bucket错误
	RespCodeOSSError              ErrCode = 60009 // OSS错误
	RespOSSNoLoginError           ErrCode = 60010 // OSS未登录错误
)

// 70000为搜索相关错误
const (
	RespCodeSearchError        ErrCode = 70000 // 搜索错误
	RespCodeShareError         ErrCode = 70001 // 分享错误
	RespCodeSearchKeywordError ErrCode = 70002 // 搜索关键词错误
)

const (
	RespCodeReserveError     ErrCode = 80000 // 预约错误
	RespCodeReserveTimeError ErrCode = 80001 // 预约时间错误
)

const (
	RespCodeBookUnAvailableError   ErrCode = 90000 // 书籍不可用
	RespCodeBookTypeNotFound       ErrCode = 90001 // 书籍类型未找到
	RespCodeUnBookUnAvailableError ErrCode = 90002 // 未预订书籍不可用
)

// 100000为内容返回（不是错误）
const (
	RespCodeBookingMessage ErrCode = iota + 100000
)

const (
	ErrLastErrCode = 110000 // 客户端错误码最大值，不能超过此值
)

// 通用错误码
const (
	RpcErrCodeBegin                   = 110000 // rpc 错误专用返回,同时可使用在后台业务和客户端业务
	RpcErrCodeInvalidParams           = 110001 // 无效参数
	RpcErrCodeServerError             = 110002 // "服务器开小差了，请稍后再试"
	RpcErrCodeCreateRuleInvalidParam  = 110003 // "生成规则参数错误"
	RpcErrCodeCreateRuleError         = 110004 // "生成规则失败"
	RpcErrCodeGenerateIDError         = 110005 // "生成 ID 失败"
	RpcErrCodeAuditOrderNotExist      = 110006 // "订单审核信息不存在"
	RpcErrCodeStringFromDecimalError  = 110007 // "sting转decimal错误"
	RpcErrCodeGenerateIDRangeConflict = 110008 // "ID 生成范围冲突"
	RpcErrCodePhoneNumberNotExist     = 110009 // "手机号不存在"
	RpcErrCodeInviteCodeExpired       = 110010 // "邀请码已失效"
	RpcErrCodeInviteCodeInvalid       = 110011 // "邀请码无效"
	RpcErrCodeAuthError               = 110012 // auth认证失败
	RpcErrCodeBatchAddError           = 110013 // 批量添加错误
	RpcErrCodeNoPermission            = 110014 // "无权限"
	RpcErrCodeServerConfigError       = 110015 // "服务器配置错误"
	RpcErrCodeUserLoggedInOtherDevice = 110016 // "该用户已在其他设备登录"
	RpcErrCodeServerBusy              = 110017 // "服务器繁忙,请稍后再试"
)

// crmuser 错误码
const (
	RpcErrCodeCrmUserNotExist                   = 120000 // "crm用户不存在"
	RpcErrCodeVerificationReviewNotExist        = 120001 // "审核记录不存在"
	RpcErrCodeDataError                         = 120002 // "数据错误,请联系管理员"
	RpcErrCodeKYCLimitExceeded                  = 120003 // "kyc审核次数已达上限"
	RpcErrCodeSubmitIdentityInfoNotSameRegister = 120004 // "审核未通过"
	RpcErrCodeHasCertified                      = 120005 // "已认证"
	RpcErrCodeUserDiabled                       = 120006 // "用户已被禁用"
	RpcErrCodePasswordError                     = 120007 // "密码错误"
	RpcErrCodeDefaultUserLevelIsNotSet          = 120008 // "默认等级未设置"
	RpcErrCodeAccountHasExist                   = 120009 // "账户已存在"
	RpcErrCodeBankCardHasDeleted                = 120010 // "银行卡已删除"
	// 用户注销相关错误码
	RpcErrCodeUserClosed                = 120011 // RpcErrCodeUserClosed 已注销的账户不能操作
	RpcErrCodeWalletBalanceNotZero      = 120012 // "钱包余额不为零"
	RpcErrCodeRealAccountBalanceNotZero = 120013 // "真实账户余额不为零"
	RpcErrCodeInvalidUserStatus         = 120014 // "用户状态不允许注销"
	RpcErrCodeDeactivationFailed        = 120015 // "用户注销失败"
	RpcErrCodeWalletBalanceNotEnough    = 120016 // "钱包余额不足"

	// 国家配置相关错误码
	RpcErrCodeCountryNotAllowRegister = 120017 // 国家不允许注册

	// 权益相关错误码
	RpcErrCodeBenefitNotExists    = 120018 // 权益不存在
	RpcErrCodeCreateBenefitFailed = 120019 // 创建权益失败
	RpcErrCodeUpdateBenefitFailed = 120020 // 更新权益失败
	RpcErrCodeDeleteBenefitFailed = 120021 // 删除权益失败
	RpcErrCodeGetBenefitFailed    = 120022 // 获取权益失败

	// 钱包相关错误码
	RpcErrCodeRealWalletNotExist = 120101 // 钱包不存在

	// 用户等级相关错误码
	RpcErrCodeUserLevelAdjusting = 120201 // 用户等级正在调整中
)

// merchant 错误码
const (
	RpcErrCodeMerchantNotExist                = 130000 // "商户不存在"
	RpcErrCodeCreditAmountinsufficientError   = 130001 // "授信额度不足"
	RpcErrCodeRepeatedAdjustCreditAmountError = 130002 // "重复调整授信额度"
	RpcErrCodeCreditAccountNotExist           = 130003 // "授信账户不存在"
	RpcErrCodeAccountSplitNotExist            = 130004 // "分账记录不存在"
	RpcErrCodeFundTransFerInNotExist          = 130005 // "入账记录不存在"
	RpcErrCodeFundTransFerOutNotExist         = 130006 // "出账记录不存在"
	RpcErrCodeMerchantDisabled                = 130007 // "用户已被禁用
)

// Finance 错误码
const (
	RpcErrCodeFinanceOrderIDError             = 140000 // "订单ID不存在"
	RpcErrCodeRepeatedInOutCurrencyError      = 140001 // "重复出入金货币错误"
	RpcErrCodeFinanceRateExitsError           = 140002 // "配置已存在"
	RpcErrCodeAmountNOTLessZeroError          = 140003 // "金额不能小于0"
	RpcErrCodeDepositApplyAlreadyAuditedError = 140004 // "已审核"
	RpcErrCodeHttpUrlNotIsHttp                = 140005 // "不是http url"
	RpcErrCodeRateNotExist                    = 140006 // "汇率不存在"
	RpcErrAccountNotExist                     = 140007 // "账户不存在"
	RpcErrCodeBankNotAuthentication           = 140008 // "银行卡未认证"
	RpcErrCodeAccountNotBelongToUser          = 140009 // "账户不属于当前用户"
	RpcErrCodeFeeGreaterThanOrderAmount       = 140010 // "手续费大于订单金额"
	RpcErrCodeEWalletHasDeleted               = 140011 // "电子钱包已删除"
	RpcErrInsufficientBalance                 = 140012 // 余额不足
	RpcErrCodeEWalletNotAuthentication        = 140013 // "电子钱包未认证"
	RpcErrCodeUserIdentityNotPass             = 140014 // "用户未实名认证"
	RpcErrCodeAccountNotBeenDebited           = 140015 // "账号未扣款"
	RpcErrCodeRealAmountNotEnough             = 140016 // "金额不足"
)

// transaction 错误码
const (
	RpcErrCodeRequestTradePlatformError   = 150000 // "请求交易平台失败"
	RpcErrCodeRequestTradePlatformTimeout = 150001 // "请求交易平台超时"
	RpcErrCodeTradeServerAddrIsEmpty      = 150002 // "交易平台地址为空"
	RpcErrCodeSymbolNotExist              = 150003 // "交易组不存在"
	RpcErrCodeTraderSymbolNotFoundSymbol  = 150004 // "交易品种不存在"
	RpcErrCodeIncorrectAccount            = 150005 // "交易账户错误"
	RpcErrRecordPendingOrder              = 150006 // "记录挂单失败"
	RpcErrRecordPositionOrder             = 150007 // "记录持仓订单失败"
	RpcErrCodeOrderNotFound               = 150008 // "订单不存在"
	RpcErrCodeTradeAccountNotActivated    = 150009 // "交易账户未激活"
	RpcErrCodeTradeAccountActivated       = 150010 // "交易账户已激活"
	RpcErrCodeStopProfitStopLossError     = 150011 // "止盈止损错误"
	RpcErrCodeStopProfitError             = 150012 // "止盈错误"
	RpcErrCodeStopLossError               = 150013 // "止损错误"
	RpcErrCodeVolumeError                 = 150014 // "交易量错误"
)

// copytrading 错误码
const (
	// 带单跟单
	RpcErrCodeCopyTradingManagementLevelNotEmpty = 160000 // "带单员等级不为空"
	RpcErrCodeNotFollow                          = 160001 // "未跟单"
	RpcErrCodeFollowError                        = 160002 // "跟单失败"
	RpcErrCodeFollowersCountExceeded             = 160003 // "跟单人数已满"
	RpcErrCodeRatioTooLarge                      = 160004 // "比例过大"
	RpcErrCodeAlreadyFollowed                    = 160005 // "已跟单"
	RpcErrCodeSignalProviderNotExist             = 160006 // "带单员不存在"
	RpcErrCodeAccountTypeError                   = 160007 // "账号类型错误"
	RpcErrCodeFollowingNotCanceledError          = 160008 // 存在未取消的跟单
	RpcErrCodeCopyTradingOrderCannotUpdate       = 160009 // "带单订单不能修改"
	RpcErrCodeCopyTradingActivityExist           = 160010 // "活动已存在"

	// 交易员考试
	RpcErrCodeExamNotExist              = 160011 // "考试不存在"
	RpcErrCodeExamInProcess             = 160012 // "正在考试中"
	RpcErrCodeExamConfigError           = 160013 // "考试配置错误"
	RpcErrCodeExamAccountNotExist       = 160014 // "考试账户不存在"
	RpcErrCodeAccountHasActivePositions = 160015 // "账户存在未平仓的持仓"

	// 跟单分享错误码
	RpcErrCodeCopyTradingPositionOrderHasBeenClose = 160016 // "跟单订单已平仓"
	RpcErrCodeCopyTradingPendinOrderHasDealed      = 160017 // "跟单挂单已成交"
)

// computecenter 错误码
const (
	// 计算中心
	RpcErrCodeGroupSymbolNotExist         = 170000 // "交易组下品种不存在"
	RpcErrCodeGroupSymbolPositionNotExist = 170001 // "交易组下品种持仓数据不存在"
)

// pkgmanager 错误码
const (
	// 包管理
	RpcErrCodePkgNotExist             = 180000 // 包不存在
	RpcErrCodeCreatePkgError          = 180001 // 创建包失败
	RpcErrCodeUpdatePkgError          = 180002 // 更新包失败
	RpcErrCodeGetPkgVersionListError  = 180003 // 获取包版本列表失败
	RpcErrCodeGetPkgUpdateListError   = 180004 // 获取包更新列表失败
	RpcErrCodeGetPkgDetailError       = 180005 // 获取包详情失败
	RpcErrCodeDeletePkgError          = 180006 // 删除包失败
	RpcErrCodeGetPkgUpdateError       = 180007 // 获取包更新失败
	RpcErrCodeAbSettingExist          = 180008 // AB 面配置已存在
	RpcErrCodeAbSettingNotExist       = 180009 // AB 面配置不存在
	RpcErrCodeUpdateAbSettingError    = 180010 // 修改 AB 面配置失败
	RpcErrCodePkgUpdateRecordNotExist = 180011 // 包更新记录不存在
	RpcErrCodeGetPkgUpdateStructError = 180012 // 获取包更新结构化配置失败
	RpcErrCodeVerifyCodeError         = 180013 // 验证码验证失败
	RpcErrCodeVersionError            = 180014 // 版本号验证错误
)

// taskcenter 错误码
const (
	//任务
	RpcErrCodeTaskNotEnable                     = 190000 // "任务未启用"
	RpcErrCodeTaskTypeHasExist                  = 190001 // "任务类型已存在"
	RpcErrCodeTaskNotExist                      = 190002 // "任务不存在"
	RpcErrCodeTaskHasCompleted                  = 190003 // "任务已完成"
	RpcErrCodeTaskHasNotEnoughPoint             = 190004 // "积分不足"
	RpcErrCodeTaskHasNotCompleted               = 190005 // "任务未完成"
	RpcErrCodeTaskRewardHasClaimed              = 190006 // "奖励已领取"
	RpcErrCodeRedeemUserPointApplyIsNotPendging = 190007 // "兑换用户积分申请状态不是待处理"
	RpcErrCodeRedeemUserPointApplyProcessing    = 190008 // "兑换用户积分申请其他人正在审核"

	// 分佣相关错误
	RpcErrCodeCommissionAmountFormatError       = 190101 // 金额格式不对
	RpcErrCodeNotConfigTransactionGroup         = 190102 // "未配置交易组"
	RpcErrCodeCommissionBalanceNotEnough        = 190103 // "佣金余额不足"
	RpcErrCodeCommissionIsWithdrawing           = 190104 // "佣金正在提现中"
	RpcErrCodeInvitationCodeInvalidReturnRatio  = 190105 // "返佣比例与回撤比例设置错误"
	RpcErrCodeCommissionVolumeAuditNotExist     = 190106 // "佣金分佣记录不存在"
	RpcErrCodeCommissionVolumeHasAudit          = 190107 // "佣金分佣记录已审核"
	RpcErrCodeCommissionVolumeAuditingByAnother = 190108 // "其他人正在审核"

	// 交易员考试分佣
	RpcErrCodeTraderExamRebateRuleHasExist = 190201 // "交易员考试分佣规则已存在"
)

// 交易平台错误码
const (
	RpcErrCodeMT5RequestCommonError               = 200000 // 请求的常规错误 (MT5错误码: 10011)
	RpcErrCodeMT5RequestTimeout                   = 200001 // 请求已超时 (MT5错误码: 10012)
	RpcErrCodeMT5InvalidRequest                   = 200002 // 无效请求 (MT5错误码: 10013)
	RpcErrCodeMT5InvalidVolume                    = 200003 // 无效量 (MT5错误码: 10014)
	RpcErrCodeMT5InvalidPrice                     = 200004 // 无效价格 (MT5错误码: 10015)
	RpcErrCodeMT5WrongStopLevelsOrPrice           = 200005 // 错误止损水平或价格 (MT5错误码: 10016)
	RpcErrCodeMT5TradeDisabled                    = 200006 // 禁用交易 (MT5错误码: 10017)
	RpcErrCodeMT5MarketClosed                     = 200007 // 关闭市场 (MT5错误码: 10018)
	RpcErrCodeMT5NotEnoughMoney                   = 200008 // 没有足够的钱款 (MT5错误码: 10019)
	RpcErrCodeMT5PriceChanged                     = 200009 // 价格已变化 (MT5错误码: 10020)
	RpcErrCodeMT5NoPrice                          = 200010 // 无价格 (MT5错误码: 10021)
	RpcErrCodeMT5InvalidOrderExpiration           = 200011 // 无效订单到期 (MT5错误码: 10022)
	RpcErrCodeMT5OrderChanged                     = 200012 // 已更改订单 (MT5错误码: 10023)
	RpcErrCodeMT5TooManyTradeRequests             = 200013 // 太多交易请求 (MT5错误码: 10024)
	RpcErrCodeMT5RequestDoesNotContainChanges     = 200014 // 请求不包含更改 (MT5错误码: 10025)
	RpcErrCodeMT5AutotradingDisabledServer        = 200015 // 服务器上禁用的自动交易 (MT5错误码: 10026)
	RpcErrCodeMT5AutotradingDisabledClient        = 200016 // 客户端上禁用的自动交易 (MT5错误码: 10027)
	RpcErrCodeMT5RequestBlockedByDealer           = 200017 // 交易员封锁的请求 (MT5错误码: 10028)
	RpcErrCodeMT5ModificationFailed               = 200018 // 未能进行的修改 (MT5错误码: 10029)
	RpcErrCodeMT5FillModeNotSupported             = 200019 // 成交模式不受支持 (MT5错误码: 10030)
	RpcErrCodeMT5NoConnection                     = 200020 // 无连接 (MT5错误码: 10031)
	RpcErrCodeMT5RealAccountsOnly                 = 200021 // 仅被允许实账户 (MT5错误码: 10032)
	RpcErrCodeMT5OrderLimitReached                = 200022 // 已达到订单数的限制 (MT5错误码: 10033)
	RpcErrCodeMT5VolumeLimitReached               = 200023 // 已达到量限制 (MT5错误码: 10034)
	RpcErrCodeMT5InvalidOrProhibitedOrderType     = 200024 // 无效或被禁止的订单类型 (MT5错误码: 10035)
	RpcErrCodeMT5PositionAlreadyClosed            = 200025 // 位置已平仓 (MT5错误码: 10036)
	RpcErrCodeMT5UsedForInternalPurposes          = 200026 // 用于内部目的 (MT5错误码: 10037)
	RpcErrCodeMT5CloseVolumeExceedsPositionVolume = 200027 // 要平仓的量超过持仓的当前量 (MT5错误码: 10038)
	RpcErrCodeMT5ExistingCloseOrder               = 200028 // 要平仓的订单已经存在 (MT5错误码: 10039)
	RpcErrCodeMT5OpenPositionLimitReached         = 200029 // 已达到开仓数的限制 (MT5错误码: 10040)
	RpcErrCodeMT5RequestRejectedOrderCanceled     = 200030 // 已拒绝的请求，已取消的订单 (MT5错误码: 10041)
	RpcErrCodeMT5OnlyLongPositionsAllowed         = 200031 // 仅允许买入持仓 (MT5错误码: 10042)
	RpcErrCodeMT5OnlyShortPositionsAllowed        = 200032 // 仅允许卖出持仓 (MT5错误码: 10043)
	RpcErrCodeMT5OnlyClosePositionsAllowed        = 200033 // 仅允许平仓 (MT5错误码: 10044)
	RpcErrCodeMT5FifoCloseRuleViolation           = 200034 // 根据FIFO规则不允许平仓 (MT5错误码: 10045)
	RpcErrCodeMT5HedgeProhibited                  = 200035 // 由于禁止锁仓持仓，不允许开仓或下挂单 (MT5错误码: 10046)
	RpcErrCodeMT5Success                          = 200036 // 成功完成 (MT5错误码: 0)
	RpcErrCodeMT5SuccessNoInfo                    = 200037 // 成功完成，没有返回的信息 (MT5错误码: 1)
	RpcErrCodeMT5GeneralError                     = 200038 // 常规错误 (MT5错误码: 2)
	RpcErrCodeMT5InvalidParams                    = 200039 // 无效参数 (MT5错误码: 3)
	RpcErrCodeMT5InvalidInfo                      = 200040 // 无效信息 (MT5错误码: 4)
	RpcErrCodeMT5HardwareError                    = 200041 // 硬盘错误 (MT5错误码: 5)
	RpcErrCodeMT5MemoryError                      = 200042 // 内存错误 (MT5错误码: 6)
	RpcErrCodeMT5NetworkError_7                   = 200043 // 网络错误 (MT5错误码: 7)
	RpcErrCodeMT5NoPermission                     = 200044 // 没有足够的权限来执行操作 (MT5错误码: 8)
	RpcErrCodeMT5Timeout                          = 200045 // 已过期的超时 (MT5错误码: 9)
	RpcErrCodeMT5NoService                        = 200046 // 没有服务 (MT5错误码: 11)
	RpcErrCodeMT5TooFrequent                      = 200047 // 过于频繁的请求 (MT5错误码: 12)
	RpcErrCodeMT5NotFound                         = 200048 // 找不到 (MT5错误码: 13)
	RpcErrCodeMT5PartialError                     = 200049 // 部分错误 (MT5错误码: 14)
	RpcErrCodeMT5ServerShutdown                   = 200050 // 进行中的服务器关闭 (MT5错误码: 15)
	RpcErrCodeMT5OperationCanceled                = 200051 // 已取消操作 (MT5错误码: 16)
	RpcErrCodeMT5ReplicaInfo                      = 200052 // 副本信息 (MT5错误码: 17)
	RpcErrCodeMT5InvalidAccount                   = 200053 // 无效账户 (MT5错误码: 1001)
	RpcErrCodeMT5AccountDisabled                  = 200054 // 禁用账户 (MT5错误码: 1002)
	RpcErrCodeMT5InvalidCertificate               = 200055 // 无效证书 (MT5错误码: 1005)
	RpcErrCodeMT5UnconfirmedCertificate           = 200056 // 未确认证书 (MT5错误码: 1006)
	RpcErrCodeMT5InvalidServer                    = 200057 // 尝试连接至不是访问服务器的服务器 (MT5错误码: 1007)
	RpcErrCodeMT5OldClientVersion                 = 200058 // 旧客户版本 (MT5错误码: 1010)
	RpcErrCodeMT5InvalidClientType                = 200059 // 程序端的无效类型 (MT5错误码: 1000)
	RpcErrCodeMT5InvalidServerVersion             = 200060 // 过时的服务器版本 (MT5错误码: 1021)
	RpcErrCodeMT5InvalidServerID                  = 200061 // 禁用无效 ID 或服务器 (MT5错误码: 1015)
	RpcErrCodeMT5InvalidServerAddress             = 200062 // 无效地址 (MT5错误码: 1016)
	RpcErrCodeMT5ServerBusy                       = 200063 // 服务器正忙 (MT5错误码: 1018)
	RpcErrCodeMT5InvalidServerCertificate         = 200064 // 无效服务器证书 (MT5错误码: 1019)
	RpcErrCodeMT5UnknownAccount                   = 200065 // 未知账户 (MT5错误码: 1020)
	RpcErrCodeMT5LicenseLimit                     = 200066 // 由于许可限制不能连接服务器 (MT5错误码: 1022)
	RpcErrCodeMT5MobileNotAllowed                 = 200067 // 许可中不允许连接移动设备 (MT5错误码: 1023)
	RpcErrCodeMT5ManagerNotAllowed                = 200068 // 不允许经理进行该类型的连接 (MT5错误码: 1024)
	RpcErrCodeMT5DemoAccountNotAllowed            = 200069 // 禁用模拟账户的创建 (MT5错误码: 1025)
	RpcErrCodeMT5PasswordChangeRequired           = 200070 // 必须更改主密码 (MT5错误码: 1026)
	RpcErrCodeMT5InvalidDynamicPassword           = 200071 // 无效的动态密码 (MT5错误码: 1027)
	RpcErrCodeMT5NoDynamicPasswordKey             = 200072 // 没有为动态密码指定密钥 (MT5错误码: 1028)
	RpcErrCodeMT5PasswordChangeRequiredMT4        = 200073 // 从MetaTrader 4服务器导入账户后，需要更改密码 (MT5错误码: 1029)
	RpcErrCodeMT5PasswordChangeRequiredMT5        = 200074 // 从MetaTrader 5服务器导入账户后，需要更改密码 (MT5错误码: 1030)
	RpcErrCodeMT5InvalidVerificationCode          = 200075 // 验证码无效或过期 (MT5错误码: 1031)
	RpcErrCodeMT5EmailVerificationFailed          = 200076 // 无法发送电子邮件验证码 (MT5错误码: 1032)
	RpcErrCodeMT5PhoneVerificationFailed          = 200077 // 无法发送电话号码验证码 (MT5错误码: 1033)
	RpcErrCodeMT5APIConnectionNotAllowed          = 200078 // 禁止通过API连接账户 (MT5错误码: 1034)
	RpcErrCodeMT5RequestProcessing                = 200079 // 正要进行请求 (MT5错误码: 10001)
	RpcErrCodeMT5RequestAccepted                  = 200080 // 已接受的请求 (MT5错误码: 10002)
	RpcErrCodeMT5RequestProcessed                 = 200081 // 已处理的请求 (MT5错误码: 10003)
	RpcErrCodeMT5RequestPriceRequest              = 200082 // 重新报价响应请求 (MT5错误码: 10004)
	RpcErrCodeMT5RequestPriceResponse             = 200083 // 价格响应请求 (MT5错误码: 10005)
	RpcErrCodeMT5RequestCanceled                  = 200084 // 已取消的请求 (MT5错误码: 10007)
	RpcErrCodeMT5RequestSubmitted                 = 200085 // 因请求而提交的订单 (MT5错误码: 10008)
	RpcErrCodeMT5RequestFilled                    = 200086 // 已实现的请求 (MT5错误码: 10009)
	RpcErrCodeMT5RequestPartiallyFilled           = 200087 // 已部分实现的请求 (MT5错误码: 10010)
	RpcErrCodeMT5RequestReturned                  = 200088 // 返回到队列的请求 (MT5错误码: 11000)
	RpcErrCodeMT5RequestPartiallyFilledCanceled   = 200089 // 已部分成交请求，已取消余数 (MT5错误码: 11001)
	RpcErrCodeMT5RequestRequoted                  = 200090 // 已重新报价并返回到具有新价格的队列的请求 (MT5错误码: 11002)
	RpcErrCodeMT5NetworkError_60001               = 200091 // 网络错误 (MT5错误码: 60001)
	RpcErrCodeMT5BusinessError                    = 200092 // 业务错误 (MT5错误码: 60002)
	RpcErrCodeMT5TradingNotAllowed                = 200093 // 当前时间不允许交易 (MT5错误码: 60003)
	RpcErrCodeMT5UnknownError                     = 200094 // 未知错误 (MT5错误码: 99999)
	RpcErrCodeMT5MtRetErrConnection               = 200095 // 无连接 (MT5错误码: 10)
	RpcErrCodeMT5MtRetAuthAdvanced                = 200096 // 所需的扩展授权 (MT5错误码: 1003)
	RpcErrCodeMT5MtRetAuthCertificate             = 200097 // 所需的证书 (MT5错误码: 1004)
	RpcErrCodeMT5MtRetAuthServerBad               = 200098 // 未授权服务器 (MT5错误码: 1008)
	RpcErrCodeMT5MtRetAuthUpdateOnly              = 200099 // 仅有更新 (MT5错误码: 1009)
	RpcErrCodeMT5MtRetAuthManagerNoconfig         = 200100 // 尚未针对经理账户创建适当的经理配置 (MT5错误码: 1011)
	RpcErrCodeMT5MtRetAuthManagerIpblock          = 200101 // IP 地址针对经理无效 (MT5错误码: 1012)
	RpcErrCodeMT5MtRetAuthGroupInvalid            = 200102 // 未初始化组 (MT5错误码: 1013)
	RpcErrCodeMT5MtRetAuthCaDisabled              = 200103 // 禁用证书的生成 (MT5错误码: 1014)
	RpcErrCodeMT5MtRetAuthInvalidType             = 200104 // 错误类型的服务器 (MT5错误码: 1017)
	RpcErrCodeMT5MtRetCfgLastAdmin                = 200105 // 删除上一个管理员配置 (MT5错误码: 2000)
	RpcErrCodeMT5MtRetCfgLastAdminGroup           = 200106 // 不能删除上组管理员 (MT5错误码: 2001)
	RpcErrCodeMT5MtRetCfgNotEmpty                 = 200107 // 组包含账户或交易操作 (MT5错误码: 2003)
	RpcErrCodeMT5MtRetCfgInvalidRange             = 200108 // 账户或交易操作的无效范围 (MT5错误码: 2004)
	RpcErrCodeMT5MtRetCfgNotManagerLogin          = 200109 // 经理账户不属于经理组 (MT5错误码: 2005)
	RpcErrCodeMT5MtRetCfgBuiltin                  = 200110 // 内置受保护的配置 (MT5错误码: 2006)
	RpcErrCodeMT5MtRetCfgDuplicate                = 200111 // 副本配置 (MT5错误码: 2007)
	RpcErrCodeMT5MtRetCfgLimitReached             = 200112 // 已达到配置数的限制 (MT5错误码: 2008)
	RpcErrCodeMT5MtRetCfgNoAccessToMain           = 200113 // 不正确的网络配置 (MT5错误码: 2009)
	RpcErrCodeMT5MtRetCfgDealerIdExist            = 200114 // 具有相同 ID（账号）的交易员已经存在 (MT5错误码: 2010)
	RpcErrCodeMT5MtRetCfgBindAddrExist            = 200115 // 连接地址已经存在 (MT5错误码: 2011)
	RpcErrCodeMT5MtRetCfgWorkingTrade             = 200116 // 尝试删除工作交易服务器 (MT5错误码: 2012)
	RpcErrCodeMT5MtRetCfgGatewayNameExist         = 200117 // 具有该名称的网关已经存在 (MT5错误码: 2013)
	RpcErrCodeMT5MtRetCfgSwitchToBackup           = 200118 // 切换到备份服务器已为交易/历史服务器启用 (MT5错误码: 2014)
	RpcErrCodeMT5MtRetCfgNoBackupModule           = 200119 // 没有备份服务器 (MT5错误码: 2015)
	RpcErrCodeMT5MtRetCfgNoTradeModule            = 200120 // 没有
	RpcErrCodeMT5MtRetCfgNoHistoryModule          = 200121 // 没有历史服务器 (MT5错误码: 2017)
	RpcErrCodeMT5MtRetCfgAnotherSwitch            = 200122 // 已经开始切换到备份服务器的过程 (MT5错误码: 2018)
	RpcErrCodeMT5MtRetCfgNoLicenseFile            = 200123 // 没有许可证文件 (MT5错误码: 2019)
	RpcErrCodeMT5MtRetCfgGatewayLoginExist        = 200124 // 无法创建一个经理配置，因为这个登录名已被网关使用 (MT5错误码: 2020)
	RpcErrCodeMT5MtRetCfgInvalidCompany           = 200125 // 公司名称与许可证或白标不符 (MT5错误码: 2021)
	RpcErrCodeMT5MtRetUsrLastAdmin                = 200126 // 已删除最后一个管理员账户 (MT5错误码: 3001)
	RpcErrCodeMT5MtRetUsrLoginExhausted           = 200127 // 已用尽登录的范围 (MT5错误码: 3002)
	RpcErrCodeMT5MtRetUsrLoginProhibited          = 200128 // 在另一服务器上预留登录 (MT5错误码: 3003)
	RpcErrCodeMT5MtRetUsrLoginExist               = 200129 // 账户已经存在 (MT5错误码: 3004)
	RpcErrCodeMT5MtRetUsrSuicide                  = 200130 // 尝试自行删除 (MT5错误码: 3005)
	RpcErrCodeMT5MtRetUsrLimitReached             = 200131 // 已达到用户数的限制 (MT5错误码: 3007)
	RpcErrCodeMT5MtRetUsrHasTrades                = 200132 // 账户有开仓 (MT5错误码: 3008)
	RpcErrCodeMT5MtRetUsrDifferentServers         = 200133 // 尝试将账户移动到另一服务器 (MT5错误码: 3009)
	RpcErrCodeMT5MtRetUsrDifferentCurrency        = 200134 // 尝试将账户移动到具有不同入金货币的组 (MT5错误码: 3010)
	RpcErrCodeMT5MtRetUsrImportBalance            = 200135 // 未能导入账户结余 (MT5错误码: 3011)
	RpcErrCodeMT5MtRetUsrImportGroup              = 200136 // 通过错误组导入账户 (MT5错误码: 3012)
	RpcErrCodeMT5MtRetUsrAccountExist             = 200137 // 外部系统中的交易品种已存在用于指定登录名 (MT5错误码: 3013)
	RpcErrCodeMT5MtRetUsrImportAccount            = 200138 // 未能导入账户交易数据 (MT5错误码: 3014)
	RpcErrCodeMT5MtRetUsrImportPositions          = 200139 // 未能导入账户交易持仓 (MT5错误码: 3015)
	RpcErrCodeMT5MtRetUsrImportOrders             = 200140 // 未能导入账户未结订单 (MT5错误码: 3016)
	RpcErrCodeMT5MtRetUsrImportDeals              = 200141 // 未能导入账户交易历史 (MT5错误码: 3017)
	RpcErrCodeMT5MtRetUsrImportHistory            = 200142 // 未能导入账户订单历史 (MT5错误码: 3018)
	RpcErrCodeMT5MtRetUsrApiLimitReached          = 200143 // 已达到允许通过API连接的用户数量限制 (MT5错误码: 3019)
	RpcErrCodeMT5MtRetTradeLimitReached           = 200144 // 已达到交易订单数的限制 (MT5错误码: 4001)
	RpcErrCodeMT5MtRetTradeOrderExist             = 200145 // 订单已经存在 (MT5错误码: 4002)
	RpcErrCodeMT5MtRetTradeOrderExhausted         = 200146 // 已用尽订单的范围 (MT5错误码: 4003)
	RpcErrCodeMT5MtRetTradeDealExhausted          = 200147 // 已用尽交易的范围 (MT5错误码: 4004)
	RpcErrCodeMT5MtRetTradeMaxMoney               = 200148 // 已达到钱款金额的限制 (MT5错误码: 4005)
	RpcErrCodeMT5MtRetTradeDealExist              = 200149 // 这个交易服务器上已存在这个单号的交易 (MT5错误码: 4006)
	RpcErrCodeMT5MtRetTradeOrderProhibited        = 200150 // 保留订单标识符以供在另一个交易服务器上使用 (MT5错误码: 4007)
	RpcErrCodeMT5MtRetTradeDealProhibited         = 200151 // 保留交易标识符以供在另一个交易服务器上使用 (MT5错误码: 4008)
	RpcErrCodeMT5MtRetTradeSplitVolume            = 200152 // 拆分操作后，持仓交易量将变为零 (MT5错误码: 4009)
	RpcErrCodeMT5MtRetReportSnapshot              = 200153 // 数据库快照错误 (MT5错误码: 5001)
	RpcErrCodeMT5MtRetReportNotsupported          = 200154 // 该报告的方法不受支持 (MT5错误码: 5002)
	RpcErrCodeMT5MtRetReportNodata                = 200155 // 无报告信息 (MT5错误码: 5003)
	RpcErrCodeMT5MtRetReportTemplateBad           = 200156 // 错误模板 (MT5错误码: 5004)
	RpcErrCodeMT5MtRetReportTemplateEnd           = 200157 // 模板结束 (MT5错误码: 5005)
	RpcErrCodeMT5MtRetReportInvalidRow            = 200158 // 无效行大小 (MT5错误码: 5006)
	RpcErrCodeMT5MtRetReportLimitRepeat           = 200159 // 已达到副本标签数的限制 (MT5错误码: 5007)
	RpcErrCodeMT5MtRetReportLimitReport           = 200160 // 已达到报告大小的限制 (MT5错误码: 5008)
	RpcErrCodeMT5MtRetHstSymbolNotfound           = 200161 // 找不到交易品种 (MT5错误码: 6001)
	RpcErrCodeMT5MtRetExecutionTraderIdNotExist   = 200162 // 交易执行id不存在 (MT5错误码: 70024)
)

// 收银台
const (
	RpcErrCashierSignError                    = 301001 // 收银台签名错误
	RpcErrCashierParamsError                  = 301002 // 收银台参数错误
	RpcErrCashierSendHttpsFasterError         = 301003 // 请求频率过高
	RpcErrCashierSendBodyIsNullError          = 301004 // 请求body为空
	RpcErrCashierParamOneIsNullError          = 301005 // 缺少某一项参数
	RpcErrCashierParamsNotCompliantError      = 301006 // 参数不合规
	RpcErrCashierPayTypeNotUsedError          = 301007 // 收银台支付方式不可用
	RpcErrCashierDepositApplyOrderNotExist    = 301008 // 收银台充值订单号不存在
	RpcErrCashierWithdrawalApplyOrderExist    = 301009 // 收银台提现订单已存在,订单号重复
	RpcErrCashierWithdrawalApplyOrderFail     = 301010 // 收银台提现订单失败,请重试或者联系管理员
	RpcErrCashierWithdrawalApplyOrderNotExist = 301011 // 收银台提现订单号不存在
	RpcErrCashierDepositOrderExist            = 301012 // 收银台充值订单已存在,订单号重复
	RpcErrCashierDepositOrderFail             = 301013 // 收银台生成充值订单失败
	RpcErrCashierWithdrawalOrderFail          = 301014 // 收银台生成提现订单失败
	RpcErrCashierWithdrawalPayTypeError       = 301015 // 收银台提现下单成功,但所选择的支付方式对应的支付渠道不可用,请更换支付方式,重新下单
)

// 推送系统错误码
const (
	RpcErrCreateInAppNotificationFailed = 400001 // 创建应用内通知失败
	RpcErrMissingLanguage               = 400002 // 缺少语言
	RpcErrGetInAppNotificationFailed    = 400003 // 获取应用内通知失败
	RpcErrDeleteInAppNotificationFailed = 400004 // 删除应用内通知失败
	RpcErrUpdateInAppNotificationFailed = 400005 // 更新应用内通知失败
	RpcErrInAppNotificationNotExist     = 400006 // 应用内通知不存在
	RpcErrCreatePushMessageFailed       = 400007 // 创建推送消息失败
	RpcErrGetPushMessageFailed          = 400008 // 获取推送消息失败
	RpcErrUpdatePushMessageFailed       = 400009 // 更新推送消息失败
	RpcErrDeletePushMessageFailed       = 400010 // 删除推送消息失败
	RpcErrMessageExist                  = 400011 // 消息已存在
	RpcErrPushMessageFailed             = 400012 // 推送消息失败
	RpcErrCreatePushProviderFailed      = 400013 // 创建推送 provider 错误
	RpcErrDeletePushProviderFailed      = 400014 // 删除推送 provider 错误
	RpcErrUpdatePushProviderFailed      = 400015 // 更新推送 provider 错误
	RpcErrGetPushProviderFailed         = 400016 // 获取推送 provider 错误
	RpcErrCreateAdPosistionFailed       = 400017 // 创建广告位置失败
	RpcErrDeleteAdPositionFailed        = 400018 // 删除广告位置失败
	RpcErrUpdateAdPositionFailed        = 400019 // 更新广告位置失败
	RpcErrGetAdPositionFailed           = 400020 // 获取广告位置失败
	RpcErrAdPositionExist               = 400021 // 广告位置已存在
	RpcErrCreateAdConfigFailed          = 400022 // 创建广告配置失败
	RpcErrUpdateAdConfigFailed          = 400023 // 更新广告配置失败
	RpcErrDeleteAdConfigFailed          = 400024 // 删除广告配置失败
	RpcErrAdConfigExist                 = 400025 // 广告配置已存在
	RpcErrGetAdConfigFailed             = 400026 // 获取广告配置失败
	RpcErrRegisterFcmTokenFailed        = 400027 // 注册 FcmToken 失败
)

// promotion 营销管理错误码
const (
	RpcErrPromotionTradeUnlockSettlementBatchAlreadyExist = 500001 // 结算批次已存在
)

// promotion 任务中心错误码
const (
	RpcErrPromotionUserIdRequired               = 510001 // 用户ID不能为空
	RpcErrPromotionTaskTypeRequired             = 510002 // 任务类型不能为空
	RpcErrPromotionTaskItemRequired             = 510003 // 任务项不能为空
	RpcErrPromotionGetTaskError                 = 510004 // 获取任务信息失败
	RpcErrPromotionNoTask                       = 510005 // 任务不存在
	RpcErrPromotionGetTaskUserError             = 510006 // 获取用户任务信息失败
	RpcErrPromotionUserForbidenTask             = 510007 // 用户禁止执行任务
	RpcErrPromotionGetCompletedTaskError        = 510008 // 获取用户完成任务信息失败
	RpcErrPromotionSignInTaskOnlyOne            = 510009 // 签到任务只有一条
	RpcErrPromotionInviteTaskOnlyOne            = 510010 // 邀请任务只有一条
	RpcErrPromotionInviteDetailIsEmpty          = 510011 // 邀请任务详情为空
	RpcErrPromotionInvitePointsIsZero           = 510012 // 邀请任务积分为0
	RpcErrPromotionBeginnerTaskOnlyOne          = 510013 // 新手任务只有一条
	RpcErrPromotionInvalidTaskType              = 510014 // 任务类型无效
	RpcErrPromotionCompletedTaskError           = 510015 // 完成任务错误
	RpcErrPromotionInvalidItemType              = 510016 // 任务项目类型无效
	RpcErrPromotionTaskExists                   = 510017 // 任务已存在
	RpcErrPromotionCreateTaskError              = 510018 // 创建任务错误
	RPcErrPromotionPointsNotEnough              = 510019 // 积分不足
	RpcErrPromotionCreateUserPointsHistoryError = 510020 // 创建用户积分记录错误
	RpcErrPromotionTaskUnvalueable              = 510021 // 任务不可用
	RpcErrPromotionUserReceiveError             = 510022 // 创建用户任务领取记录错误
	RpcErrPromotionGetTaskListError             = 510023 // 获取任务列表错误
	RpcErrPromotionGetTaskDetailError           = 510024 // 获取任务详情错误
	RpcErrPromotionGetLanguageContentError      = 510025 // 获取语言内容错误
	RpcErrPromotionTaskIdRequired               = 510026 // 任务ID不能为空
	RpcErrPromotionGetGlobalSettingsError       = 510027 // 获取全局设置错误
	RpcErrPromotionGetTaskSettingError          = 510028 // 获取任务设置错误
	RpcErrPromotionUserPointHistoryIdRequired   = 510029 // 用户积分记录ID不能为空
	RpcErrPromotionGetUserPointAuditListError   = 510030 // 获取用户积分审核列表错误
	RpcErrPromotionGetUserPointHistoryListError = 510031 // 获取用户积分记录列表错误
	RpcErrPromotionGetUserReceivedTaskListError = 510032 // 获取用户任务领取列表错误
	RpcErrPromotionTccCancelExchangeError       = 510033 // TCC取消积分兑换错误
	RpcErrPromotionTccConfirmExchangeError      = 510034 // TCC确认积分兑换错误
	RpcErrPromotionAuditStatusNotMatch          = 510035 // 审核状态不匹配
	RpcErrPromotionTaskTypeNotMatch             = 510036 // 任务类型不匹配
	RpcErrPromotionTccTryExchangeError          = 510037 // TCC尝试积分兑换错误
	RpcErrPromotionUpdateTaskError              = 510038 // 更新任务错误
	RpcErrPromotionUpdateSettingError           = 510039 // 更新任务设置错误
	RpcErrPromotionTaskUserIdRequired           = 510040 // 任务用户ID不能为空
	RpcErrPromotionUpdateTaskUserError          = 510041 // 更新任务用户错误
	RpcErrPromotionActionTypeError              = 510042 // 动作类型错误
	RpcErrPromotionTaskDailyTypeError           = 510043 // 日常任务类型错误
	RpcErrPromotionTaskBeginnerTypeError        = 510044 // 新手任务类型错误
	RpcErrPromotionPointsMustGreaterThanZero    = 510045 // 积分必须大于0
	RpcErrPromotionPointsOutOfStock             = 510046 // 积分库存不足
	RpcErrPromotionSignInTaskRepeated           = 510047 // 签到任务重复完成
	RpcErrPromotionBeginnerTaskRepeated         = 510048 // 新手任务重复完成
	RpcErrPromotionPeriodType                   = 510049 // 任务周期类型错误
	RpcErrPromotionBonusActivityRepeated        = 510050 // 你已经参加过该活动，无法重复参加
	RpcErrPromotionBonusActivityConflict        = 510051 // 你已经参与其他活动，无法参加该活动
)

// MT5 错误码 @not export excel
const (
	MT5ErrSuccess                        = 0     // 成功完成
	MT5ErrSuccessNoInfo                  = 1     // 成功完成，没有返回的信息
	MT5ErrGeneralError                   = 2     // 常规错误
	MT5ErrInvalidParams                  = 3     // 无效参数
	MT5ErrInvalidInfo                    = 4     // 无效信息
	MT5ErrHardwareError                  = 5     // 硬盘错误
	MT5ErrMemoryError                    = 6     // 内存错误
	MT5ErrNetworkError_7                 = 7     // 网络错误
	MT5ErrNoPermission                   = 8     // 没有足够的权限来执行操作
	MT5ErrTimeout                        = 9     // 已过期的超时
	MT5ErrMtRetErrConnection             = 10    // 无连接。
	MT5ErrNoService                      = 11    // 没有服务
	MT5ErrTooFrequent                    = 12    // 过于频繁的请求
	MT5ErrNotFound                       = 13    // 找不到
	MT5ErrPartialError                   = 14    // 部分错误
	MT5ErrServerShutdown                 = 15    // 进行中的服务器关闭
	MT5ErrOperationCanceled              = 16    // 已取消操作
	MT5ErrReplicaInfo                    = 17    // 副本信息
	MT5ErrInvalidClientType              = 1000  // 程序端的无效类型
	MT5ErrInvalidAccount                 = 1001  // 无效账户
	MT5ErrAccountDisabled                = 1002  // 禁用账户
	MT5ErrMtRetAuthAdvanced              = 1003  // 所需的扩展授权。
	MT5ErrMtRetAuthCertificate           = 1004  // 所需的证书。
	MT5ErrInvalidCertificate             = 1005  // 无效证书
	MT5ErrUnconfirmedCertificate         = 1006  // 未确认证书
	MT5ErrInvalidServer                  = 1007  // 尝试连接至不是访问服务器的服务器
	MT5ErrMtRetAuthServerBad             = 1008  // 未授权服务器。
	MT5ErrMtRetAuthUpdateOnly            = 1009  // 仅有更新。
	MT5ErrOldClientVersion               = 1010  // 旧客户版本
	MT5ErrMtRetAuthManagerNoconfig       = 1011  // 尚未针对经理账户创建适当的经理配置。
	MT5ErrMtRetAuthManagerIpblock        = 1012  // IP 地址针对经理无效。
	MT5ErrMtRetAuthGroupInvalid          = 1013  // 未初始化组（您必须重启服务器）。
	MT5ErrMtRetAuthCaDisabled            = 1014  // 禁用证书的生成。
	MT5ErrInvalidServerID                = 1015  // 禁用无效 ID 或服务器（应该检查服务器 ID）
	MT5ErrInvalidServerAddress           = 1016  // 无效地址（应该检查服务器 IP 地址）
	MT5ErrMtRetAuthInvalidType           = 1017  // 错误类型的服务器（应该检查服务器 ID 和类型）。
	MT5ErrServerBusy                     = 1018  // 服务器正忙
	MT5ErrInvalidServerCertificate       = 1019  // 无效服务器证书
	MT5ErrUnknownAccount                 = 1020  // 未知账户
	MT5ErrInvalidServerVersion           = 1021  // 过时的服务器版本
	MT5ErrLicenseLimit                   = 1022  // 由于许可限制不能连接服务器
	MT5ErrMobileNotAllowed               = 1023  // 许可中不允许连接移动设备
	MT5ErrManagerNotAllowed              = 1024  // 不允许经理进行该类型的连接
	MT5ErrDemoAccountNotAllowed          = 1025  // 禁用模拟账户的创建
	MT5ErrPasswordChangeRequired         = 1026  // 必须更改主密码
	MT5ErrInvalidDynamicPassword         = 1027  // 无效的动态密码
	MT5ErrNoDynamicPasswordKey           = 1028  // 没有为动态密码指定密钥
	MT5ErrPasswordChangeRequiredMT4      = 1029  // 从MetaTrader 4服务器导入账户后，需要更改密码
	MT5ErrPasswordChangeRequiredMT5      = 1030  // 从MetaTrader 5服务器导入账户后，需要更改密码
	MT5ErrInvalidVerificationCode        = 1031  // 验证码无效或过期
	MT5ErrEmailVerificationFailed        = 1032  // 无法发送电子邮件验证码
	MT5ErrPhoneVerificationFailed        = 1033  // 无法发送电话号码验证码
	MT5ErrAPIConnectionNotAllowed        = 1034  // 禁止通过API连接账户
	MT5ErrMtRetCfgLastAdmin              = 2000  // 删除上一个管理员配置。
	MT5ErrMtRetCfgLastAdminGroup         = 2001  // 不能删除上组管理员。
	MT5ErrMtRetCfgNotEmpty               = 2003  // 组包含账户或交易操作。
	MT5ErrMtRetCfgInvalidRange           = 2004  // 账户或交易操作的无效范围。
	MT5ErrMtRetCfgNotManagerLogin        = 2005  // 经理账户不属于经理组。
	MT5ErrMtRetCfgBuiltin                = 2006  // 内置受保护的配置。
	MT5ErrMtRetCfgDuplicate              = 2007  // 副本配置。
	MT5ErrMtRetCfgLimitReached           = 2008  // 已达到配置数的限制。
	MT5ErrMtRetCfgNoAccessToMain         = 2009  // 不正确的网络配置。
	MT5ErrMtRetCfgDealerIdExist          = 2010  // 具有相同 ID（账号）的交易员已经存在。
	MT5ErrMtRetCfgBindAddrExist          = 2011  // 连接地址已经存在。
	MT5ErrMtRetCfgWorkingTrade           = 2012  // 尝试删除工作交易服务器。
	MT5ErrMtRetCfgGatewayNameExist       = 2013  // 具有该名称的网关已经存在。
	MT5ErrMtRetCfgSwitchToBackup         = 2014  // 切换到备份服务器已为交易/历史服务器启用。
	MT5ErrMtRetCfgNoBackupModule         = 2015  // 没有备份服务器。
	MT5ErrMtRetCfgNoTradeModule          = 2016  // 没有交易服务器。
	MT5ErrMtRetCfgNoHistoryModule        = 2017  // 没有历史服务器。
	MT5ErrMtRetCfgAnotherSwitch          = 2018  // 已经开始切换到备份服务器的过程。
	MT5ErrMtRetCfgNoLicenseFile          = 2019  // 没有许可证文件。
	MT5ErrMtRetCfgGatewayLoginExist      = 2020  // 无法创建一个经理配置，因为这个登录名已被网关使用。
	MT5ErrMtRetCfgInvalidCompany         = 2021  // 公司名称与许可证或白标不符。当您尝试添加或保持 公司名称 与平台许可证中指定的公司名称不同的组作为主标或附加白标时，将返回错误。
	MT5ErrMtRetUsrLastAdmin              = 3001  // 已删除最后一个管理员账户。
	MT5ErrMtRetUsrLoginExhausted         = 3002  // 已用尽登录的范围。
	MT5ErrMtRetUsrLoginProhibited        = 3003  // 在另一服务器上预留登录。
	MT5ErrMtRetUsrLoginExist             = 3004  // 账户已经存在。
	MT5ErrMtRetUsrSuicide                = 3005  // 尝试自行删除。
	MT5ErrMtRetUsrLimitReached           = 3007  // 已达到用户数的限制。
	MT5ErrMtRetUsrHasTrades              = 3008  // 账户有开仓。
	MT5ErrMtRetUsrDifferentServers       = 3009  // 尝试将账户移动到另一服务器。
	MT5ErrMtRetUsrDifferentCurrency      = 3010  // 尝试将账户移动到具有不同入金货币的组。
	MT5ErrMtRetUsrImportBalance          = 3011  // 未能导入账户结余。
	MT5ErrMtRetUsrImportGroup            = 3012  // 通过错误组导入账户。
	MT5ErrMtRetUsrAccountExist           = 3013  // 外部系统中的交易品种已存在用于指定登录名。
	MT5ErrMtRetUsrImportAccount          = 3014  // 未能导入账户交易数据。
	MT5ErrMtRetUsrImportPositions        = 3015  // 未能导入账户交易持仓。
	MT5ErrMtRetUsrImportOrders           = 3016  // 未能导入账户未结订单。
	MT5ErrMtRetUsrImportDeals            = 3017  // 未能导入账户交易历史。
	MT5ErrMtRetUsrImportHistory          = 3018  // 未能导入账户订单历史。
	MT5ErrMtRetUsrApiLimitReached        = 3019  // 已达到允许通过API连接的用户数量限制(IMTUser::USER_RIGHT_API_ENABLED)。当前限制为100个用户。
	MT5ErrMtRetTradeLimitReached         = 4001  // 已达到交易订单数的限制。
	MT5ErrMtRetTradeOrderExist           = 4002  // 订单已经存在。
	MT5ErrMtRetTradeOrderExhausted       = 4003  // 已用尽订单的范围。
	MT5ErrMtRetTradeDealExhausted        = 4004  // 已用尽交易的范围。
	MT5ErrMtRetTradeMaxMoney             = 4005  // 已达到钱款金额的限制。
	MT5ErrMtRetTradeDealExist            = 4006  // 这个交易服务器上已存在这个单号的交易。
	MT5ErrMtRetTradeOrderProhibited      = 4007  // 保留订单标识符以供在另一个交易服务器上使用。
	MT5ErrMtRetTradeDealProhibited       = 4008  // 保留交易标识符以供在另一个交易服务器上使用。
	MT5ErrMtRetTradeSplitVolume          = 4009  // 拆分操作后，持仓交易量将变为零。错误用于以下方法：
	MT5ErrMtRetReportSnapshot            = 5001  // 数据库快照错误。
	MT5ErrMtRetReportNotsupported        = 5002  // 该报告的方法不受支持。
	MT5ErrMtRetReportNodata              = 5003  // 无报告信息。
	MT5ErrMtRetReportTemplateBad         = 5004  // 错误模板。
	MT5ErrMtRetReportTemplateEnd         = 5005  // 模板结束。
	MT5ErrMtRetReportInvalidRow          = 5006  // 无效行大小。
	MT5ErrMtRetReportLimitRepeat         = 5007  // 已达到副本标签数的限制。
	MT5ErrMtRetReportLimitReport         = 5008  // 已达到报告大小的限制。
	MT5ErrMtRetHstSymbolNotfound         = 6001  // 找不到交易品种，请试图重启历史服务器。
	MT5ErrRequestProcessing              = 10001 // 正要进行请求
	MT5ErrRequestAccepted                = 10002 // 已接受的请求
	MT5ErrRequestProcessed               = 10003 // 已处理的请求
	MT5ErrRequestPriceRequest            = 10004 // 重新报价响应请求
	MT5ErrRequestPriceResponse           = 10005 // 价格响应请求
	MT5ErrMtRetRequestReject             = 10006 // 已拒绝的请求。
	MT5ErrRequestCanceled                = 10007 // 已取消的请求
	MT5ErrRequestSubmitted               = 10008 // 因请求而提交的订单
	MT5ErrRequestFilled                  = 10009 // 已实现的请求
	MT5ErrRequestPartiallyFilled         = 10010 // 已部分实现的请求
	MT5ErrCommonError                    = 10011 // 请求的常规错误 (MT5错误码: 10011)
	MT5ErrRequestTimeout                 = 10012 // 请求已超时 (MT5错误码: 10012)
	MT5ErrInvalidRequest                 = 10013 // 无效请求 (MT5错误码: 10013)
	MT5ErrInvalidVolume                  = 10014 // 无效量 (MT5错误码: 10014)
	MT5ErrInvalidPrice                   = 10015 // 无效价格 (MT5错误码: 10015)
	MT5ErrWrongStopLevel                 = 10016 // 错误止损水平或价格 (MT5错误码: 10016)
	MT5ErrTradeDisabled                  = 10017 // 禁用交易 (MT5错误码: 10017)
	MT5ErrMarketClosed                   = 10018 // 关闭市场 (MT5错误码: 10018)
	MT5ErrNotEnoughMoney                 = 10019 // 没有足够的钱款 (MT5错误码: 10019)
	MT5ErrPriceChanged                   = 10020 // 价格已变化 (MT5错误码: 10020)
	MT5ErrNoPrice                        = 10021 // 无价格 (MT5错误码: 10021)
	MT5ErrInvalidOrderExpire             = 10022 // 无效订单到期 (MT5错误码: 10022)
	MT5ErrOrderChanged                   = 10023 // 已更改订单 (MT5错误码: 10023)
	MT5ErrTooManyRequests                = 10024 // 太多交易请求 (MT5错误码: 10024)
	MT5ErrRequestNoChanges               = 10025 // 请求不包含更改 (MT5错误码: 10025)
	MT5ErrAutotradingDisabledServer      = 10026 // 服务器上禁用的自动交易 (MT5错误码: 10026)
	MT5ErrAutotradingDisabledClient      = 10027 // 客户端上禁用的自动交易 (MT5错误码: 10027)
	MT5ErrRequestBlockedByDealer         = 10028 // 交易员封锁的请求 (MT5错误码: 10028)
	MT5ErrModificationFailed             = 10029 // 未能进行的修改 (MT5错误码: 10029)
	MT5ErrFillModeNotSupported           = 10030 // 成交模式不受支持 (MT5错误码: 10030)
	MT5ErrNoConnection                   = 10031 // 无连接 (MT5错误码: 10031)
	MT5ErrRealAccountsOnly               = 10032 // 仅被允许实账户 (MT5错误码: 10032)
	MT5ErrOrderLimitReached              = 10033 // 已达到订单数的限制 (MT5错误码: 10033)
	MT5ErrVolumeLimitReached             = 10034 // 已达到量限制 (MT5错误码: 10034)
	MT5ErrInvalidOrderType               = 10035 // 无效或被禁止的订单类型 (MT5错误码: 10035)
	MT5ErrPositionClosed                 = 10036 // 位置已平仓 (MT5错误码: 10036)
	MT5ErrInternalPurpose                = 10037 // 用于内部目的 (MT5错误码: 10037)
	MT5ErrCloseVolumeExceeds             = 10038 // 要平仓的量超过持仓的当前量 (MT5错误码: 10038)
	MT5ErrExistingCloseOrder             = 10039 // 要平仓的订单已经存在 (MT5错误码: 10039)
	MT5ErrOpenPositionLimit              = 10040 // 已达到开仓数的限制 (MT5错误码: 10040)
	MT5ErrRequestRejected                = 10041 // 已拒绝的请求，已取消的订单 (MT5错误码: 10041)
	MT5ErrOnlyLongAllowed                = 10042 // 仅允许买入持仓 (MT5错误码: 10042)
	MT5ErrOnlyShortAllowed               = 10043 // 仅允许卖出持仓 (MT5错误码: 10043)
	MT5ErrOnlyCloseAllowed               = 10044 // 仅允许平仓 (MT5错误码: 10044)
	MT5ErrFifoCloseViolation             = 10045 // 根据FIFO规则不允许平仓 (MT5错误码: 10045)
	MT5ErrHedgeProhibited                = 10046 // 由于禁止锁仓持仓，不允许开仓或下挂单 (MT5错误码: 10046)
	MT5ErrRequestReturned                = 11000 // 返回到队列的请求
	MT5ErrRequestPartiallyFilledCanceled = 11001 // 已部分成交请求，已取消余数
	MT5ErrRequestRequoted                = 11002 // 已重新报价并返回到具有新价格的队列的请求
	MT5ErrMtRetErrNotimplement           = 12000 // 尚未执行。
	MT5ErrMtRetErrNotmain                = 12001 // 应该在主交易服务器上执行操作。
	MT5ErrMtRetErrNotsupported           = 12002 // 命令不受该服务器支持。
	MT5ErrMtRetErrDeadlock               = 12003 // 因可能的死锁已取消操作。
	MT5ErrMtRetErrLocked                 = 12004 // 处理已冻结对象。
	MT5ErrMtRetMessengerInvalidPhone     = 14000 // 指定无效电话号码。电话号码格式如下：+[country code][number]，例如：+74951113594。号码不得有空格。
	MT5ErrMtRetMessengerNotMobile        = 14001 // 指定固定电话号码，而不是移动电话号码。发送信息时必须指定移动电话号码。信息不能传送到其他电话号码。
	MT5ErrMtRetSubsNotFound              = 15000 // 订阅未找到。
	MT5ErrMtRetSubsNotFoundCfg           = 15001 // 订阅配置未找到。
	MT5ErrMtRetSubsNotFoundUser          = 15002 // 无法找到订阅用户。
	MT5ErrMtRetSubsDisabled              = 15003 // 订阅禁用。当前状态通过IMTSubscription::Status获得。
	MT5ErrMtRetSubsPermissionUser        = 15004 // 用户不允许订阅。
	MT5ErrMtRetSubsPermissionSubscribe   = 15005 // 不允许订阅。订阅选项的可用性可通过IMTConSubscription::ControlMode属性来决定。
	MT5ErrMtRetSubsPermissionUnsubscribe = 15006 // 不允许取消订阅。取消订阅的功能由IMTConSubscription::ControlMode属性来决定。
	MT5ErrMtRetSubsRealOnly              = 15007 // 仅允许真实账户订阅。
	MT5ErrNetworkError_60001             = 60001 // 网络错误
	MT5ErrBusinessError                  = 60002 // 业务错误
	MT5ErrTradingNotAllowed              = 60003 // 当前时间不允许交易
	MT5ErrBusiErr                        = 70000 // 通用业务错误
	MT5ErrBusiErrParam                   = 70001 // 业务参数错误
	MT5ErrBusiErrPath                    = 70002 // path不能为*
	MT5ErrBusiErrTimer                   = 70003 // 当前时间不允许交易
	MT5ErrBusiErrHolidayTime             = 70004 // 假期不允许交易
	MT5ErrBusiErrPermissions             = 70005 // 账号无交易权限
	MT5ErrBusiErrPermissionsDemo         = 70006 // 无权访问虚拟组接口
	MT5ErrBusiErrPermissionsReal         = 70007 // 无权访问真实组接口
	MT5ErrBusiErrGroupNotExist           = 70008 // 用户所在组不存在
	MT5ErrBusiErrGroupSymbolNotExist     = 70009 // 用户所在组不存在该品种
	MT5ErrBusiErrPositionNotExist        = 70010 // 用户持仓不存在
	MT5ErrBusiErrTpGt                    = 70011 // 新止盈价必须 > 当前价
	MT5ErrBusiErrSlLt                    = 70012 // 新止损价必须 < 当前价
	MT5ErrBusiErrSlGlZero                = 70013 // 新止损价不能小于0
	MT5ErrBusiErrTpLt                    = 70014 // 新止盈价必须 < 当前价
	MT5ErrBusiErrTpGl                    = 70015 // 新止损价必须 > 当前价
	MT5ErrBusiErrTpGlZero                = 70016 // 新止盈价不能小于0
	MT5ErrBusiErrPositionType            = 70017 // 不支持的持仓类型
	MT5ErrBusiErrHolidateErr             = 70018 // 假期设置不存在
	MT5ErrBusiErrGroupNotMatch           = 70019 // 账户组不匹配
	MT5ErrBusiErrGroupApiNotMatch        = 70020 // 接口与组不匹配
	MT5ErrBusiErrMarginInitialBuy        = 70021 // 预付款比率要大于0
	MT5ErrBusiErrFixHoldingBuy           = 70022 // 固定模式下，锁仓保证金=合约大小
	MT5ErrBusiErrGroupAddMarginInitral   = 70023 // 追加预付款水平 < 强制平仓水平
	MT5ErrMtRetExecutionTraderIdNotExist = 70024 // 交易员ID不存在
	MT5ErrUnknownError                   = 99999 // 未知错误
)

// mt5 错误码映射 map
var errorCodeMap = map[int]ErrCode{
	MT5ErrCommonError:                    RpcErrCodeMT5RequestCommonError,
	MT5ErrRequestTimeout:                 RpcErrCodeMT5RequestTimeout,
	MT5ErrInvalidRequest:                 RpcErrCodeMT5InvalidRequest,
	MT5ErrInvalidVolume:                  RpcErrCodeMT5InvalidVolume,
	MT5ErrInvalidPrice:                   RpcErrCodeMT5InvalidPrice,
	MT5ErrWrongStopLevel:                 RpcErrCodeMT5WrongStopLevelsOrPrice,
	MT5ErrTradeDisabled:                  RpcErrCodeMT5TradeDisabled,
	MT5ErrMarketClosed:                   RpcErrCodeMT5MarketClosed,
	MT5ErrNotEnoughMoney:                 RpcErrCodeMT5NotEnoughMoney,
	MT5ErrPriceChanged:                   RpcErrCodeMT5PriceChanged,
	MT5ErrNoPrice:                        RpcErrCodeMT5NoPrice,
	MT5ErrInvalidOrderExpire:             RpcErrCodeMT5InvalidOrderExpiration,
	MT5ErrOrderChanged:                   RpcErrCodeMT5OrderChanged,
	MT5ErrTooManyRequests:                RpcErrCodeMT5TooManyTradeRequests,
	MT5ErrRequestNoChanges:               RpcErrCodeMT5RequestDoesNotContainChanges,
	MT5ErrAutotradingDisabledServer:      RpcErrCodeMT5AutotradingDisabledServer,
	MT5ErrAutotradingDisabledClient:      RpcErrCodeMT5AutotradingDisabledClient,
	MT5ErrRequestBlockedByDealer:         RpcErrCodeMT5RequestBlockedByDealer,
	MT5ErrModificationFailed:             RpcErrCodeMT5ModificationFailed,
	MT5ErrFillModeNotSupported:           RpcErrCodeMT5FillModeNotSupported,
	MT5ErrNoConnection:                   RpcErrCodeMT5NoConnection,
	MT5ErrRealAccountsOnly:               RpcErrCodeMT5RealAccountsOnly,
	MT5ErrOrderLimitReached:              RpcErrCodeMT5OrderLimitReached,
	MT5ErrVolumeLimitReached:             RpcErrCodeMT5VolumeLimitReached,
	MT5ErrInvalidOrderType:               RpcErrCodeMT5InvalidOrProhibitedOrderType,
	MT5ErrPositionClosed:                 RpcErrCodeMT5PositionAlreadyClosed,
	MT5ErrInternalPurpose:                RpcErrCodeMT5UsedForInternalPurposes,
	MT5ErrCloseVolumeExceeds:             RpcErrCodeMT5CloseVolumeExceedsPositionVolume,
	MT5ErrExistingCloseOrder:             RpcErrCodeMT5ExistingCloseOrder,
	MT5ErrOpenPositionLimit:              RpcErrCodeMT5OpenPositionLimitReached,
	MT5ErrRequestRejected:                RpcErrCodeMT5RequestRejectedOrderCanceled,
	MT5ErrOnlyLongAllowed:                RpcErrCodeMT5OnlyLongPositionsAllowed,
	MT5ErrOnlyShortAllowed:               RpcErrCodeMT5OnlyShortPositionsAllowed,
	MT5ErrOnlyCloseAllowed:               RpcErrCodeMT5OnlyClosePositionsAllowed,
	MT5ErrFifoCloseViolation:             RpcErrCodeMT5FifoCloseRuleViolation,
	MT5ErrHedgeProhibited:                RpcErrCodeMT5HedgeProhibited,
	MT5ErrSuccess:                        RpcErrCodeMT5Success,
	MT5ErrSuccessNoInfo:                  RpcErrCodeMT5SuccessNoInfo,
	MT5ErrGeneralError:                   RpcErrCodeMT5GeneralError,
	MT5ErrInvalidParams:                  RpcErrCodeMT5InvalidParams,
	MT5ErrInvalidInfo:                    RpcErrCodeMT5InvalidInfo,
	MT5ErrHardwareError:                  RpcErrCodeMT5HardwareError,
	MT5ErrMemoryError:                    RpcErrCodeMT5MemoryError,
	MT5ErrNetworkError_7:                 RpcErrCodeMT5NetworkError_7,
	MT5ErrNoPermission:                   RpcErrCodeMT5NoPermission,
	MT5ErrTimeout:                        RpcErrCodeMT5Timeout,
	MT5ErrNoService:                      RpcErrCodeMT5NoService,
	MT5ErrTooFrequent:                    RpcErrCodeMT5TooFrequent,
	MT5ErrNotFound:                       RpcErrCodeMT5NotFound,
	MT5ErrPartialError:                   RpcErrCodeMT5PartialError,
	MT5ErrServerShutdown:                 RpcErrCodeMT5ServerShutdown,
	MT5ErrOperationCanceled:              RpcErrCodeMT5OperationCanceled,
	MT5ErrReplicaInfo:                    RpcErrCodeMT5ReplicaInfo,
	MT5ErrInvalidAccount:                 RpcErrCodeMT5InvalidAccount,
	MT5ErrAccountDisabled:                RpcErrCodeMT5AccountDisabled,
	MT5ErrInvalidCertificate:             RpcErrCodeMT5InvalidCertificate,
	MT5ErrUnconfirmedCertificate:         RpcErrCodeMT5UnconfirmedCertificate,
	MT5ErrInvalidServer:                  RpcErrCodeMT5InvalidServer,
	MT5ErrOldClientVersion:               RpcErrCodeMT5OldClientVersion,
	MT5ErrInvalidClientType:              RpcErrCodeMT5InvalidClientType,
	MT5ErrInvalidServerVersion:           RpcErrCodeMT5InvalidServerVersion,
	MT5ErrInvalidServerID:                RpcErrCodeMT5InvalidServerID,
	MT5ErrInvalidServerAddress:           RpcErrCodeMT5InvalidServerAddress,
	MT5ErrServerBusy:                     RpcErrCodeMT5ServerBusy,
	MT5ErrInvalidServerCertificate:       RpcErrCodeMT5InvalidServerCertificate,
	MT5ErrUnknownAccount:                 RpcErrCodeMT5UnknownAccount,
	MT5ErrLicenseLimit:                   RpcErrCodeMT5LicenseLimit,
	MT5ErrMobileNotAllowed:               RpcErrCodeMT5MobileNotAllowed,
	MT5ErrManagerNotAllowed:              RpcErrCodeMT5ManagerNotAllowed,
	MT5ErrDemoAccountNotAllowed:          RpcErrCodeMT5DemoAccountNotAllowed,
	MT5ErrPasswordChangeRequired:         RpcErrCodeMT5PasswordChangeRequired,
	MT5ErrInvalidDynamicPassword:         RpcErrCodeMT5InvalidDynamicPassword,
	MT5ErrNoDynamicPasswordKey:           RpcErrCodeMT5NoDynamicPasswordKey,
	MT5ErrPasswordChangeRequiredMT4:      RpcErrCodeMT5PasswordChangeRequiredMT4,
	MT5ErrPasswordChangeRequiredMT5:      RpcErrCodeMT5PasswordChangeRequiredMT5,
	MT5ErrInvalidVerificationCode:        RpcErrCodeMT5InvalidVerificationCode,
	MT5ErrEmailVerificationFailed:        RpcErrCodeMT5EmailVerificationFailed,
	MT5ErrPhoneVerificationFailed:        RpcErrCodeMT5PhoneVerificationFailed,
	MT5ErrAPIConnectionNotAllowed:        RpcErrCodeMT5APIConnectionNotAllowed,
	MT5ErrRequestProcessing:              RpcErrCodeMT5RequestProcessing,
	MT5ErrRequestAccepted:                RpcErrCodeMT5RequestAccepted,
	MT5ErrRequestProcessed:               RpcErrCodeMT5RequestProcessed,
	MT5ErrRequestPriceRequest:            RpcErrCodeMT5RequestPriceRequest,
	MT5ErrRequestPriceResponse:           RpcErrCodeMT5RequestPriceResponse,
	MT5ErrRequestCanceled:                RpcErrCodeMT5RequestCanceled,
	MT5ErrRequestSubmitted:               RpcErrCodeMT5RequestSubmitted,
	MT5ErrRequestFilled:                  RpcErrCodeMT5RequestFilled,
	MT5ErrRequestPartiallyFilled:         RpcErrCodeMT5RequestPartiallyFilled,
	MT5ErrRequestReturned:                RpcErrCodeMT5RequestReturned,
	MT5ErrRequestPartiallyFilledCanceled: RpcErrCodeMT5RequestPartiallyFilledCanceled,
	MT5ErrRequestRequoted:                RpcErrCodeMT5RequestRequoted,
	MT5ErrNetworkError_60001:             RpcErrCodeMT5NetworkError_60001,
	MT5ErrBusinessError:                  RpcErrCodeMT5BusinessError,
	MT5ErrTradingNotAllowed:              RpcErrCodeMT5TradingNotAllowed,
	MT5ErrUnknownError:                   RpcErrCodeMT5UnknownError,
	MT5ErrMtRetErrConnection:             RpcErrCodeMT5MtRetErrConnection,
	MT5ErrMtRetAuthAdvanced:              RpcErrCodeMT5MtRetAuthAdvanced,
	MT5ErrMtRetAuthCertificate:           RpcErrCodeMT5MtRetAuthCertificate,
	MT5ErrMtRetAuthServerBad:             RpcErrCodeMT5MtRetAuthServerBad,
	MT5ErrMtRetAuthUpdateOnly:            RpcErrCodeMT5MtRetAuthUpdateOnly,
	MT5ErrMtRetAuthManagerNoconfig:       RpcErrCodeMT5MtRetAuthManagerNoconfig,
	MT5ErrMtRetAuthManagerIpblock:        RpcErrCodeMT5MtRetAuthManagerIpblock,
	MT5ErrMtRetAuthGroupInvalid:          RpcErrCodeMT5MtRetAuthGroupInvalid,
	MT5ErrMtRetAuthCaDisabled:            RpcErrCodeMT5MtRetAuthCaDisabled,
	MT5ErrMtRetAuthInvalidType:           RpcErrCodeMT5MtRetAuthInvalidType,
	MT5ErrMtRetCfgLastAdmin:              RpcErrCodeMT5MtRetCfgLastAdmin,
	MT5ErrMtRetCfgLastAdminGroup:         RpcErrCodeMT5MtRetCfgLastAdminGroup,
	MT5ErrMtRetCfgNotEmpty:               RpcErrCodeMT5MtRetCfgNotEmpty,
	MT5ErrMtRetCfgInvalidRange:           RpcErrCodeMT5MtRetCfgInvalidRange,
	MT5ErrMtRetCfgNotManagerLogin:        RpcErrCodeMT5MtRetCfgNotManagerLogin,
	MT5ErrMtRetCfgBuiltin:                RpcErrCodeMT5MtRetCfgBuiltin,
	MT5ErrMtRetCfgDuplicate:              RpcErrCodeMT5MtRetCfgDuplicate,
	MT5ErrMtRetCfgLimitReached:           RpcErrCodeMT5MtRetCfgLimitReached,
	MT5ErrMtRetCfgNoAccessToMain:         RpcErrCodeMT5MtRetCfgNoAccessToMain,
	MT5ErrMtRetCfgDealerIdExist:          RpcErrCodeMT5MtRetCfgDealerIdExist,
	MT5ErrMtRetCfgBindAddrExist:          RpcErrCodeMT5MtRetCfgBindAddrExist,
	MT5ErrMtRetCfgWorkingTrade:           RpcErrCodeMT5MtRetCfgWorkingTrade,
	MT5ErrMtRetCfgGatewayNameExist:       RpcErrCodeMT5MtRetCfgGatewayNameExist,
	MT5ErrMtRetCfgSwitchToBackup:         RpcErrCodeMT5MtRetCfgSwitchToBackup,
	MT5ErrMtRetCfgNoBackupModule:         RpcErrCodeMT5MtRetCfgNoBackupModule,
	MT5ErrMtRetCfgNoTradeModule:          RpcErrCodeMT5MtRetCfgNoTradeModule,
	MT5ErrMtRetCfgNoHistoryModule:        RpcErrCodeMT5MtRetCfgNoHistoryModule,
	MT5ErrMtRetCfgAnotherSwitch:          RpcErrCodeMT5MtRetCfgAnotherSwitch,
	MT5ErrMtRetCfgNoLicenseFile:          RpcErrCodeMT5MtRetCfgNoLicenseFile,
	MT5ErrMtRetCfgGatewayLoginExist:      RpcErrCodeMT5MtRetCfgGatewayLoginExist,
	MT5ErrMtRetCfgInvalidCompany:         RpcErrCodeMT5MtRetCfgInvalidCompany,
	MT5ErrMtRetUsrLastAdmin:              RpcErrCodeMT5MtRetUsrLastAdmin,
	MT5ErrMtRetUsrLoginExhausted:         RpcErrCodeMT5MtRetUsrLoginExhausted,
	MT5ErrMtRetUsrLoginProhibited:        RpcErrCodeMT5MtRetUsrLoginProhibited,
	MT5ErrMtRetUsrLoginExist:             RpcErrCodeMT5MtRetUsrLoginExist,
	MT5ErrMtRetUsrSuicide:                RpcErrCodeMT5MtRetUsrSuicide,
	MT5ErrMtRetUsrLimitReached:           RpcErrCodeMT5MtRetUsrLimitReached,
	MT5ErrMtRetUsrHasTrades:              RpcErrCodeMT5MtRetUsrHasTrades,
	MT5ErrMtRetUsrDifferentServers:       RpcErrCodeMT5MtRetUsrDifferentServers,
	MT5ErrMtRetUsrDifferentCurrency:      RpcErrCodeMT5MtRetUsrDifferentCurrency,
	MT5ErrMtRetUsrImportBalance:          RpcErrCodeMT5MtRetUsrImportBalance,
	MT5ErrMtRetUsrImportGroup:            RpcErrCodeMT5MtRetUsrImportGroup,
	MT5ErrMtRetUsrAccountExist:           RpcErrCodeMT5MtRetUsrAccountExist,
	MT5ErrMtRetUsrImportAccount:          RpcErrCodeMT5MtRetUsrImportAccount,
	MT5ErrMtRetUsrImportPositions:        RpcErrCodeMT5MtRetUsrImportPositions,
	MT5ErrMtRetUsrImportOrders:           RpcErrCodeMT5MtRetUsrImportOrders,
	MT5ErrMtRetUsrImportDeals:            RpcErrCodeMT5MtRetUsrImportDeals,
	MT5ErrMtRetUsrImportHistory:          RpcErrCodeMT5MtRetUsrImportHistory,
	MT5ErrMtRetUsrApiLimitReached:        RpcErrCodeMT5MtRetUsrApiLimitReached,
	MT5ErrMtRetTradeLimitReached:         RpcErrCodeMT5MtRetTradeLimitReached,
	MT5ErrMtRetTradeOrderExist:           RpcErrCodeMT5MtRetTradeOrderExist,
	MT5ErrMtRetTradeOrderExhausted:       RpcErrCodeMT5MtRetTradeOrderExhausted,
	MT5ErrMtRetTradeDealExhausted:        RpcErrCodeMT5MtRetTradeDealExhausted,
	MT5ErrMtRetTradeMaxMoney:             RpcErrCodeMT5MtRetTradeMaxMoney,
	MT5ErrMtRetTradeDealExist:            RpcErrCodeMT5MtRetTradeDealExist,
	MT5ErrMtRetTradeOrderProhibited:      RpcErrCodeMT5MtRetTradeOrderProhibited,
	MT5ErrMtRetTradeDealProhibited:       RpcErrCodeMT5MtRetTradeDealProhibited,
	MT5ErrMtRetTradeSplitVolume:          RpcErrCodeMT5MtRetTradeSplitVolume,
	MT5ErrMtRetReportSnapshot:            RpcErrCodeMT5MtRetReportSnapshot,
	MT5ErrMtRetReportNotsupported:        RpcErrCodeMT5MtRetReportNotsupported,
	MT5ErrMtRetReportNodata:              RpcErrCodeMT5MtRetReportNodata,
	MT5ErrMtRetReportTemplateBad:         RpcErrCodeMT5MtRetReportTemplateBad,
	MT5ErrMtRetReportTemplateEnd:         RpcErrCodeMT5MtRetReportTemplateEnd,
	MT5ErrMtRetReportInvalidRow:          RpcErrCodeMT5MtRetReportInvalidRow,
	MT5ErrMtRetReportLimitRepeat:         RpcErrCodeMT5MtRetReportLimitRepeat,
	MT5ErrMtRetReportLimitReport:         RpcErrCodeMT5MtRetReportLimitReport,
	MT5ErrMtRetHstSymbolNotfound:         RpcErrCodeMT5MtRetHstSymbolNotfound,
	MT5ErrMtRetExecutionTraderIdNotExist: RpcErrCodeMT5MtRetExecutionTraderIdNotExist,
}

// ConvertMT5ToCustomCode 交易平台 错误码转换函数
func ConvertMT5ToCustomCode(mt5Code int) ErrCode {
	if customCode, exists := errorCodeMap[mt5Code]; exists {
		return customCode
	}
	return RpcErrCodeRequestTradePlatformError // 默认错误码
}

func GetMessage(code int, message string) string {
	// if message, OK := errorMessages[code]; OK {
	//	return message
	// }
	// return ""

	return fmt.Sprintf("MT平台错误 code: %d, message %s", code, message)
}

// 收银台错误码 @not export excel
const (
	CashierSuccess                      = 200  // success
	CashierSignError                    = 1001 // 收银台签名错误
	CashierParamsError                  = 1002 // 收银台参数错误
	CashierSendHttpsFasterError         = 1101 // 请求频率过高
	CashierSendBodyIsNullError          = 1102 // 请求body为空
	CashierParamOneIsNullError          = 1103 // 缺少某一项参数
	CashierParamsNotCompliantError      = 1104 // 参数不合规
	CashierChannelNotAvailableError     = 1105 // 暂无可用通道,请更换支付方式或稍后再试
	CashierPayTypeNotUsedError          = 1201 // 收银台支付方式不可用
	CashierDepositApplyOrderNotExist    = 1202 // 收银台充值订单号不存在
	CashierWithdrawalApplyOrderExist    = 1203 // 收银台提现订单已存在,订单号重复
	CashierWithdrawalApplyOrderFail     = 1204 // 收银台提现订单失败,请重试或者联系管理员
	CashierWithdrawalApplyOrderNotExist = 1205 // 收银台提现订单号不存在
	CashierDepositOrderExist            = 1206 // 收银台充值订单已存在,订单号重复
	CashierDepositOrderFail             = 1207 // 生成充值订单失败
	CashierWithdrawalOrderFail          = 1208 // 生成提现订单失败
	CashierWithdrawalPayTypeError       = 1209 // 提现下单成功,但所选择的支付方式对应的支付渠道不可用,请更换支付方式,重新下单
)

var CashierErrorCode = map[int64]int{
	CashierSuccess:                      RpcSuccess,
	CashierSignError:                    RpcErrCashierSignError,
	CashierParamsError:                  RpcErrCashierParamsError,
	CashierSendHttpsFasterError:         RpcErrCashierSendHttpsFasterError,
	CashierSendBodyIsNullError:          RpcErrCashierSendBodyIsNullError,
	CashierParamOneIsNullError:          RpcErrCashierParamOneIsNullError,
	CashierParamsNotCompliantError:      RpcErrCashierParamsNotCompliantError,
	CashierPayTypeNotUsedError:          RpcErrCashierPayTypeNotUsedError,
	CashierDepositApplyOrderNotExist:    RpcErrCashierDepositApplyOrderNotExist,
	CashierWithdrawalApplyOrderExist:    RpcErrCashierWithdrawalApplyOrderExist,
	CashierWithdrawalApplyOrderFail:     RpcErrCashierWithdrawalApplyOrderFail,
	CashierWithdrawalApplyOrderNotExist: RpcErrCashierWithdrawalApplyOrderNotExist,
	CashierDepositOrderExist:            RpcErrCashierDepositOrderExist,
	CashierDepositOrderFail:             RpcErrCashierDepositOrderFail,
	CashierWithdrawalOrderFail:          RpcErrCashierWithdrawalOrderFail,
	CashierWithdrawalPayTypeError:       RpcErrCashierWithdrawalPayTypeError,
}

var CashierErrorMsg = map[int64]string{
	CashierSuccess:                      "success",
	CashierSignError:                    "收银台签名错误",
	CashierParamsError:                  "收银台参数错误",
	CashierSendHttpsFasterError:         "请求频率过高",
	CashierSendBodyIsNullError:          "请求body为空",
	CashierParamOneIsNullError:          "缺少某一项参数",
	CashierParamsNotCompliantError:      "参数不合规",
	CashierPayTypeNotUsedError:          "收银台支付方式不可用",
	CashierDepositApplyOrderNotExist:    "收银台充值订单号不存在",
	CashierWithdrawalApplyOrderExist:    "收银台提现订单已存在,订单号重复",
	CashierWithdrawalApplyOrderFail:     "收银台提现订单失败,请重试或者联系管理员",
	CashierWithdrawalApplyOrderNotExist: "收银台提现订单号不存在",
	CashierDepositOrderExist:            "收银台充值订单已存在,订单号重复",
	CashierDepositOrderFail:             "收银台生成充值订单失败",
	CashierWithdrawalOrderFail:          "收银台生成提现订单失败",
	CashierWithdrawalPayTypeError:       "收银台提现下单成功,但所选择的支付方式对应的支付渠道不可用,请更换支付方式,重新下单",
}

// SMS 错误码映射表 - 将 SMS 服务器返回的错误码转换为系统内部错误码
var SMSErrorCodeMapping = map[string]ErrCode{
	"1":  RespCodeSMSSuccess,                   // 操作成功
	"2":  RespCodeSMSOperationFailed,           // 操作失败 -> 系统繁忙提示
	"3":  RespCodeSMSInternalError,             // 内部配置错误 -> 系统繁忙提示（敏感错误）
	"4":  RespCodeSMSConfigError,               // 内部配置错误 -> 系统繁忙提示（敏感错误）
	"5":  RespCodeSMSConfigError,               // 数据错误 -> 系统繁忙提示（敏感错误）
	"6":  RespCodeSMSBusinessError,             // 手机类型，区号和号码不能为空
	"7":  RespCodeSMSBusinessError,             // 邮箱类型，邮箱不能为空
	"8":  RespCodeSMSTemplateEmpty,             // 模板为空 -> 系统繁忙提示（敏感错误）
	"9":  RespCodeSMSDataError,                 // 数据错误 -> 系统繁忙提示（敏感错误）
	"10": RespCodeSMSDataError,                 // 数据错误 -> 系统繁忙提示（敏感错误）
	"11": RespCodeSMSVerifyCodeNotUsed,         // 验证码未使用
	"12": RespCodeSMSVerifyCodeUsed,            // 验证码已使用
	"13": RespCodeSMSFrequencyLimit,            // 验证码请求频繁
	"14": RespCodeSMSFrequencyLimit,            // 验证码超过一天请求量
	"15": RespCodeAccountCodeError,             // 验证码错误或已失效
	"16": RespCodeSMSVerifyCodeExpired,         // 验证码已过期
	"17": RespCodeSMSVerifyCodeUsed,            // 验证码已使用
	"18": RespCodeSMSGoogleJWTVerifyFailed,     // 谷歌JWT验证失败
	"19": RespCodeSMSGoogleJWTFailedToOpenAuth, // 谷歌JWT打开授权json失败
}

// MError defined error
type MError struct {
	Code int
	Msg  string
	Data interface{}
}

func (e *MError) Error() string {
	return e.Msg
}

// New ...
func NewErr(code int, msg string) *MError {
	return &MError{Code: code, Msg: msg}
}
