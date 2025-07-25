// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountMagicConnectorTelemetrySnapshotService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicConnectorTelemetrySnapshotService] method instead.
type AccountMagicConnectorTelemetrySnapshotService struct {
	Options []option.RequestOption
}

// NewAccountMagicConnectorTelemetrySnapshotService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountMagicConnectorTelemetrySnapshotService(opts ...option.RequestOption) (r *AccountMagicConnectorTelemetrySnapshotService) {
	r = &AccountMagicConnectorTelemetrySnapshotService{}
	r.Options = opts
	return
}

// List Snapshots
func (r *AccountMagicConnectorTelemetrySnapshotService) List(ctx context.Context, accountID string, connectorID string, query AccountMagicConnectorTelemetrySnapshotListParams, opts ...option.RequestOption) (res *AccountMagicConnectorTelemetrySnapshotListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/snapshots", accountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Snapshot
func (r *AccountMagicConnectorTelemetrySnapshotService) Get(ctx context.Context, accountID string, connectorID string, snapshotT float64, opts ...option.RequestOption) (res *AccountMagicConnectorTelemetrySnapshotGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/snapshots/%v", accountID, connectorID, snapshotT)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountMagicConnectorTelemetrySnapshotListResponse struct {
	Result   AccountMagicConnectorTelemetrySnapshotListResponseResult `json:"result,required"`
	Success  bool                                                     `json:"success,required"`
	Errors   []MconnCodedMessage                                      `json:"errors"`
	Messages []MconnCodedMessage                                      `json:"messages"`
	JSON     accountMagicConnectorTelemetrySnapshotListResponseJSON   `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotListResponseJSON contains the JSON
// metadata for the struct [AccountMagicConnectorTelemetrySnapshotListResponse]
type accountMagicConnectorTelemetrySnapshotListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetrySnapshotListResponseResult struct {
	Count  float64                                                        `json:"count,required"`
	Items  []AccountMagicConnectorTelemetrySnapshotListResponseResultItem `json:"items,required"`
	Cursor string                                                         `json:"cursor"`
	JSON   accountMagicConnectorTelemetrySnapshotListResponseResultJSON   `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotListResponseResultJSON contains the JSON
// metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotListResponseResult]
type accountMagicConnectorTelemetrySnapshotListResponseResultJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetrySnapshotListResponseResultItem struct {
	// Time the Snapshot was collected (seconds since the Unix epoch)
	A float64 `json:"a,required"`
	// Time the Snapshot was recorded (seconds since the Unix epoch)
	T    float64                                                          `json:"t,required"`
	JSON accountMagicConnectorTelemetrySnapshotListResponseResultItemJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotListResponseResultItemJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotListResponseResultItem]
type accountMagicConnectorTelemetrySnapshotListResponseResultItemJSON struct {
	A           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotListResponseResultItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotListResponseResultItemJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetrySnapshotGetResponse struct {
	// Snapshot
	Result   AccountMagicConnectorTelemetrySnapshotGetResponseResult `json:"result,required"`
	Success  bool                                                    `json:"success,required"`
	Errors   []MconnCodedMessage                                     `json:"errors"`
	Messages []MconnCodedMessage                                     `json:"messages"`
	JSON     accountMagicConnectorTelemetrySnapshotGetResponseJSON   `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseJSON contains the JSON metadata
// for the struct [AccountMagicConnectorTelemetrySnapshotGetResponse]
type accountMagicConnectorTelemetrySnapshotGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseJSON) RawJSON() string {
	return r.raw
}

// Snapshot
type AccountMagicConnectorTelemetrySnapshotGetResponseResult struct {
	// Count of failures to reclaim space
	CountReclaimFailures float64 `json:"count_reclaim_failures,required"`
	// Count of reclaimed paths
	CountReclaimedPaths float64 `json:"count_reclaimed_paths,required"`
	// Count of failed snapshot recordings
	CountRecordFailed float64 `json:"count_record_failed,required"`
	// Count of failed snapshot transmissions
	CountTransmitFailures float64 `json:"count_transmit_failures,required"`
	// Time the Snapshot was recorded (seconds since the Unix epoch)
	T float64 `json:"t,required"`
	// Version
	V string `json:"v,required"`
	// Count of processors/cores
	CPUCount float64 `json:"cpu_count"`
	// Percentage of time over a 10 second window that tasks were stalled
	CPUPressure10s float64 `json:"cpu_pressure_10s"`
	// Percentage of time over a 5 minute window that tasks were stalled
	CPUPressure300s float64 `json:"cpu_pressure_300s"`
	// Percentage of time over a 1 minute window that tasks were stalled
	CPUPressure60s float64 `json:"cpu_pressure_60s"`
	// Total stall time (microseconds)
	CPUPressureTotalUs float64 `json:"cpu_pressure_total_us"`
	// Time spent running a virtual CPU or guest OS (milliseconds)
	CPUTimeGuestMs float64 `json:"cpu_time_guest_ms"`
	// Time spent running a niced guest (milliseconds)
	CPUTimeGuestNiceMs float64 `json:"cpu_time_guest_nice_ms"`
	// Time spent in idle state (milliseconds)
	CPUTimeIdleMs float64 `json:"cpu_time_idle_ms"`
	// Time spent wait for I/O to complete (milliseconds)
	CPUTimeIowaitMs float64 `json:"cpu_time_iowait_ms"`
	// Time spent servicing interrupts (milliseconds)
	CPUTimeIrqMs float64 `json:"cpu_time_irq_ms"`
	// Time spent in low-priority user mode (milliseconds)
	CPUTimeNiceMs float64 `json:"cpu_time_nice_ms"`
	// Time spent servicing softirqs (milliseconds)
	CPUTimeSoftirqMs float64 `json:"cpu_time_softirq_ms"`
	// Time stolen (milliseconds)
	CPUTimeStealMs float64 `json:"cpu_time_steal_ms"`
	// Time spent in system mode (milliseconds)
	CPUTimeSystemMs float64 `json:"cpu_time_system_ms"`
	// Time spent in user mode (milliseconds)
	CPUTimeUserMs float64                                                            `json:"cpu_time_user_ms"`
	DhcpLeases    []AccountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLease `json:"dhcp_leases"`
	Disks         []AccountMagicConnectorTelemetrySnapshotGetResponseResultDisk      `json:"disks"`
	// Name of high availability state
	HaState string `json:"ha_state"`
	// Numeric value associated with high availability state (0 = disabled, 1 = active,
	// 2 = standby, 3 = stopped, 4 = fault)
	HaValue    float64                                                            `json:"ha_value"`
	Interfaces []AccountMagicConnectorTelemetrySnapshotGetResponseResultInterface `json:"interfaces"`
	// Percentage of time over a 10 second window that all tasks were stalled
	IoPressureFull10s float64 `json:"io_pressure_full_10s"`
	// Percentage of time over a 5 minute window that all tasks were stalled
	IoPressureFull300s float64 `json:"io_pressure_full_300s"`
	// Percentage of time over a 1 minute window that all tasks were stalled
	IoPressureFull60s float64 `json:"io_pressure_full_60s"`
	// Total stall time (microseconds)
	IoPressureFullTotalUs float64 `json:"io_pressure_full_total_us"`
	// Percentage of time over a 10 second window that some tasks were stalled
	IoPressureSome10s float64 `json:"io_pressure_some_10s"`
	// Percentage of time over a 3 minute window that some tasks were stalled
	IoPressureSome300s float64 `json:"io_pressure_some_300s"`
	// Percentage of time over a 1 minute window that some tasks were stalled
	IoPressureSome60s float64 `json:"io_pressure_some_60s"`
	// Total stall time (microseconds)
	IoPressureSomeTotalUs float64 `json:"io_pressure_some_total_us"`
	// Boot time (seconds since Unix epoch)
	KernelBtime float64 `json:"kernel_btime"`
	// Number of context switches that the system underwent
	KernelCtxt float64 `json:"kernel_ctxt"`
	// Number of forks since boot
	KernelProcesses float64 `json:"kernel_processes"`
	// Number of processes blocked waiting for I/O
	KernelProcessesBlocked float64 `json:"kernel_processes_blocked"`
	// Number of processes in runnable state
	KernelProcessesRunning float64 `json:"kernel_processes_running"`
	// The fifteen-minute load average
	LoadAverage15m float64 `json:"load_average_15m"`
	// The one-minute load average
	LoadAverage1m float64 `json:"load_average_1m"`
	// The five-minute load average
	LoadAverage5m float64 `json:"load_average_5m"`
	// Number of currently runnable kernel scheduling entities
	LoadAverageCur float64 `json:"load_average_cur"`
	// Number of kernel scheduling entities that currently exist on the system
	LoadAverageMax float64 `json:"load_average_max"`
	// Memory that has been used more recently
	MemoryActiveBytes float64 `json:"memory_active_bytes"`
	// Non-file backed huge pages mapped into user-space page tables
	MemoryAnonHugepagesBytes float64 `json:"memory_anon_hugepages_bytes"`
	// Non-file backed pages mapped into user-space page tables
	MemoryAnonPagesBytes float64 `json:"memory_anon_pages_bytes"`
	// Estimate of how much memory is available for starting new applications
	MemoryAvailableBytes float64 `json:"memory_available_bytes"`
	// Memory used for block device bounce buffers
	MemoryBounceBytes float64 `json:"memory_bounce_bytes"`
	// Relatively temporary storage for raw disk blocks
	MemoryBuffersBytes float64 `json:"memory_buffers_bytes"`
	// In-memory cache for files read from the disk
	MemoryCachedBytes float64 `json:"memory_cached_bytes"`
	// Free CMA (Contiguous Memory Allocator) pages
	MemoryCmaFreeBytes float64 `json:"memory_cma_free_bytes"`
	// Total CMA (Contiguous Memory Allocator) pages
	MemoryCmaTotalBytes float64 `json:"memory_cma_total_bytes"`
	// Total amount of memory currently available to be allocated on the system
	MemoryCommitLimitBytes float64 `json:"memory_commit_limit_bytes"`
	// Amount of memory presently allocated on the system
	MemoryCommittedAsBytes float64 `json:"memory_committed_as_bytes"`
	// Memory which is waiting to get written back to the disk
	MemoryDirtyBytes float64 `json:"memory_dirty_bytes"`
	// The sum of LowFree and HighFree
	MemoryFreeBytes float64 `json:"memory_free_bytes"`
	// Amount of free highmem
	MemoryHighFreeBytes float64 `json:"memory_high_free_bytes"`
	// Total amount of highmem
	MemoryHighTotalBytes float64 `json:"memory_high_total_bytes"`
	// The number of huge pages in the pool that are not yet allocated
	MemoryHugepagesFree float64 `json:"memory_hugepages_free"`
	// Number of huge pages for which a commitment has been made, but no allocation has
	// yet been made
	MemoryHugepagesRsvd float64 `json:"memory_hugepages_rsvd"`
	// Number of huge pages in the pool above the threshold
	MemoryHugepagesSurp float64 `json:"memory_hugepages_surp"`
	// The size of the pool of huge pages
	MemoryHugepagesTotal float64 `json:"memory_hugepages_total"`
	// The size of huge pages
	MemoryHugepagesizeBytes float64 `json:"memory_hugepagesize_bytes"`
	// Memory which has been less recently used
	MemoryInactiveBytes float64 `json:"memory_inactive_bytes"`
	// Kernel allocations that the kernel will attempt to reclaim under memory pressure
	MemoryKReclaimableBytes float64 `json:"memory_k_reclaimable_bytes"`
	// Amount of memory allocated to kernel stacks
	MemoryKernelStackBytes float64 `json:"memory_kernel_stack_bytes"`
	// Amount of free lowmem
	MemoryLowFreeBytes float64 `json:"memory_low_free_bytes"`
	// Total amount of lowmem
	MemoryLowTotalBytes float64 `json:"memory_low_total_bytes"`
	// Files which have been mapped into memory
	MemoryMappedBytes float64 `json:"memory_mapped_bytes"`
	// Amount of memory dedicated to the lowest level of page tables
	MemoryPageTablesBytes float64 `json:"memory_page_tables_bytes"`
	// Memory allocated to the per-cpu alloctor used to back per-cpu allocations
	MemoryPerCPUBytes float64 `json:"memory_per_cpu_bytes"`
	// Percentage of time over a 10 second window that all tasks were stalled
	MemoryPressureFull10s float64 `json:"memory_pressure_full_10s"`
	// Percentage of time over a 5 minute window that all tasks were stalled
	MemoryPressureFull300s float64 `json:"memory_pressure_full_300s"`
	// Percentage of time over a 1 minute window that all tasks were stalled
	MemoryPressureFull60s float64 `json:"memory_pressure_full_60s"`
	// Total stall time (microseconds)
	MemoryPressureFullTotalUs float64 `json:"memory_pressure_full_total_us"`
	// Percentage of time over a 10 second window that some tasks were stalled
	MemoryPressureSome10s float64 `json:"memory_pressure_some_10s"`
	// Percentage of time over a 5 minute window that some tasks were stalled
	MemoryPressureSome300s float64 `json:"memory_pressure_some_300s"`
	// Percentage of time over a 1 minute window that some tasks were stalled
	MemoryPressureSome60s float64 `json:"memory_pressure_some_60s"`
	// Total stall time (microseconds)
	MemoryPressureSomeTotalUs float64 `json:"memory_pressure_some_total_us"`
	// Part of slab that can be reclaimed on memory pressure
	MemorySReclaimableBytes float64 `json:"memory_s_reclaimable_bytes"`
	// Part of slab that cannot be reclaimed on memory pressure
	MemorySUnreclaimBytes float64 `json:"memory_s_unreclaim_bytes"`
	// Amount of memory dedicated to the lowest level of page tables
	MemorySecondaryPageTablesBytes float64 `json:"memory_secondary_page_tables_bytes"`
	// Amount of memory consumed by tmpfs
	MemoryShmemBytes float64 `json:"memory_shmem_bytes"`
	// Memory used by shmem and tmpfs, allocated with huge pages
	MemoryShmemHugepagesBytes float64 `json:"memory_shmem_hugepages_bytes"`
	// Shared memory mapped into user space with huge pages
	MemoryShmemPmdMappedBytes float64 `json:"memory_shmem_pmd_mapped_bytes"`
	// In-kernel data structures cache
	MemorySlabBytes float64 `json:"memory_slab_bytes"`
	// Memory swapped out and back in while still in swap file
	MemorySwapCachedBytes float64 `json:"memory_swap_cached_bytes"`
	// Amount of swap space that is currently unused
	MemorySwapFreeBytes float64 `json:"memory_swap_free_bytes"`
	// Total amount of swap space available
	MemorySwapTotalBytes float64 `json:"memory_swap_total_bytes"`
	// Total usable RAM
	MemoryTotalBytes float64 `json:"memory_total_bytes"`
	// Largest contiguous block of vmalloc area which is free
	MemoryVmallocChunkBytes float64 `json:"memory_vmalloc_chunk_bytes"`
	// Total size of vmalloc memory area
	MemoryVmallocTotalBytes float64 `json:"memory_vmalloc_total_bytes"`
	// Amount of vmalloc area which is used
	MemoryVmallocUsedBytes float64 `json:"memory_vmalloc_used_bytes"`
	// Memory which is actively being written back to the disk
	MemoryWritebackBytes float64 `json:"memory_writeback_bytes"`
	// Memory used by FUSE for temporary writeback buffers
	MemoryWritebackTmpBytes float64 `json:"memory_writeback_tmp_bytes"`
	// Memory consumed by the zswap backend, compressed
	MemoryZSwapBytes float64 `json:"memory_z_swap_bytes"`
	// Amount of anonymous memory stored in zswap, uncompressed
	MemoryZSwappedBytes float64                                                         `json:"memory_z_swapped_bytes"`
	Mounts              []AccountMagicConnectorTelemetrySnapshotGetResponseResultMount  `json:"mounts"`
	Netdevs             []AccountMagicConnectorTelemetrySnapshotGetResponseResultNetdev `json:"netdevs"`
	// Number of ICMP Address Mask Reply messages received
	SnmpIcmpInAddrMaskReps float64 `json:"snmp_icmp_in_addr_mask_reps"`
	// Number of ICMP Address Mask Request messages received
	SnmpIcmpInAddrMasks float64 `json:"snmp_icmp_in_addr_masks"`
	// Number of ICMP messages received with bad checksums
	SnmpIcmpInCsumErrors float64 `json:"snmp_icmp_in_csum_errors"`
	// Number of ICMP Destination Unreachable messages received
	SnmpIcmpInDestUnreachs float64 `json:"snmp_icmp_in_dest_unreachs"`
	// Number of ICMP Echo Reply messages received
	SnmpIcmpInEchoReps float64 `json:"snmp_icmp_in_echo_reps"`
	// Number of ICMP Echo (request) messages received
	SnmpIcmpInEchos float64 `json:"snmp_icmp_in_echos"`
	// Number of ICMP messages received with ICMP-specific errors
	SnmpIcmpInErrors float64 `json:"snmp_icmp_in_errors"`
	// Number of ICMP messages received
	SnmpIcmpInMsgs float64 `json:"snmp_icmp_in_msgs"`
	// Number of ICMP Parameter Problem messages received
	SnmpIcmpInParmProbs float64 `json:"snmp_icmp_in_parm_probs"`
	// Number of ICMP Redirect messages received
	SnmpIcmpInRedirects float64 `json:"snmp_icmp_in_redirects"`
	// Number of ICMP Source Quench messages received
	SnmpIcmpInSrcQuenchs float64 `json:"snmp_icmp_in_src_quenchs"`
	// Number of ICMP Time Exceeded messages received
	SnmpIcmpInTimeExcds float64 `json:"snmp_icmp_in_time_excds"`
	// Number of ICMP Address Mask Request messages received
	SnmpIcmpInTimestampReps float64 `json:"snmp_icmp_in_timestamp_reps"`
	// Number of ICMP Timestamp (request) messages received
	SnmpIcmpInTimestamps float64 `json:"snmp_icmp_in_timestamps"`
	// Number of ICMP Address Mask Reply messages sent
	SnmpIcmpOutAddrMaskReps float64 `json:"snmp_icmp_out_addr_mask_reps"`
	// Number of ICMP Address Mask Request messages sent
	SnmpIcmpOutAddrMasks float64 `json:"snmp_icmp_out_addr_masks"`
	// Number of ICMP Destination Unreachable messages sent
	SnmpIcmpOutDestUnreachs float64 `json:"snmp_icmp_out_dest_unreachs"`
	// Number of ICMP Echo Reply messages sent
	SnmpIcmpOutEchoReps float64 `json:"snmp_icmp_out_echo_reps"`
	// Number of ICMP Echo (request) messages sent
	SnmpIcmpOutEchos float64 `json:"snmp_icmp_out_echos"`
	// Number of ICMP messages which this entity did not send due to ICMP-specific
	// errors
	SnmpIcmpOutErrors float64 `json:"snmp_icmp_out_errors"`
	// Number of ICMP messages attempted to send
	SnmpIcmpOutMsgs float64 `json:"snmp_icmp_out_msgs"`
	// Number of ICMP Parameter Problem messages sent
	SnmpIcmpOutParmProbs float64 `json:"snmp_icmp_out_parm_probs"`
	// Number of ICMP Redirect messages sent
	SnmpIcmpOutRedirects float64 `json:"snmp_icmp_out_redirects"`
	// Number of ICMP Source Quench messages sent
	SnmpIcmpOutSrcQuenchs float64 `json:"snmp_icmp_out_src_quenchs"`
	// Number of ICMP Time Exceeded messages sent
	SnmpIcmpOutTimeExcds float64 `json:"snmp_icmp_out_time_excds"`
	// Number of ICMP Timestamp Reply messages sent
	SnmpIcmpOutTimestampReps float64 `json:"snmp_icmp_out_timestamp_reps"`
	// Number of ICMP Timestamp (request) messages sent
	SnmpIcmpOutTimestamps float64 `json:"snmp_icmp_out_timestamps"`
	// Default value of the Time-To-Live field of the IP header
	SnmpIPDefaultTtl float64 `json:"snmp_ip_default_ttl"`
	// Number of datagrams forwarded to their final destination
	SnmpIPForwDatagrams float64 `json:"snmp_ip_forw_datagrams"`
	// Set when acting as an IP gateway
	SnmpIPForwardingEnabled bool `json:"snmp_ip_forwarding_enabled"`
	// Number of datagrams generated by fragmentation
	SnmpIPFragCreates float64 `json:"snmp_ip_frag_creates"`
	// Number of datagrams discarded because fragmentation failed
	SnmpIPFragFails float64 `json:"snmp_ip_frag_fails"`
	// Number of datagrams successfully fragmented
	SnmpIPFragOks float64 `json:"snmp_ip_frag_oks"`
	// Number of input datagrams discarded due to errors in the IP address
	SnmpIPInAddrErrors float64 `json:"snmp_ip_in_addr_errors"`
	// Number of input datagrams successfully delivered to IP user-protocols
	SnmpIPInDelivers float64 `json:"snmp_ip_in_delivers"`
	// Number of input datagrams otherwise discarded
	SnmpIPInDiscards float64 `json:"snmp_ip_in_discards"`
	// Number of input datagrams discarded due to errors in the IP header
	SnmpIPInHdrErrors float64 `json:"snmp_ip_in_hdr_errors"`
	// Number of input datagrams received from interfaces
	SnmpIPInReceives float64 `json:"snmp_ip_in_receives"`
	// Number of input datagrams discarded due unknown or unsupported protocol
	SnmpIPInUnknownProtos float64 `json:"snmp_ip_in_unknown_protos"`
	// Number of output datagrams otherwise discarded
	SnmpIPOutDiscards float64 `json:"snmp_ip_out_discards"`
	// Number of output datagrams discarded because no route matched
	SnmpIPOutNoRoutes float64 `json:"snmp_ip_out_no_routes"`
	// Number of datagrams supplied for transmission
	SnmpIPOutRequests float64 `json:"snmp_ip_out_requests"`
	// Number of failures detected by the reassembly algorithm
	SnmpIPReasmFails float64 `json:"snmp_ip_reasm_fails"`
	// Number of datagrams successfully reassembled
	SnmpIPReasmOks float64 `json:"snmp_ip_reasm_oks"`
	// Number of fragments received which needed to be reassembled
	SnmpIPReasmReqds float64 `json:"snmp_ip_reasm_reqds"`
	// Number of seconds fragments are held while awaiting reassembly
	SnmpIPReasmTimeout float64 `json:"snmp_ip_reasm_timeout"`
	// Number of times TCP transitions to SYN-SENT from CLOSED
	SnmpTcpActiveOpens float64 `json:"snmp_tcp_active_opens"`
	// Number of times TCP transitions to CLOSED from SYN-SENT or SYN-RCVD, plus
	// transitions to LISTEN from SYN-RCVD
	SnmpTcpAttemptFails float64 `json:"snmp_tcp_attempt_fails"`
	// Number of TCP connections in ESTABLISHED or CLOSE-WAIT
	SnmpTcpCurrEstab float64 `json:"snmp_tcp_curr_estab"`
	// Number of times TCP transitions to CLOSED from ESTABLISHED or CLOSE-WAIT
	SnmpTcpEstabResets float64 `json:"snmp_tcp_estab_resets"`
	// Number of TCP segments received with checksum errors
	SnmpTcpInCsumErrors float64 `json:"snmp_tcp_in_csum_errors"`
	// Number of TCP segments received in error
	SnmpTcpInErrs float64 `json:"snmp_tcp_in_errs"`
	// Number of TCP segments received
	SnmpTcpInSegs float64 `json:"snmp_tcp_in_segs"`
	// Limit on the total number of TCP connections
	SnmpTcpMaxConn float64 `json:"snmp_tcp_max_conn"`
	// Number of TCP segments sent with RST flag
	SnmpTcpOutRsts float64 `json:"snmp_tcp_out_rsts"`
	// Number of TCP segments sent
	SnmpTcpOutSegs float64 `json:"snmp_tcp_out_segs"`
	// Number of times TCP transitions to SYN-RCVD from LISTEN
	SnmpTcpPassiveOpens float64 `json:"snmp_tcp_passive_opens"`
	// Number of TCP segments retransmitted
	SnmpTcpRetransSegs float64 `json:"snmp_tcp_retrans_segs"`
	// Maximum value permitted by a TCP implementation for the retransmission timeout
	// (milliseconds)
	SnmpTcpRtoMax float64 `json:"snmp_tcp_rto_max"`
	// Minimum value permitted by a TCP implementation for the retransmission timeout
	// (milliseconds)
	SnmpTcpRtoMin float64 `json:"snmp_tcp_rto_min"`
	// Number of UDP datagrams delivered to UDP applications
	SnmpUdpInDatagrams float64 `json:"snmp_udp_in_datagrams"`
	// Number of UDP datagrams failed to be delivered for reasons other than lack of
	// application at the destination port
	SnmpUdpInErrors float64 `json:"snmp_udp_in_errors"`
	// Number of UDP datagrams received for which there was not application at the
	// destination port
	SnmpUdpNoPorts float64 `json:"snmp_udp_no_ports"`
	// Number of UDP datagrams sent
	SnmpUdpOutDatagrams float64 `json:"snmp_udp_out_datagrams"`
	// Boottime of the system (seconds since the Unix epoch)
	SystemBootTimeS float64                                                          `json:"system_boot_time_s"`
	Thermals        []AccountMagicConnectorTelemetrySnapshotGetResponseResultThermal `json:"thermals"`
	Tunnels         []AccountMagicConnectorTelemetrySnapshotGetResponseResultTunnel  `json:"tunnels"`
	// Sum of how much time each core has spent idle
	UptimeIdleMs float64 `json:"uptime_idle_ms"`
	// Uptime of the system, including time spent in suspend
	UptimeTotalMs float64                                                     `json:"uptime_total_ms"`
	JSON          accountMagicConnectorTelemetrySnapshotGetResponseResultJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultJSON contains the JSON
// metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResult]
type accountMagicConnectorTelemetrySnapshotGetResponseResultJSON struct {
	CountReclaimFailures           apijson.Field
	CountReclaimedPaths            apijson.Field
	CountRecordFailed              apijson.Field
	CountTransmitFailures          apijson.Field
	T                              apijson.Field
	V                              apijson.Field
	CPUCount                       apijson.Field
	CPUPressure10s                 apijson.Field
	CPUPressure300s                apijson.Field
	CPUPressure60s                 apijson.Field
	CPUPressureTotalUs             apijson.Field
	CPUTimeGuestMs                 apijson.Field
	CPUTimeGuestNiceMs             apijson.Field
	CPUTimeIdleMs                  apijson.Field
	CPUTimeIowaitMs                apijson.Field
	CPUTimeIrqMs                   apijson.Field
	CPUTimeNiceMs                  apijson.Field
	CPUTimeSoftirqMs               apijson.Field
	CPUTimeStealMs                 apijson.Field
	CPUTimeSystemMs                apijson.Field
	CPUTimeUserMs                  apijson.Field
	DhcpLeases                     apijson.Field
	Disks                          apijson.Field
	HaState                        apijson.Field
	HaValue                        apijson.Field
	Interfaces                     apijson.Field
	IoPressureFull10s              apijson.Field
	IoPressureFull300s             apijson.Field
	IoPressureFull60s              apijson.Field
	IoPressureFullTotalUs          apijson.Field
	IoPressureSome10s              apijson.Field
	IoPressureSome300s             apijson.Field
	IoPressureSome60s              apijson.Field
	IoPressureSomeTotalUs          apijson.Field
	KernelBtime                    apijson.Field
	KernelCtxt                     apijson.Field
	KernelProcesses                apijson.Field
	KernelProcessesBlocked         apijson.Field
	KernelProcessesRunning         apijson.Field
	LoadAverage15m                 apijson.Field
	LoadAverage1m                  apijson.Field
	LoadAverage5m                  apijson.Field
	LoadAverageCur                 apijson.Field
	LoadAverageMax                 apijson.Field
	MemoryActiveBytes              apijson.Field
	MemoryAnonHugepagesBytes       apijson.Field
	MemoryAnonPagesBytes           apijson.Field
	MemoryAvailableBytes           apijson.Field
	MemoryBounceBytes              apijson.Field
	MemoryBuffersBytes             apijson.Field
	MemoryCachedBytes              apijson.Field
	MemoryCmaFreeBytes             apijson.Field
	MemoryCmaTotalBytes            apijson.Field
	MemoryCommitLimitBytes         apijson.Field
	MemoryCommittedAsBytes         apijson.Field
	MemoryDirtyBytes               apijson.Field
	MemoryFreeBytes                apijson.Field
	MemoryHighFreeBytes            apijson.Field
	MemoryHighTotalBytes           apijson.Field
	MemoryHugepagesFree            apijson.Field
	MemoryHugepagesRsvd            apijson.Field
	MemoryHugepagesSurp            apijson.Field
	MemoryHugepagesTotal           apijson.Field
	MemoryHugepagesizeBytes        apijson.Field
	MemoryInactiveBytes            apijson.Field
	MemoryKReclaimableBytes        apijson.Field
	MemoryKernelStackBytes         apijson.Field
	MemoryLowFreeBytes             apijson.Field
	MemoryLowTotalBytes            apijson.Field
	MemoryMappedBytes              apijson.Field
	MemoryPageTablesBytes          apijson.Field
	MemoryPerCPUBytes              apijson.Field
	MemoryPressureFull10s          apijson.Field
	MemoryPressureFull300s         apijson.Field
	MemoryPressureFull60s          apijson.Field
	MemoryPressureFullTotalUs      apijson.Field
	MemoryPressureSome10s          apijson.Field
	MemoryPressureSome300s         apijson.Field
	MemoryPressureSome60s          apijson.Field
	MemoryPressureSomeTotalUs      apijson.Field
	MemorySReclaimableBytes        apijson.Field
	MemorySUnreclaimBytes          apijson.Field
	MemorySecondaryPageTablesBytes apijson.Field
	MemoryShmemBytes               apijson.Field
	MemoryShmemHugepagesBytes      apijson.Field
	MemoryShmemPmdMappedBytes      apijson.Field
	MemorySlabBytes                apijson.Field
	MemorySwapCachedBytes          apijson.Field
	MemorySwapFreeBytes            apijson.Field
	MemorySwapTotalBytes           apijson.Field
	MemoryTotalBytes               apijson.Field
	MemoryVmallocChunkBytes        apijson.Field
	MemoryVmallocTotalBytes        apijson.Field
	MemoryVmallocUsedBytes         apijson.Field
	MemoryWritebackBytes           apijson.Field
	MemoryWritebackTmpBytes        apijson.Field
	MemoryZSwapBytes               apijson.Field
	MemoryZSwappedBytes            apijson.Field
	Mounts                         apijson.Field
	Netdevs                        apijson.Field
	SnmpIcmpInAddrMaskReps         apijson.Field
	SnmpIcmpInAddrMasks            apijson.Field
	SnmpIcmpInCsumErrors           apijson.Field
	SnmpIcmpInDestUnreachs         apijson.Field
	SnmpIcmpInEchoReps             apijson.Field
	SnmpIcmpInEchos                apijson.Field
	SnmpIcmpInErrors               apijson.Field
	SnmpIcmpInMsgs                 apijson.Field
	SnmpIcmpInParmProbs            apijson.Field
	SnmpIcmpInRedirects            apijson.Field
	SnmpIcmpInSrcQuenchs           apijson.Field
	SnmpIcmpInTimeExcds            apijson.Field
	SnmpIcmpInTimestampReps        apijson.Field
	SnmpIcmpInTimestamps           apijson.Field
	SnmpIcmpOutAddrMaskReps        apijson.Field
	SnmpIcmpOutAddrMasks           apijson.Field
	SnmpIcmpOutDestUnreachs        apijson.Field
	SnmpIcmpOutEchoReps            apijson.Field
	SnmpIcmpOutEchos               apijson.Field
	SnmpIcmpOutErrors              apijson.Field
	SnmpIcmpOutMsgs                apijson.Field
	SnmpIcmpOutParmProbs           apijson.Field
	SnmpIcmpOutRedirects           apijson.Field
	SnmpIcmpOutSrcQuenchs          apijson.Field
	SnmpIcmpOutTimeExcds           apijson.Field
	SnmpIcmpOutTimestampReps       apijson.Field
	SnmpIcmpOutTimestamps          apijson.Field
	SnmpIPDefaultTtl               apijson.Field
	SnmpIPForwDatagrams            apijson.Field
	SnmpIPForwardingEnabled        apijson.Field
	SnmpIPFragCreates              apijson.Field
	SnmpIPFragFails                apijson.Field
	SnmpIPFragOks                  apijson.Field
	SnmpIPInAddrErrors             apijson.Field
	SnmpIPInDelivers               apijson.Field
	SnmpIPInDiscards               apijson.Field
	SnmpIPInHdrErrors              apijson.Field
	SnmpIPInReceives               apijson.Field
	SnmpIPInUnknownProtos          apijson.Field
	SnmpIPOutDiscards              apijson.Field
	SnmpIPOutNoRoutes              apijson.Field
	SnmpIPOutRequests              apijson.Field
	SnmpIPReasmFails               apijson.Field
	SnmpIPReasmOks                 apijson.Field
	SnmpIPReasmReqds               apijson.Field
	SnmpIPReasmTimeout             apijson.Field
	SnmpTcpActiveOpens             apijson.Field
	SnmpTcpAttemptFails            apijson.Field
	SnmpTcpCurrEstab               apijson.Field
	SnmpTcpEstabResets             apijson.Field
	SnmpTcpInCsumErrors            apijson.Field
	SnmpTcpInErrs                  apijson.Field
	SnmpTcpInSegs                  apijson.Field
	SnmpTcpMaxConn                 apijson.Field
	SnmpTcpOutRsts                 apijson.Field
	SnmpTcpOutSegs                 apijson.Field
	SnmpTcpPassiveOpens            apijson.Field
	SnmpTcpRetransSegs             apijson.Field
	SnmpTcpRtoMax                  apijson.Field
	SnmpTcpRtoMin                  apijson.Field
	SnmpUdpInDatagrams             apijson.Field
	SnmpUdpInErrors                apijson.Field
	SnmpUdpNoPorts                 apijson.Field
	SnmpUdpOutDatagrams            apijson.Field
	SystemBootTimeS                apijson.Field
	Thermals                       apijson.Field
	Tunnels                        apijson.Field
	UptimeIdleMs                   apijson.Field
	UptimeTotalMs                  apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Snapshot DHCP lease
type AccountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLease struct {
	// Client ID of the device the IP Address was leased to
	ClientID string `json:"client_id,required"`
	// Expiry time of the DHCP lease (seconds since the Unix epoch)
	ExpiryTime float64 `json:"expiry_time,required"`
	// Hostname of the device the IP Address was leased to
	Hostname string `json:"hostname,required"`
	// Name of the network interface
	InterfaceName string `json:"interface_name,required"`
	// IP Address that was leased
	IPAddress string `json:"ip_address,required"`
	// MAC Address of the device the IP Address was leased to
	MacAddress string `json:"mac_address,required"`
	// Connector identifier
	ConnectorID string                                                               `json:"connector_id"`
	JSON        accountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLeaseJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLeaseJSON contains
// the JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLease]
type accountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLeaseJSON struct {
	ClientID      apijson.Field
	ExpiryTime    apijson.Field
	Hostname      apijson.Field
	InterfaceName apijson.Field
	IPAddress     apijson.Field
	MacAddress    apijson.Field
	ConnectorID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultDhcpLeaseJSON) RawJSON() string {
	return r.raw
}

// Snapshot Disk
type AccountMagicConnectorTelemetrySnapshotGetResponseResultDisk struct {
	// I/Os currently in progress
	InProgress float64 `json:"in_progress,required"`
	// Device major number
	Major float64 `json:"major,required"`
	// Reads merged
	Merged float64 `json:"merged,required"`
	// Device minor number
	Minor float64 `json:"minor,required"`
	// Device name
	Name string `json:"name,required"`
	// Reads completed successfully
	Reads float64 `json:"reads,required"`
	// Sectors read successfully
	SectorsRead float64 `json:"sectors_read,required"`
	// Sectors written successfully
	SectorsWritten float64 `json:"sectors_written,required"`
	// Time spent doing I/Os (milliseconds)
	TimeInProgressMs float64 `json:"time_in_progress_ms,required"`
	// Time spent reading (milliseconds)
	TimeReadingMs float64 `json:"time_reading_ms,required"`
	// Time spent writing (milliseconds)
	TimeWritingMs float64 `json:"time_writing_ms,required"`
	// Weighted time spent doing I/Os (milliseconds)
	WeightedTimeInProgressMs float64 `json:"weighted_time_in_progress_ms,required"`
	// Writes completed
	Writes float64 `json:"writes,required"`
	// Writes merged
	WritesMerged float64 `json:"writes_merged,required"`
	// Connector identifier
	ConnectorID string `json:"connector_id"`
	// Discards completed successfully
	Discards float64 `json:"discards"`
	// Discards merged
	DiscardsMerged float64 `json:"discards_merged"`
	// Flushes completed successfully
	Flushes float64 `json:"flushes"`
	// Sectors discarded
	SectorsDiscarded float64 `json:"sectors_discarded"`
	// Time spent discarding (milliseconds)
	TimeDiscardingMs float64 `json:"time_discarding_ms"`
	// Time spent flushing (milliseconds)
	TimeFlushingMs float64                                                         `json:"time_flushing_ms"`
	JSON           accountMagicConnectorTelemetrySnapshotGetResponseResultDiskJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultDiskJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultDisk]
type accountMagicConnectorTelemetrySnapshotGetResponseResultDiskJSON struct {
	InProgress               apijson.Field
	Major                    apijson.Field
	Merged                   apijson.Field
	Minor                    apijson.Field
	Name                     apijson.Field
	Reads                    apijson.Field
	SectorsRead              apijson.Field
	SectorsWritten           apijson.Field
	TimeInProgressMs         apijson.Field
	TimeReadingMs            apijson.Field
	TimeWritingMs            apijson.Field
	WeightedTimeInProgressMs apijson.Field
	Writes                   apijson.Field
	WritesMerged             apijson.Field
	ConnectorID              apijson.Field
	Discards                 apijson.Field
	DiscardsMerged           apijson.Field
	Flushes                  apijson.Field
	SectorsDiscarded         apijson.Field
	TimeDiscardingMs         apijson.Field
	TimeFlushingMs           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultDisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultDiskJSON) RawJSON() string {
	return r.raw
}

// Snapshot Interface
type AccountMagicConnectorTelemetrySnapshotGetResponseResultInterface struct {
	// Name of the network interface
	Name string `json:"name,required"`
	// UP/DOWN state of the network interface
	Operstate string `json:"operstate,required"`
	// Connector identifier
	ConnectorID string                                                                       `json:"connector_id"`
	IPAddresses []AccountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddress `json:"ip_addresses"`
	// Speed of the network interface (bits per second)
	Speed float64                                                              `json:"speed"`
	JSON  accountMagicConnectorTelemetrySnapshotGetResponseResultInterfaceJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultInterfaceJSON contains
// the JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultInterface]
type accountMagicConnectorTelemetrySnapshotGetResponseResultInterfaceJSON struct {
	Name        apijson.Field
	Operstate   apijson.Field
	ConnectorID apijson.Field
	IPAddresses apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultInterface) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultInterfaceJSON) RawJSON() string {
	return r.raw
}

// Snapshot Interface Address
type AccountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddress struct {
	// Name of the network interface
	InterfaceName string `json:"interface_name,required"`
	// IP address of the network interface
	IPAddress string `json:"ip_address,required"`
	// Connector identifier
	ConnectorID string                                                                         `json:"connector_id"`
	JSON        accountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddressJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddressJSON
// contains the JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddress]
type accountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddressJSON struct {
	InterfaceName apijson.Field
	IPAddress     apijson.Field
	ConnectorID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultInterfacesIPAddressJSON) RawJSON() string {
	return r.raw
}

// Snapshot Mount
type AccountMagicConnectorTelemetrySnapshotGetResponseResultMount struct {
	// File system on disk (EXT4, NTFS, etc.)
	FileSystem string `json:"file_system,required"`
	// Kind of disk (HDD, SSD, etc.)
	Kind string `json:"kind,required"`
	// Path where disk is mounted
	MountPoint string `json:"mount_point,required"`
	// Name of the disk mount
	Name string `json:"name,required"`
	// Available disk size (bytes)
	AvailableBytes float64 `json:"available_bytes"`
	// Connector identifier
	ConnectorID string `json:"connector_id"`
	// Determines whether the disk is read-only
	IsReadOnly bool `json:"is_read_only"`
	// Determines whether the disk is removable
	IsRemovable bool `json:"is_removable"`
	// Total disk size (bytes)
	TotalBytes float64                                                          `json:"total_bytes"`
	JSON       accountMagicConnectorTelemetrySnapshotGetResponseResultMountJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultMountJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultMount]
type accountMagicConnectorTelemetrySnapshotGetResponseResultMountJSON struct {
	FileSystem     apijson.Field
	Kind           apijson.Field
	MountPoint     apijson.Field
	Name           apijson.Field
	AvailableBytes apijson.Field
	ConnectorID    apijson.Field
	IsReadOnly     apijson.Field
	IsRemovable    apijson.Field
	TotalBytes     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultMount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultMountJSON) RawJSON() string {
	return r.raw
}

// Snapshot Netdev
type AccountMagicConnectorTelemetrySnapshotGetResponseResultNetdev struct {
	// Name of the network device
	Name string `json:"name,required"`
	// Total bytes received
	RecvBytes float64 `json:"recv_bytes,required"`
	// Compressed packets received
	RecvCompressed float64 `json:"recv_compressed,required"`
	// Packets dropped
	RecvDrop float64 `json:"recv_drop,required"`
	// Bad packets received
	RecvErrs float64 `json:"recv_errs,required"`
	// FIFO overruns
	RecvFifo float64 `json:"recv_fifo,required"`
	// Frame alignment errors
	RecvFrame float64 `json:"recv_frame,required"`
	// Multicast packets received
	RecvMulticast float64 `json:"recv_multicast,required"`
	// Total packets received
	RecvPackets float64 `json:"recv_packets,required"`
	// Total bytes transmitted
	SentBytes float64 `json:"sent_bytes,required"`
	// Number of packets not sent due to carrier errors
	SentCarrier float64 `json:"sent_carrier,required"`
	// Number of collisions
	SentColls float64 `json:"sent_colls,required"`
	// Number of compressed packets transmitted
	SentCompressed float64 `json:"sent_compressed,required"`
	// Number of packets dropped during transmission
	SentDrop float64 `json:"sent_drop,required"`
	// Number of transmission errors
	SentErrs float64 `json:"sent_errs,required"`
	// FIFO overruns
	SentFifo float64 `json:"sent_fifo,required"`
	// Total packets transmitted
	SentPackets float64 `json:"sent_packets,required"`
	// Connector identifier
	ConnectorID string                                                            `json:"connector_id"`
	JSON        accountMagicConnectorTelemetrySnapshotGetResponseResultNetdevJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultNetdevJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultNetdev]
type accountMagicConnectorTelemetrySnapshotGetResponseResultNetdevJSON struct {
	Name           apijson.Field
	RecvBytes      apijson.Field
	RecvCompressed apijson.Field
	RecvDrop       apijson.Field
	RecvErrs       apijson.Field
	RecvFifo       apijson.Field
	RecvFrame      apijson.Field
	RecvMulticast  apijson.Field
	RecvPackets    apijson.Field
	SentBytes      apijson.Field
	SentCarrier    apijson.Field
	SentColls      apijson.Field
	SentCompressed apijson.Field
	SentDrop       apijson.Field
	SentErrs       apijson.Field
	SentFifo       apijson.Field
	SentPackets    apijson.Field
	ConnectorID    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultNetdev) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultNetdevJSON) RawJSON() string {
	return r.raw
}

// Snapshot Thermal
type AccountMagicConnectorTelemetrySnapshotGetResponseResultThermal struct {
	// Sensor identifier for the component
	Label string `json:"label,required"`
	// Connector identifier
	ConnectorID string `json:"connector_id"`
	// Critical failure temperature of the component (degrees Celsius)
	CriticalCelcius float64 `json:"critical_celcius"`
	// Current temperature of the component (degrees Celsius)
	CurrentCelcius float64 `json:"current_celcius"`
	// Maximum temperature of the component (degrees Celsius)
	MaxCelcius float64                                                            `json:"max_celcius"`
	JSON       accountMagicConnectorTelemetrySnapshotGetResponseResultThermalJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultThermalJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultThermal]
type accountMagicConnectorTelemetrySnapshotGetResponseResultThermalJSON struct {
	Label           apijson.Field
	ConnectorID     apijson.Field
	CriticalCelcius apijson.Field
	CurrentCelcius  apijson.Field
	MaxCelcius      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultThermal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultThermalJSON) RawJSON() string {
	return r.raw
}

// Snapshot Tunnels
type AccountMagicConnectorTelemetrySnapshotGetResponseResultTunnel struct {
	// Name of tunnel health state (unknown, healthy, degraded, down)
	HealthState string `json:"health_state,required"`
	// Numeric value associated with tunnel state (0 = unknown, 1 = healthy, 2 =
	// degraded, 3 = down)
	HealthValue float64 `json:"health_value,required"`
	// The tunnel interface name (i.e. xfrm1, xfrm3.99, etc.)
	InterfaceName string `json:"interface_name,required"`
	// Tunnel identifier
	TunnelID string `json:"tunnel_id,required"`
	// Connector identifier
	ConnectorID string                                                            `json:"connector_id"`
	JSON        accountMagicConnectorTelemetrySnapshotGetResponseResultTunnelJSON `json:"-"`
}

// accountMagicConnectorTelemetrySnapshotGetResponseResultTunnelJSON contains the
// JSON metadata for the struct
// [AccountMagicConnectorTelemetrySnapshotGetResponseResultTunnel]
type accountMagicConnectorTelemetrySnapshotGetResponseResultTunnelJSON struct {
	HealthState   apijson.Field
	HealthValue   apijson.Field
	InterfaceName apijson.Field
	TunnelID      apijson.Field
	ConnectorID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicConnectorTelemetrySnapshotGetResponseResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicConnectorTelemetrySnapshotGetResponseResultTunnelJSON) RawJSON() string {
	return r.raw
}

type AccountMagicConnectorTelemetrySnapshotListParams struct {
	From   param.Field[float64] `query:"from,required"`
	To     param.Field[float64] `query:"to,required"`
	Cursor param.Field[string]  `query:"cursor"`
	Limit  param.Field[float64] `query:"limit"`
}

// URLQuery serializes [AccountMagicConnectorTelemetrySnapshotListParams]'s query
// parameters as `url.Values`.
func (r AccountMagicConnectorTelemetrySnapshotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
