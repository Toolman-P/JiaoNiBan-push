package tpns

import (
	"fmt"
)

// audience type
type AudienceType string

const (
	AudienceAll            AudienceType = "all"
	AudienceTag            AudienceType = "tag"
	AudienceToken          AudienceType = "token"
	AudienceTokenList      AudienceType = "token_list"
	AudienceAccount        AudienceType = "account"
	AudienceAccountList    AudienceType = "account_list"
	AudienceAccountPackage AudienceType = "package_account_push"
	AudienceTokenPackage   AudienceType = "package_token_push"
)

//tag operation
type TagOperationType string

const (
	TagOperationAnd TagOperationType = "AND"
	TagOperationOr  TagOperationType = "OR"
)

//deprecated
// platform type
type PlatformType string

const (
	PlatformAndroid PlatformType = "android"
	PlatformIOS     PlatformType = "ios"
)

// message type
type MessageType string

const (
	Notify  MessageType = "notify"
	Message MessageType = "message"
)

// environment type
type EnvironmentType string

const (
	Product EnvironmentType = "product"
	Develop EnvironmentType = "dev"
)

type AndroidActionActivityAttribute struct {
	If int32 `json:"if"`
	Pf int32 `json:"pf"`
}

type AndroidActionBrowserAttribute struct {
	Url     string `json:"url"`
	Confirm int32  `json:"confirm,omitempty"`
}

type AndroidAction struct {
	ActionType   int                             `json:"action_type"`
	Activity     string                          `json:"activity,omitempty"`
	ActivityAttr *AndroidActionActivityAttribute `json:"aty_attr,omitempty"`
	Browser      *AndroidActionBrowserAttribute  `json:"browser,omitempty"`
	Intent       string                          `json:"intent,omitempty"`
}

func DefaultAndroidAction() *AndroidAction {
	return &AndroidAction{
		ActionType: 1, //set to default 1
	}
}

type AndroidMessage struct {
	ChannelId       string         `json:"n_ch_id,omitempty"`
	ChannelName     string         `json:"n_ch_name,omitempty"`
	XiaomiChannelId string         `json:"xm_ch_id,omitempty"`
	HuaweiChannelId string         `json:"hw_ch_id,omitempty"`
	OppoChannelId   string         `json:"oppo_ch_id,omitempty"`
	VivoChannelId   string         `json:"vivo_ch_id,omitempty"`
	BuilderId       int            `json:"builder_id"`
	BadgeType       int            `json:"badge_type"` //default -1
	Ring            int            `json:"ring"`       //default 1
	RingRaw         string         `json:"ring_raw"`
	Vibrate         int            `json:"vibrate"`   //default 1
	Lights          int            `json:"lights"`    //default 1
	Clearable       int            `json:"clearable"` //default 1
	IconType        int            `json:"icon_type"`
	IconRes         string         `json:"icon_res,omitempty"`
	StyleId         int            `json:"style_id"` //default 1
	SmallIcon       string         `json:"small_icon,omitempty"`
	Action          *AndroidAction `json:"action,omitempty"`
	CustomContent   string         `json:"custom_content,omitempty"` //serialize from map<string,string>
	ShowType        int            `json:"show_type"`                //default 2
	IconColor       uint32         `json:"icon_color"`
}

func DefaultAndroidMessage() *AndroidMessage {
	return &AndroidMessage{
		BadgeType: -1,
		Ring:      1,
		Vibrate:   1,
		Lights:    1,
		Clearable: 1,
		StyleId:   1,
		ShowType:  2,
	}
}

type IOSAps struct {
	Alert            interface{} `json:"alert,omitempty"` //map or string
	BadgeType        int         `json:"badge_type"`
	Category         string      `json:"category,omitempty"`
	ContentAvailable int         `json:"content-available"`
	Sound            string      `json:"sound,omitempty"`
	MutableContent   int         `json:"mutable-content"` //default 1
}

func DefaultIOSAps() *IOSAps {
	return &IOSAps{
		MutableContent: 1,
	}
}

type IOSMessage struct {
	Aps    *IOSAps `json:"aps,omitempty"`
	Custom string  `json:"custom,omitempty"`
}

type TagItem struct {
	Tags          []string         `json:"tags,omitempty"`
	IsNot         bool             `json:"is_not"`
	TagsOperator  TagOperationType `json:"tags_operator,omitempty"`
	ItemsOperator TagOperationType `json:"items_operator,omitempty"`
	TagType       string           `json:"tag_type,omitempty"`
}

type TagRule struct {
	TagItems []TagItem        `json:"tag_items,omitempty"`
	IsNot    bool             `json:"is_not"`
	Operator TagOperationType `json:"operator,omitempty"`
}

type AcceptTime struct {
	Hour   string `json:"hour,omitempty"`
	Minute string `json:"min,omitempty"`
}

type AcceptTimeRange struct {
	Start *AcceptTime `json:"start,omitempty"`
	End   *AcceptTime `json:"end,omitempty"`
}

type TPNsMessage struct {
	Title                 string            `json:"title,omitempty"`
	Content               string            `json:"content,omitempty"`
	AcceptTime            []AcceptTimeRange `json:"accept_time,omitempty"`
	ThreadId              string            `json:"thread_id,omitempty"`
	ThreadSumText         string            `json:"thread_sumtext,omitempty"`
	XGMediaResources      string            `json:"xg_media_resources,omitempty"`
	XGMediaAudioResources string            `json:"xg_media_audio_resources,omitempty"`
	Android               *AndroidMessage   `json:"android,omitempty"`
	IOS                   *IOSMessage       `json:"ios,omitempty"`
}

type ChannelRule struct {
	Channel string `json:"channel,omitempty"`
	Disable bool   `json:"disable"`
}

type LoopParameter struct {
	StartDate     string   `json:"startDate,omitempty"`
	EndDate       string   `json:"endDate,omitempty"`
	LoopType      int      `json:"loopType"`
	LoopDayIndexs []uint32 `json:"loopDayIndexs,omitempty"`
	DayTimes      []string `json:"DayTimes,omitempty"`
}

type Request struct {
	Audience AudienceType `json:"audience_type"`
	//deprecated
	Platform    PlatformType `json:"platform"`
	Message     *TPNsMessage `json:"message"`
	MessageType MessageType  `json:"message_type"`

	TagRules    []TagRule `json:"tag_rules,omitempty"`
	TokenList   []string  `json:"token_list,omitempty"`
	AccountList []string  `json:"account_list,omitempty"`

	Environment EnvironmentType `json:"environment"`

	UploadId           int    `json:"upload_id"`
	ExpireTime         int    `json:"expire_time"` //default 259200
	SendTime           string `json:"send_time,omitempty"`
	MultiPkg           bool   `json:"multi_pkg"`
	PlanId             string `json:"plan_id,omitempty"`
	AccountPushType    int    `json:"account_push_type"`
	CollapseId         int    `json:"collapse_id"`
	PushSpeed          int    `json:"push_speed"`
	TPNsOnlinePushType int    `json:"tpns_online_push_type"`
	ForceCollapse      bool   `json:"force_collapse"`
	IgnoreInvalidToken int    `json:"ignore_invalid_token"`
	TraceId            string `json:"trace_id"`

	ChannelRules []ChannelRule  `json:"channel_rules,omitempty"`
	LoopParam    *LoopParameter `json:"loop_param,omitempty"`
}

func (r *Request) Validate() error {
	if len(r.Audience) == 0 {
		return fmt.Errorf("missing audience type")
	}
	if r.Audience != AudienceAll && r.Audience != AudienceTag &&
		r.Audience != AudienceToken && r.Audience != AudienceTokenList &&
		r.Audience != AudienceAccount && r.Audience != AudienceAccountList &&
		r.Audience != AudienceAccountPackage && r.Audience != AudienceTokenPackage {
		return fmt.Errorf("invalid audience type: %s", r.Audience)
	}
	if r.Audience == AudienceToken || r.Audience == AudienceTokenList {
		if len(r.TokenList) == 0 {
			return fmt.Errorf("missing token info")
		}
	}

	if r.Audience == AudienceAccount || r.Audience == AudienceAccountList {
		if len(r.AccountList) == 0 {
			return fmt.Errorf("missing account info")
		}
	}

	if r.Audience == AudienceTag && len(r.TagRules) == 0 {
		return fmt.Errorf("missing tag rule info")
	}

	if r.Audience == AudienceAccountPackage || r.Audience == AudienceTokenPackage {
		if r.UploadId == 0 {
			return fmt.Errorf("missing UploadId")
		}
	}

	//if len(r.Platform) == 0 {
	//	return fmt.Errorf("missing platform type")
	//} else if r.Platform != PlatformAndroid && r.Platform != PlatformIOS {
	//	return fmt.Errorf("invalid platform type: %s", r.Platform)
	//}

	if len(r.MessageType) == 0 {
		return fmt.Errorf("missing message type")
	} else if r.MessageType != Notify && r.MessageType != Message {
		return fmt.Errorf("invalid message type: %s", r.MessageType)
	}

	//if r.Platform == PlatformIOS {
	//	if len(r.Environment) == 0 {
	//		return fmt.Errorf("missing environment")
	//	} else if r.Environment != Product && r.Environment != Develop {
	//		return fmt.Errorf("invalid environment: %v", r.Environment)
	//	}
	//}
	if r.Message != nil && r.Message.IOS != nil {
		if r.Environment != Product && r.Environment != Develop {
			return fmt.Errorf("invalid environment: %v", r.Environment)
		}
	}
	return nil
}

type Response struct {
	Seq         int64       `json:"seq"`
	PushId      interface{} `json:"push_id,omitempty"` //may be return string or array of string
	RetCode     int         `json:"ret_code"`
	ErrMsg      string      `json:"err_msg,omitempty"`
	Environment string      `json:"environment,omitempty"`
	Result      string      `json:"result,omitempty"`
}

func (r *Response) GetPushId() string {
	if v, ok := r.PushId.(string); ok {
		return v
	}
	return ""
}

func (r *Response) GetPushIds() []string {
	if v, ok := r.PushId.([]string); ok {
		return v
	}
	return nil
}

type UploadResponse struct {
	RetCode  int    `json:"retCode"`
	ErrMsg   string `json:"errMsg"`
	UploadId int    `json:"uploadId"`
}
