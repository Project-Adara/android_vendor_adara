package generator

import (
	"fmt"

	"android/soong/android"
)

func adaraExpandVariables(ctx android.ModuleContext, in string) string {
	adaraVars := ctx.Config().VendorConfig("adaraVarsPlugin")

	out, err := android.Expand(in, func(name string) (string, error) {
		if adaraVars.IsSet(name) {
			return adaraVars.String(name), nil
		}
		// This variable is not for us, restore what the original
		// variable string will have looked like for an Expand
		// that comes later.
		return fmt.Sprintf("$(%s)", name), nil
	})

	if err != nil {
		ctx.PropertyErrorf("%s: %s", in, err.Error())
		return ""
	}

	return out
}