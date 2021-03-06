package app_test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/Tapjoy/dynamiq/app"
	"github.com/hashicorp/memberlist"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tpjg/goriakpbc"
	"github.com/tpjg/goriakpbc/pb"
)

var cfg *app.Config
var core app.Core
var queues *app.Queues
var duration time.Duration
var memberList *memberlist.Memberlist
var testQueueName = "test_queue"
var RDtMap *riak.RDtMap

func TestPartitions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")

}

var _ = BeforeSuite(func() {
	cfg = &app.Config{}
	// Create the basic Configuration object
	// later tests can change these values as needed
	core = app.Core{
		Name:                  "john",
		Port:                  8000,
		SeedServer:            "steve",
		SeedServers:           []string{"steve"},
		SeedPort:              8001,
		HTTPPort:              8003,
		RiakNodes:             "127.0.0.1",
		BackendConnectionPool: 16,
		SyncConfigInterval:    duration,
	}

	queueMap := make(map[string]*app.Queue)
	configRDtMap := riak.RDtMap{
		Values:   make(map[riak.MapKey]interface{}),
		ToAdd:    make([]*pb.MapUpdate, 1),
		ToRemove: make([]*pb.MapField, 1),
	}

	configRDtMap.Values[riak.MapKey{Key: "max_partitions", Type: pb.MapField_REGISTER}] = &riak.RDtRegister{Value: []byte(app.DefaultSettings[app.MaxPartitions])}
	configRDtMap.Values[riak.MapKey{Key: "min_partitions", Type: pb.MapField_REGISTER}] = &riak.RDtRegister{Value: []byte(app.DefaultSettings[app.MinPartitions])}
	configRDtMap.Values[riak.MapKey{Key: "visibility_timeout", Type: pb.MapField_REGISTER}] = &riak.RDtRegister{Value: []byte(app.DefaultSettings[app.VisibilityTimeout])}

	queue := &app.Queue{
		Name:   testQueueName,
		Config: &configRDtMap,
	}
	queueMap[testQueueName] = queue

	queues = &app.Queues{
		QueueMap: queueMap,
	}

	cfg.Core = core
	cfg.Queues = queues

	// Create a memberlist, aka the list of possible RiaQ processes to communicate with
	memberList, _, _ = app.InitMemberList(core.Name, core.Port, core.SeedServers, core.SeedPort)

	// Disable log output during tests
	logrus.SetOutput(ioutil.Discard)
})

var _ = AfterSuite(func() {

	// Shut this down incase another suite of tests needs the port, or it's own instance
	memberList.Shutdown()
})
