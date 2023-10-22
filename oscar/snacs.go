package oscar

//
// 0x01: OService
//

const (
	OServiceErr               uint16 = 0x0001
	OServiceClientOnline      uint16 = 0x0002
	OServiceHostOnline        uint16 = 0x0003
	OServiceServiceRequest    uint16 = 0x0004
	OServiceServiceResponse   uint16 = 0x0005
	OServiceRateParamsQuery   uint16 = 0x0006
	OServiceRateParamsReply   uint16 = 0x0007
	OServiceRateParamsSubAdd  uint16 = 0x0008
	OServiceRateDelParamSub   uint16 = 0x0009
	OServiceRateParamChange   uint16 = 0x000A
	OServicePauseReq          uint16 = 0x000B
	OServicePauseAck          uint16 = 0x000C
	OServiceResume            uint16 = 0x000D
	OServiceUserInfoQuery     uint16 = 0x000E
	OServiceUserInfoUpdate    uint16 = 0x000F
	OServiceEvilNotification  uint16 = 0x0010
	OServiceIdleNotification  uint16 = 0x0011
	OServiceMigrateGroups     uint16 = 0x0012
	OServiceMotd              uint16 = 0x0013
	OServiceSetPrivacyFlags   uint16 = 0x0014
	OServiceWellKnownUrls     uint16 = 0x0015
	OServiceNoop              uint16 = 0x0016
	OServiceClientVersions    uint16 = 0x0017
	OServiceHostVersions      uint16 = 0x0018
	OServiceMaxConfigQuery    uint16 = 0x0019
	OServiceMaxConfigReply    uint16 = 0x001A
	OServiceStoreConfig       uint16 = 0x001B
	OServiceConfigQuery       uint16 = 0x001C
	OServiceConfigReply       uint16 = 0x001D
	OServiceSetUserInfoFields uint16 = 0x001E
	OServiceProbeReq          uint16 = 0x001F
	OServiceProbeAck          uint16 = 0x0020
	OServiceBartReply         uint16 = 0x0021
	OServiceBartQuery2        uint16 = 0x0022
	OServiceBartReply2        uint16 = 0x0023

	OServiceTLVTagsReconnectHere uint16 = 0x05
	OServiceTLVTagsLoginCookie   uint16 = 0x06
	OServiceTLVTagsGroupID       uint16 = 0x0D
	OServiceTLVTagsSSLCertName   uint16 = 0x8D
	OServiceTLVTagsSSLState      uint16 = 0x8E
)

type SNAC_0x01_0x02_OServiceClientOnline struct {
	GroupVersions []struct {
		FoodGroup   uint16
		Version     uint16
		ToolID      uint16
		ToolVersion uint16
	}
}

type SNAC_0x01_0x03_OServiceHostOnline struct {
	FoodGroups []uint16
}

type SNAC_0x01_0x04_OServiceServiceRequest struct {
	FoodGroup uint16
	TLVRestBlock
}

type SNAC_0x01_0x04_TLVRoomInfo struct {
	Exchange       uint16
	Cookie         []byte `len_prefix:"uint8"`
	InstanceNumber uint16
}

type SNAC_0x01_0x05_OServiceServiceResponse struct {
	TLVRestBlock
}

type SNAC_0x01_0x07_OServiceRateParamsReply struct {
	RateClasses []struct {
		ID              uint16
		WindowSize      uint32
		ClearLevel      uint32
		AlertLevel      uint32
		LimitLevel      uint32
		DisconnectLevel uint32
		CurrentLevel    uint32
		MaxLevel        uint32
		LastTime        uint32
		CurrentState    uint8
	} `count_prefix:"uint16"`
	RateGroups []struct {
		ID    uint16
		Pairs []struct {
			FoodGroup uint16
			SubGroup  uint16
		} `count_prefix:"uint16"`
	}
}

type SNAC_0x01_0x08_OServiceRateParamsSubAdd struct {
	TLVRestBlock
}

type SNAC_0x01_0x0F_OServiceUserInfoUpdate struct {
	TLVUserInfo
}

type SNAC_0x01_0x10_OServiceEvilNotification struct {
	NewEvil uint16
	TLVUserInfo
}

type SNAC_0x01_0x10_OServiceEvilNotificationAnon struct {
	NewEvil uint16
}

type SNAC_0x01_0x11_OServiceIdleNotification struct {
	IdleTime uint32
}

type SNAC_0x01_0x17_OServiceClientVersions struct {
	Versions []uint16
}

type SNAC_0x01_0x18_OServiceHostVersions struct {
	Versions []uint16
}

type SNAC_0x01_0x1E_OServiceSetUserInfoFields struct {
	TLVRestBlock
}

//
// 0x02: Locate
//

const (
	LocateErr                  uint16 = 0x0001
	LocateRightsQuery          uint16 = 0x0002
	LocateRightsReply          uint16 = 0x0003
	LocateSetInfo              uint16 = 0x0004
	LocateUserInfoQuery        uint16 = 0x0005
	LocateUserInfoReply        uint16 = 0x0006
	LocateWatcherSubRequest    uint16 = 0x0007
	LocateWatcherNotification  uint16 = 0x0008
	LocateSetDirInfo           uint16 = 0x0009
	LocateSetDirReply          uint16 = 0x000A
	LocateGetDirInfo           uint16 = 0x000B
	LocateGetDirReply          uint16 = 0x000C
	LocateGroupCapabilityQuery uint16 = 0x000D
	LocateGroupCapabilityReply uint16 = 0x000E
	LocateSetKeywordInfo       uint16 = 0x000F
	LocateSetKeywordReply      uint16 = 0x0010
	LocateGetKeywordInfo       uint16 = 0x0011
	LocateGetKeywordReply      uint16 = 0x0012
	LocateFindListByEmail      uint16 = 0x0013
	LocateFindListReply        uint16 = 0x0014
	LocateUserInfoQuery2       uint16 = 0x0015

	LocateType2Sig          uint32 = 0x00000001
	LocateType2Unavailable  uint32 = 0x00000002
	LocateType2Capabilities uint32 = 0x00000004
	LocateType2Certs        uint32 = 0x00000008
	LocateType2HtmlInfo     uint32 = 0x00000400

	LocateTLVTagsInfoSigMime         uint16 = 0x01
	LocateTLVTagsInfoSigData         uint16 = 0x02
	LocateTLVTagsInfoUnavailableMime uint16 = 0x03
	LocateTLVTagsInfoUnavailableData uint16 = 0x04
	LocateTLVTagsInfoCapabilities    uint16 = 0x05
	LocateTLVTagsInfoCerts           uint16 = 0x06
	LocateTLVTagsInfoSigTime         uint16 = 0x0A
	LocateTLVTagsInfoUnavailableTime uint16 = 0x0B
	LocateTLVTagsInfoSupportHostSig  uint16 = 0x0C
	LocateTLVTagsInfoHtmlInfoData    uint16 = 0x0E
	LocateTLVTagsInfoHtmlInfoType    uint16 = 0x0D
)

type SNAC_0x02_0x03_LocateRightsReply struct {
	TLVRestBlock
}

type SNAC_0x02_0x04_LocateSetInfo struct {
	TLVRestBlock
}

type SNAC_0x02_0x06_LocateUserInfoReply struct {
	TLVUserInfo
	LocateInfo TLVRestBlock
}

type SNAC_0x02_0x09_LocateSetDirInfo struct {
	TLVRestBlock
}

type SNAC_0x02_0x0A_LocateSetDirReply struct {
	Result uint16
}

type SNAC_0x02_0x0B_LocateGetDirInfo struct {
	WatcherScreenNames string `len_prefix:"uint8"`
}

type SNAC_0x02_0x0F_LocateSetKeywordInfo struct {
	TLVRestBlock
}

type SNAC_0x02_0x10_LocateSetKeywordReply struct {
	Unknown uint16
}

type SNAC_0x02_0x15_LocateUserInfoQuery2 struct {
	Type2      uint32
	ScreenName string `len_prefix:"uint8"`
}

func (s SNAC_0x02_0x15_LocateUserInfoQuery2) RequestProfile() bool {
	return s.Type2&LocateType2Sig == LocateType2Sig
}

func (s SNAC_0x02_0x15_LocateUserInfoQuery2) RequestAwayMessage() bool {
	return s.Type2&LocateType2Unavailable == LocateType2Unavailable
}

//
// 0x03: Buddy
//

type SNAC_0x03_0x02_BuddyRightsQuery struct {
	TLVRestBlock
}

type SNAC_0x03_0x03_BuddyRightsReply struct {
	TLVRestBlock
}

type SNAC_0x03_0x0A_BuddyArrived struct {
	TLVUserInfo
}

type SNAC_0x03_0x0B_BuddyDeparted struct {
	TLVUserInfo
}

//
// 0x04: ICBM
//

const (
	ICBMErr                uint16 = 0x0001
	ICBMAddParameters      uint16 = 0x0002
	ICBMDelParameters      uint16 = 0x0003
	ICBMParameterQuery     uint16 = 0x0004
	ICBMParameterReply     uint16 = 0x0005
	ICBMChannelMsgToHost   uint16 = 0x0006
	ICBMChannelMsgToclient uint16 = 0x0007
	ICBMEvilRequest        uint16 = 0x0008
	ICBMEvilReply          uint16 = 0x0009
	ICBMMissedCalls        uint16 = 0x000A
	ICBMClientErr          uint16 = 0x000B
	ICBMHostAck            uint16 = 0x000C
	ICBMSinStored          uint16 = 0x000D
	ICBMSinListQuery       uint16 = 0x000E
	ICBMSinListReply       uint16 = 0x000F
	ICBMSinRetrieve        uint16 = 0x0010
	ICBMSinDelete          uint16 = 0x0011
	ICBMNotifyRequest      uint16 = 0x0012
	ICBMNotifyReply        uint16 = 0x0013
	ICBMClientEvent        uint16 = 0x0014
	ICBMSinReply           uint16 = 0x0017

	ICBMTLVTagRequestHostAck uint16 = 0x03
	ICBMTLVTagsWantEvents    uint16 = 0x0B
)

type SNAC_0x04_0x02_ICBMAddParameters struct {
	Channel              uint16
	ICBMFlags            uint32
	MaxIncomingICBMLen   uint16
	MaxSourceEvil        uint16
	MaxDestinationEvil   uint16
	MinInterICBMInterval uint32
}

type SNAC_0x04_0x05_ICBMParameterReply struct {
	MaxSlots             uint16
	ICBMFlags            uint32
	MaxIncomingICBMLen   uint16
	MaxSourceEvil        uint16
	MaxDestinationEvil   uint16
	MinInterICBMInterval uint32
}

type SNAC_0x04_0x06_ICBMChannelMsgToHost struct {
	Cookie     [8]byte
	ChannelID  uint16
	ScreenName string `len_prefix:"uint8"`
	TLVRestBlock
}

type SNAC_0x04_0x07_ICBMChannelMsgToClient struct {
	Cookie    [8]byte
	ChannelID uint16
	TLVUserInfo
	TLVRestBlock
}

type SNAC_0x04_0x08_ICBMEvilRequest struct {
	SendAs     uint16
	ScreenName string `len_prefix:"uint8"`
}

type SNAC_0x04_0x09_ICBMEvilReply struct {
	EvilDeltaApplied uint16
	UpdatedEvilValue uint16
}

type SNAC_0x04_0x0B_ICBMClientErr struct {
	Cookie     [8]byte
	ChannelID  uint16
	ScreenName string `len_prefix:"uint8"`
	Code       uint16
	ErrInfo    []byte
}

type SNAC_0x04_0x0C_ICBMHostAck struct {
	Cookie     [8]byte
	ChannelID  uint16
	ScreenName string `len_prefix:"uint8"`
}

type SNAC_0x04_0x14_ICBMClientEvent struct {
	Cookie     [8]byte
	ChannelID  uint16
	ScreenName string `len_prefix:"uint8"`
	Event      uint16
}

//
// 0x09: PD
//

type SNAC_0x09_0x03_PDRightsReply struct {
	TLVRestBlock
}

//
// 0x0D: ChatNav
//

const (
	ChatNavTLVRedirect     uint16 = 0x01
	ChatNavTLVMaxRooms     uint16 = 0x02
	ChatNavTLVExchangeInfo uint16 = 0x03
	ChatNavTLVRoomInfo     uint16 = 0x04
)

type SNAC_0x0D_0x04_ChatNavRequestRoomInfo struct {
	Exchange       uint16
	Cookie         []byte `len_prefix:"uint8"`
	InstanceNumber uint16
	DetailLevel    uint8
}

type SNAC_0x0D_0x09_ChatNavNavInfo struct {
	TLVRestBlock
}

type SNAC_0x0D_0x09_TLVExchangeInfo struct {
	Identifier uint16
	TLVBlock
}

//
// 0x0E: Chat
//

const (
	ChatTLVPublicWhisperFlag    uint16 = 0x01
	ChatTLVSenderInformation    uint16 = 0x03
	ChatTLVEnableReflectionFlag uint16 = 0x06
	ChatTLVRoomName             uint16 = 0xD3
)

type SNAC_0x0E_0x02_ChatRoomInfoUpdate struct {
	Exchange       uint16
	Cookie         string `len_prefix:"uint8"`
	InstanceNumber uint16
	DetailLevel    uint8
	TLVBlock
}

type SNAC_0x0E_0x03_ChatUsersJoined struct {
	Users []TLVUserInfo
}

type SNAC_0x0E_0x04_ChatUsersLeft struct {
	Users []TLVUserInfo
}

type SNAC_0x0E_0x05_ChatChannelMsgToHost struct {
	Cookie  uint64
	Channel uint16
	TLVRestBlock
}

type SNAC_0x0E_0x06_ChatChannelMsgToClient struct {
	Cookie  uint64
	Channel uint16
	TLVRestBlock
}

//
// 0x13: Feedbag
//

type SNAC_0x13_0x02_FeedbagRightsQuery struct {
	TLVRestBlock
}

type SNAC_0x13_0x03_FeedbagRightsReply struct {
	TLVRestBlock
}

type SNAC_0x13_0x05_FeedbagQueryIfModified struct {
	LastUpdate uint32
	Count      uint8
}

type SNAC_0x13_0x06_FeedbagReply struct {
	Version    uint8
	Items      []FeedbagItem `count_prefix:"uint16"`
	LastUpdate uint32
}

type SNAC_0x13_0x08_FeedbagInsertItem struct {
	Items []FeedbagItem
}

type SNAC_0x13_0x09_FeedbagUpdateItem struct {
	Items []FeedbagItem
}

type SNAC_0x13_0x0A_FeedbagDeleteItem struct {
	Items []FeedbagItem
}

type SNAC_0x13_0x0E_FeedbagStatus struct {
	Results []uint16
}

//
// 0x17: BUCP
//

type SNAC_0x17_0x02_BUCPLoginRequest struct {
	TLVRestBlock
}

type SNAC_0x17_0x03_BUCPLoginResponse struct {
	TLVRestBlock
}

type SNAC_0x17_0x06_BUCPChallengeRequest struct {
	TLVRestBlock
}

type SNAC_0x17_0x07_BUCPChallengeResponse struct {
	AuthKey string `len_prefix:"uint16"`
}

type SnacOServiceErr struct {
	Code uint16
}

type TLVUserInfo struct {
	ScreenName   string `len_prefix:"uint8"`
	WarningLevel uint16
	TLVBlock
}

type FeedbagItem struct {
	Name    string `len_prefix:"uint16"`
	GroupID uint16
	ItemID  uint16
	ClassID uint16
	TLVLBlock
}
