package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Run("Test with valid config", func(t *testing.T) {
		viper.Set("server.port", 8080)
		viper.Set("db.host", "localhost")
		viper.Set("db.port", 5432)
		viper.Set("db.user", "testuser")
		viper.Set("db.password", "testpassword")
		viper.Set("db.dbname", "testdb")
		viper.Set("db.sslmode", "disable")
		viper.Set("db.timezone", "UTC")

		// Get the config instance
		config := GetConfig()

		// Assertions
		assert.NotNil(t, config)
		assert.NotNil(t, config.Server)
		assert.NotNil(t, config.Db)
		assert.Equal(t, 8080, config.Server.Port)
		assert.Equal(t, "localhost", config.Db.Host)
		assert.Equal(t, 5432, config.Db.Port)
		assert.Equal(t, "testuser", config.Db.User)
		assert.Equal(t, "testpassword", config.Db.Password)
		assert.Equal(t, "testdb", config.Db.DBName)
		assert.Equal(t, "disable", config.Db.SSLMode)
		assert.Equal(t, "UTC", config.Db.TimeZone)
	})
}
