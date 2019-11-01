#!/bin/sh

SIZES=(128 256)

docker build -t svgconvertor:latest -f tools/Dockerfile .

total=$(ls svg/*/*/* | wc -l)
counter=1

for size in "${SIZES[@]}"; do 
    for svg in $(ls svg/*/*/* | cut -b 5-); do
        dir=$(echo $svg | cut -d / -f-2)
        mkdir -p png/$dir
        output=png/${svg%%.*}-$size.png
        echo "[$counter/$total] Generating $output"
        docker run -v $(pwd)/svg:/convertor svgconvertor:latest  $svg -h $size -w $size > $output
        counter=$[$counter +1]
    done
done
