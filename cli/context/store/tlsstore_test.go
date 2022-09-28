package store

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestTlsCreateUpdateGetRemove(t *testing.T) {
	testee := tlsStore{root: t.TempDir()}

	const contextName = "test-ctx"
	contextID := contextdirOf(contextName)

	_, err := testee.getData(contextName, "test-ep", "test-data")
	assert.Equal(t, true, IsErrTLSDataDoesNotExist(err))

	err = testee.createOrUpdate(contextName, "test-ep", "test-data", []byte("data"))
	assert.NilError(t, err)
	data, err := testee.getData(contextName, "test-ep", "test-data")
	assert.NilError(t, err)
	assert.Equal(t, string(data), "data")
	err = testee.createOrUpdate(contextName, "test-ep", "test-data", []byte("data2"))
	assert.NilError(t, err)
	data, err = testee.getData(contextName, "test-ep", "test-data")
	assert.NilError(t, err)
	assert.Equal(t, string(data), "data2")

	err = testee.remove(contextID, "test-ep", "test-data")
	assert.NilError(t, err)
	err = testee.remove(contextID, "test-ep", "test-data")
	assert.NilError(t, err)

	_, err = testee.getData(contextName, "test-ep", "test-data")
	assert.Equal(t, true, IsErrTLSDataDoesNotExist(err))
}

func TestTlsListAndBatchRemove(t *testing.T) {
	testee := tlsStore{root: t.TempDir()}

	all := map[string]EndpointFiles{
		"ep1": {"f1", "f2", "f3"},
		"ep2": {"f1", "f2", "f3"},
		"ep3": {"f1", "f2", "f3"},
	}

	ep1ep2 := map[string]EndpointFiles{
		"ep1": {"f1", "f2", "f3"},
		"ep2": {"f1", "f2", "f3"},
	}

	const contextName = "test-ctx"
	contextID := contextdirOf(contextName)
	for name, files := range all {
		for _, file := range files {
			err := testee.createOrUpdate(contextName, name, file, []byte("data"))
			assert.NilError(t, err)
		}
	}

	resAll, err := testee.listContextData(contextID)
	assert.NilError(t, err)
	assert.DeepEqual(t, resAll, all)

	err = testee.removeAllEndpointData(contextName, "ep3")
	assert.NilError(t, err)
	resEp1ep2, err := testee.listContextData(contextID)
	assert.NilError(t, err)
	assert.DeepEqual(t, resEp1ep2, ep1ep2)

	err = testee.removeAllContextData(contextName)
	assert.NilError(t, err)
	resEmpty, err := testee.listContextData(contextID)
	assert.NilError(t, err)
	assert.DeepEqual(t, resEmpty, map[string]EndpointFiles{})
}
