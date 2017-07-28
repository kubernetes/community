#!/bin/bash

export CRTDIR=$(pwd)
export TMPDIR=/tmp/testgendocs
mkdir $TMPDIR

cp -r sig* Makefile generator $TMPDIR

cd $TMPDIR
make 1>/dev/null

mismatches=0
break=$(printf "=%.0s" $(seq 1 68))

for file in $(ls $CRTDIR/sig-*/README.md $CRTDIR/sig-list.md); do
  real=${file#$CRTDIR/}
  if ! diff -q $file $TMPDIR/$real &>/dev/null; then
    echo "$file does not match $TMPDIR/$real";
    mismatches=$((mismatches+1))
  fi;
done

if [ $mismatches -gt "0" ]; then
  echo ""
  echo $break
  noun="mismatch was"
  if [ $mismatches -gt "0" ]; then
    noun="mismatches were"
  fi
  echo "$mismatches $noun detected."
  echo "Do not manually edit sig-list.md or README.md files inside the sig folders."
  echo "Instead make your changes to sigs.yaml and then run \`make\`.";
  echo $break
  exit 1;
fi

rm -rf $TMPDIR
exit 0
