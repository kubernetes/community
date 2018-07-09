#!/bin/sh
SIZE=256

docker build -t svgconvertor:latest -f tools/Dockerfile .

total=$(ls svg/*/*/* | wc -l)
counter=1

for svg in $(ls svg/*/*/* | cut -b 5-); do
	dir=$(echo $svg | cut -d / -f-2)
	mkdir -p png/$dir
	output=png/${svg%%.*}-$SIZE.png
	echo "[$counter/$total] Generating $output"
	docker run -v $(pwd)/svg:/convertor svgconvertor:latest  $svg -h $SIZE -w $SIZE > $output
	counter=$[$counter +1]
done

