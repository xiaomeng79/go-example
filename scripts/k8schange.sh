#!/bin/bash
#镜像变动后，版本没变，k8s不会拉取镜像，需要改变一个随机数
sed -i "/- name: RAND_NUM/{ n;s/\(value: \).*/\1num$RANDOM/ }" $1.yaml
