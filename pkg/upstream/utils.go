/*
 * Copyright (C) 2020-2022, IrineSistiana
 *
 * This file is part of mosdns.
 *
 * mosdns is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * mosdns is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package upstream

import (
	"context"
	"fmt"
	"golang.org/x/net/proxy"
	"net"
	"net/url"
)

type socketOpts struct {
	so_mark        int
	bind_to_device string
}

func dialTCP(ctx context.Context, addr, proxyURLStr string, dialer *net.Dialer) (net.Conn, error) {
 
	if len(proxyURLStr) > 0 {
		proxyURL, err := url.Parse(proxyURLStr)
		proxyDialer, err := proxy.FromURL(proxyURL, dialer)
		if err != nil {
			return nil, fmt.Errorf("failed to init socks5 dialer: %w", err)
		}
		return proxyDialer.(proxy.ContextDialer).DialContext(ctx, "tcp", addr)
	}

	return dialer.DialContext(ctx, "tcp", addr)
}
