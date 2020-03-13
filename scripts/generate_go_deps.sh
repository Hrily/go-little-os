#!/bin/sh

get_dependencies () {
	cat $1/*.go | \
		tr '\n' ' ' | \
		tr '\t' ' ' | \
		tr -s ' ' ' ' | \
		# grep -oP '(?<=import \()[^\)]+|(?<=import \")[^\"]+' | \
		perl -nle'print $& while m{(?<=import \()[^\)]+|(?<=import \")[^\"]+}g' | \
		tr ' ' '\n' | \
		tr -d '"' | \
		grep "kernel/" | \
		# xargs -n1 dirname | \
		sed -e 's/$/.o #/g' | \
		sed '$ s/#$//g' | \
		sed 's|#|\\ \\n|g' | \
		sort | uniq
}

get_dependency_rule () {
	echo $1 ":" $2
}

DIR=$1
DIR=${DIR%/}

OUTPUT=$DIR/deps.mk

OBJECT=$DIR.o

deps=`get_dependencies $DIR`

echo $OBJECT ":" $deps | \
	awk '{$1=$1;print}' | \
	sed $'s|^|\t|g'  | \
	sed $'1 s|\t||g' > $OUTPUT
