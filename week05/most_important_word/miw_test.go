package mostimportantword

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMostImportantWord(t *testing.T) {
	require.Equal(t, "sad", MostImportantWord("the eel the eel sad the sad the sad", []string{"the"}))
	assert.Equal(t, "", MostImportantWord("", []string{}))
}
