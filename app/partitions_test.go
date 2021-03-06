package app_test

import (
	"github.com/Tapjoy/dynamiq/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Partition", func() {

	var (
		partitions        *app.Partitions
		err               error
		partitionTopID    int
		partitionBottomID int
		partition         *app.Partition
	)

	BeforeEach(func() {
		// Load up a list of partitions
		partitions = app.InitPartitions(cfg, testQueueName)
	})

	Context("InitPartitions", func() {
		It("should return a number of partitions equal to the configured amount", func() {
			minParts, _ := cfg.GetMinPartitions(testQueueName)
			Expect(partitions.PartitionCount()).To(Equal(minParts))
		})
	})

	Context("GetPartition", func() {
		BeforeEach(func() {
			// Get the partition ids, and any errors
			partitionBottomID, partitionTopID, partition, err = partitions.GetPartition(cfg, testQueueName, memberList)
		})

		It("should get a partitionTopId", func() {
			Expect(partitionTopID).ToNot(BeNil())
		})

		It("should get a partitionBottomId", func() {
			Expect(partitionBottomID).ToNot(BeNil())
		})

		It("should not get an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
