package integration_test

import (
	dbHelper "code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport/ports"
	"code.cloudfoundry.org/lager/lagertest"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"policy-server/config"
	"policy-server/db"
	"policy-server/integration/helpers"
	"policy-server/store/migrations"
	"test-helpers"
	"time"
	"math"
)

const TimeoutShort = 20 * time.Second

var _ = Describe("Migrate DB Binary", func() {
	var (
		dbConf dbHelper.Config
		conf   config.Config
	)

	BeforeEach(func() {
		dbConf = testsupport.GetDBConfig()
		dbConf.DatabaseName = fmt.Sprintf("migrate_test_node_%d", ports.PickAPort())

		conf, _ = helpers.DefaultTestConfig(dbConf, "unused", "fixtures")
		conf.Database = dbConf

		testhelpers.CreateDatabase(dbConf)
	})

	AfterEach(func() {
		testhelpers.RemoveDatabase(dbConf)
	})

	It("runs the migrations and seeds the groups table", func() {
		session := helpers.RunMigrations(migrateDbPath, conf)
		Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(0))

		conn := db.NewConnectionPool(
			dbConf,
			1,
			1,
			"test-db",
			"test-job-prefix",
			lagertest.NewTestLogger("test"),
		)

		m := migrations.MigrationsToPerform
		lastMigration := m[len(m)-1]

		var highestID string
		conn.QueryRow("SELECT MAX(ID) FROM gorp_migrations LIMIT 1").Scan(&highestID)
		Expect(highestID).To(Equal(lastMigration.Id))

		var groupCount int
		conn.QueryRow("SELECT COUNT(*) FROM groups").Scan(&groupCount)
		Expect(groupCount).To(Equal(int(math.Exp2(float64(conf.TagLength*8)))-1))
	})

	Context("when the tag length is invalid", func() {
		It("should exit non zero", func() {
			conf.TagLength = 5
			session := helpers.RunMigrations(migrateDbPath, conf)
			Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(1))
		})
	})
})