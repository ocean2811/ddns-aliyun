package ddns

import (
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

	"github.com/pkg/errors"
)

const (
	gcConnectTimeout = time.Second * 5
	gcReadTimeout    = time.Second * 5
)

// DDNS provide DDNS modification and query
type DDNS struct {
	accessKeyID     string
	accessKeySecret string

	client *alidns.Client
}

// NewDDNSAndConnect new DDNS and connect aliyun API Server
func NewDDNSAndConnect(accessKeyID, accessKeySecret string) (*DDNS, error) {
	ddns := &DDNS{accessKeyID: accessKeyID, accessKeySecret: accessKeySecret}

	client, err := alidns.NewClientWithAccessKey("", ddns.accessKeyID, ddns.accessKeySecret)
	if err != nil {
		return nil, errors.Wrap(err, "alidns NewClientWithAccessKey has error")
	}
	ddns.client = client

	return ddns, nil
}

// DescribeDomainRecords Describe all subDomain list
func (ddns *DDNS) DescribeDomainRecords(domain string) ([]alidns.Record, error) {
	req := alidns.CreateDescribeDomainRecordsRequest()

	records := []alidns.Record{}

	pageNum := 1
	for {
		req.DomainName = domain
		req.PageNumber = requests.NewInteger(pageNum)
		req.SetConnectTimeout(gcConnectTimeout)
		req.SetReadTimeout(gcReadTimeout)

		resp, err := ddns.client.DescribeDomainRecords(req)
		if err != nil {
			return nil, errors.Wrap(err, "alidns DescribeSubDomainRecords has error")
		}

		if resp.PageSize <= 0 || len(resp.DomainRecords.Record) == 0 {
			break
		}

		records = append(records, resp.DomainRecords.Record...)
		pageNum++
	}

	return records, nil
}

// DescribeSubDomainRecords Describe subDomain
func (ddns *DDNS) DescribeSubDomainRecords(subDomain string) ([]alidns.Record, error) {
	req := alidns.CreateDescribeSubDomainRecordsRequest()
	req.SetConnectTimeout(gcConnectTimeout)
	req.SetReadTimeout(gcReadTimeout)

	req.SubDomain = subDomain
	if strings.HasPrefix(subDomain, "*.") {
		req.SubDomain = subDomain
		req.DomainName = subDomain[2:]
	}

	resp, err := ddns.client.DescribeSubDomainRecords(req)
	if err != nil {
		return nil, errors.Wrap(err, "alidns DescribeSubDomainRecords has error")
	}

	// fmt.Println(resp.String())

	if !resp.IsSuccess() {
		return nil, errors.New("alidns DescribeSubDomainRecords is not success: " + resp.String())
	}

	records := resp.DomainRecords.Record

	// for _, r := range records {
	// 	return r.RecordId, r.Type, r.Value, nil
	// }

	return records, nil
}

// UpdateDomainRecord Update Domain Record
func (ddns *DDNS) UpdateDomainRecord(recordID, typ, rr, value string) error {
	req := alidns.CreateUpdateDomainRecordRequest()
	req.Type = typ
	req.RecordId = recordID
	req.RR = rr
	req.Value = value
	req.SetConnectTimeout(gcConnectTimeout)
	req.SetReadTimeout(gcReadTimeout)

	resp, err := ddns.client.UpdateDomainRecord(req)
	if err != nil {
		return errors.Wrap(err, "alidns UpdateDomainRecord has error")
	}

	// fmt.Println(resp.String())

	if !resp.IsSuccess() {
		return errors.New("alidns UpdateDomainRecord is not success: " + resp.String())
	}

	return nil
}

// AddDomainRecord Add Domain Record
func (ddns *DDNS) AddDomainRecord(typ, domain, value string) error {
	req := alidns.CreateAddDomainRecordRequest()
	req.Type = typ
	req.Value = value
	req.SetConnectTimeout(gcConnectTimeout)
	req.SetReadTimeout(gcReadTimeout)

	req.RR = "@"
	req.DomainName = domain
	domainSlice := strings.Split(domain, ".")
	if len(domainSlice) > 2 {
		req.RR = strings.Join(domainSlice[:len(domainSlice)-2], ".")
		req.DomainName = strings.Join(domainSlice[len(domainSlice)-2:], ".")
	}

	resp, err := ddns.client.AddDomainRecord(req)
	if err != nil {
		return errors.Wrap(err, "alidns AddDomainRecord has error")
	}

	// fmt.Println(resp.String())

	if !resp.IsSuccess() {
		return errors.New("alidns AddDomainRecord is not success: " + resp.String())
	}

	return nil
}
