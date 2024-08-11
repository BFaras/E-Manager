package cli

import (
   "github.com/spf13/cobra"
   "github.com/spf13/viper"
)

var (
   config string
   cfg    *viper.Viper
)

var cmd = &cobra.Command{
   Use:   "cmd",
   Short: "ShortDescription ",
   Long:  `Long Description`,
   PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
      cfg = viper.New()
      cfg.AutomaticEnv()

      cfg.SetConfigFile(config)
      cfg.ReadInConfig()
      return cfg.ReadInConfig()
   },
}

func Execute() error {
   return cmd.Execute()
}

func init() {
   cmd.PersistentFlags().StringVarP(
      &config, "config", "c", "../../.bin/config.dev.yaml",
   "path to file")
}