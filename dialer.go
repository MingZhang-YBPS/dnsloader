package main

import (
	"context"
	"net"
)

type Dialer struct {
	dialer     *net.Dialer
	resolver   *net.Resolver
	nameserver string
}

// NewDialer create a Dialer with user's nameserver.
func NewDialer(dialer *net.Dialer, nameserver string, protocol string) (*Dialer, error) {
	conn, err := dialer.Dial(protocol, nameserver)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return &Dialer{
		dialer: dialer,
		resolver: &net.Resolver{
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return dialer.DialContext(ctx, protocol, nameserver)
			},
		},
		nameserver: nameserver, // 用户设置的nameserver
	}, nil
}

// DialContext connects to the address on the named network using
// the provided context.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	ips, err := d.resolver.LookupHost(ctx, host) // 通过自定义nameserver查询域名
	for _, ip := range ips {
		// 创建链接
		conn, err := d.dialer.DialContext(ctx, network, ip+":"+port)
		if err == nil {
			return conn, nil
		}
	}

	return d.dialer.DialContext(ctx, network, address)
}
