// +build !darwin,windows !darwin,linux

package sqlite

import _ "modernc.org/sqlite"

const DriverName = "sqlite"
