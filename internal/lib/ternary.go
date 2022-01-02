/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

*/
package lib

func Ternary(cond bool, truthy interface{}, falsy interface{}) interface{} {
	if cond {
		return truthy
	} else {
		return falsy
	}
}
