{% if pillar.get('network_provider', '').lower() == 'calico' %}

calicoctl:
  file.managed:
    - name: /usr/bin/calicoctl
    - source: https://github.com/projectcalico/calico-docker/releases/download/v0.10.0/calicoctl
    - source_hash: sha512=5dd8110cebfc00622d49adddcccda9d4906e6bca8a777297e6c0ffbcf0f7e40b42b0d6955f2e04b457b0919cb2d5ce39d2a3255d34e6ba36e8350f50319b3896
    - makedirs: True
    - mode: 744

calico-node:
  cmd.run:
    - name: calicoctl node
    - env:
      - ETCD_AUTHORITY: "127.0.0.1:6666"
    - require:
      - kmod: ip6_tables
      - kmod: xt_set
      - service: docker
      - cmd: docker-available 
      - file: calicoctl
      - cmd: etcd

remove-stale-etcd:
  cmd.run:
    - name: docker ps -a | grep calico-etcd && ! docker ps | grep calico-etcd && docker rm calico-etcd || true
    - require:
      - cmd: docker-available

etcd:
  cmd.run:
    - require:
      - cmd: remove-stale-etcd
    - name: >
               docker ps | grep calico-etcd || 
               docker run --name calico-etcd -d --restart=always --net=host 
               -v /varetcd:/var/etcd
               gcr.io/google_containers/etcd:2.2.1
               /usr/local/bin/etcd --name calico
               --data-dir /var/etcd/calico-data
               --advertise-client-urls http://{{grains.kubelet_api_servers}}:6666
               --listen-client-urls http://0.0.0.0:6666
               --listen-peer-urls http://0.0.0.0:6667
               --initial-advertise-peer-urls http://{{grains.kubelet_api_servers}}:6667
               --initial-cluster calico=http://{{grains.kubelet_api_servers}}:6667

ip6_tables:
  kmod.present

xt_set:
  kmod.present

{% endif %}
