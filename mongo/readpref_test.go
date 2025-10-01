package mongo

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestResolveReadPrefUsesConfiguredForFind(t *testing.T) {
	secondary, err := readpref.New(readpref.SecondaryMode)
	require.NoError(t, err)

	m := Mongo{defaultReadPref: secondary}

	rp := m.resolveReadPref(nil, false, Find, &opMsg{})
	require.Equal(t, readpref.SecondaryMode, rp.Mode())
}

func TestResolveReadPrefPrefersPrimaryForWrites(t *testing.T) {
	secondary, err := readpref.New(readpref.SecondaryMode)
	require.NoError(t, err)

	m := Mongo{defaultReadPref: secondary}

	rp := m.resolveReadPref(nil, false, Insert, &opMsg{})
	require.Equal(t, readpref.PrimaryMode, rp.Mode())
}

func TestResolveReadPrefHonorsProvidedPreference(t *testing.T) {
	secondary, err := readpref.New(readpref.SecondaryMode)
	require.NoError(t, err)
	m := Mongo{defaultReadPref: secondary}

	nearest, err := readpref.New(readpref.NearestMode)
	require.NoError(t, err)

	rp := m.resolveReadPref(nearest, true, Find, &opMsg{})
	require.Same(t, nearest, rp)
}

func TestResolveReadPrefSupportsLegacyQueries(t *testing.T) {
	secondary, err := readpref.New(readpref.SecondaryMode)
	require.NoError(t, err)
	m := Mongo{defaultReadPref: secondary}

	rp := m.resolveReadPref(nil, false, Unknown, &opQuery{})
	require.Equal(t, readpref.SecondaryMode, rp.Mode())
}
