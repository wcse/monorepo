package logging

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wcse/monorepo/backend/api/config/base"
)

func TestNewLogger(t *testing.T) {
	testConfigs := []*base.Logger{
		{
			Level:  base.Logger_INFO,
			Format: &base.Logger_Pretty{Pretty: true},
		},
		{
			Level: base.Logger_WARN,
		},
	}

	for idx, tc := range testConfigs {
		tc := tc
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()
			l, err := NewLogger(tc)
			assert.NotNil(t, l)
			assert.NoError(t, err)
		})
	}
}
