<!-----



Conversion time: 0.359 seconds.


Using this Markdown file:

1. Paste this output into your source file.
2. See the notes and action items below regarding this conversion run.
3. Check the rendered output (headings, lists, code blocks, tables) for proper
   formatting and use a linkchecker before you publish this page.

Conversion notes:

* Docs to Markdown version 1.0β35
* Sun Mar 03 2024 15:52:16 GMT-0800 (PST)
* Source doc: Machine Database
----->



# Machine Database

<p style="text-align: right">
Brotherlogic</p>


<p style="text-align: right">
2024-02-29</p>


<p style="text-align: right">
Draft</p>



### Abstract

Define a DB that detects machines on the network, assists with setup and provisioning automatically. Part of the everything should be automated plan


### Goals



1. Identify new machines on the network
2. Detect when we can’t provision said machine
3. Add machines into the ansible DB automatically and have them run and provision
4. Build out cluster machines automatically
5. [Stretch] Automatically reserve IP for new machines / existing machines
6. Use the MDB as the source of truth for setting up machines


### Definition

Machines have:



* Name
* IP
* Use
    * Cluster
        * Part of Kubernetes - should explicit specify which cluster they belong to
    * Discover
        * Part of the old discover system
    * Development
        * Used for development - should specify if this is a server or visual machine
    * Games
        * Used for playing games
    * Device
        * Supports an iOT device - e.g. the receipt printer
* Other features
    * Provisioner
        * Device is able to provision other devices (usually a primary and secondary provisioner)


### Process

MDB is a distributed datastore that lives in a github private repo, and gets infrequent updates. We also have a local kubernetes service which exposes the db internally - mdb.brotherlogic-backend.com read only. Kubernetes services reads updates from github and then updates itself and hard syncs every hour.

Kubernetes mdb job runs a ip broadcast and performs the assessment of connected devices, raising high priority issues if it detects new devices. Calls into other services in the local network to determine if other devices are ready. Also read by the ansible job when building out the infra page, and read by the cluster builder when building out the kubernetes cluster


### Graphs



1. Number of total machines in the database
2. Machines per use


### Milestones



1. Created mdb repo in github
2. Added this proposal
3. Added proto definition
4. Active IPs are detected in the local network
5. Active IPs are resolved down to actual machine types
6. Machine types are populated in the database
7. Skeleton MDB job is ready for kube
8. Flux loads the MDB server
9. MDB job loads the database into memory
10. Grafana graphs are displayed
11. Updates the ansible database from the MDB
12. Cluster builder source of truth is pulled from MDB
13. Clusters automatically rebuild