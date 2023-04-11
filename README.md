# 百度营销 MarketingAPI Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/baidu-marketing.svg)](https://pkg.go.dev/github.com/bububa/baidu-marketing)
[![Go](https://github.com/bububa/baidu-marketing/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/baidu-marketing/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/baidu-marketing/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/baidu-marketing/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/baidu-marketing.svg)](https://github.com/bububa/baidu-marketing)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/baidu-marketing)](https://goreportcard.com/report/github.com/bububa/baidu-marketing)
[![GitHub license](https://img.shields.io/github/license/bububa/baidu-marketing.svg)](https://github.com/bububa/baidu-marketing/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/baidu-marketing.svg)](https://GitHub.com/bububa/baidu-marketing/releases/)

- 账户管理
  - 财务管理 (api/account/balance)
    - 查询账户余额成分 [ GetBalanceInfo(clt \*core.SDKClient, auth model.RequestHeader, productIds[]int64) ([]balance.BalanceInfo, error) ]
    - 查询转账记录 [ GetAccountTransferHistory(clt \*core.SDKClient, auth model.RequestHeader, startTime time.Time, endTime time.Time) ([]balance.AccountTransferHistory, error) ]
    - 查询待加款信息 [ GetPaymentHistory(clt \*core.SDK, auth model.RequestHeader, reqBody balance.GetPaymentHistoryRequest) ([]balance.PaymentHistory, error) ]
    - 查询付款信息与待加款信息 [ GetPaymentRecord(clt *core.SDK, auth model.RequestHeader, reqBody balance.GetPaymentRecordRequest) (*balance.GetPaymentRecordResponse, error) ]
  - 账户管家管理 [ GetUserListByMccid(clt \*core.SDK, auth model.RequestHeader) ([]account.MccUser, error) ]
- 搜索广告投放 (api/search)
  - 账户 (api/search/account)
    - 查询账户 [ GetAccountInfo(clt \*core.SDKClient, auth model.RequestHeader, accountFields []string) ([]account.Account, error) ]
    - 更新账户 [ UpdateAccountInfo(clt *core.SDKClient, auth model.RequestHeader, account *account.Account) ([]account.Account, error) ]
  - 计划 (api/search/campaign)
    - 查询计划 [ GetCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.GetCampaignRequest) ([]campaign.Campaign, error) ]
    - 添加计划 [ AddCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.AddCampaignRequest) ([]campaign.Campaign, error) ]
    - 更新计划 [ UpdateCampaign(clt \*core.SDKClient, auth model.RequestHeader, campaigns []campaign.Campaign) ([]campaign.Campaign, error) ]
    - 删除计划 [ DeleteCampaign(clt \*core.SDKClient, auth model.RequestHeader, campaignIds []int64) error ]
  - 单元 (api/search/adgroup)
    - 查询单元 [ GetAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *adgroup.GetAdgroupRequest) ([]adgroup.Adgroup, error) ]
    - 添加单元 [ AddAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *adgroup.AddAdgroupRequest) ([]adgroup.Adgroup, error) ]
    - 更新单元 [ UpdateAdgroup(clt \*core.SDKClient, auth model.RequestHeader, adgroups []adgroup.Adgroup) ([]adgroup.Adgroup, error) ]
    - 删除单元 [ DeleteAdgroup(clt \*core.SDKClient, auth model.RequestHeader, adgroupIds []int64) error ]
  - 创意 (api/search/creative)
    - 查询创意 [ GetCreative(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.GetCreativeRequest) ([]creative.Creative, error) ]
    - 添加创意 [ AddCreative(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.AddCreativeRequest) ([]creative.Creative, error) ]
    - 更新创意 [ UpdateCreative(clt \*core.SDKClient, auth model.RequestHeader, creatives []creative.Creative) ([]creative.Creative, error) ]
    - 删除创意 [ DeleteCreative(clt \*core.SDKClient, auth model.RequestHeader, creativeIds []int64) error ]
- 信息流广告投放 (api/feed)
  - 账户 (api/feed/account)
    - 查询账户 [ GetAccountFeed(clt \*core.SDKClient, auth model.RequestHeader, accountFields []string) ([]account.Account, error) ]
    - 更新账户 [ UpdateAccountFeed(clt \*core.SDKClient, auth model.RequestHeader, budget float64) ([]account.Account, error) ]
  - 计划 (api/feed/campaign)
    - 查询计划 [ GetCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.GetCampaignFeedRequest) ([]campaign.Campaign, error) ]
    - 添加计划 [ AddCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.AddCampaignRequest) ([]campaign.Campaign, error) ]
    - 更新计划 [ UpdateCampaign(clt \*core.SDKClient, auth model.RequestHeader, campaigns []campaign.Campaign) ([]campaign.Campaign, error) ]
    - 删除计划 [ DeleteCampaign(clt \*core.SDKClient, auth model.RequestHeader, campaignIds []int64) error ]
  - 单元
    - 原生推广单元 (api/feed/adgroup/native)
      - 查询原生推广单元 [ GetAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *native.GetAdgroupFeedRequest) ([]native.Adgroup, error) ]
      - 添加单元 [ AddAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *native.AddAdgroupRequest) ([]native.Adgroup, error) ]
      - 更新单元 [ UpdateAdgroup(clt \*core.SDKClient, auth model.RequestHeader, adgroups []native.Adgroup) (*model.ResponseHeader, []native.Adgroup, error) ]
      - 删除单元 [ DeleteAdgroup(clt \*core.SDKClient, auth model.RequestHeader, adgroupFeedIds []int64) error ]
  - 商品推广单元 (api/feed/adgroup/dpa)
      - 查询商品推广单元 [ GetAdgroup(clt *core.SDKClient, auth model.RequestHeader, reqBody *dpa.GetAdgroupFeedRequest) ([]dpa.Adgroup, error) ]
  - 创意 (api/feed/creative)
    - 查询创意 [ GetCreativeFeed(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.GetCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) ]
    - 添加创意 [ AddCreative(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.AddCreativeRequest) ([]creative.Creative, error) ]
    - 更新创意 [ UpdateCreative(clt \*core.SDKClient, auth model.RequestHeader, creatives []creative.Creative) ([]creative.Creative, error) ]
    - 删除创意 [ DeleteCreative(clt \*core.SDKClient, auth model.RequestHeader, creativeFeedIds []int64) error ]
- 搜索报告 (api/search/report)
  - 推广报告 [ GetRealTimeData(clt *core.SDKClient, auth model.RequestHeader, realTimeRequest *report.ReqlTimeRequest) ([]report.RealTimeResult, error) ]
  - 账户实时数据 [ GetAccountLiveData(clt \*core.SDKClient, auth model.RequestHeader, dataType int, device int) ([]report.AccountLiveData, error) ]
  - 关键词实时数据 [ GetKeywordLiveData(clt *core.SDKClient, auth model.RequestHeader, reqBody *report.GetKeywordLiveDataRequest) ([]report.KeywordLiveData, error) ]
  - 创建异步报告，获取报告 ID(reportId) [ GetProfessionalReportId(clt *core.SDKClient, auth model.RequestHeader, reqBody *report.ReportRequest) (string, error) ]
  - 获取异步报告状态 [ GetReportState(clt *core.SDKClient, auth model.RequestHeader, reportId string) (int, error) ]
  - 获取异步报告文件 URL [ GetReportFileUrl(clt *core.SDKClient, auth model.RequestHeader, reportId string) (string, error) ]
- 信息流报告 (api/feed/report)
  - 推广报告 [ GetRealTimeData(clt *core.SDKClient, auth model.RequestHeader, realTimeRequest *report.ReqlTimeRequest) ([]report.RealTimeResult, error) ]
  - 创建异步报告，获取报告 ID(reportId) [ GetReportFeedId(clt *core.SDKClient, auth model.RequestHeader, reqBody *report.ReportRequest) (string, error) ]
  - 获取异步报告状态 [ GetReportFeedState(clt *core.SDKClient, auth model.RequestHeader, reportId string) (int, error) ]
  - 获取异步报告文件 URL [ GetReportFeedFileUrl(clt *core.SDKClient, auth model.RequestHeader, reportId string) (string, error) ]
- 资产管理 (api/asset)
  - 图片 (api/asset/image)
    - 查询图片 [ GetImage(clt *core.SDKClient, auth model.RequestHeader, reqBody *image.GetImageRequest) ([]image.Image, error) ]
  - 视频 (api/asset/video)
    - 查询图片 [ GetVideo(clt *core.SDKClient, auth model.RequestHeader, reqBody *image.GetVideoRequest) ([]video.Video, error) ]
- 转化上报 (api/ocpc)
  - 广告主回传转化数据接口 [ UploadConvertData(clt *core.SDKClient, req *ocpc.UploadConvertDataRequest) error ]
  - 广告主回传无效转化数据接口 [ UploadInvalidConvertData(clt *core.SDKClient, req *ocpc.UploadInvalidConvertDataRequest) error ]
  - APP 转化数据收集 [ ActionCb(req *ocpc.ActionCbRequest) error ]
- 数据报告（api/report）
  - 信息流报告
    - 推广报告
      - 单元报告 [ GetAdgroupReportData(clt *core.SDKClient, auth model.RequestHeader, reportRequest report.GetReportDataRequest) (*model.ResponseHeader, []report.GetAdgroupReportData, error) ]