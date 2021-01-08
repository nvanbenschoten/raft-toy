package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/nvanbenschoten/rafttoy/peer"
	"github.com/nvanbenschoten/rafttoy/util"
	"github.com/spf13/pflag"
	"go.etcd.io/etcd/raft/v3"
)

var raftID = pflag.Uint64("id", 1, "ID of this process in the Raft replication group")
var raftPeers = pflag.StringSlice("peers", []string{"localhost:1234"}, "IP address of each peer in the Raft replication group")
var runLoad = pflag.Bool("load", false, "Propose changes to Raft")
var verbose = pflag.Bool("verbose", false, "Verbose logging")
var pprof = pflag.Int("pprof", 0, "Port to serve pprof on. Negative to disable")
var recordMetrics = pflag.Bool("metrics", false, "Record metrics and print before exiting")
var dataDir = pflag.String("data-dir", "", "Directory to store persistent data")
var pipelineImpl = pflag.String("pipeline", "basic", "Raft proposal pipeline implementation")

func init() {
	// Seed rng.
	rand.Seed(time.Now().UTC().UnixNano())

	// Setup logging.
	log.SetOutput(os.Stderr)
	log.SetPrefix("rafttoy | ")
	log.SetFlags(log.LstdFlags)
	util.SetRaftLoggingVerbosity(*verbose)
}

func initFlags() {
	// Add the set of pflags to Go's flag package so that they are usable
	// in tests and benchmarks.
	pflag.CommandLine.VisitAll(func(f *pflag.Flag) {
		switch f.Value.Type() {
		case "bool":
			def, err := strconv.ParseBool(f.DefValue)
			if err != nil {
				panic(err)
			}
			_ = flag.Bool(f.Name, def, f.Usage)
		default:
			_ = flag.String(f.Name, f.DefValue, f.Usage)
		}
	})
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}

func cfgFromFlags() peer.Config {
	cfg := peer.Config{ID: *raftID}
	if cfg.ID == 0 {
		log.Fatalf("invalid ID (%d); must be > 0", cfg.ID)
	}
	if len(*raftPeers) < int(cfg.ID) {
		log.Fatalf("missing own ID (%d) in peers (%v)", cfg.ID, raftPeers)
	}
	cfg.Peers = make([]raft.Peer, len(*raftPeers))
	cfg.PeerAddrs = make(map[uint64]string, len(*raftPeers))
	for i, addr := range *raftPeers {
		pID := uint64(i + 1)
		cfg.Peers[i].ID = pID
		cfg.PeerAddrs[pID] = addr
	}
	cfg.SelfAddr = cfg.PeerAddrs[cfg.ID]
	return cfg
}
