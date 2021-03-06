package vars

const (
	DefaultMaxConcurrentIndexers = 2
	PageStep                     = 5
	SearchNum                    = 25
	Source                       = "gshark"
	GITHUB                       = "github"
	GITLAB                       = "gitlab"
)

var (
	REPO_PATH    string
	MAX_INDEXERS int

	HTTP_HOST string
	HTTP_PORT int

	MAX_Concurrency_REPOS int

	DEBUG_MODE bool

	PAGE_SIZE = 10
	SCKEY     string
)
