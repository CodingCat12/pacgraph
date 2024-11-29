#! /usr/bin/fish

function pacjson
    set -l pkg $argv[1]
    set -l pkg_info (pacman -Si $pkg 2>/dev/null)

    set -l out $argv[2]

    set -l repo (string collect $pkg_info | awk '/^Repository/ {print $3}')
    set -l arch (string collect $pkg_info | awk '/^Architecture/ {print $3}')

    curl -so $out "https://archlinux.org/packages/$repo/$arch/$pkg/json/"
end

set packages (pacman -Slq)
set -l iteration 0
for package in $packages
    set iteration (math $iteration + 1)
    set -l outfile packages/json/$package.json

    if test -s $outfile
        continue
    end

    pacjson $package $outfile
    echo $iteration/(echo $packages | wc -w)
end