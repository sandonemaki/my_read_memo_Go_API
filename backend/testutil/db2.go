package testutil

import (
	_ "embed"

	"github.com/DATA-DOG/go-txdb"
	"github.com/spf13/viper"
)

func init() {
	txdb.Register("txdb", "postgres", "postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable")
	viper.Set("env.name", "test")

	// logger := slog.New(logger.NewLogger(config.Prepare().Logger))
	// dbConn := db.NewPSQL(config.Prepare().Postgres, logger, "yondeco")
	// prepareTestDatabase(dbConn, "default")
	//	boil.DebugMode = true
	// }

	// type TestFunc func(ctx context.Context, sqlDB *sql.DB, ctrl *gomock.Controller) (actual, expected interface{}, err, wantErr error, options cmp.Options)

	// func TxDB(name string, t *testing.T, f TestFunc) {
	// 	dbConn, err := sql.Open("txdb", fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	defer dbConn.Close()

	// 	// テストケースと同じ名前のディレクトリがあれば、それをfixtureとして追加する
	// 	var extraFixturePath []string
	// 	if _, _err := os.Stat(path.Join(GetTestDataPath(), "fixture", t.Name())); _err == nil {
	// 		extraFixturePath = []string{t.Name()}
	// 	}

	// 	prepareTestDatabase(dbConn, extraFixturePath...)
	// 	ctx := context.Background()
	// 	ctrl := gomock.NewController(t)
	// 	actual, expected, err, wantErr, ignores := f(ctx, dbConn, ctrl)

	// 	if !errors.Is(err, wantErr) {
	// 		t.Errorf("test %s, %s = %+v, want %v", name, t.Name(), err, wantErr)
	// 	}
	// 	if err != nil {
	// 		return
	// 	}

	// 	if diff := cmp.Diff(actual, expected, ignores); diff != "" {
	// 		t.Errorf("test: %s, diff %s", name, diff)
	// 	}
	// }

	// PrepareTxDB :
	// func PrepareTxDB(fixturePath ...string) *sql.DB {
	// 	dbConn, err := sql.Open("txdb", fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	prepareTestDatabase(dbConn, fixturePath...)
	// 	return dbConn
	// }

	// func prepareTestDatabase(dbConn *sql.DB, fixturePath ...string) {
	// 	if len(fixturePath) == 0 {
	// 		return
	// 	}
	// 	var err error
	// 	if err = dbConn.Ping(); err != nil {
	// 		panic(err)
	// 	}

	// 	var paths []string
	// 	for _, p := range fixturePath {
	// 		paths = append(paths, path.Join(GetTestDataPath(), "fixture", p))
	// 	}

	// var fixtures *testfixtures.Loader
	// if fixtures, err = testfixtures.New(
	// 	testfixtures.Database(dbConn),
	// 	testfixtures.Dialect("postgres"),
	// 	testfixtures.Paths(paths...),
	// ); err != nil {
	// 	pp.Println(err)
	// 	log.Fatal(err)
	// }

	// if err = fixtures.Load(); err != nil {
	// 	pp.Println(err)
	// 	log.Fatal(err)
	// }
}

// GetTestDataPath :
// func GetTestDataPath() string {
// 	_, b, _, _ := runtime.Caller(0)
// 	testutilPath := filepath.Dir(b)
// 	internalPath := filepath.Dir(testutilPath)
// 	return filepath.Join(internalPath, "testdata")
// }
