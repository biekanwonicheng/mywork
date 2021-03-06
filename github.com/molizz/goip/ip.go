package goip

import (
	"errors"
	"fmt"

	"github.com/molizz/goip/ipdb"
	"github.com/molizz/goip/platform"
)

type Platform interface {
	GetLocation(ip string) (*platform.Location, error)
}

type Locator interface {
	GetCountry() string
	GetRegion() string
	GetCity() string

	ToString() string
}

type Address struct {
	plts []Platform
}

var address *Address

func init() {
	err := ipdb.RestoreAsset("ipdb", "17monipdb.datx")
	if err != nil {
		fmt.Println("无法释放ip数据库(17monipdb.datx), 将无法正常使用ip获取服务: ", err)
	}

	address = &Address{
		plts: make([]Platform, 0, 1),
	}
}

func addPlatform(p Platform) {
	address.plts = append(address.plts, p)
}

func AddTaobao() *Address {
	addPlatform(platform.NewTaobao())
	return address
}

func AddAmap(key string) *Address {
	addPlatform(platform.NewAmap(key))
	return address
}

func AddTencent(key string) *Address {
	addPlatform(platform.NewTencent(key))
	return address
}

func AddChinaz() *Address {
	addPlatform(platform.NewChinaz())
	return address
}

func AddLocal() *Address {
	addPlatform(platform.NewLocal())
	return address
}

func GetLocation(ip string) (locator Locator, err error) {
	length := len(address.plts)
	if length == 0 {
		return nil, errors.New("Not found Platform")
	}

	for _, plt := range address.plts {
		locator, err = plt.GetLocation(ip)
		if err != nil {
			continue
		}
		return locator, nil
	}

	return nil, err
}
