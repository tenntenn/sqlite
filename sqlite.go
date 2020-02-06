// +build !darwin
// +build windows linux

package sqlite

import _ "modernc.org/sqlite"

const DriverName = "sqlite"
