package openidmass

const msgToUserCountLimit = 10000

type commonMsgHead struct {
	ToUser  []string `json:"touser"`
	MsgType string   `json:"msgtype"`
}

func (msg *commonMsgHead) AppendToUser(touser ...string) {
	if len(touser) <= 0 {
		return
	}
	if len(msg.ToUser) >= msgToUserCountLimit {
		return
	}

	if n := msgToUserCountLimit - len(msg.ToUser); len(touser) > n {
		touser = touser[:n]
	}

	msg.ToUser = append(msg.ToUser, touser...)
}

// news ========================================================================

type NewsMsg struct {
	commonMsgHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNewsMsg(touser []string, mediaId string) *NewsMsg {
	if len(touser) > msgToUserCountLimit {
		touser = touser[:msgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg NewsMsg
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

// text ========================================================================

type TextMsg struct {
	commonMsgHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewTextMsg(touser []string, content string) *TextMsg {
	if len(touser) > msgToUserCountLimit {
		touser = touser[:msgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg TextMsg
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// voice =======================================================================

type VoiceMsg struct {
	commonMsgHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoiceMsg(touser []string, mediaId string) *VoiceMsg {
	if len(touser) > msgToUserCountLimit {
		touser = touser[:msgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg VoiceMsg
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// image =======================================================================

type ImageMsg struct {
	commonMsgHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImageMsg(touser []string, mediaId string) *ImageMsg {
	if len(touser) > msgToUserCountLimit {
		touser = touser[:msgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg ImageMsg
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// voice =======================================================================

type VideoMsg struct {
	commonMsgHead

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

func NewVideoMsg(touser []string, mediaId, title, description string) *VideoMsg {
	if len(touser) > msgToUserCountLimit {
		touser = touser[:msgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg VideoMsg
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}