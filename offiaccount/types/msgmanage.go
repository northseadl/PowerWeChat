package types

// SetIndustryRequest 设置行业请求
type SetIndustryRequest struct {
	// 公众号模板消息所属行业编号
	IndustryID1 string `json:"industry_id1"`
	// 公众号模板消息所属行业编号
	IndustryID2 string `json:"industry_id2"`
}

// SetIndustryResponse 设置行业响应
type SetIndustryResponse struct {
	*Error
}

// Industry 行业
type Industry struct {
	// 主行业
	FirstClass string `json:"first_class"`
	// 副行业
	SecondClass string `json:"second_class"`
}

// GetIndustryResponse 获取设置的行业信息响应
type GetIndustryResponse struct {
	// 账号设置的主营行业
	PrimaryIndustry *Industry `json:"primary_industry"`
	// 账号设置的副营行业
	SecondaryIndustry *Industry `json:"secondary_industry"`
	*Error
}

// AddTemplateIDRequest 获取模板ID请求
type AddTemplateIDRequest struct {
	// 模板库中模板的编号，有“TM**”和“OPENTMTM**”等形式,对于类目模板，为纯数字ID
	TemplateIDShort string `json:"template_id_short"`
	// 选用的类目模板的关键词,按顺序传入,如果为空，或者关键词不在模板库中，会返回40246错误码
	KeywordIDList []int `json:"keyword_id_list"`
}

// AddTemplateIDResponse 获取模板ID响应
type AddTemplateIDResponse struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	*Error
}

// Template 模板
type Template struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板标题
	Title string `json:"title"`
	// 模板所属行业的一级行业
	PrimaryIndustry string `json:"primary_industry"`
	// 模板所属行业的二级行业
	DeputyIndustry string `json:"deputy_industry"`
	// 模板内容
	Content string `json:"content"`
	// 模板示例
	Example string `json:"example"`
}

// GetTemplateListResponse 获取模板列表响应
type GetAllTemplateListResponse struct {
	TemplateList []Template `json:"template_list"`
	*Error
}

// DelTemplateRequest 删除模板请求
type DelTemplateRequest struct {
	TemplateID string `json:"template_id"`
}

// DelTemplateResponse 删除模板响应
type DelTemplateResponse struct {
	*Error
}

// TemplateMessageRequest 发送模板消息请求
type SendTemplateMessageRequest struct {
	// 接收者openid
	Touser string `json:"touser"`
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板跳转链接（海外帐号没有跳转能力）
	URL string `json:"url,omitempty"`
	// 跳小程序所需数据，不需跳小程序可不用传该数据
	Miniprogram *TemplateMiniprogram `json:"miniprogram,omitempty"`
	// 模板数据
	Data map[string]KeywordData `json:"data"`
	// 防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填
	ClientMsgID string `json:"client_msg_id,omitempty"`
}

// TemplateMiniprogram 跳小程序所需数据，不需跳小程序可不用传该数据
type TemplateMiniprogram struct {
	// 所需跳转到的小程序appid
	AppID string `json:"appid"`
	// 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	PagePath string `json:"pagepath"`
}

// KeywordData 模板数据
type KeywordData struct {
	Value string `json:"value"`
}

// TemplateMessageResponse 发送模板消息响应
type SendTemplateMessageResponse struct {
	MsgID int64 `json:"msgid"`
	*Error
}
