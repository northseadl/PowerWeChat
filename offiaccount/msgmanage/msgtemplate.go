package msgmanage

import (
	"context"

	"github.com/ArtisanCloud/PowerWeChat/pro/offiaccount/core"
	"github.com/ArtisanCloud/PowerWeChat/pro/offiaccount/types"
)

const (
	setIndustryURI         = "/template/api_set_industry"
	getIndustryURI         = "/template/get_industry"
	addTemplateURI         = "/template/api_add_template"
	getAllTemplateListURI  = "/template/get_all_private_template"
	delTemplateURI         = "/template/del_private_template"
	sendTemplateMessageURI = "/message/template/send"
)

type MessageTemplate struct {
	*core.Provider
}

func New(app *core.Provider) *MessageTemplate {
	return &MessageTemplate{
		app,
	}
}

// SetIndustry 设置所属行业
func (comp *MessageTemplate) SetIndustry(ctx context.Context, req *types.SetIndustryRequest) (result *types.SetIndustryResponse, err error) {
	result = &types.SetIndustryResponse{}
	err = comp.H.Df().WithContext(ctx).Method("POST").Uri(setIndustryURI).Json(req).Result(result)
	if types.HandelError(err, result.Error); err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndustry 获取设置的行业信息
func (comp *MessageTemplate) GetIndustry(ctx context.Context) (result *types.GetIndustryResponse, err error) {
	result = &types.GetIndustryResponse{}
	err = comp.H.Df().WithContext(ctx).Method("GET").Uri(getIndustryURI).Result(result)
	if types.HandelError(err, result.Error); err != nil {
		return nil, err
	}
	return result, nil
}

// AddTemplate 获得模板ID
func (comp *MessageTemplate) AddTemplate(ctx context.Context, req *types.AddTemplateIDRequest) (result *types.AddTemplateIDResponse, err error) {
	result = &types.AddTemplateIDResponse{}
	err = comp.H.Df().WithContext(ctx).Method("POST").Uri(addTemplateURI).Json(req).Result(result)
	if types.HandelError(err, result.Error); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllTemplateList 获取所有模板列表
func (comp *MessageTemplate) GetAllTemplateList(ctx context.Context) (result *types.GetAllTemplateListResponse, err error) {
	result = &types.GetAllTemplateListResponse{}
	err = comp.H.Df().WithContext(ctx).Method("GET").Uri(getAllTemplateListURI).Result(result)
	if types.HandelError(err, result.Error); err != nil {
		return nil, err
	}
	return result, nil
}

// DelTemplate 删除模板
func (comp *MessageTemplate) DelTemplate(ctx context.Context, req *types.DelTemplateRequest) (result *types.DelTemplateResponse, err error) {
	result = &types.DelTemplateResponse{}
	err = comp.H.Df().WithContext(ctx).Method("POST").Uri(delTemplateURI).Json(req).Result(result)
	if types.HandelError(err, result.Error); err != nil {
		return nil, err
	}
	return result, nil
}

// SendTemplateMessage 发送模板消息
func (comp *MessageTemplate) SendTemplateMessage(ctx context.Context, req *types.SendTemplateMessageRequest) (result *types.SendTemplateMessageResponse, err error) {
	result = &types.SendTemplateMessageResponse{}
	err = comp.H.Df().WithContext(ctx).Method("POST").Uri(sendTemplateMessageURI).Json(req).Result(result)
	if types.HandelError(err, result.Error); err != nil {
		return nil, err
	}
	return result, nil
}
