package PowerWeChat

import (
	"encoding/json"
	"github.com/artisancloud/httphelper/dataflow"
	"github.com/pkg/errors"
)

type Result interface {
	Error() error
}

type ResHelper[T Result] struct {
	df *dataflow.Dataflow
}

func (r *ResHelper[T]) Result() (result *T, err error) {
	res, err := r.df.RequestResHelper()
	if err != nil {
		return nil, err
	}
	bodyBytes, err := res.GetBodyBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get response body bytes")
	}
	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body")
	}
	if (*result).Error() != nil {
		return result, (*result).Error()
	}
	return result, nil
}

func (r *ResHelper[T]) ToMap() (result map[string]any, err error) {
	res, err := r.df.RequestResHelper()
	if err != nil {
		return nil, err
	}
	return res.GetBodyJsonAsMap()
}
