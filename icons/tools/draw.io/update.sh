# ! /bin/sh

# From git root, run:`make generate-icons-drawio` or
# `generate-icons-drawio-dockerized`.

SIZE=80

LTOTAL=$(ls -d svg/*/* | wc -l | sed 's/^ *//')
LCOUNT=1
for DIR in $(ls -d svg/*/* | cut -b 5-); do
    # draw.io library names default to the imported file name, so make them
    # human readable.
    OUTPUT="tools/draw.io/K8S $(echo $DIR | sed -e 's/[_/]/ /g' | awk '{for(i=1;i<=NF;i++){ $i=toupper(substr($i,1,1)) substr($i,2) }}1')"
    LNAME=$(basename "$OUTPUT")

    [ ! -f "$OUTPUT" ] || rm "$OUTPUT"
    echo "Library [$LCOUNT/$LTOTAL] Generating '$LNAME'"

    FTOTAL=$(ls svg/$DIR/* | wc -l | sed 's/^ *//')
    FCOUNT=1
    DATA=
    for SVG in svg/$DIR/*; do
        FNAME=$(basename "$SVG" | awk -F'.' '{print $1}')
        echo "- File [$FCOUNT/$FTOTAL] Adding '$FNAME' data"

        # Construct JSON obect data. We remove the first 3 lines to base64 just
        # the SVG XML object.
        BASE64=$(tail -n +4 "$SVG" | base64 | tr -d \\n)
        DATA="${DATA}{\"data\":\"data:image/svg+xml;base64,${BASE64}\",\"w\":$SIZE,\"h\":$SIZE,\"title\":\"$FNAME\",\"aspect\":\"fixed\"},"

        FCOUNT=$((FCOUNT+1))
    done

    # The last object should not have a trailing comma.
    DATA=${DATA%?}

    # Reproduce draw.io library file structure.
    echo "<mxlibrary>[${DATA}]</mxlibrary>" >> "$OUTPUT"
    LCOUNT=$((LCOUNT+1))
done
