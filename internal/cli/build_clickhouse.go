// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/tooolbox/migrate/v4/database/clickhouse"
)
