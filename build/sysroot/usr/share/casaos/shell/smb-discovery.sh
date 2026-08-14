#!/bin/bash

set -euo pipefail

readonly CONFIG_PATH="/etc/casaos/smb-discovery.conf"
readonly SMB_CONFIG_PATH="/etc/samba/smb.conf"

find_configured_interfaces() {
    if [[ ! -r "${CONFIG_PATH}" ]]; then
        return 0
    fi

    awk -F= '
        /^[[:space:]]*[Ii]nterfaces[[:space:]]*=/ {
            value = $2
            gsub(/[[:space:]]*,[[:space:]]*/, " ", value)
            gsub(/^[[:space:]]+|[[:space:]]+$/, "", value)
            print value
            exit
        }
    ' "${CONFIG_PATH}"
}

find_default_interfaces() {
    if command -v ip >/dev/null 2>&1; then
        ip -o route show default 2>/dev/null | awk '
            $1 == "default" {
                for (i = 1; i < NF; i++) {
                    if ($i == "dev") {
                        print $(i + 1)
                    }
                }
            }
        ' | awk '!seen[$0]++'
        return 0
    fi

    if command -v route >/dev/null 2>&1; then
        route -n 2>/dev/null | awk '$1 == "0.0.0.0" { print $8 }' | awk '!seen[$0]++'
    fi
}

find_fallback_interfaces() {
    local path interface

    for path in /sys/class/net/*; do
        [[ -e "${path}" ]] || continue
        interface=${path##*/}
        case "${interface}" in
            lo|docker*|br-*|veth*|virbr*|tun*|tap*|wg*)
                continue
                ;;
        esac
        [[ -e "${path}/device" ]] && printf '%s\n' "${interface}"
    done
}

find_workgroup() {
    local workgroup=""

    if command -v testparm >/dev/null 2>&1; then
        workgroup=$(testparm -s --parameter-name=workgroup 2>/dev/null | tail -n 1 | tr -d '[:space:]' || true)
    fi

    if [[ -z "${workgroup}" && -r "${SMB_CONFIG_PATH}" ]]; then
        workgroup=$(awk -F= '
            $1 ~ /^[[:space:]]*[Ww]orkgroup[[:space:]]*$/ {
                value = $2
                gsub(/^[[:space:]]+|[[:space:]]+$/, "", value)
                print value
                exit
            }
        ' "${SMB_CONFIG_PATH}" || true)
    fi

    printf '%s\n' "${workgroup:-WORKGROUP}"
}

find_hostname() {
    local hostname_value=""

    if command -v hostname >/dev/null 2>&1; then
        hostname_value=$(hostname -s 2>/dev/null || hostname 2>/dev/null || true)
    fi

    printf '%s\n' "${hostname_value}"
}

wsdd_bin=$(command -v wsdd || true)
if [[ -z "${wsdd_bin}" ]]; then
    echo "CasaOS SMB discovery: wsdd is not installed; Windows discovery is disabled." >&2
    exit 0
fi

interfaces=$(find_configured_interfaces || true)
if [[ -z "${interfaces}" ]]; then
    interfaces=$(find_default_interfaces || true)
fi
if [[ -z "${interfaces}" ]]; then
    interfaces=$(find_fallback_interfaces || true)
fi

wsdd_args=(--shortlog --workgroup "$(find_workgroup)")
hostname_value=$(find_hostname)
if [[ -n "${hostname_value}" ]]; then
    wsdd_args+=(--hostname "${hostname_value}")
fi
for interface in ${interfaces}; do
    [[ -n "${interface}" ]] && wsdd_args+=(--interface "${interface}")
done

exec "${wsdd_bin}" "${wsdd_args[@]}"
