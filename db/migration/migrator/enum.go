package migration

type Migration int

const (
	MigrationUp Migration = iota
	MigrationDown
)

var keys = []string{"up", "down"}

func (m Migration) String() string {
	return keys[m]
}
