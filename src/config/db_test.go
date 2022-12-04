package config

import (
	"os"
	"testing"
)

func Test_Database_System_Env_Variables(t *testing.T) {
	requiredEnvs := []string{
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
		"POSTGRES_DRIVER",
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

func Test_Database_Get_Env_Variables(t *testing.T) {
	t.Run(envVar, func(t *testing.T) {
		config := getDatabaseConfig()

		if config.DRIVER == "" {
			t.Errorf("Env variable DRIVER does not exist")
			return
		}
		if config.POSTGRES_USER == "" {
			t.Errorf("Env variable DRIVER does not exist")
			return
		}
		if config.DRIVER == "" {
			t.Errorf("Env variable DRIVER does not exist")
			return
		}
		if config.DRIVER == "" {
			t.Errorf("Env variable DRIVER does not exist")
			return
		}
		if config.DRIVER == "" {
			t.Errorf("Env variable DRIVER does not exist")
			return
		}
	})
}
