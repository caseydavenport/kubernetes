Bare Metal CoreOS with Kubernetes and Project Calico
------------------------------------------
This guide explains how to deploy a bare-metal Kubernetes cluster on CoreOS using [Calico networking](http://www.projectcalico.org).

Specifically, this guide will have you do the following:
- Deploy a Kubernetes master node on CoreOS using cloud-config
- Deploy two Kubernetes compute hosts with Calico Networking using cloud-config
- Provision pods on your Kubernetes cluster and check connectivity.

## Prerequisites
1. At least three bare-metal machines (or VMs) to work with.
	- 1 Kubernetes Master
	- 2 Kubernetes compute hosts
2. Your nodes should have IP connectivity over a Layer 2 network.
	- Connectivity over a L3 network is also supported, but not discussed in this guide.

## Cloud-config
This guide will use [cloud-config](https://coreos.com/docs/cluster-management/setup/cloudinit-cloud-config/) to configure each of the nodes in our Kubernetes cluster.

For ease of distribution, the cloud-config files required for this demonstration can be found on [GitHub](https://github.com/Metaswitch/calico-kubernetes-demo/tree/master/coreos).  

This repo includes two cloud config files:
- `master-config.yaml`: Cloud-config for the Kubernetes master
- `node-config.yaml`: Cloud-config for each Kubernetes compute host

## Download CoreOS
First, lets download the CoreOS bootable ISO.  We'll use this image to boot and install CoreOS on each server.
```
wget http://stable.release.core-os.net/amd64-usr/current/coreos_production_iso_image.iso
```
You can also download the ISO from the [CoreOS website](https://coreos.com/docs/running-coreos/platforms/iso/).

## Configure the Kubernetes Master
Once you've downloaded the image, use it to boot your Kubernetes Master server.  Once booted, you should be automatically logged in as the `core` user.

Let's get the master-config.yaml and fill in the necessary variables.  You can pull the config files from the git repo directly to the Kubernetes Master server.
```
git clone https://github.com/Metaswitch/calico-kubernetes-demo.git
cd calico-kubernetes-demo/coreos
``` 

You'll need to replace the following variables in the `master-config.yaml` file to match your deployment.
- `<SSH_PUBLIC_KEY>`: The public key you will use for SSH access to this server.
- `<DEFAULT_IPV4>`: The Master node's IPv4 address.
- `<KUBERNETES_LOC>`: The address used to get the kubernetes binaries over HTTP. 

The CoreOS bootable ISO comes with a tool called `coreos-install` which will allow us to install CoreOS to disk and configure the install using cloud-config.  The following command will download and install stable CoreOS, using the master-config.yaml file for configuration.
```
sudo coreos-install -d /dev/sda -C stable -c master-config.yaml
```

Once complete, restart the server.  When it comes back up, you should have SSH access as the `core` user using the public key provided in the master-config.yaml file.

## Configure the compute hosts
The following steps will set up a Kubernetes node for use as a compute host.  This demo uses two compute hosts, so you should run the following steps on each.

First, boot up your node using the bootable ISO we downloaded earlier.  You should be automatically logged in as the `core` user.

Let's get the `node-config.yaml` cloud-config file and fill in the necessary variables.  You can pull from the git repo directly to the Kubernetes node.
```
git clone https://github.com/Metaswitch/calico-kubernetes-demo.git
cd calico-kubernetes-demo/coreos
``` 

You'll need to replace the following variables in the `node-config.yaml` file to match your deployment.
- `<SSH_PUBLIC_KEY>`: The public key you will use for SSH access to this server.
- `<DEFAULT_IPV4>`: This node's IPv4 address.
- `<MASTER_IP>`: The IPv4 address of the Kubernetes master.
- `<KUBERNETES_IP>`: The IP address to use in order to get the kubernetes binaries over HTTP.

Install and configure CoreOS on the node using the following command.
```
sudo coreos-install -d /dev/sda -C stable -c node-config.yaml
```

Once complete, restart the server.  When it comes back up, you should have SSH access as the `core` user using the public key provided in the master-config.yaml file.

## Testing the Cluster
You should now have a functional bare-metal Kubernetes cluster with one master and two compute hosts.

TODO: Testing.
