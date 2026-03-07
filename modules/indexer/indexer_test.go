// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package indexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZoektSearchModes(t *testing.T) {
	assert.Contains(t, ZoektSearchModes(), SearchMode{
		ModeValue:    SearchModeZoekt,
		TooltipTrKey: "search.zoekt_tooltip",
		TitleTrKey:   "search.zoekt",
	})
}
