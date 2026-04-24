// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_26

import (
	"xorm.io/xorm"
)

func AddOrgParentID(x *xorm.Engine) error {
	type User struct {
		ParentID int64 `xorm:"INDEX DEFAULT 0"`
	}
	_, err := x.SyncWithOptions(xorm.SyncOptions{IgnoreDropIndices: true}, new(User))
	return err
}
