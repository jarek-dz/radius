// Code generated by radius-dict-gen. DO NOT EDIT.

package cryptotest

import (
	"crypto/rand"
	"net"
	"strconv"

	"github.com/holgermetschulat/radius"
)

const (
	CTInt_Type     radius.Type = 11
	CTOctects_Type radius.Type = 12
	CTIPADDR_Type  radius.Type = 13
)

type CTInt uint32

var CTInt_Strings = map[CTInt]string{}

func (a CTInt) String() string {
	if str, ok := CTInt_Strings[a]; ok {
		return str
	}
	return "CTInt(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CTInt_Add(p *radius.Packet, value CTInt) (err error) {
	a := radius.NewInteger(uint32(value))
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword(a, salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Add(CTInt_Type, a)
	return
}

func CTInt_Get(p *radius.Packet) (value CTInt) {
	value, _ = CTInt_Lookup(p)
	return
}

func CTInt_Gets(p *radius.Packet) (values []CTInt, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != CTInt_Type {
			continue
		}
		attr := avp.Attribute
		attr, _, err = radius.TunnelPassword(attr, p.Secret, p.CryptoAuthenticator[:])
		if err != nil {
			return
		}
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CTInt(i))
	}
	return
}

func CTInt_Lookup(p *radius.Packet) (value CTInt, err error) {
	a, ok := p.Lookup(CTInt_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	a, _, err = radius.TunnelPassword(a, p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CTInt(i)
	return
}

func CTInt_Set(p *radius.Packet, value CTInt) (err error) {
	a := radius.NewInteger(uint32(value))
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword(a, salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Set(CTInt_Type, a)
	return
}

func CTInt_Del(p *radius.Packet) {
	p.Attributes.Del(CTInt_Type)
}

func CTOctects_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword(value, salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Add(CTOctects_Type, a)
	return
}

func CTOctects_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword([]byte(value), salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Add(CTOctects_Type, a)
	return
}

func CTOctects_Get(p *radius.Packet) (value []byte) {
	value, _ = CTOctects_Lookup(p)
	return
}

func CTOctects_GetString(p *radius.Packet) (value string) {
	value, _ = CTOctects_LookupString(p)
	return
}

func CTOctects_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != CTOctects_Type {
			continue
		}
		attr := avp.Attribute
		i, _, err = radius.TunnelPassword(attr, p.Secret, p.CryptoAuthenticator[:])
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CTOctects_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != CTOctects_Type {
			continue
		}
		attr := avp.Attribute
		var up []byte
		up, _, err = radius.TunnelPassword(attr, p.Secret, p.CryptoAuthenticator[:])
		if err == nil {
			i = string(up)
		}
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CTOctects_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(CTOctects_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, _, err = radius.TunnelPassword(a, p.Secret, p.CryptoAuthenticator[:])
	return
}

func CTOctects_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(CTOctects_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var b []byte
	b, _, err = radius.TunnelPassword(a, p.Secret, p.CryptoAuthenticator[:])
	if err == nil {
		value = string(b)
	}
	return
}

func CTOctects_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword(value, salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Set(CTOctects_Type, a)
	return
}

func CTOctects_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword([]byte(value), salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Set(CTOctects_Type, a)
	return
}

func CTOctects_Del(p *radius.Packet) {
	p.Attributes.Del(CTOctects_Type)
}

func CTIPADDR_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword(a, salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Add(CTIPADDR_Type, a)
	return
}

func CTIPADDR_Get(p *radius.Packet) (value net.IP) {
	value, _ = CTIPADDR_Lookup(p)
	return
}

func CTIPADDR_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, avp := range p.Attributes {
		if avp.Type != CTIPADDR_Type {
			continue
		}
		attr := avp.Attribute
		attr, _, err = radius.TunnelPassword(attr, p.Secret, p.CryptoAuthenticator[:])
		if err != nil {
			return
		}
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CTIPADDR_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := p.Lookup(CTIPADDR_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	a, _, err = radius.TunnelPassword(a, p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func CTIPADDR_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	var salt [2]byte
	_, err = rand.Read(salt[:])
	if err != nil {
		return
	}
	salt[0] |= 1 << 7
	a, err = radius.NewTunnelPassword(a, salt[:], p.Secret, p.CryptoAuthenticator[:])
	if err != nil {
		return
	}
	p.Set(CTIPADDR_Type, a)
	return
}

func CTIPADDR_Del(p *radius.Packet) {
	p.Attributes.Del(CTIPADDR_Type)
}