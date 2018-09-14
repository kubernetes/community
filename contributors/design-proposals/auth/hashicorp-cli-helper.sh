#!/bin/bash

set -e

function usage {
	echo "Usage: $0 [ mount | get <path> <parameters> ]" >&2
	exit 1
}

case "$1" in
	mount)
		echo "{"
		echo "\"mount-param\": [ \"kubernetes.io/pod.namespace\", \"kubernetes.io/pod.name\" ],"
		echo "\"enable-dirs\": [ \"/db\" ],"
		echo "}"
		;;
	get)
		. /etc/hashicorp/hashicorp-cli.conf
		/usr/bin/hashicorp-vault read -field="${2##*/}" "secret/pod-secrets/$3/$4/${2%/*}"
		;;
	*)
		usage
esac
exit
