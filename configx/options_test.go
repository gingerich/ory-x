// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	ctx := context.Background()

	t.Run("case=does not load env if disabled", func(t *testing.T) {
		schema := `{"type": "object", "properties": {"path": {"type": "string"}}}`

		envP, err := New(ctx, []byte(schema))
		require.NoError(t, err)
		assert.NotZero(t, envP.String("path"))

		nonEnvP, err := New(ctx, []byte(schema), DisableEnvLoading())
		require.NoError(t, err)
		assert.Nil(t, nonEnvP.Get("path"))
	})
}
