#! /usr/bin/fish

set template template.json

set non_arrays (jq -r 'to_entries | map(select(.value | type != "array")) | .[].key' $template)
set arrays (jq -r 'to_entries | map(select(.value | type == "array")) | .[].key' $template)

set header (string replace -a ' ' ',' (echo $non_arrays))

echo $header > packages/csv/packages.csv

for file in packages/json/*
    set i (math $i + 1)
    set -l result ""

    for property in $non_arrays
        set value (jq -r .$property $file)
        set value (string replace -a \" \\\" $value)
        set result "$result\"$value\","
    end

    set result (string trim -r -c ',' $result)
    echo $result >> packages/csv/packages.csv
end

#for property in $arrays
#    echo package,$property > packages/csv/$property.csv
#    for file in packages/json/*
#        for value in (jq -r .$property\[\] $file)
#            set value (string replace -a \" \\\" $value)
#            set result (jq .pkgname $file)
#            set result "$result,\"$value\""
#            echo $result >> packages/csv/$property.csv
#        end
#    end
#end
