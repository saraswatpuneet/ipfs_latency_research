# IPFS Gateway Research

## Context and Scope

Frequency enables network partners to exchange data in a decentralized manner. This entail IPFS services, which are used to store and retrieve data. The purpose of this research document is to perform a study on the IPFS gateway service, given pinned files on the IPFS network can take a while before they are available via gateway services like infura, cloudflare, or other services. In this document we will study the following aspects:
    - Average throughput of the gateway service.
    - Average latency of the gateway service.
    - Results for pinned files on the IPFS network for smaller and larger files.

## Background

Various http gateways backing IPFS network can be found [here](https://ipfs.github.io/public-gateway-checker/) and [here](https://luke.lol/ipfs.php). While the scope of this study is not to evaluate the thoroughness of the gateway services, it is to evaluate the performance of the gateway services.
For this research we choose to use the [infura](https://infura.io/) and [cloudflare](https://www.cloudflare.com/) gateways.

## Criteria

At time of this study following active gateways were evaluated from the list obtained from [here](https://ipfs.github.io/public-gateway-checker/). The list only included gateways with verified origin IP addresses.

| Gateway|
|-------|
| [CF](cf-ipfs.com)|
| [CloudFare]([infura.io](https://cloudflare-ipfs.com/ipfs/))|
| [Infura](infura-ipfs.io)|

Following measurements were performed on the following gateways:

* Average latency i.e. time to first byte.
* Average throughput i.e. bytes per second to download a file which includes latency.
* Check for content-length in response header.
* Any download failures were observed when test script was run few times.

### IPFS Node Setup

Created two IPFS node as digital ocean [droplet](https://www.digitalocean.com/droplets)s. One in Bangalore India and another node in San Francisco California.

Following configuration of droplets were to run IPFS node:

![Node_Bangalore](ipfs1.png)
![Node_San_Francisco](ipfs2.png)

Follow this perfect [tutorial](https://medium.com/pinata/how-to-deploy-an-ipfs-node-on-digital-ocean-c59b9e83098e) to setup ipfs node on digital ocean droplet.

### Files to Pin

Four files  named a, b, c and d with sizes 10MB, 50MB, 150MB and 400MB were pinned on the IPFS network.

## Results

Complete results for respective pinned files are recorded on a_file.csv, b_file.csv, c_file.csv and d_file.csv. Following are numbers recorded for different file sizes

## Smaller File A; 10 MB and 50 MB
