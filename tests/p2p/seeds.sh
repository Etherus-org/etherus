#! /bin/bash
set -eu

N=$1

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
P2PDIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"


seeds="$($P2PDIR/ip.sh 2):26656"

for i in $(seq 2 $N); do
	index=$(($i*2))
	TENDERMINT_IP=$($P2PDIR/ip.sh $index)

	seeds="$seeds,$TENDERMINT_IP:26656"
done

echo "$seeds"
