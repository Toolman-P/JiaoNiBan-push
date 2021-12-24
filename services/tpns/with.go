package tpns

type Option func(r *Request)

func WithAudience(audience AudienceType) Option {
	return func(r *Request) {
		r.Audience = audience
	}
}

//deprecated
func WithPlatform(platform PlatformType) Option {
	return func(r *Request) {
		r.Platform = platform
	}
}

func WithMessage(message *TPNsMessage) Option {
	return func(r *Request) {
		r.Message = message
	}
}

func WithTitle(title string) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.Title = title
	}
}

func WithContent(content string) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.Content = content
	}
}

func WithAcceptTime(acceptTime []AcceptTimeRange) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.AcceptTime = acceptTime
	}
}

func WithThreadId(threadId string) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.ThreadId = threadId
	}
}

func WithThreadSumText(threadSumText string) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.ThreadSumText = threadSumText
	}
}

func WithXGMediaResources(xgMediaResources string) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.XGMediaResources = xgMediaResources
	}
}

func WithXGMediaAudioResources(xgMediaAudioResources string) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.XGMediaAudioResources = xgMediaAudioResources
	}
}

func WithAndroidMessage(android *AndroidMessage) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.Android = android
	}
}

func WithIOSMessage(ios *IOSMessage) Option {
	return func(r *Request) {
		if r.Message == nil {
			r.Message = &TPNsMessage{}
		}
		r.Message.IOS = ios
	}
}

func WithMessageType(messageType MessageType) Option {
	return func(r *Request) {
		r.MessageType = messageType
	}
}

func WithTagRules(tagRules []TagRule) Option {
	return func(r *Request) {
		r.TagRules = tagRules
	}
}

func WithTokenList(tokenList []string) Option {
	return func(r *Request) {
		r.TokenList = tokenList
	}
}

func WithAccountList(accountList []string) Option {
	return func(r *Request) {
		r.AccountList = accountList
	}
}

func WithEnvironment(environment EnvironmentType) Option {
	return func(r *Request) {
		r.Environment = environment
	}
}

func WithUploadId(uploadId int) Option {
	return func(r *Request) {
		r.UploadId = uploadId
	}
}

func WithExpireTime(expireTime int) Option {
	return func(r *Request) {
		r.ExpireTime = expireTime
	}
}

func WithSendTime(sendTime string) Option {
	return func(r *Request) {
		r.SendTime = sendTime
	}
}

func WithMultiPkg(multiPkg bool) Option {
	return func(r *Request) {
		r.MultiPkg = multiPkg
	}
}

func WithPlanId(planId string) Option {
	return func(r *Request) {
		r.PlanId = planId
	}
}

func WithAccountPushType(accountPushType int) Option {
	return func(r *Request) {
		r.AccountPushType = accountPushType
	}
}

func WithCollapseId(collapseId int) Option {
	return func(r *Request) {
		r.CollapseId = collapseId
	}
}

func WithPushSpeed(pushSpeed int) Option {
	return func(r *Request) {
		r.PushSpeed = pushSpeed
	}
}

func WithTPNsOnlinePushType(tpnsOnlinePushType int) Option {
	return func(r *Request) {
		r.TPNsOnlinePushType = tpnsOnlinePushType
	}
}

func WithForceCollapse(forceCollapse bool) Option {
	return func(r *Request) {
		r.ForceCollapse = forceCollapse
	}
}

func WithChannelRules(channelRules []ChannelRule) Option {
	return func(r *Request) {
		r.ChannelRules = channelRules
	}
}

func WithLoopParam(loopParam *LoopParameter) Option {
	return func(r *Request) {
		r.LoopParam = loopParam
	}
}

func WithIgnoreInvalidToken(ignoreInvalidToken bool) Option {
	return func(r *Request) {
		if ignoreInvalidToken {
			r.IgnoreInvalidToken = 1
		} else {
			r.IgnoreInvalidToken = 0
		}
	}
}

func WithTraceId(traceId string) Option {
	return func(r *Request) {
		r.TraceId = traceId
	}
}

func NewRequest(opts ...Option) *Request {
	var r = Request{
		ExpireTime: 259200,
	}

	for _, o := range opts {
		o(&r)
	}
	return &r
}
