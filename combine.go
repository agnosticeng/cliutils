package cliutils

import (
	"github.com/urfave/cli/v2"
)

func CombineBeforeFuncs(funcs ...cli.BeforeFunc) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		for _, f := range funcs {
			if err := f(ctx); err != nil {
				return err
			}
		}

		return nil
	}
}

func CombineAfterFuncs(funcs ...cli.AfterFunc) cli.AfterFunc {
	return func(ctx *cli.Context) error {
		for _, f := range funcs {
			if err := f(ctx); err != nil {
				return err
			}
		}

		return nil
	}
}

func CombineFlags(flagSets [][]cli.Flag) []cli.Flag {
	var res []cli.Flag

	for _, flagSet := range flagSets {
		res = append(res, flagSet...)
	}

	return res
}
