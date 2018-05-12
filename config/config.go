package config

import (
    "github.com/BurntSushi/toml"
    "os"
    "bytes"
)

type Config struct {
    ID         string `toml:"id"`
    DbUser     string `toml:"dbuser"`
    DbPassword string `toml:"dbpassword"`
}

func Parse(f string) (Config, error) {
    file, err := os.Open(f)

    if err != nil {
        return Config{}, err
    }

    if err != nil {
        return Config{}, err
    }

    conf := Config{}

    buf := bytes.NewBufferString("")

    _, err = buf.ReadFrom(file)

    if err != nil {
        return Config{}, err
    }

    _, err = toml.Decode(buf.String(), &conf)

    if err != nil {
        return Config{}, err
    }

    return conf, nil
}
