Bare Metal CoreOS with Kubernetes and Project Calico
------------------------------------------
Deploy a Kubernetes cluster on CoreOS using [Calico networking]().  This guide provisions a simple cluster on bare-metal to test Kubernetes with Calico networking.

In this guide, you will do the following:
- Deploy a Kubernetes master node on CoreOS using cloud-config
- Deploy two Kubernetes compute hosts with Calico Networking using cloud-config
- Provision pods on your Kubernetes cluster and check connectivity.

**Table of Contents**

- [Prerequisites](#prerequisites)

## Prerequisites
1. At least three bare-metal nodes (or VMs) to work with.
	- 1 Kubernetes Master
	- 2 Kubernetes compute hosts
2. Your nodes should have IP connectivity over a Layer 2 network.
	- Connectivity over a L3 network is also supported, but not discussed in this guide.


## Configuring the Kubernetes Master node 
This guide will use [cloud-config] to configure each of the nodes in our Kubernetes cluster.

For the Kubernetes Master node, we'll use the following cloud-config file.  Replace the following variables with the correct values for your cluster:
- <SSH_PUBLIC_KEY>: The public key you will use for SSH access to this server.
- <DEFAULT_IPV4>: The Master's IPv4 address.
- <KUBERNETES_IP>: The IP address to use in order to get the kubernetes binaries over HTTP.
```
#cloud-config
---
hostname: kube-master
users:
  - name: core
    ssh-authorized-keys:
      - ssh-rsa <SSH_PUBLIC_KEY> 
    sudo: ['ALL=(ALL) NOPASSWD:ALL']
    groups: sudo
    shell: /bin/bash

write_files:
  - path: /etc/network-environment
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      # The master's IPv4 address - reachable by the kubernetes nodes.
      DEFAULT_IPV4=<DEFAULT_IPV4>

      # Location of the kubernetes binaries.
      KUBERNETES_IP=<KUBERNETES_IP>

  - path: /home/core/kubernetes-download.sh
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      # Download kubernetes binaries from the following location.
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kubectl"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kubernetes"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kube-controller-manager"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kube-apiserver"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kube-scheduler"
      chmod +x /home/core/*
      exit 0

  - path: /etc/profile.d/opt-path.sh
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      PATH=$PATH/home/core

coreos:
  update:
    reboot-strategy: off
  units:
    - name: get-kube-tools.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=Installs Kubernetes tools
        After=network-online.service
        [Service]
        EnvironmentFile=/etc/network-environment
        ExecStart=/home/core/kubernetes-download.sh
        RemainAfterExit=yes
        Type=oneshot

    - name: etcd.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=etcd
        After=network-online.service
        [Service]
        EnvironmentFile=/etc/network-environment
        User=etcd
        PermissionsStartOnly=trueExecStart=/usr/bin/etcd \
        --name ${DEFAULT_IPV4} \
        --addr ${DEFAULT_IPV4}:4001 \
        --bind-addr 0.0.0.0 \
        --cluster-active-size 1 \
        --data-dir /var/lib/etcd \
        --http-read-timeout 86400 \
        --peer-addr ${DEFAULT_IPV4}:7001 \
        --snapshot true
        Restart=always
        RestartSec=10s

    - name: kube-apiserver.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=Kubernetes API Server
        Documentation=https://github.com/GoogleCloudPlatform/kubernetes
        Requires=etcd.service
        After=etcd.service
        [Service]
        ExecStart=/home/core/kube-apiserver \
        --address=0.0.0.0 \
        --port=8080 \
        --etcd_servers=http://127.0.0.1:4001 \
        --logtostderr=true
        Restart=always
        RestartSec=10

    - name: kube-controller-manager.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=Kubernetes Controller Manager
        Documentation=https://github.com/GoogleCloudPlatform/kubernetes
        Requires=etcd.service
        After=etcd.service
        [Service]
        ExecStart=/home/core/kube-controller-manager \
        --master=127.0.0.1:8080 \
        --logtostderr=true
        Restart=always
        RestartSec=10

    - name: kube-scheduler.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=Kubernetes Scheduler
        Documentation=https://github.com/GoogleCloudPlatform/kubernetes
        Requires=etcd.service
        After=etcd.service
        [Service]
        ExecStart=/home/core/kube-scheduler --master=127.0.0.1:8080
        Restart=always
        RestartSec=10
```

## Configuring the Kubernetes compute nodes
For the Kubernetes compute nodes, we'll use the following cloud-config file.  Replace the following variables with the correct values for your cluster:
- <SSH_PUBLIC_KEY>: The public key you will use for SSH access to this server.
- <DEFAULT_IPV4>: The Master's IPv4 address.
- <KUBERNETES_IP>: The IP address to use in order to get the kubernetes binaries over HTTP.
```
#cloud-config
---
hostname: kube-node1
users:
  - name: core
    ssh-authorized-keys:
      - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDa6B24q6vqsarrwEhHQEkrIgxR5iPm6Ay9grMehZouPd1zqRB4S6zU13k5uphUj3TT5z2Pv6r0Rj026tuPofKfNA4SEidBRqmWniWusBDRY/AFc+6lCfn/jd+d+Jgw0RWwQUt/g27Qm2cbiO8b/yqtScIGr/pCkm52rRj+CLe4/u7YbVk36UHvDt7WYKeVBPh8S3YtSqbjhOJeARMBMQrGVl/eSmaVwpI5lGPtqCS8c76ovrGZbAIEVHPfQWHBewzEFYLQo/XUQolNaFVzLpNKHtna8+v/XLFxk/eLZe2tW41nQQd42GYzJ9tMiCi5ZSeSfLINjKPU/0bswivyiG1H cd4@illium
    sudo: ['ALL=(ALL) NOPASSWD:ALL']
    groups: sudo
    shell: /bin/bash

write_files:
  - path: /etc/network-environment
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      # This node's IPv4 address
      DEFAULT_IPV4=172.24.114.126

      # The kubernetes master IP
      KUBERNETES_MASTER=172.24.114.11

      # Location of etcd cluster used by Calico.  By default, this assumes etcd is
      # running locally on the compute host
      ETCD_AUTHORITY=127.0.0.1:4001

      # Location from which to download the Kubernetes binaries
      KUBERNETES_IP=172.24.114.228

  - path: /home/core/kubernetes-download.sh
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      # Download kubernetes binaries
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kubectl"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kubernetes"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kube-proxy"
      /usr/bin/wget -N -P "/home/core" "http://$KUBERNETES_IP/kubelet"
      chmod +x /home/core/*
      exit 0

  - path: /home/core/calico-download.sh
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      # Download calicoctl and the calico kubernetes networking plugin.
      /usr/bin/wget -N -P "/home/core" "https://github.com/Metaswitch/calico-docker/releases/download/v0.4.7/calicoctl"

      # TODO: Currently using my own fork of the plugin.
      #/usr/bin/wget -N -P "/home/core" "https://github.com/Metaswitch/calico-docker/releases/download/v0.4.2/calico_kubernetes"
      /usr/bin/wget -N -P "/home/core" "http://172.24.114.228/calico_kubernetes"
      chmod +x /home/core/calicoctl

      # Move the networking plugin into the networking plugins directory.  This directory is passed to the kubelet so it
      # can find the plugin.
      mkdir -p /home/core/kubelet-net-plugins/calico
      mv /home/core/calico_kubernetes /home/core/kubelet-net-plugins/calico/calico
      chmod +x /home/core/kubelet-net-plugins/calico/calico

      # FIXME: Temporary fix to get around hardcoded calicoctl location in the calico networking plugin.
      sudo mkdir -p /home/vagrant
      ln -s /home/core/calicoctl /home/vagrant/calicoctl

      exit 0

  - path: /home/core/calico-node.sh
    owner: root
    permissions: 0755
    content: |#! /usr/bin/bash
      sudo modprobe xt_set
      sudo modprobe ip6_tables
      sudo /home/core/calicoctl node --ip=$1
      touch /home/calico-node-$1
      exit 0

  - path: /etc/profile.d/opt-path.sh
    owner: root
    permissions: 0755
    content: |
      #! /usr/bin/bash
      PATH=$PATH:/home/core

coreos:
  update:
    reboot-strategy: off
  etcd2:
    name: calico-etcd
    discovery: https://discovery.etcd.io/292dc93f4919a44327e0f22d54afd09f
    advertise-client-urls: http://172.24.114.126:2379
    initial-advertise-peer-urls: http://172.24.114.126:2380
    listen-client-urls: http://0.0.0.0:2379,http://0.0.0.0:4001
    listen-peer-urls: http://172.24.114.126:2380,http://172.24.114.126:7001

  units:
    - name: get-kube-tools.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=Downloads Kubernetes binaries
        After=network-online.service
        [Service]
        EnvironmentFile=/etc/network-environment
        ExecStartPre=-/usr/bin/mkdir -p /home/core
        ExecStart=/home/core/kubernetes-download.sh
        RemainAfterExit=yes
        Type=oneshot

    - name: get-calico-tools.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=Downloads the Calico binaries
        After=network-online.service
        [Service]
        ExecStartPre=-/usr/bin/mkdir -p /home/core
        ExecStart=/home/core/calico-download.sh
        RemainAfterExit=yes
        Type=oneshot

    - name: etcd2.service
      command: start

    - name: calico-node.service
      runtime: true
      command: start
      content: |
        [Unit]
        Description=calicoctl node
        After=etcd2.service
        [Service]
        EnvironmentFile=/etc/network-environment
        User=root
        PermissionsStartOnly=true
        ExecStart=/home/core/calico-node.sh $DEFAULT_IPV4
        RemainAfterExit=yes
        Type=oneshot

- name: kube-proxy.service
      command: start
      content: |
        [Unit]
        Description=Kubernetes Proxy
        Documentation=https://github.com/GoogleCloudPlatform/kubernetes
        After=calico-node.service
        [Service]
        EnvironmentFile=/etc/network-environment
        ExecStart=/home/core/kube-proxy --master=http://${KUBERNETES_MASTER}:8080 --logtostderr=true
        Restart=always
        RestartSec=10

    - name: kube-kubelet.service
      command: start
      content: |
        [Unit]
        Description=Kubernetes Kubelet
        Documentation=https://github.com/GoogleCloudPlatform/kubernetes
        After=calico-node.service
        [Service]
        EnvironmentFile=/etc/network-environment
        ExecStart=/home/core/kubelet \
        --v=5 \
        --address=0.0.0.0 \
        --port=10250 \
        --hostname_override=${DEFAULT_IPV4} \
        --api_servers=${KUBERNETES_MASTER}:8080 \
        --healthz_bind_address=0.0.0.0 \
        --healthz_port=10248 \
        --network_plugin=calico \
        --network_plugin_dir=/home/core/kubelet-net-plugins \
        --logtostderr=true
        Restart=always
        RestartSec=10
```
