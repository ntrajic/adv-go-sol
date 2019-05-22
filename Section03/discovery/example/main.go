package main

import (
	".."
	consul "github.com/hashicorp/consul/api"
)
/*
c:\SRC\GO\AaronTorres\adv-go-sol\Section03\discovery>consul agent -dev -node=localhost
==> Starting Consul agent...
==> Consul agent running!
           Version: 'v1.5.0'
           Node ID: 'a6a3ef08-2c9c-6767-9b3b-a89a6870c1c3'
         Node name: 'localhost'
        Datacenter: 'dc1' (Segment: '<all>')
            Server: true (Bootstrap: false)
       Client Addr: [127.0.0.1] (HTTP: 8500, HTTPS: -1, gRPC: 8502, DNS: 8600)
      Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
           Encrypt: Gossip: false, TLS-Outgoing: false, TLS-Incoming: false

==> Log data will now stream in as it occurs:

    2019/05/20 22:37:14 [DEBUG] agent: Using random ID "a6a3ef08-2c9c-6767-9b3b-a89a6870c1c3" as node ID
    2019/05/20 22:37:14 [DEBUG] tlsutil: Update with version 1
    2019/05/20 22:37:14 [DEBUG] tlsutil: OutgoingRPCWrapper with version 1
    2019/05/20 22:37:14 [DEBUG] tlsutil: IncomingRPCConfig with version 1
    2019/05/20 22:37:14 [DEBUG] tlsutil: OutgoingRPCWrapper with version 1
    2019/05/20 22:37:14 [INFO] raft: Initial configuration (index=1): [{Suffrage:Voter ID:a6a3ef08-2c9c-6
	05/20 22:37:14 [INFO] agent: Started DNS server 127.0.0.1:8600 (udp)
	05/20 22:37:14 [INFO] raft: Node at 127.0.0.1:8300 [Follower] entering Follower state (Leader: "")
	05/20 22:37:14 [INFO] consul: Adding LAN server localhost (Addr: tcp/127.0.0.1:8300) (DC: dc1)
	05/20 22:37:14 [INFO] consul: Handled member-join event for server "localhost.dc1" in area "wan"
	05/20 22:37:14 [DEBUG] agent/proxy: managed Connect proxy manager started
	05/20 22:37:14 [INFO] agent: Started DNS server 127.0.0.1:8600 (tcp)
	05/20 22:37:14 [INFO] agent: Started HTTP server on 127.0.0.1:8500 (tcp)
	05/20 22:37:14 [INFO] agent: started state syncer
	05/20 22:37:14 [INFO] agent: Started gRPC server on 127.0.0.1:8502 (tcp)
	05/20 22:37:14 [WARN] raft: Heartbeat timeout from "" reached, starting election
	05/20 22:37:14 [INFO] raft: Node at 127.0.0.1:8300 [Candidate] entering Candidate state in term 2
	05/20 22:37:14 [DEBUG] raft: Votes needed: 1
	05/20 22:37:14 [DEBUG] raft: Vote granted from a6a3ef08-2c9c-6767-9b3b-a89a6870c1c3 in term 2. Tally: 1
	05/20 22:37:14 [INFO] raft: Election won. Tally: 1
	05/20 22:37:14 [INFO] raft: Node at 127.0.0.1:8300 [Leader] entering Leader state

	It is necessary to download following libs to %GOPATH% which is c:\go-work
	C:\Users\nik>go get github.com/hashicorp/go-cleanhttp

	C:\Users\nik>
	C:\Users\nik>go get github.com/hashicorp/go-rootcerts

	C:\Users\nik>go get github.com/hashicorp/serf/coordinate

	C:\Users\nik>
	C:\Users\nik>go get github.com/mitchellh/mapstructure
	GOROOT=C:\Go #gosetup
	GOPATH=C:\go-work #gosetup
	C:\Go\bin\go.exe build -o C:\Users\nik\AppData\Local\Temp\___go_build_main_go__5_.exe
	C:/SRC/GO/AaronTorres/adv-go-sol/Section03/discovery/example/main.go #gosetup
	C:\Users\nik\AppData\Local\Temp\___go_build_main_go__5_.exe #gosetup

	OUTPUT:
	&api.AgentService{Kind:"", ID:"discovery", Service:"discovery", Tags:[]string{"Go", "Awesome"}, Meta:map[string]string(nil),
	Port:8080, Address:"127.0.0.1", Weights:api.AgentWeights{Passing:1, Warning:1}, EnableTagOverride:false, CreateIndex:0x1e,
	ModifyIndex:0x1e, ContentHash:"", ProxyDestination:"", Proxy:(*api.AgentServiceConnectProxyConfig)(0xc0000381e0),
	Connect:(*api.AgentServiceConnect)(0xc000004320)}

	Process finished with exit code 0

*/
func main() {
	config := consul.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	cli, err := discovery.NewClient(config, "127.0.0.1", "discovery", 8080)
	if err != nil {
		panic(err)
	}

	if err := discovery.Exec(cli); err != nil {
		panic(err)
	}
}
