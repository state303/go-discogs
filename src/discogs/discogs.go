package discogs

import (
	"context"
	"fmt"
	"github.com/knadh/koanf"
	"github.com/state303/go-discogs/src/batch"
	"github.com/state303/go-discogs/src/data"
	"github.com/state303/go-discogs/src/database"
	"time"
)

// Discogs will be removed after refactor.
// Currently serves as an entrypoint of batch process.
type Discogs struct{}

func (*Discogs) Run(ctx context.Context, config *koanf.Koanf) error {
	begin := time.Now()
	if err := database.Connect(config.String("dsn")); err != nil {
		return err
	}

	if config.Bool("new") {
		fmt.Println("execute DDL update...")
		if err := RunDDL(database.DB); err != nil {
			return err
		}
	}

	dataRepo := data.NewDataRepository(database.DB)

	if config.Bool("update") {
		fmt.Println("begin update...")
		if updated, err := data.UpdateData(ctx, dataRepo); err != nil {
			return err
		} else {
			fmt.Printf("update affected: %+v rows\n", updated)
		}
	}

	typeResourceMap, err := data.FetchFiles(config, dataRepo)
	if err != nil {
		return err
	}

	var (
		b            = batch.New()
		totalUpdates = 0
		chunk        = config.Int("chunk")
		db           = database.DB
		steps        = make([]batch.Step, 0)
	)

	if hasArtist(config) {
		order := batch.NewOrder(ctx, chunk, typeResourceMap["artists"], db)
		steps = append(steps, b.UpdateArtist(order))
	}

	if hasLabel(config) {
		order := batch.NewOrder(ctx, chunk, typeResourceMap["labels"], db)
		steps = append(steps, b.UpdateLabel(order))
	}

	if hasMaster(config) {
		order := batch.NewOrder(ctx, chunk, typeResourceMap["masters"], db)
		steps = append(steps, b.UpdateMaster(order))
	}

	if hasRelease(config) {
		order := batch.NewOrder(ctx, chunk, typeResourceMap["releases"], db)
		steps = append(steps, b.UpdateRelease(order))
	}

	for i := range steps {
		r := steps[i]()
		totalUpdates += r.Count()
		if r.IsErr() {
			err = r.Err()
			break
		}
	}

	printResult(begin, totalUpdates, err)
	return err
}

func printResult(begin time.Time, total int, err error) {
	took := time.Since(begin).Truncate(time.Second).String()
	s := fmt.Sprintf("updated %+v records in %+v.", total, took)
	if err != nil {
		s += fmt.Sprintf(" [error: %+v]", err)
	}
	fmt.Println(s)
}
