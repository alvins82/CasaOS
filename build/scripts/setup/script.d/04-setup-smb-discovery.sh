#!/bin/bash

# SMB discovery is an optional enhancement. This script must never prevent
# the core CasaOS services from being installed when a distribution does not
# provide Avahi or wsdd.
set -u

readonly SMB_CONFIG_PATH="/etc/samba/smb.conf"
readonly DISCOVERY_CONFIG_PATH="/etc/casaos/smb-discovery.conf"
readonly DISCOVERY_CONFIG_SAMPLE="/etc/casaos/smb-discovery.conf.sample"
readonly CASAOS_SMB_MARKER="# Copyright (c) 2021-2022 CasaOS Inc. All rights reserved."
readonly DISCOVERY_SERVICE="casaos-wsdd.service"

warning() {
    echo "CasaOS SMB discovery: $1" >&2
}

ensure_discovery_config() {
    if [[ ! -e "${DISCOVERY_CONFIG_PATH}" && -f "${DISCOVERY_CONFIG_SAMPLE}" ]]; then
        cp -p "${DISCOVERY_CONFIG_SAMPLE}" "${DISCOVERY_CONFIG_PATH}" || \
            warning "could not create ${DISCOVERY_CONFIG_PATH}"
    fi
}

ensure_samba_mdns() {
    local first_line backup_path

    [[ -f "${SMB_CONFIG_PATH}" ]] || return 0
    first_line=$(sed -n '1p' "${SMB_CONFIG_PATH}")
    [[ "${first_line}" == *"${CASAOS_SMB_MARKER}"* ]] || return 0

    if grep -Eiq '^[[:space:]]*multicast[[:space:]]+dns[[:space:]]+register[[:space:]]*=' "${SMB_CONFIG_PATH}"; then
        return 0
    fi

    backup_path=$(mktemp) || {
        warning "could not create a temporary Samba configuration backup"
        return 0
    }
    if ! cp -p "${SMB_CONFIG_PATH}" "${backup_path}"; then
        rm -f "${backup_path}"
        warning "could not back up ${SMB_CONFIG_PATH} before enabling mDNS"
        return 0
    fi

    if ! sed -i '/^[[:space:]]*\[global\][[:space:]]*$/a\   multicast dns register = yes' "${SMB_CONFIG_PATH}" || \
        ! grep -Eiq '^[[:space:]]*multicast[[:space:]]+dns[[:space:]]+register[[:space:]]*=' "${SMB_CONFIG_PATH}"; then
        cp -p "${backup_path}" "${SMB_CONFIG_PATH}" || true
        rm -f "${backup_path}"
        warning "could not enable Samba mDNS registration in ${SMB_CONFIG_PATH}"
        return 0
    fi

    if command -v testparm >/dev/null 2>&1 && ! testparm -s "${SMB_CONFIG_PATH}" >/dev/null 2>&1; then
        cp -p "${backup_path}" "${SMB_CONFIG_PATH}" || warning "could not restore ${SMB_CONFIG_PATH}"
        rm -f "${backup_path}"
        warning "Samba configuration validation failed after enabling mDNS"
        return 0
    fi

    rm -f "${backup_path}"

    if systemctl is-active --quiet smbd 2>/dev/null; then
        systemctl restart smbd >/dev/null 2>&1 || warning "could not restart smbd after enabling mDNS"
    fi
}

enable_avahi() {
    if ! command -v avahi-daemon >/dev/null 2>&1; then
        warning "avahi-daemon is unavailable; mDNS SMB discovery is disabled"
        return 0
    fi

    systemctl enable --now avahi-daemon >/dev/null 2>&1 || \
        warning "could not enable/start avahi-daemon"
}

enable_wsdd() {
    if ! command -v wsdd >/dev/null 2>&1; then
        systemctl disable --now "${DISCOVERY_SERVICE}" >/dev/null 2>&1 || true
        warning "wsdd is unavailable; Windows SMB discovery is disabled"
        return 0
    fi

    systemctl daemon-reload >/dev/null 2>&1 || warning "could not reload systemd"
    systemctl enable "${DISCOVERY_SERVICE}" >/dev/null 2>&1 || {
        warning "could not enable ${DISCOVERY_SERVICE}"
        return 0
    }

    if systemctl is-active --quiet smbd 2>/dev/null; then
        systemctl start "${DISCOVERY_SERVICE}" >/dev/null 2>&1 || \
            warning "could not start ${DISCOVERY_SERVICE}"
    else
        warning "smbd is not active; ${DISCOVERY_SERVICE} will start with smbd"
    fi
}

configure_firewall() {
    local rule zone

    if command -v ufw >/dev/null 2>&1 && ufw status 2>/dev/null | grep -qi '^Status: active'; then
        for rule in 5353/udp 3702/udp 5357/tcp; do
            if ! ufw status 2>/dev/null | grep -Fq "${rule}"; then
                ufw allow "${rule}" >/dev/null 2>&1 || warning "could not add UFW rule for ${rule}"
            fi
        done
    fi

    if command -v firewall-cmd >/dev/null 2>&1 && [[ "$(firewall-cmd --state 2>/dev/null || true)" == "running" ]]; then
        zone=$(firewall-cmd --get-default-zone 2>/dev/null || true)
        if [[ -z "${zone}" ]]; then
            warning "could not determine the firewalld default zone"
            return 0
        fi

        for rule in 5353/udp 3702/udp 5357/tcp; do
            firewall-cmd --quiet --permanent --zone="${zone}" --add-port="${rule}" || \
                warning "could not add firewalld rule for ${rule}"
        done
        firewall-cmd --quiet --reload || warning "could not reload firewalld"
    fi
}

ensure_discovery_config
ensure_samba_mdns
enable_avahi
enable_wsdd
configure_firewall
