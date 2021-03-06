package serve

import (
	"context"
	"strconv"
	"time"

	"github.com/matthewmueller/joy/api"
	"github.com/matthewmueller/joy/internal/mains"
	"github.com/matthewmueller/joy/internal/stats"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("serve", "start a development server with livereload")
	pkgs := cmd.Arg("packages", "packages to bundle").Required().Strings()
	port := cmd.Flag("port", "port to serve from").Short('p').Default("8080").String()

	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		start := time.Now()

		// stats
		defer stats.TrackError("serve", time.Now(), &err)

		packages, err := mains.Find(*pkgs)
		if err != nil {
			return err
		}

		port, e := strconv.Atoi(*port)
		if e != nil {
			return errors.Wrap(e, "invalid port")
		}

		if err := api.Serve(ctx, &api.ServeSettings{
			Packages: packages,
			Port:     port,
		}); err != nil {
			return errors.Wrapf(err, "error serving")
		}

		// stats
		stats.Track("serve", map[string]interface{}{
			"duration": time.Since(start).Round(time.Millisecond).String(),
			"packages": len(packages),
		})

		return nil
	})
}
