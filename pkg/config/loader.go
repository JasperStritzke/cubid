package config

import (
	"github.com/jasperstritzke/cubid/pkg/util/fileutil"
	"os"
)

func WrapExistingConfig(cfg interface{}) func() interface{} {
	return func() interface{} {
		return cfg
	}
}

func InitConfigIfNotExists(path string, configCallback func() interface{}) error {
	cfg := configCallback()

	file, err, created := fileutil.OpenFileOrCreate(path)

	if err != nil {
		return err
	}

	if created {
		encoder := fileutil.NewPrettyEncoder(file)

		encodeErr := encoder.Encode(cfg)

		if encodeErr != nil {
			return encodeErr
		}

		closeErr := file.Close()
		if closeErr != nil {
			return closeErr
		}

		return nil
	}

	return nil
}

func WriteConfig(path string, cfg interface{}) error {
	bytes, err := fileutil.MarshalIndent(cfg)

	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, os.ModePerm)
}

func LoadConfig(path string, config interface{}) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	decoder := fileutil.NewDecoder(file)

	decodeErr := decoder.Decode(config)

	if decodeErr != nil {
		return nil
	}

	closeErr := file.Close()
	if closeErr != nil {
		return closeErr
	}

	return nil
}
