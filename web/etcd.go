// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/coreos/go-etcd/etcd"
	"github.com/golang/glog"
)

func Init_etcd() error {
	// Connect the etcd.
	client := etcd.NewClient(Conf.Etcd_addr)
	_, err := client.Set(Conf.Etcd_dir+"/"+Conf.Listen_addr, "running", Conf.Etcd_interval)
	if err != nil {
		return err
	}
	c_time := time.After(time.Duration(Conf.Etcd_interval / 2))
	go func() {
		for {
			select {
			case <-c_time:
				// Flush the etcd node time.
				if _, err = client.Update(Conf.Etcd_dir+"/"+Conf.Listen_addr, "running", Conf.Etcd_interval); err != nil {
					glog.Fatalf("Comet system will be closed ,err:%v\n", err)
				}
			}
		}
	}()
	return nil
}
