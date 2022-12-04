package config

import (
	"os"
	"testing"
)

func Test_Database_Env_Variables(t *testing.T) {
	requiredEnvs := []string{
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
	}

	for _, envVar := range requiredEnvs {
		t.Run(envVar, func(t *testing.T) {
			_, exists := os.LookupEnv(envVar)
			if !exists {
				t.Errorf("Env variable %v, does not exist", envVar)
				return
			}
		})
	}
}
