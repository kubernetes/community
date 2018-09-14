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
		echo "\"enable-dirs\": [ \"/certs/HTTP\" ],"
		echo "}"
		;;
	get)
		. /etc/custodia/custodia-cli.conf
		mkdir -p "$CUSTODIA_KRB5CCDIR"
		export KRB5CCNAME="$CUSTODIA_KRB5CCDIR/krb5cc_%{uid}"
		klist -s || kinit -kt "$CUSTODIA_CLIENT_KEYTAB"
		/usr/bin/custodia-cli --cafile "$CUSTODIA_CAFILE" --gssapi --server "https://${CUSTODIA_SERVER}/custodia/" get "kubernetes-secrets/$3/$4/$2"
		;;
	*)
		usage
esac
exit
