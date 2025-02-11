// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package dockertest

import (
	"testing"

	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/stretchr/testify/require"
)

func NewLocalTestCRDBServer(t testing.TB) string {
	ts, err := testserver.NewTestServer()
	require.NoError(t, err)
	t.Cleanup(ts.Stop)

	require.NoError(t, ts.WaitForInit())

	ts.PGURL().Scheme = "cockroach"
	return ts.PGURL().String()
}
