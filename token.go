package token

// token generator package
// Copyright (C) 2016  <geosoft1@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import (
	"crypto/md5"
	"fmt"
	"time"
)

//GetToken generate random strings of characters of a given size. GetToken
//use current time to avoid overlapping probability.
func GetToken(index int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(time.Now()))))[0:index]
}
