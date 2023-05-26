#!/bin/bash

SIZES=(128 256)

docker build -t svgconvertor:latest -f tools/Dockerfile .

total=$(($(ls svg/*/*/* | wc -l) * "${#SIZES[@]}"))
counter=1

for size in "${SIZES[@]}"; do 
    for svg in $(ls svg/*/*/* | cut -b 5-); do
        dir=$(echo $svg | cut -d / -f-2)
        mkdir -p png/$dir
        output=png/${svg%%.*}-$size.png
        echo "[$counter/$total] Generating $output"
        # Only specify the width, since we know heptagons are wider than they
        # are tall. The tool will automatically retain the source aspect ratio.
        docker run --rm -v $(pwd)/svg:/convertor svgconvertor:latest $svg -w $size > $output
        counter=$[$counter +1]
    done
done
