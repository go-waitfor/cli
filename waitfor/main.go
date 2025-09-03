package main

import (
	"fmt"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-fs"
	"github.com/go-waitfor/waitfor-http"
	"github.com/go-waitfor/waitfor-mongodb"
	"github.com/go-waitfor/waitfor-mysql"
	"github.com/go-waitfor/waitfor-postgres"
	"github.com/go-waitfor/waitfor-proc"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var version string

func main() {
	app := &cli.App{
		Name:        "waitfor",
		Usage:       "Tests and waits on the availability of a remote resource",
		Description: "Tests and waits on the availability of a remote resource before executing a command with exponential backoff",
		Version:     version,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "resource",
				Aliases:  []string{"r"},
				Usage:    "-r http://localhost:8080",
				EnvVars:  []string{"WAITFOR_RESOURCE"},
				Required: true,
			},
			&cli.Uint64Flag{
				Name:    "attempts",
				Aliases: []string{"a"},
				Usage:   "amount of attempts",
				EnvVars: []string{"WAITFOR_ATTEMPTS"},
				Value:   5,
			},
			&cli.Uint64Flag{
				Name:    "interval",
				Usage:   "interval between attempts (sec)",
				EnvVars: []string{"WAITFOR_INTERVAL"},
				Value:   5,
			},
			&cli.Uint64Flag{
				Name:    "max-interval",
				Usage:   "maximum interval between attempts (sec)",
				EnvVars: []string{"WAITFOR_MAX_INTERVAL"},
				Value:   60,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "enable verbose progress output",
				EnvVars: []string{"WAITFOR_VERBOSE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				return cli.Exit("executable is required", 1)
			}

			runner := waitfor.New(
				fs.Use(),
				proc.Use(),
				http.Use(),
				postgres.Use(),
				mongodb.Use(),
				mysql.Use(),
			)

			program := waitfor.Program{
				Executable: "",
				Args:       nil,
				Resources:  ctx.StringSlice("resource"),
			}

			args := ctx.Args().Slice()

			program.Executable = args[0]

			if len(args) > 1 {
				program.Args = args[1:]
			}

			// Build waitfor options
			options := []waitfor.Option{
				waitfor.WithAttempts(ctx.Uint64("attempts")),
				waitfor.WithInterval(ctx.Uint64("interval")),
				waitfor.WithMaxInterval(ctx.Uint64("max-interval")),
			}

			// If verbose mode is enabled, show progress
			if ctx.Bool("verbose") {
				fmt.Printf("waitfor: checking %d resource(s) with %d max attempts\n", len(program.Resources), ctx.Uint64("attempts"))
				for i, resource := range program.Resources {
					fmt.Printf("waitfor: [%d/%d] checking %s\n", i+1, len(program.Resources), resource)
				}
				fmt.Printf("waitfor: retry interval: %ds, max interval: %ds\n", ctx.Uint64("interval"), ctx.Uint64("max-interval"))
				fmt.Println("waitfor: starting resource availability tests...")
				
				start := time.Now()
				
				// Test resources first to show progress
				err := runner.Test(ctx.Context, program.Resources, options...)
				
				duration := time.Since(start)
				
				if err != nil {
					fmt.Printf("waitfor: resource tests failed after %v: %v\n", duration.Round(time.Millisecond), err)
					return err
				}
				
				fmt.Printf("waitfor: all resources available after %v\n", duration.Round(time.Millisecond))
				fmt.Printf("waitfor: executing command: %s\n", program.Executable)
			}

			out, err := runner.Run(
				ctx.Context,
				program,
				options...,
			)

			if out != nil {
				fmt.Println(string(out))
			}

			return err
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
