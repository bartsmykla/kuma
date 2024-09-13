package parameters

// Owner
//       This module attempts to match various characteristics of the packet creator,
//       for locally generated packets. This match is only valid in the OUTPUT and POSTROUTING
//       chains. Forwarded packets do not have any socket associated with them.
//       Packets from kernel threads do have a socket, but usually no owner.
//
//       [!] --uid-owner userid
//              Matches if the packet socket's file structure (if it has one) is owned by
//              the given user.
//
//       [!] --gid-owner groupid
//              Matches if the packet socket's file structure is owned by the given group.
//
// ref. iptables-extensions(8) > owner

import (
	"strconv"

	"github.com/kumahq/kuma/pkg/transparentproxy/config"
)

var _ ParameterBuilder = &OwnerParameter{}

type OwnerParameter struct {
	flag     string
	value    string
	negative bool
}

func (p *OwnerParameter) Negate() ParameterBuilder {
	p.negative = !p.negative

	return p
}

func (p *OwnerParameter) Build(bool) []string {
	if p.negative {
		return []string{"!", p.flag, p.value}
	}

	return []string{p.flag, p.value}
}

func uid(id uint64, negative bool) *OwnerParameter {
	return &OwnerParameter{
		flag:     "--uid-owner",
		value:    strconv.FormatUint(id, 10),
		negative: negative,
	}
}

// Uid matches if the packet socket's file structure (if it has one) is owned by the user
// with given UID
func Uid(id uint64) *OwnerParameter {
	return uid(id, false)
}

// UidRangeOrValue matches if the packet socket's file structure (if it has one) is owned by the user
// with given list of UID values or ranges
func UidRangeOrValue(exclusion config.Exclusion) *OwnerParameter {
	return &OwnerParameter{
		flag:  "--uid-owner",
		value: string(exclusion.UIDs),
	}
}

func NotUid(id uint64) *OwnerParameter {
	return uid(id, true)
}

// Owner attempts to match various characteristics of the packet creator,for locally generated
// packets. This match is only valid in the OUTPUT and POSTROUTING chains. Forwarded packets
// do not have any socket associated with them. Packets from kernel threads do have a socket,
// but usually no owner
func Owner(ownerParameters ...*OwnerParameter) *MatchParameter {
	var parameters []ParameterBuilder

	for _, parameter := range ownerParameters {
		parameters = append(parameters, parameter)
	}

	return &MatchParameter{
		name:       "owner",
		parameters: parameters,
	}
}
