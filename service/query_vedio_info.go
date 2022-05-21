package service

import (
	"github.com/yintamaa/998244353Plus1/repository"
	"github.com/yintamaa/998244353Plus1/utils"
)

type VedioInfo struct {
	vedios []*repository.Vedio
}

var next_time int64

func QueryVedioInfoFlow() ([]*repository.Vedio, error) {
	var feed_query_flow VedioInfo
	feed_query_flow.Do()
	return feed_query_flow.vedios, nil
}

func (f *VedioInfo) Do() (*VedioInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	return f, nil
}
func (f *VedioInfo) checkParam() error {
	return nil
}
func (f *VedioInfo) prepareInfo() error {
	vedios, err := repository.NewVedioDaoInstance().QueryVedioRecent(1)
	if err != nil {
		utils.GetLoggerInstance().Error(":" + err.Error())
	}
	f.vedios = vedios
	next_time = vedios[0].CreateTime
	return nil
}
