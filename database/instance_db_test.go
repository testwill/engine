package database

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/instance"
	"github.com/stretchr/testify/require"
)

func instancedb(t *testing.T, dir string) InstanceDB {
	db, err := NewInstanceDB(dir)
	require.NoError(t, err)
	return db
}

func TestFindInstance(t *testing.T) {
	dir, _ := ioutil.TempDir("", "TestFindInstance")
	defer os.RemoveAll(dir)
	db := instancedb(t, dir)
	defer db.Close()

	i := &instance.Instance{Hash: hash.Int(1)}
	db.Save(i)
	tests := []struct {
		hash     hash.Hash
		hasError bool
	}{
		{hash: i.Hash, hasError: false},
		{hash: hash.Int(2), hasError: true},
	}

	for _, test := range tests {
		instance, err := db.Get(test.hash)
		if test.hasError {
			require.Error(t, err)
			continue
		}

		require.NoError(t, err)
		require.Equal(t, instance, i)
	}
}

func TestSaveInstance(t *testing.T) {
	dir, _ := ioutil.TempDir("", "TestSaveInstance")
	defer os.RemoveAll(dir)
	db := instancedb(t, dir)
	defer db.Close()
	tests := []struct {
		instance *instance.Instance
		hasError bool
	}{
		{&instance.Instance{Hash: hash.Int(1)}, false},
		{&instance.Instance{}, true},
	}
	for _, test := range tests {
		err := db.Save(test.instance)
		if test.hasError {
			require.Error(t, err)
			continue
		}
		require.NoError(t, err)
	}
}

func TestDeleteInstance(t *testing.T) {
	dir, _ := ioutil.TempDir("", "TestDeleteInstance")
	defer os.RemoveAll(dir)
	db := instancedb(t, dir)
	defer db.Close()
	i := &instance.Instance{Hash: hash.Int(1)}
	db.Save(i)
	require.NoError(t, db.Delete(i.Hash))
	_, err := db.Get(i.Hash)
	require.Error(t, err)

	require.NoError(t, db.Delete(i.Hash))
}