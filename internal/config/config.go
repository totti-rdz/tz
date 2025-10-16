package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the structure of ~/.tz/config.json
type Config struct {
	Projects map[string]ProjectConfig `json:"projects"`
}

// ProjectConfig holds command mappings for a specific project
type ProjectConfig struct {
	Install string            `json:"install,omitempty"`
	Dev     string            `json:"dev,omitempty"`
	Test    string            `json:"test,omitempty"`
	Build   string            `json:"build,omitempty"`
	Clear   string            `json:"clear,omitempty"`
	Custom  map[string]string `json:"custom,omitempty"` // Custom user-defined commands
}

// configPath returns the path to the config file
func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ".tz", "config.json"), nil
}

// ensureConfigDir creates the ~/.tz directory if it doesn't exist
func ensureConfigDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	tzDir := filepath.Join(home, ".tz")
	if err := os.MkdirAll(tzDir, 0755); err != nil {
		return fmt.Errorf("failed to create .tz directory: %w", err)
	}
	return nil
}

// Load reads the config file and returns the Config
func Load() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	// If config doesn't exist, return empty config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &Config{Projects: make(map[string]ProjectConfig)}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Projects == nil {
		cfg.Projects = make(map[string]ProjectConfig)
	}

	return &cfg, nil
}

// Save writes the config to ~/.tz/config.json
func (c *Config) Save() error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

// GetCommand retrieves the command mapping for the current project
func (c *Config) GetCommand(projectPath, commandName string) (string, error) {
	projectCfg, exists := c.Projects[projectPath]
	if !exists {
		return "", fmt.Errorf("no configuration found for project: %s", projectPath)
	}

	var cmd string
	switch commandName {
	case "install":
		cmd = projectCfg.Install
	case "dev":
		cmd = projectCfg.Dev
	case "test":
		cmd = projectCfg.Test
	case "build":
		cmd = projectCfg.Build
	case "clear":
		cmd = projectCfg.Clear
	default:
		// Check custom commands
		if projectCfg.Custom != nil {
			if customCmd, ok := projectCfg.Custom[commandName]; ok {
				cmd = customCmd
			}
		}
		if cmd == "" {
			return "", fmt.Errorf("unknown command: %s", commandName)
		}
	}

	if cmd == "" {
		return "", fmt.Errorf("no mapping found for '%s' in project: %s", commandName, projectPath)
	}

	return cmd, nil
}

// SetCommand sets a command mapping for a project
func (c *Config) SetCommand(projectPath, commandName, command string) error {
	if c.Projects == nil {
		c.Projects = make(map[string]ProjectConfig)
	}

	projectCfg := c.Projects[projectPath]

	switch commandName {
	case "install":
		projectCfg.Install = command
	case "dev":
		projectCfg.Dev = command
	case "test":
		projectCfg.Test = command
	case "build":
		projectCfg.Build = command
	case "clear":
		projectCfg.Clear = command
	default:
		// Custom command
		if projectCfg.Custom == nil {
			projectCfg.Custom = make(map[string]string)
		}
		projectCfg.Custom[commandName] = command
	}

	c.Projects[projectPath] = projectCfg
	return nil
}

// GetCurrentProjectPath returns the absolute path of the current working directory
func GetCurrentProjectPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}
	return path, nil
}
